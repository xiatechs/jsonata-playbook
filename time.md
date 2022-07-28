# time

[<< back](readme.md)

you can use various functions to get timestamps in jsonata.

[https://docs.jsonata.org/date-time-functions](time docs for 1.8.0)

the one used in this example is $millis() which returns the number of milliseconds since
unix epoch (1 January 1970 UTC).

you can use time in interesting ways - such as controlling the flow of data. In the example 
let's say you are sending data to a system that's only interested in data that has been provided
within the last 5 seconds of now. Any data outside that boundary is ignored.

Using jsonata you can split up an object into different blocks like below using time:

# INPUT
```
{
    "cutoff": 0,
    "input": "player scored a point",
    "itemsInLast5seconds": [
    ],
    "itemsOutside5seconds": [
    ]
}

```

# JSONATA
```
(
 /* get current time */
 $timeNow := $string($millis());
 
 /* get cutoff point - 5 seconds ago */
 $5secondsAgo := $millis()-5000;
 
 /* update cutoff & itemsInLast5minutes [actually items on or in last 5 minutes] */
 $state1 := $ ~> | $ | {
 "cutoff": $10secondsAgo,
 "itemsInLast5seconds": $append([$.itemsInLast5seconds], { "item": {"timestamp": $timeNow, "object": $$.input}})
}|;

/* remove items outside cutoff from the itemsInLast5seconds array and place in itemsOutside5seconds*/
 $state1 := $state1 ~> | $ | {
 "itemsInLast5seconds": [$.itemsInLast5seconds[$number($.item.timestamp) >= $5secondsAgo]],
 "itemsOutside5seconds": $append($.itemsOutside5seconds, [$.itemsInLast5seconds[$number($.item.timestamp) < $5secondsAgo]])
 }|
)
```

# OUTPUT
```
{
    "cutoff": 1659010660131,
    "input": "player scored a point",
    "itemsInLast5seconds": [
        {
            "item": {
                "object": "player scored a point",
                "timestamp": "1659010864776"
            }
        }
    ],
    "itemsOutside5seconds": []
}
```

if you feed this output as input, it will continue to update and place items in their specific category.