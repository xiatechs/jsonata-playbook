# query composition

[<< back](readme.md)

In jsonata you can write singular commands like this:

```
$$.field.subfield
```

But it's easier to break down your jsonata into blocks of code
(similar to CTE's in SQL) and labelling each section as a variable.

```
( /* the key is to use these brackets outside your jsonata command */

$example1 := $$.field.subfield;

$example2 := $$.anotherfield;

$finalResult := $append($example1, example2);

$finalResult

)
```

This is called query composition, and can be used in various ways.

[original documentation](https://docs.jsonata.org/composition)

It can be put to great use as shown in the practical examples below:

example 1 shows an enrichment of data - it can be done without composition
if it's a non-array input. However if both sets are arrays, composition will be useful.

# INPUT
```
{
   "new": {
     "ID": 12345,
     "AdditionalData": 9
   },
   "old": [
      {
     "ID": 12345,
     "Data": [1,2,3,4,5,6,7,8]
   	  },
      {
     "ID": 234567,
     "Data": [1,2,3,4,5,6,7,8]
   	  }
    ]
}
```

# JSONATA
```
$$.old.{
	"ID": $.ID,
    "Data": $append($.Data, $filter($$.new, function($i){
    	$.ID = $$.new.ID
    }).AdditionalData)
}
```

# OUTPUT
```
[
    {
        "Data": [
            9,
            1,
            2,
            3,
            4,
            5,
            6,
            7,
            8
        ],
        "ID": 12345
    },
    {
        "Data": [
            1,
            2,
            3,
            4,
            5,
            6,
            7,
            8
        ],
        "ID": 234567
    }
]
```

the second example shows how you can do the same thing to two arrays, but via
the use of query composition, it can be broken down into easier to understand parts.
It can obviously be done as one large code block but then it would be harder to follow.

through the use of composition & functions you can perform
SQL like queries on event-based data.

# INPUT
```
{
   "new": [
     {
     "ID": 12345,
     "AdditionalData": 9
     },
     {
     "ID": 234567,
     "AdditionalData": 10
     },
     {
     "ID": 159,
     "AdditionalData": 8
     },
     {
     "ID": 141,
     "AdditionalData": 2
     }
          ],
   "old": [
      {
     "ID": 12345,
     "Data": [1,3,5]
   	  },
      {
     "ID": 234567,
     "Data": [2,4,8]
   	  }
    ]
}
```

# JSONATA
```
(
/* 
	below are a series of jsonata cte-esque transforms
    that will consistently enrich & append data to the 
    old block from data in the new block.
    
    this is a cyclic set of transformations that will 
    constantly enrich & append new data to the old block
*/

/* yes == item exists in old data */
$notExistInOldData := $$.new.{
"ID": $.ID,
"Data": [$.AdditionalData],
"yes": $filter($$.old, function($i){
				$.ID = $i.ID
		})
};

/* list of items NOT in old data */
$toAdd := $filter($notExistInOldData, function($i){
	$not($i."yes")
}).{
	"ID": $.ID,
    "Data": $.Data
};

/* append new data to items in old data where exist item in new data */
$appendedOldData := $$.new.{
    	"ID": $.ID,
    	"NewData": $append($filter($$.old, function($i){
				$.ID = $i.ID
		}).Data, $.AdditionalData)
    };

/* append list of items not in old data to the old list with enriched data */
$finalResult := $append($$.old.{
    	"Data": $distinct($filter($appendedOldData, function($i){
			$.ID = $i.ID
		}).NewData),
        "ID": $.ID
}, $toAdd);

$finalResult

/*
	try replacing the old block in the input data with the new data
    and adding new data to the new block
*/
)
```

the result is that the old data set is both populated with new objects
from the new data set, and any old objects with matching fields were
enriched with additional data from the new data set.

# OUTPUT

```
[
    {
        "Data": [
            1,
            3,
            5,
            9
        ],
        "ID": 12345
    },
    {
        "Data": [
            2,
            4,
            8,
            10
        ],
        "ID": 234567
    },
    {
        "Data": [
            8
        ],
        "ID": 159
    },
    {
        "Data": [
            2
        ],
        "ID": 141
    }
]

```