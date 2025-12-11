# Day 08

## Part 1

### Part 1 TDD

First things first, can we parse the input into a list of items which have the 3 positons?  
Then given we have two items, each with an X Y and Z value, can we calculate the distance?  
Then calculate the distance for all values relative to each other.  
Then sort so we can see who are the closest to each other.
We then need to test the adding to circuits, especially where one junction box is already in a circuit. The example input doesn't suggest 2 separate circuits will intersect, but I'm not convinced.
Once we have a number of circuits, we need to be able to sort them, and return X amount.
For X amount of circuits, multiple the number of JunctionBoxes in the circuits

### Part 1 Impl

Using a struct to create a custom JunctionBox type seems the most sensible.
On first reading in the input, we'll get a list of junction boxes with just their position, and later set the circuit id value as they are connected.
I wish the example input was so long, but creating my own simplified test data will likely create more work.
To work out the distance we need some pythagoras.
We need another object to hold all the distance values, naming things is hard.
For the test I'm using a contains to check the output rather than an exact equals because there's too many permutations to create by hand.
During the sorting exercise it looks like I'm getting duplicates:
`{{425 690 689 0} {162 817 812 0} 100427} {{162 817 812 0} {425 690 689 0}`
Note how the first and second element have switched places.
That'll be a limitation in CalculateDistanceForAllJunctionBoxes; updated loops.  
Implementing the circuits is proving to be a little trickier; is it enough to track the id in the JunctionBox object or do I need a separate entity for a given circuit?
Let's create the separate entity, put JBs into a list. I may end up taking that circuitId out of the JunctionBox defition.
Turns out having the circuit IDs start at 1 is tricker than I thought, for a list of structs. So setting to 0, and updating the blank circuitId in a JunctionBox to be -1.  
As I'm pulling it together I find that the list of pairs is not being updated as each JunctionBox is connected, which means going back through that list and updating it.  
And now I'm hitting the print statement for the combining of two networks.
I'm not sure how to approach that given my current set up as the connection function is only aware of the list of circuits and the current pair being connected; how to update the other pairs? Passing the list in, and moving the call to UpdateListOfPairs inside the connect function

#### Gotcha: remove from slice
Following the advice on `https://go.dev/wiki/SliceTricks` which says to do `a = append(a[:i], a[i+1:]...)` however the +1 ends up with an index out of range error if the item you're removing is the last one in the list.
What's that about?!  

## Part 2

instead of stopping at a fixed number, keep going until all pairs are in the same circuit, then do a quick calculation before break.

### Part 2 TDD

What happens if we run out of junction box pairs?
Create a test with the example input.


### Part 2 Impl
The slice operations are biting me, looks like removing is not a good idea as we run into index issues.
I need a way of knowing that there a circuit has all the junction boxes, which also means I need to know how many junction boxes there are.
Count of how many junction boxes there are is easy enough.
However I've ended up looping through each circuit to compare the number of boxes it has, on every iteration.
End result is this is the longest running task so far.

## Take aways

Stucts are useful but instead of using a slice of structs, I should have considered a map so that the circuitId was the key and I could get/set using it.
It also turns out that removing an item from a slice is tricky too.
