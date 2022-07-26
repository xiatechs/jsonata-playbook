# reduce

[<< back](readme.md)

Reduce is a higher order function in jsonata that enables
you to apply a function of type function($i, $j) on an array field

[original docs for higher order functions](https://docs.jsonata.org/higher-order-functions)

Reduce allows you to do a lot, and when combined with map & filter, many magical things can be done!

But reduce alone lets you do things like aggregate a field in an array of objects.

# INPUT
```
[
  {
    "id": 10,
    "name": "Poe Dameron",
    "years": 14
  },
  {
    "id": 2,
    "name": "Temmin 'Snap' Wexley",
    "years": 30
  },
  {
    "id": 41,
    "name": "Tallissan Lintra",
    "years": 16
  },
  {
    "id": 99,
    "name": "Ello Asty",
    "years": 22
  }
]
```

# JSONATA
```
/* here we calculate the total number of years */
$reduce($$.years, function($i, $j){$i + $j})
```

# OUTPUT
```
82
```
