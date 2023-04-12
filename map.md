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

Here's a more sophisticated example:

# INPUT
```
{
 "top": {
    "array": [
      {
        "jim": "jones"
      }
    ]
   }
}
```

# JSONATA
```
( 
$mapfunc := function($z){$map($z, function($v, $i, $a) {

	$v.$keys().{
    "needsunescape": $type($lookup($v, $[0])) = "array" ? 1 : $type($lookup($v, $[0])) = "object" ? 1 : 0,
    "value": $string($lookup($v, $[0])), 
    "key": $[0],
    "type": $type($lookup($v, $[0]))}
})};

$topmapfunc := function($z){ $mapfunc($z)[0].{
	  "key": $.key,
      "value": $.needsunescape = 1 ? $topmapfunc($unescape($.value)) : $.value,
      "type": $.type
}};

{
	"fully-denormalised-data": $topmapfunc($$)
}
)
```

# OUTPUT
```
{
    "fully-denormalised-data": {
        "key": "top",
        "type": "object",
        "value": {
            "key": "array",
            "type": "array",
            "value": {
                "key": "jim",
                "type": "string",
                "value": "jones"
            }
        }
    }
}
```
