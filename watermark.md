# timestamps as watermarks

[<< back](readme.md)

you can statelessly allocate data to specific time ranges
within a block using jsonata using a combination of previously taught subject matter
which - if used correctly - can take off cognitive load on systems.

event-time aggregation is a big deal in the event streaming world, and jsonata
steps up to this challenge too.

if an event appears outside any recognised existent time window, the code below generates a bespoke
time window for a range that satisfies the incoming event & which doesn't collide with existent time windows.

additionally, new events coming in will either fit into an existent time window, or a new window will be 
generated and the item will be placed in that window.

this kind of code can be ran in a stateless lambda, and the state can be captured in a key value store.
another lambda, or containerised microservice, could routine drain the data and provide time-based data 
for a variety of sources of information.

# INPUT
```
{
    "block": [
    ],
    "cutoff": true,
    "data": {
        "timestamp": 16590122584,
        "value": "hello"
    },
    "start": true
}


```

# JSONATA
```
(
/* get current time */
	$timeNow := $millis();
 /* get cut off time */   
    $cutOff := $timeNow - 5000;
  /* begin creating data time blocks */  
    $start := $count($$.block) = 0;
 
    $block := $start ? {"start": $start, "block": [{"datum": [], "fixed": false, "start": $timeNow, "end": 9999999999999}]} : $$;
    
    $sortedItems := $sort($block[0].block, function($l, $r) {
  		$l.start > $r.start
	});
    
    $block := $block ~> | $ |{
    	"block": $sortedItems
    }|;
    
    $fixup := $block.block[-1].start > $cutOff;
    
    /* here we add another block if we need to */ 
    $bb := $block ~> | $ |{
        "data": $$.data,
        "cutoff": $block.block[-1].start > $cutOff,
        "start": $start,
    	"block": $block.block[-1].start > $cutOff ? $block.block : [$append($$.block, {"datum": [], "fixed": false, "start": $timeNow, "end": 9999999999999})]
    }|;
   
   /* here we update the previous block so that the data range is frozen */
   $end := $bb ~> | $.block[-2] |{
     "end": $.fixed ? $.end : $timeNow-1,
     "fixed": true
   }|;
   
   /* here we add the data to the correct time range */
   $end := $end ~> | $.block |{
   	"datum": $.start <= $$.data.timestamp and $.end > $$.data.timestamp ? $append($.datum, $$.data) : $.datum
   }|;
   
   /* if there is _no_ time block for the incoming data i.e it was was generated outside the range */
   $outsideRange := [$$.block.{
    "yes": $.start <= $$.data.timestamp and $.end > $$.data.timestamp
   }[$.yes = true]] = [];
   
   /* create a time block for period between the data timestamp and next data start point */
   $final := $outsideRange ? $end ~> | $ |{
      "block": $append($.block, {"datum": [$$.data], "fixed": true, "start": $$.data.timestamp, "end": $min($.block.start)-1})
	}| : $end;
    
    /* update blocks again to ensure correctly ordered */
    $sortedFinal := $sort($final.block, function($l, $r) {
  		$l.start > $r.start
	});
    
   
   $fixedBlocks := $sortedFinal ~> | $.block[-2] |{
     "end": $.fixed ? $.end : $timeNow-1,
     "fixed": true
   }|;
   
   $final ~> | $ |{
    "block": $fixedBlocks
   }|
    
)
```

# OUTPUT
```
{
    "block": [
        {
            "datum": [
                {
                    "timestamp": 16590122584,
                    "value": "hello"
                }
            ],
            "end": 1659035154054,
            "fixed": true,
            "start": 16590122584
        },
        {
            "datum": [],
            "end": 9999999999999,
            "fixed": false,
            "start": 1659035154055
        }
    ],
    "cutoff": true,
    "data": {
        "timestamp": 16590122584,
        "value": "hello"
    },
    "start": true
}
```

try placing the output back in the input
