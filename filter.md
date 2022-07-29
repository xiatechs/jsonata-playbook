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

go-jsonata 1.5.4 does not yet have the '#' and '@' operators which enable useful functionalities such as JOIN.

In the meantime, you can use filter to create join functions - this one is a simple 'full' join on one field in
each object:

# INPUT
```
{
    "ids": [1,2,3],
    "new": {
        "additionalfeatures": [{
                "colour": "blue",
                "id": 1
            },
{
                "colour": "red",
                "id": 2
            }
        ],
        "id": 1
    },
    "old": {
        "id": 1,
        "items": [{
                "id": 1,
                "name": "table"
            }, {
                "id": 2,
                "name": "chair"
            }
        ]
    }
}
```

# JSONATA
```
(
	$join := function($obj1, $obj2, $join1, $join2){
      (
      	$block1 := $filter($obj1, function($index1){
             $index1.$join1 = $obj2.$join2
        });
        
      	$block2 := $filter($obj2, function($index2){
             $index2.$join2 = $obj1.$join1
        });
        
        $append($block1, $block2)
      )
    };
    
    $join($$.old, $$.new, "id", "id")
)
```

# OUTPUT
```
[
    {
        "id": 1,
        "items": [
            {
                "id": 1,
                "name": "table"
            },
            {
                "id": 2,
                "name": "chair"
            }
        ]
    },
    {
        "additionalfeatures": [
            {
                "colour": "blue",
                "id": 1
            },
            {
                "colour": "red",
                "id": 2
            }
        ],
        "id": 1
    }
]

```