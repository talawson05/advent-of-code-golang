# Day 12

## Part 1

What a strange input layout.
I need a structure to hold the shape of the present, and then another about the tree regions.
I'll worry about fitting the presents into the region once we have the input parsed, might help visualise it better.
The real input is 1000+ lines long, finish with a bang.
As I'm creating the Present struct, I thought I'd used a map of coordinate keys and then for the value I can used the Present ID. That'll help if I print it out so Present 0 is using all 0s then Present 1 is using 1s. Not the worst idea I've had.
I think we're going to need a couple of area functions for both the Present and Tree Region.
I can use the total present area vs the region area as a quick check i.e. if we know they won't all fit from the start, move on.
Ok looking at the example input, all 3 tree regions have a total area which could take the total area of the presents; but we know that the 3rd region can't fit the shapes. Now we need to actually deal with placing and rotating...
Let's create a rotate and flip function, and then we'll need to store those details somewhere so that shape 0 has this list of permutations. Given all inputs are a 3x3 shape, we can manually translate it.
Ok we can now rotate and flip our presents, let's take a look how others on Reddit are placing in the tree area.
It's a trap!
Turns out we don't need to bother, the real input is set so we can get away with just the area comparison, without worrying about rotating and flipping; Huzzar!

## Part 2

There is no part 2!

## Take Away

It's been fun, but I'm glad it was only 12 days.
I think it's time to get back to that Golang course and move beyond scripting.
