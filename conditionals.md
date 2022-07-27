# conditionals

[<< back](readme.md)

in jsonata you can implement conditionals via the tertiary operator on single instances:
```
condition ? if true : if false.
```

it's as simple as that!

# INPUT
```
{ "Name": "Tom" }
```

# JSONATA
```
$$.Name = "Tom" ? "yes" : "no"
```

# OUTPUT
```
"yes"
```

you can also apply conditionals on arrays and edit items in an array depending on the condition:

# INPUT
```
{"items": [1,2,3,4,5]}
```

# JSONATA
```
{
"items": $$.items.[
$ = 4 ? $*100 : $
][0]
}
```

# OUTPUT
```
{
    "items": [
        1,
        2,
        3,
        400,
        5
    ]
}
```