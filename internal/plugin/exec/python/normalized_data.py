import argparse
import os


def str_to_float(data):
    nums = []
    for i in range(len(data)):
        nums.append(float(data[i]))
    return nums


parser = argparse.ArgumentParser()

# Add parameters
parser.add_argument('params', nargs='+', help='参数列表')

# Parsing parameters
args = parser.parse_args()

# Accessing parameter arrays
disease = args.params[0]
disease = disease.replace(" ", "_")
params = args.params[1:]
params = str_to_float(params)

with open(f'{os.getcwd()}/internal/plugin/exec/python/column_uniformization_info/{disease}.tsv', 'r') as f:
    lines = f.readlines()


max = str_to_float(lines[0].strip().split('\t')[1:])
min = str_to_float(lines[1].strip().split('\t')[1:])


for index in range(0, len(params)):
    if params[index] > max[index]:
        params[index] = max[index]

result = []

for index in range(0, len(params)):
    # Less than 0, set to 0, greater than 1, set to 1
    molecule = params[index] - min[index]
    if molecule == 0:
        result.append(0)
        continue
    denominator = max[index] - min[index]
    result.append((params[index] - min[index]) / (max[index] - min[index]))

for index in range(0, len(result)):
    if result[index] < 0: result[index] = 0
    if result[index] > 1: result[index] = 1

print(result)
