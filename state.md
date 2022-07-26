# state & transactional data

Below is an example where we can feed a jsonata block a series of events
that dictate the current state.

you would ideally want to capture the time that items are purchased / cancelled
at the source and then keep track of the time the event is received.

You might receive an event that socks were purchased 3 minutes after an event
that declared that the socks were actually cancelled in some wierd circumstance involving
latency.

the simple example below shows how you can keep track of data at a rudimentary level coming in as events.

it could probably be improved to take into account more information - but i'll leave that to you to play around with.

this example utilises things you've learned previously including:

- composition
- transformation
- functions
- filter
- tertiary operators

# INPUT
```
{
    "Bought": "socks",
    "Cancelled": "hat",
    "state": [
    ]
}

```

# JSONATA
```
(
    /* capture current time of jsonata being ran */
$t := $now();

/* if event includes a purchased item, and state doesn't contain that item - let's add it */
$insertNew := $$.Bought != "" ? $$.state[$.Item = $$.Bought] ? $$.state : $append($$.state, {
  "Item": $$.Bought,
  "Number": 0,
  "PurchaseTimes": []
}): $$.state;

/* if event contains a cancelled item, and state doesn't contain that item, let's add it */
$insertCancelled := $$.Cancelled != "" ? $$.state[$.Item = $$.Cancelled] ? $insertNew : $append($insertNew, {
  "Item": $$.Cancelled,
  "Number": 0,
  "CancelTimes": []
}): $insertNew;

/* filter out empty objects */
$firstPhase := $$ ~> | $ | {
 	"state": $filter($insertCancelled, function($i){$i})
}|;

/* update purchased items */
$updateNew := $firstPhase ~> | $.state | {
    "PurchaseTimes": $.Item = $$.Bought ? $append($.PurchaseTimes, $t),
	"Number": $$.Bought != "" and $.Item = $$.Bought and $.Item != $$.Cancelled ? $.Number + 1
}|;

/* update cancelled items */
$decreaseOld := $$.Cancelled != "" ? $updateNew ~> | $.state | {
	"Number": $.Item = $$.Cancelled and $.Item != $$.Bought ? $.Number - 1,
    "CancelTimes": $.Item = $$.Cancelled ? $append($.CancelTimes, $t)
}|: $updateNew
)
```

# OUTPUT

If you feed this output as input, it will continue to update the state of the data based
on the values you place in 'Bought' and 'Cancelled'.
```
{
    "Bought": "socks",
    "Cancelled": "hat",
    "state": [
        {
            "Item": "socks",
            "Number": 1,
            "PurchaseTimes": "2022-07-26T15:45:44.160Z"
        },
        {
            "CancelTimes": "2022-07-26T15:45:44.160Z",
            "Item": "hat",
            "Number": -1
        }
    ]
}

```