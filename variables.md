# variables

[<< back](readme.md)

in jsonata - any command or query can be made into a variable.
below is a simple example showing a variable 'result' created
from appending two 1D array fields together.

# INPUT
```
{
  "block1": [1,2,3,4,5],
  "block2": [6,7,8,9,10]
}
```

# JSONATA
```
(
$result := $append($$.block1, $$.block2);

{
 "result": $result
}
)
```

# OUTPUT
```
{
    "result": [
        1,
        2,
        3,
        4,
        5,
        6,
        7,
        8,
        9,
        10
    ]
}
```