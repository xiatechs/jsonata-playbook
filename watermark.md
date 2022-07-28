# timestamps as watermarks

[<< back](readme.md)

you can statelessly allocate data to specific time ranges
within a block using jsonata using a combination of previously taught subject matter
which - if used correctly - can take off cognitive load on systems

# INPUT
```
{
    "block": [
    ],
    "cutoff": false,
    "data": {
        "timestamp": 1659024212181,
        "value": "hello"
    },
    "start": false
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
   $end ~> | $.block |{
   	"datum": $.start <= $$.data.timestamp and $.end > $$.data.timestamp ? $append($.datum, $$.data) : $.datum
   }|
)
```

# OUTPUT
```
{
    "block": [
        {
            "datum": [],
            "end": 9999999999999,
            "fixed": false,
            "start": 1659024390914
        }
    ],
    "cutoff": true,
    "data": {
        "timestamp": 1659024212181,
        "value": "hello"
    },
    "start": true
}
```
