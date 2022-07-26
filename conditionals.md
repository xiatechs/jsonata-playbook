# conditionals

[<< back](readme.md)

in jsonata you can implement conditionals via the tertiary operator:
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