# time

[<< back](readme.md)

you can use various functions to get timestamps in jsonata.

[time docs for 1.8.0](https://docs.jsonata.org/date-time-functions)

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

if you feed the output above as input, it will continue to update and place items in their specific category.

## Source data timestamps

You can also refer to timestamps on the source data to allocate them to specific categories.
Jsonata allows you manage data based on source timestamps - for example if the source data is over X seconds old
you can treat it differently to when it is brand new data.

Maybe you'll receive data that's over X seconds old due to latency - and you don't want that data to cause issues.
Jsonata allows this kind of logic to be implemented.

# INPUT
```
{
    "cutoff": 1659012931109,
    "dataTimeStamp": 1659012918692,
    "input": "player scored a point"
}
```

# JSONATA
```
(
/* below we will store data where the data source timestamp is within 60 seconds 

you may want to get an updated UNIX epoch timestamp for the incoming data to see this working properly
*/

 /* get current time */
 $timeNow := $millis();
 
 /* get cutoff point - 60 seconds ago */
 $cutOffTime := $timeNow-60000;
 
 /* update cutoff & itemsInLast5minutes */
 $state1 := $ ~> | $ | {
 "cutoff": $cutOffTime,
 "itemsWithinCutOff": $append([$.itemsWithinCutOff], { "item": {"ingestedTimestamp": $timeNow, "sourceTimeStamp": $$.dataTimeStamp, "object": $$.input}})
}|;

/* remove items outside cutoff from itemsWithinCutOff */
$state2 := $state1 ~> | $ | {
 "itemsOutsideCutoff": [$append($.itemsOutsideCutoff, $.itemsWithinCutOff[$.item.sourceTimeStamp < $cutOffTime])]
 }|;


/* itemsWithinCutOff = items within the cutoff */
 $state3 := $state2 ~> | $ | {
 "itemsWithinCutOff": [$.itemsWithinCutOff[$.item.sourceTimeStamp >= $cutOffTime]]
 }|;
 

 $state3
)
```

# OUTPUT
```
{
    "cutoff": 1659013088036,
    "dataTimeStamp": 1659012918692,
    "input": "player scored a point",
    "itemsOutsideCutoff": [
        {
            "item": {
                "ingestedTimestamp": 1659013148036,
                "object": "player scored a point",
                "sourceTimeStamp": 1659012918692
            }
        }
    ],
    "itemsWithinCutOff": []
}
```
