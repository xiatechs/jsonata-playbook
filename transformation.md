# transformation

[<< back](readme.md)

there's a specific bit of syntax in jsonata that lets you
edit / transforms sections of an object. Using composition
in tandem, you can do things like edit an object prior to transforming it.

the syntax generally is:

```
Root ~> objectToTransform | {transformation logic} |
```

for example, let's say we had slightly different requirements from
the requirements in the 'errors' example.

let's say a service has sent us some data asking us
to enrich a suite of products with some additional unique identifiers.

what if the data says that product A has unique identifier 500, but
product B already has that unique identifier? The new requirements now 
state that we need to actively _remove_ the unique identifier from the offending
product, and then append that identifier to the new product.

for example, let's say it's a barcode. Barcodes are frequently assigned new products
and removed from old products - we don't want a duplicate barcode on multiple products.

# INPUT
```
{
	"new": {
		"OptionCode": 1000,
		"Product": 505,
		"UniqueID": 3
	},
	"old": {
		"OptionCode": 1000,
		"Products": [{
				"ID": 505,
				"UniqueIDs": [6, 8]
			},
			{
				"ID": 707,
				"UniqueIDs": [1, 3]
			}
		]
	}
}
```

# JSONATA
```
( 
/* first - let's make sure the option codes match */
$e := $$.new.OptionCode = $$.old.OptionCode ? "" : $error("optioncodes do not match");

/* now let's remove the identifier from all products where it exists */
$preTransformed := $$ ~> | $.old.Products | {
	"UniqueIDs": [$.UniqueIDs[$ != $$.new.UniqueID]]
}|;

/* now let's append it to the correct product */
$preTransformed ~> | $.old.Products | {
  "UniqueIDs": $append($.UniqueIDs, $filter($$.new, function($i){
        	$.ID = $i.Product
        }).UniqueID)
}|
)
```

# OUTPUT
```
{
    "new": {
        "OptionCode": 1000,
        "Product": 505,
        "UniqueID": 3
    },
    "old": {
        "OptionCode": 1000,
        "Products": [
            {
                "ID": 505,
                "UniqueIDs": [
                    6,
                    8,
                    3
                ]
            },
            {
                "ID": 707,
                "UniqueIDs": [
                    1
                ]
            }
        ]
    }
}
```