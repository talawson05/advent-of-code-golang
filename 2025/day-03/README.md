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
