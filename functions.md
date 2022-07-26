# functions

in jsonata, you can create functions and use recursion to 
do cursor-esque processing of data. jsonata is both declarative and procedural.

Jsonata is a turing complete functional programming language.

There are some basic fibonacci examples of jsonata on the web that
use a simple array & looping through array an apply to index.

But here's a more interesting example from me:

# INPUT
```
{}
```

# JSONATA
```
(
 /* change this figure will change the output array of numbers
 there is an upper limit to numbers in jsonata - somewhere around 4.8149675025e+298 (that's a big number!)
 */
 $m := 10;
 
 /* the basic fibonacci function */
 $fib := function($first, $second){
 	$first + $second
 };
 
 /* here we use composition to calculate the next result, and then a tertiary operator to perform the next step */
 $arr := function($input, $prev, $next, $counter){
    (
    $result := $fib($prev, $next);
    $counter >= $m ? $input : $arr($append($input, $result), $result, $prev, $counter+1)
    )
 };
 
 /* the arguments for the function are: input array, first number, second number and the counter (up to value of $m) */
 $arr([], 0, 1, 0)
)
```
