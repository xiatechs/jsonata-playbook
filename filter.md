# filter

[<< back](readme.md)

Filter is a higher order function in jsonata that enables
you to filter an object based on a set of validation criteria

[original docs for higher order functions](https://docs.jsonata.org/higher-order-functions)

The key functions are map, filter & reduce. Together they form a very powerful suite of data transformation logic.

You can use filter to great effect on it's own however - it enables many transformations
that are normally done via usage of JOIN / WHERE in SQL

Filter can be used either using conditions not related to other objects,
or the opposite as shown below.

Example below is filtering object based on field in an array equalling '5'.

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
/* general usage of filter:

$filter(objectToFilter, function($arrayIndex){
    $arrayIndex.Field = 5 and $arrayIndex.OtherField = 8
})
*/


$filter($$.Other, function($index){
  $index.Misc = 5
})
```

# OUTPUT
```
[
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

```

