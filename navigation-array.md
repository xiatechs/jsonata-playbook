# json array navigation

[<< back](readme.md)

for navigating an array of objects, you need to use .{} logic.

.{
    inside an array - each item is accessible via '$'
}

which enables a large assortment of transformation steps

# INPUT
```
{
"Other": [
  {
    "Over 18 ?": true,
    "Misc": 5,
    "Alternative.Address": {
      "Street": "Brick Lane",
      "City": "London",
      "Postcode": "E1 6RF"
    }
  },
  {
    "Over 18 ?": false,
    "Misc": 6,
    "Alternative.Address": {
      "Street": "Rock Lane",
      "City": "London",
      "Postcode": "E2 7RF"
    }
  }
         ]
}
```

# JSONATA
```
$$.Other.{
	"item": $
}
```

# OUTPUT
```
[
    {
        "item": [
            {
                "Alternative.Address": {
                    "City": "London",
                    "Postcode": "E1 6RF",
                    "Street": "Brick Lane"
                },
                "Misc": 5,
                "Over 18 ?": true
            }
        ]
    },
    {
        "item": [
            {
                "Alternative.Address": {
                    "City": "London",
                    "Postcode": "E2 7RF",
                    "Street": "Rock Lane"
                },
                "Misc": 6,
                "Over 18 ?": false
            }
        ]
    }
]
```