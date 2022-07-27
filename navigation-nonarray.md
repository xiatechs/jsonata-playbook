# json non-array navigation

[<< back](readme.md)

You can work your way down a json object via the usage of the . operator

# INPUT
```
{
"Other": {
    "Over 18 ?": true,
    "Misc": 5,
    "Alternative.Address": {
      "Street": "Brick Lane",
      "City": "London",
      "Postcode": "E1 6RF"
    }
  }
}
```

# JSONATA
```
/* $$ means the 'root' of the entire JSON object */
$$.Other.Misc
```

# OUTPUT
```
5
```

You can use quotes around fields for when they include spaces i.e for the input above:

# JSONATA
```
$$.Other."Over 18 ?"
```

# OUTPUT
```
true
```