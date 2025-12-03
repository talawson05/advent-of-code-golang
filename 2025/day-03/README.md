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
