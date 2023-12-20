import argparse
import math
import os

import torch

INT = ["int8", "int16", "int32", "int64"]
UINT = ["uint8", "uint16", "uint32", "uint64"]
Q_MAX = "Q_MAX"
Q_MIN = "Q_MIN"


def get_uint_type_range(uint_type: str) -> int:
    res = ""
    for i in uint_type:
        if str.isdigit(i):
            res += i
    return int(res)


QUANTIZE_RANGE = {UINT[0]: {Q_MIN: 0, Q_MAX: (1 << get_uint_type_range(UINT[0])) - 1},
                  UINT[1]: {Q_MIN: 0, Q_MAX: (1 << get_uint_type_range(UINT[1])) - 1},
                  UINT[2]: {Q_MIN: 0, Q_MAX: (1 << get_uint_type_range(UINT[2])) - 1},
                  UINT[3]: {Q_MIN: 0, Q_MAX: (1 << get_uint_type_range(UINT[3])) - 1}}


def get_max(x: list):
    x_max = x[0]
    for ele in x:
        if x_max < ele:
            x_max = ele
    return x_max


def get_min(x: list):
    x_min = x[0]
    for ele in x:
        if x_min > ele:
            x_min = ele
    return x_min


def calc_scale(x: list, target_type):
    q_min, q_max = get_quantize_range(target_type)
    # Calculate value range (denominator)
    x_range = get_max(x) - get_min(x)
    x_range = 1 if x_range == 0 else x_range

    # Calculate scale
    scale_molecule = q_max - q_min
    scale_denominator = x_range
    return scale_molecule, scale_denominator


def get_quantize_range(target_type):
    assert target_type in UINT or target_type in INT
    q_max = QUANTIZE_RANGE[target_type][Q_MAX]
    q_min = QUANTIZE_RANGE[target_type][Q_MIN]
    return q_min, q_max


def calc_zero_point(x: list, scale, target_type: str) -> int:
    q_min, q_max = get_quantize_range(target_type)
    # zero_point = (-scale_molecule / scale_denominator * get_min(x) + q_min).round()
    zero_point = (-scale * get_max(x) + q_max).__round__()
    return zero_point


def quantize(values: list, q_scale, q_zero_point: int, target_type: str):
    # detect input arg "target_type" in uint or int list
    q_min, q_max = get_quantize_range(target_type)
    x_quantize = []
    for value in values:
        x_quantize.append((value * math.ceil(q_scale) + q_zero_point).__round__())
    x_quantize = torch.clip(torch.tensor(x_quantize), min=q_min, max=q_max)

    return [int(x_ele) for x_ele in x_quantize]


def str_to_float(data):
    nums = []
    for i in range(len(data)):
        nums.append(float(data[i]))
    return nums


parser = argparse.ArgumentParser()

parser.add_argument('params', nargs='+', help='参数列表')

args = parser.parse_args()

disease = args.params[0]
disease = disease.replace(" ", "_")
params = args.params[1:]

x = [float(x) for x in params]

_type = "uint8"

# scale_molecule, scale_denominator = calc_scale(x, _type)
# scale = math.ceil(scale_molecule / scale_denominator)
# _zero_point = calc_zero_point(x, scale, _type)

with open(f'{os.getcwd()}/internal/plugin/exec/python/column_uniformization_info/{disease}.tsv', 'r') as f:
    lines = f.readlines()

scale = int(lines[2].strip().split(' ')[0])
_zero_point = int(lines[2].strip().split(' ')[1])

print("quantized: ", quantize(x, scale, _zero_point, _type))
print("scale: ", scale)
print("zero_point: ", _zero_point)
