# errors

[<< back](readme.md)

for event based data transformations, if we are going to use jsonata
we also want to accommodate for exceptions.

For example, let's say a service has sent us some data asking us
to enrich a suite of products with some additional unique identifiers.

what if the data says that product A has unique identifier 500, but
product B already has that unique identifier? The original requirements
state that if we encounter a duplicate, we need to error out and let
the client know. 

But how can we do this with jsonata? Here's a similar example:

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
$e := $$.new.OptionCode != $$.old.OptionCode ? $error("optioncodes do not match");
 
/* now - let's make sure that the unique ID is not already assigned to a different product */
$e2 := $$.old.Products.{
 	"isCorrect": $.ID != $$.new.Product and $.UniqueIDs[$ = $$.new.UniqueID] ? $error("unique ID already assigned to: " & $.ID)
};
 
/* the final result */
 $final := $$.old.{
   "OptionCode": $.OptionCode,
   "Products": $$.old.Products.{
   		"ID": $.ID,
        "UniqueIDs": $append($.UniqueIDs, $filter($$.new, function($i){
        	$.ID = $i.Product
        }).UniqueID)
   }
 };
 
 $final
)
```

# OUTPUT
```
depending on the input:

eval error: optioncodes do not match

or

eval error: unique ID already assigned to: 707

or something like below:

{
    "OptionCode": 1000,
    "Products": [
        {
            "ID": 505,
            "UniqueIDs": [
                6,
                8,
                4
            ]
        },
        {
            "ID": 707,
            "UniqueIDs": [
                1,
                3
            ]
        }
    ]
}
```

