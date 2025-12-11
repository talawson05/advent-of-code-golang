# Day 04

## Part 1

We're traversing a grid! That's going to be a nested for loop.  
We need to parse the input into a structure, parse the current location in the grid to determine what character it is, and query the surrounding area.

### Part 1 TDD

Starting small, let's test we can parse a given character. It's only a simple check but having a function for it should help readability.  
Next let's work on converting a 9x9 multiline string into a grid. Turns out you use backticks.  
Now a grid has x and y coordinates, plus a value for the current coordiate. I think a map where the x/y are the key should suffice.  
Finally we need a way to check the surrounding coordinates so that for a given x/y we can check x-1 to x+1, and y-1 to y+1. Ignoring the x/y = 0 as thats our current location. And we need to make sure we stay in the grid; if current coord is at the edge we don't try to access something out of range.

#### Part 1 Impl

The task is to update any coordinate which is currently @ and has fewer than 4 neighbours.  
But the neighbours can't be updated, and already updated values also count towards the neighbour count.
During a test I found the updates were causing a problem with the character parsing check, added condition to handle the updated values.

## Part 2

Recursion!  
Keep going over the grid and removing the paper rolls; until no more updates are being made.  

### Part 2 TDD

One thing is for sure is we now need to remove the updated values so they aren't picked up on the next run.

### Part 2 Impl

I realise I'm returning and reassigning maps, whilst the course I've been following says maps are passed by reference so the underlying grid should be being updated anyway but it just doesn't seem clear and obvious.  
Once upon a time I was told there are two types of developers:

- those who prefer concise efficient code over readability
- those who prefer readability and maintainability of execution speed.

I guess unless I'm in the situation where every millisecond counts, I prefer readable code.

## Takeaway

First opportunity to work with the two value assignment from a map. Nice way of knowing the cell I was looking for was not in the grid.  
