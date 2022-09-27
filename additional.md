# additional functions

in the xiatech fork of jsonata - we've been busy
adding some additional functions to go-jsonata [xiatech fork]

# sjoin

a simple join function that enables joins like below:

# input
```
{
    "ids": [1,2,3],
    "a": {
        "one": [{
                "colour": "blue",
                "id": 1
            },
{
                "colour": "red",
                "id": 2
            }
        ]
    },
    "b": {
        "two": [{
                "id": 1,
                "name": "table"
            }, {
                "id": 2,
                "name": "chair"
            }
        ]
    },
    "c": {
        "three": [{
                "age": {
  					"furnished": 17,
  					"id": 1
				}
            }, {
                "age": {
  					"furnished": 28,
  					"id": 2
				}
            }
        ]
    },
    "d": {
        "four": {
                    "id": 1,
					"eans": [12,24,36,48]
				}
    }
}
```

# jsonata
```
(
$i1 := $sjoin($$.a.one, $$.b.two, "id", "id");

$i2 := $sjoin($i1, $$.c.three, "id", "age¬id");

$i3 := $sjoin($i2, [$$.d], "id", "four¬id")
)
```

# output
```
[
    {
        "age": {
            "furnished": 17,
            "id": 1
        },
        "colour": "blue",
        "four": {
            "eans": [
                12,
                24,
                36,
                48
            ],
            "id": 1
        },
        "id": 1,
        "name": "table"
    },
    {
        "age": {
            "furnished": 28,
            "id": 2
        },
        "colour": "red",
        "id": 2,
        "name": "chair"
    }
]
```
