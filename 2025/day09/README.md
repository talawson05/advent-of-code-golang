# Day 09

## Part 1

Another day, another grid.

### Part 1 TDD

Test we can generate the list of tiles from the input.
Then for a given two tiles, calculate the area.
Test the biggest area for the example input and check we get the expected result.

### Part 1 Impl

Playing with structs again. We'll go with slices instead of maps for now, but I want to make sure I've learnt my lesson from yesterday.  
Nice use of Sscanf to extract the values from the input string.
Added a condition to step over blank rows in entry, to make the test readable.
Calculating the area is straight forward enough.  
Once we have a list of the rectangles, sorting by the area value is simple too.

## Part 2

Things took a turn for the worse.  
First figure out the green rectangles.
Then we're going to iterate over both the original (red) rectangles and green rectangles whilst looking for a red rectangle that is inside a green rectangle without it overlapping the edges.
Of those which are true, sort by largest area.  
Looking up how others attempted this, I found a similar approach but they also had an if which tracked the largest area so far and skipped anything smaller than it. Seemed to speed things up.


## Take Aways

Learnt a nice little trick with Sscanf to pull data out of a string.  
Used a loop anotation, to differ between an inner and outer.
This was the trickiest day so far.
