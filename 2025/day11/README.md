# Day 11

## Part 1

Toroidal reactor; mmmm, donut...
We're parsing a tree, recursively going over child nodes.

### Part 1 TDD

As always, let's parse the input into something meaningful we can work with.  
Test the node/child parsing
Test the sample input.

## Part 2

Ok so now the starting label changes, and we must only count paths which contain specific nodes.

### Part 2 TDD

First test the new input with new label returns the total path count.
Then add the functionality to check the contains.

### Part 2 Impl

After succeeding with the same input, tried on the main input; program was still running 3 minutes later.  
It's not enough to just check all possibilities, we need an early exit for nodes we've already visited.


## Take Away
Must remember to include ! when checking equality in tests.
Linter suggested a tagged switch instead of an if/else.
When dealing with recursively iterating over some data, consider caching from the start to avoid long runtime.
I miss having nullable default parameters, could have combined parts 1 & 2 recursive function instead of creating a new one. There is probably a Go-idiomatic way to do it, I just haven't reached that part on the course.
