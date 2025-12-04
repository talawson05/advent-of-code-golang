# Day 03

## Part 1

Find the biggest number in the range, without sorting.  
Then find the next biggest number after the first big number, unless it was at the end, then just find the next biggest number.
Going over the list twice seems necessary so let's create an inner function to do the basic return of the biggest number, and then have a wrapping function which calls it a second time with the sublist.

### TDD

Inner tests

| input | expected result |
| --- | --- |
|123|3|
|132|3|
|321|3|
|312|3|

Wrapping tests

| input | expected result |
| --- | --- |
| 321 | 32 |
| 312 | 32 |
| 231 | 31 |
| 123 | 23 |
| 213 | 23 |

### Implementation

For the basic inner, casting from rune to int to make sure the numerical comparison is legit was the main drama. After that, a simple loop and if was enough.
Now we need to take the index and only use the subset of character after the index for the second iteration. Unless the index is at the end, in which case we want all but the last number.  
I don't enjoy how I'm doing the same IF twice, does Go have ternaries? Apparently not.  

## Part 2

Instead of just finding the biggest 2 numbers, it's now 12.  
A quick check of the input reveals there are no lines shorter than 12 characters so we don't need to be worried about that.  
First thing to do is to skip my short tests, then update the expected values on the longer tests.  

### Part 2 impl

This got tricky.  
Clearly I need a nested loop to find the biggest number, but then try again when that number wasn't 12 digits long.  
And it was with great sadness that I ended up deleting the function used to solve part 1.  
Added helpers to convert from string to slice of ints and back.
First time using strings.Builder.  

#### Gotcha: rune to int

Got caught out trying to cast a rune to an int. StackOverflow to the rescue: <https://stackoverflow.com/questions/21322173/convert-rune-to-int#comment112055881_21322694>
