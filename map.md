# map

[<< back](readme.md)

Map is a higher order function in jsonata that enables
you to apply a function of type function($value, $index, $array){} on an array

[original docs for higher order functions](https://docs.jsonata.org/higher-order-functions)

here's an example:

# INPUT
```
[
  {
   	"Name": "Tim",
    "Likes": "Cats"
  },
  {
   	"Name": "Joe",
    "Likes": "Dogs"
  },
  {
   	"Name": "Simon",
    "Likes": "Dogs",
    "Dislikes": "Cats"
  }
]
```

# JSONATA
```
/* let's convert the incoming irregular schema into a set of data of regular schema using map */
$map($$, function($v, $i, $a) {

	$v.$keys().{
    "name": $string($lookup($v, $[0])),
    "key": $[0]}

})
```

# OUTPUT
```
[
    [
        {
            "key": "Name",
            "name": "Tim"
        },
        {
            "key": "Likes",
            "name": "Cats"
        }
    ],
    [
        {
            "key": "Name",
            "name": "Joe"
        },
        {
            "key": "Likes",
            "name": "Dogs"
        }
    ],
    [
        {
            "key": "Name",
            "name": "Simon"
        },
        {
            "key": "Likes",
            "name": "Dogs"
        },
        {
            "key": "Dislikes",
            "name": "Cats"
        }
    ]
]

```