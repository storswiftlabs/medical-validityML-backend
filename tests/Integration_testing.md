# MEDICAL-ZKML-BACKEND Integration Testing
## Overview of Integration Testing
    Testing the entire process from user registration to prediction and validation
## test process
### 1. User Registration
Call Interface: /api/user/register      
Input:
```json
{
    "address": "0x01f5BB073e893d334FF9b0e239939982c124AF97"
}
```
Output:
```json
{
    "ok": true,
    "msg": "login was successful",
    "data": null
}
```
### 2. Get disease information
Call Interface: /api/diseases/diseases      
Input: null     
Output:
```json
{
    "ok": true,
    "count": 9,
    "msg": "Query completed",
    "data": [
        {
            "Name": "Lymphography",
            "Description": ""
        },
        {
            "Name": "Parkinsons",
            "Description": ""
        },
        {
            "Name": "Liver_Disorders",
            "Description": ""
        },
        {
            "Name": "Heart_Disease",
            "Description": ""
        },
        {
            "Name": "Primary_Tumor",
            "Description": ""
        },
        {
            "Name": "Chronic_Kidney_Disease",
            "Description": ""
        },
        {
            "Name": "Acute_Inflammations",
            "Description": ""
        },
        {
            "Name": "Heart_Failure_Clinical_Records",
            "Description": ""
        },
        {
            "Name": "Breast_Cancer",
            "Description": ""
        }
    ]
}
```
Call the interface again        
Input： Lymphography     
Output:
```json
{
    "ok": true,
    "msg": "",
    "data": {
        "Name": "Lymphography",
        "Description": "",
        "Inputs": [
            {
                "Name": "class",
                "Description": "normal find, metastases, malign lymph, fibrosis",
                "Index": 0,
                "InputMethod": "select",
                "Select": [
                    {
                        "Key": "normal find",
                        "Value": 1
                    },
                    {
                        "Key": "metastases",
                        "Value": 2
                    },
                    {
                        "Key": "malign lymph",
                        "Value": 3
                    },
                    {
                        "Key": "fibrosis",
                        "Value": 4
                    }
                ]
            },
            {
                "Name": "lymphatics",
                "Description": "normal, arched, deformed, displaced",
                "Index": 1,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "normal",
                        "Value": 1
                    },
                    {
                        "Key": "arched",
                        "Value": 2
                    },
                    {
                        "Key": "deformed",
                        "Value": 3
                    },
                    {
                        "Key": "displaced",
                        "Value": 4
                    }
                ]
            },
            {
                "Name": "block of affere",
                "Description": "no, yes",
                "Index": 2,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 0
                    },
                    {
                        "Key": "yes",
                        "Value": 1
                    }
                ]
            },
            {
                "Name": "bl of lymph c",
                "Description": "no, yes",
                "Index": 3,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 0
                    },
                    {
                        "Key": "yes",
                        "Value": 1
                    }
                ]
            },
            {
                "Name": "bl of lymph s",
                "Description": "no, yes",
                "Index": 4,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 0
                    },
                    {
                        "Key": "yes",
                        "Value": 1
                    }
                ]
            },
            {
                "Name": "by pass",
                "Description": "no, yes",
                "Index": 5,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 0
                    },
                    {
                        "Key": "yes",
                        "Value": 1
                    }
                ]
            },
            {
                "Name": "extravasates",
                "Description": "no, yes",
                "Index": 6,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 0
                    },
                    {
                        "Key": "yes",
                        "Value": 1
                    }
                ]
            },
            {
                "Name": "regeneration of",
                "Description": "no, yes",
                "Index": 7,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 0
                    },
                    {
                        "Key": "yes",
                        "Value": 1
                    }
                ]
            },
            {
                "Name": "early uptake in",
                "Description": "no, yes",
                "Index": 8,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 0
                    },
                    {
                        "Key": "yes",
                        "Value": 1
                    }
                ]
            },
            {
                "Name": "lym nodes dimin",
                "Description": "0-3",
                "Index": 9,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "0",
                        "Value": 0
                    },
                    {
                        "Key": "1",
                        "Value": 1
                    },
                    {
                        "Key": "2",
                        "Value": 2
                    },
                    {
                        "Key": "3",
                        "Value": 3
                    }
                ]
            },
            {
                "Name": "lym nodes enlar",
                "Description": "1-4",
                "Index": 10,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "1",
                        "Value": 1
                    },
                    {
                        "Key": "2",
                        "Value": 2
                    },
                    {
                        "Key": "3",
                        "Value": 3
                    },
                    {
                        "Key": "4",
                        "Value": 4
                    }
                ]
            },
            {
                "Name": "changes in lym",
                "Description": "bean, oval, round",
                "Index": 11,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "bean",
                        "Value": 1
                    },
                    {
                        "Key": "oval",
                        "Value": 2
                    },
                    {
                        "Key": "round",
                        "Value": 3
                    }
                ]
            },
            {
                "Name": "defect in node",
                "Description": "no, lacunar, lac. marginal, lac. central",
                "Index": 12,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 1
                    },
                    {
                        "Key": "lacunar",
                        "Value": 2
                    },
                    {
                        "Key": "lac. marginal",
                        "Value": 3
                    },
                    {
                        "Key": "lac. central",
                        "Value": 4
                    }
                ]
            },
            {
                "Name": "changes in node",
                "Description": "no, lacunar, lac. margin, lac. central",
                "Index": 13,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 1
                    },
                    {
                        "Key": "lacunar",
                        "Value": 2
                    },
                    {
                        "Key": "lac. marginal",
                        "Value": 3
                    },
                    {
                        "Key": "lac. central",
                        "Value": 4
                    }
                ]
            },
            {
                "Name": "changes in stru",
                "Description": "no, grainy, drop-like, coarse, diluted, reticular, stripped, faint,",
                "Index": 14,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 1
                    },
                    {
                        "Key": "grainy",
                        "Value": 2
                    },
                    {
                        "Key": "drop-like",
                        "Value": 3
                    },
                    {
                        "Key": "coarse",
                        "Value": 4
                    },
                    {
                        "Key": "diluted",
                        "Value": 5
                    },
                    {
                        "Key": "reticular",
                        "Value": 6
                    },
                    {
                        "Key": "stripped",
                        "Value": 7
                    },
                    {
                        "Key": "faint",
                        "Value": 8
                    }
                ]
            },
            {
                "Name": "special forms",
                "Description": "no, chalices, vesicles",
                "Index": 15,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 1
                    },
                    {
                        "Key": "chalices",
                        "Value": 2
                    },
                    {
                        "Key": "vesicles",
                        "Value": 3
                    }
                ]
            },
            {
                "Name": "dislocation of",
                "Description": "no, yes",
                "Index": 16,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 0
                    },
                    {
                        "Key": "yes",
                        "Value": 1
                    }
                ]
            },
            {
                "Name": "exclusion of no",
                "Description": "no, yes",
                "Index": 17,
                "InputMethod": "",
                "Select": [
                    {
                        "Key": "no",
                        "Value": 0
                    },
                    {
                        "Key": "yes",
                        "Value": 1
                    }
                ]
            }
        ],
        "Output": {
            "Description": "no. of nodes in",
            "Result": [
                {
                    "Key": "0-9",
                    "Value": 0
                },
                {
                    "Key": "10-19",
                    "Value": 1
                },
                {
                    "Key": "20-29",
                    "Value": 2
                },
                {
                    "Key": "30-39",
                    "Value": 3
                },
                {
                    "Key": "40-49",
                    "Value": 4
                },
                {
                    "Key": "50-59",
                    "Value": 5
                },
                {
                    "Key": "60-69",
                    "Value": 6
                },
                {
                    "Key": ">=70",
                    "Value": 7
                }
            ]
        }
    }
}
```
### 3. Making disease predictions
Call Interface: /api/operator/disease_prediction        
input:
```json
{
    "user": "0x01f5BB073e893d334FF9b0e239939982c124AF97",
    "name": "Lymphography",
    "module": "decision_tree",
    "inputs": [
        {
            "name": "class",
            "index": 0,
            "select": "3"
        },
        {
            "name": "lymphatics",
            "index": 1,
            "select": "4"
        },
        {
            "name": "block of affere",
            "index": 2,
            "select": "2"
        },
        {
            "name": "bl of lymph c",
            "index": 3,
            "select": "1"
        },
        {
            "name": "bl of lymph s",
            "index": 4,
            "select": "1"
        },
        {
            "name": "by pass",
            "index": 5,
            "select": "1"
        },
        {
            "name": "extravasates",
            "index": 6,
            "select": "1"
        },
        {
            "name": "regeneration of",
            "index": 7,
            "select": "1"
        },
        {
            "name": "early uptake in",
            "index": 8,
            "select": "2"
        },
                {
            "name": "lym nodes dimin",
            "index": 9,
            "select": "1"
        },
        {
            "name": "lym nodes enlar",
            "index": 10,
            "select": "2"
        },
        {
            "name": "changes in lym",
            "index": 11,
            "select": "2"
        },
        {
            "name": "defect in node",
            "index": 12,
            "select": "2"
        },
        {
            "name": "changes in node",
            "index": 13,
            "select": "4"
        },
        {
            "name": "changes in stru",
            "index": 14,
            "select": "8"
        },
        {
            "name": "special forms",
            "index": 15,
            "select": "1"
        },
        {
            "name": "dislocation of",
            "index": 16,
            "select": "1"
        },
        {
            "name": "exclusion of no",
            "index": 17,
            "select": "2"
        }
    ]
}
```
Output:
```json
{
    "ok": true,
    "msg": "ok",
    "data": null
}
```
### 4. Get prediction results
Call Interface: /api/operator/predicting_outcomes       
Input: 
```json
{
    "user": "0x01f5BB073e893d334FF9b0e239939982c124AF97"
}
```
Output:
```json
{
  "ok": true,
  "count": 4,
  "msg": "ok",
  "data": [
    {
      "ID": 24,
      "Disease": "Acute_Inflammations",
      "Description": "",
      "EndTime": 1695553239,
      "Output": "-1"
    },
    {
      "ID": 25,
      "Disease": "Acute_Inflammations",
      "Description": "",
      "EndTime": 1695553242,
      "Output": "normal"
    },
    {
      "ID": 26,
      "Disease": "Acute_Inflammations",
      "Description": "",
      "EndTime": 1695873392,
      "Output": "normal"
    },
    {
      "ID": 27,
      "Disease": "Lymphography",
      "Description": "",
      "EndTime": 1696646605,
      "Output": "10-19"
    }
  ]
}
```
### 5. Get disease advice
Call Interface: /api/operator/recommend   
Input:
```json
{
    "disease": "Lymphography"
}
```
Output:
```json
{
    "ok": true,
    "count": 10,
    "msg": "ok",
    "data": [
        "bafkreibqfyqsxyqyzplpsdefxfqtvxxj4m4bgpe4qsmhkeoqoz54zvrede",
        "bafkreiheuc46hmngxxtkwk666lx4gcb4k6deoxkl45ed5glakpgdonp6ja",
        "bafkreiapb32g6cjhwttysg3cr6hbng4jpiirojb6hq6ylftfjj3hhdhpaa",
        "bafkreihdq3vf5ewcqkjng7wi4cx3lhma2c5h4n7rktmxovgzsm26xfmt7m",
        "bafkreiduiwmqknh4ghdyk2uwcyuawvpd7eltddyistanve2w6mhok7lygu",
        "bafkreihquuuvhkhstoj3vpfcp4ky2nvghlcumgb42epv664nljpoxphy7a",
        "bafkreiaqsrzyz57rg3rtkpsjtc7rehe5l7sya5l46bcorqtr66ckvz4koq",
        "bafkreig6yf5tdls6paweqflvt5f5urwpgtm4h3fyzndbisph4f3wq5r3n4",
        "bafkreid5xj3afuvnjcv7ibankdwpggdbtlhgpubhvolwamgwha6o74pvd4",
        "bafkreif4ec4jpaubuhqlxyg5drvj4wyjywtcdrwba2nmy7tyxyqcfvjmou"
    ]
}
```
### 6. Perform result validation
Call Interface: /api/operator/verify_prediction_results     
Input:
```json
{
    "user": "0x01f5BB073e893d334FF9b0e239939982c124AF97",
    "id": 27
}
```
Output:
```json
{
    "ok": true,
    "msg": "ok",
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "contract_address": "0xFB063Ffe686C254573e4d5CC87C87BA10515D016",
        "contract_function": "verify",
        "proof": "0x0009e74a05cd12574135cdade3f3ab2eaa5a6a5b907c16467ccba45aa9a09eac2c80b85bad14cc7bc7401216b05fe80d8f4c74b088da95637b9fe275a4393c2f1676ad12b49f90c08a8c4ec1f86939fffabd9599d2ae873cdb052bc65d1339b62e0784b4d7efacdb39aeb35543130fa601bb91290f6ac5b8ba4ffa6f7850a25128d35710751b2123cd7ca85434b7ee2c7aced52020afef0e0b8714f0b6bc2eec03e489d35fdb20857fe4f6f64db045ac9bd87e3777a49068e3930a9b3fa6f46a2ef3a0a49a1bbc43cfa9ec9c5feb0d721a0bde9719f1e91ab7baa0412608861c2e0e6c65c4b81baffaef638e03b6b37851fe7cc007b7d90228181db1b7609cbc1be60521a044c6b7751439e9b28aa349f3b9a815806308cf410f8b8e37fcf98923304ad45c9bca173d94b83cce115f18b231eefcd74e83c8a48b76b4dc7e5db30d99ee7ae10b06d087028c504d19c2635bfa2564f2beb58b850157cb214339d32d543d711db01aeabedf42b84e7ef6b510a0b20a52f1ba1bd3d45d88f6ce29e81e70be95b081a69ae44d9388beccbd3c083493f808184cced6a68c167b2a09b31d3a0c87c9c30602daa518b7d87b58cbb35a55eb98bcf8c96440c491ea0f97070bb0dc2f9a4cb4fd54f139f2fc56642de6a4302544a510ff118b56c20b46fdce17647ea5535e281219ab29487efd3a829d5b55216e7d92943102021b6aa6873918984d35c29d4a7f41488e970df7aba6bfd53374aae034f7bfe7a18b140576ae14974fd4aa1b2921e21619d179ca70ed4628523e7f39edeabad34ff4498ddba9142f5508d6bf112ce41da93e13bda9e2b5c68f18129ab6ec32f12d8f0d2337d611a3a7359735b4a9dbe3e1cc327367fb908430a055295aff0ae9fcd351f45f2a0adb7dedf45d768f8f62ec98d55b28ffec615b12d8e4086497b480a133b24eb70133c606730ad08fbfb6e79144a134562a7b291b0ec664010a87f667d8b9828f139fc873bd5d6fc4e0edc1349281a04cea7e99b871604f5f4a71ed001a731a6a21f20d49f01d9f6d1597a3c0690801c24e656929f0c92c98e56f7de1e5520e042b0062e6c3d100115a0f85e570acabba35afcfe2d978e314a624dd57b21f04150f9f61fc678517883c9a567c756aea18ff615e60cf1ae1e4d6fe2a107699052602e16d65c68efccb03c61a1cdba7d2c07b94b0268b0cbba9a242f7254bea34c6074149216be0282f52bf6bc3cf410bca67dd8cb0bed759a219654e8798ea69180a28d208cd6898471678847b8439cdab21596c9c434ed3e6d3efbfce4e0d501e1e7e8fe6ec010b9f94a35da82f3961138fdf906abbcc9c15480b5c3a6fe8e7972e61bb99293cf122d6bc27285765d6546728a4be28549da49afec582d3296e9c2cee92f657d0d556ad3f00aba6af6ea5bbef147b0408f650e54639acaaf35b9015a03b21eb2bd6e1507d68e5c814a262f156395c1642c909156c19ad2afa0d5a0086c0fe8e15099d4729082bc39dd40d17aa2f9e7cb1fbfb9f0be573c1b4cc5a19b77cde653ae59b459ee6879e95dfd09d9bf14b1f6a949442f857aa78c4032214f7522d6c6d045ec20b78e2670271464fcdb113a6f7dfbb0552bc268a570eae2b8837b8df6ac3761f95cd43e0a4d64206c3d6d52fcc74209444f55347da0f961b26ff81084c5ad53bc720709d58b9645d513873746037558469879b1ee6ca1a2e4cc443e281d333b8bf1ed668385d393c32a89870fe0dabe2a4216be41ce08f1e2d04cae2889f772251d82a41b112ec2ec01095971b97a9de7e091981f7b8e4055b776d8d5c1fcc9758d52101ffcd8cdea5a65a0e38d4ee6dbc3d1e05c434b6014476d53a105aa8cfb0b5e8cec736307e2d1472787fe70b72338764df42258e2d88bb42575c7ec53f75cc39a039c1a70a4ae556b2b52c2b867452ba2fd4e0712433ff56b5bb23c6fa5ed5eb7c764c85d1c1a099f08049c75b375f797e890d7806f575a6aebefbfbbf068e9ac5d497fd886f28767364af8c759003b653bf23ec1a1b3a6988f4745a3bfe8d0090b43bd26750989b700285e2d3ca9d8718f53a612d40ff2c6329ecb8b8f68b665b93dfa7463208c06ca05c3932053757de2b50d6110e3a93db85ab687d66d785b196a8b0f2e03074f3e27370fcfcc5a8b952f70319db35c6454ae941e807654107a3ca4dbbdfe79eb973c06ff723fb624b6cdbd00d1cd239036f374bc932293131f8d0fd9c6aa38674b5c7a2ad1b1da8bf72f7d6203a1e68d24bdf12eb91790ba4188b287ac0fe47d44db62f7e8ca4961a4aeb432762fc61b7706d6da012ec0236482cac9aab920c4d0e5d776729b429cccb412828025a036e1cde0c76d363c1dc2aacd214834f61b1c582a92680d69687e5d00d1d32755b58687dc195f80fe35b946d2e28bbd7c03485b8b23a74d9de12a6184514fb1743621f9e58127d50e5bf08af07734716fe6fdcf6ccee2adfe05d8df8c52564b58f196966e960af2269598441616e68d2a533e8c499fde54bc09b8153c612573eb53c0f00446bb57b8746c89a46d23ca33bd3b63c4815b659f2a589985b162ad458ebb98a8fc09bd6974f42612e174f4810b868079af036db1653d861c32e27fb0e779ba24ab8e4f20774da984b7fedda459314897f2b5802b7b47ed44800f71134a1f67d22fe6ab4eedaf77bc05ab9bd2f37d7b4673a513c9fe721f2411d462b115d7d4fa8e9cd1757e36ce446b9bcd472c3808519e6c67b98dbba1bc70930f67b37d282051cdf340a6a60f46ff08c036dd56fe53b4f59c4fde052454c25801057f359548b0841967372d65cf64f8f1ab16118b5edfbcf03f6d4ea6ed2243fdaaed50c914f33e65653b182c3e7a4666c9d1b2b07f4b91414176546c3f801d1936af271356a38d878f5515163904756a3fdb51d111554df91df3740d6a70013b152f410314b8ddeea136e49cbfa12e95b6310cb0edb5630fe0c25c7bdeb1aa6ca3a4b1f345794e9a64ccb683da6311eddf081264bbe56cfac309a78b111",
        "result": "0x0000000000000000000000000000000000000000000000000000000000000001"
    }
}
```