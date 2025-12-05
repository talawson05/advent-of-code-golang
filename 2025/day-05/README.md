# Day 05

## Part 1

We're expanding ranges again. We can pull that in from a previous day.  
But this time it looks like we're then trying to find if we have a number in that range; with multiple ranges.  
I think a set would be useful.

### Sets in Go

Turns out go does not have a set class, like Java/Kotlin.
[Stack Overflow suggests a map](https://stackoverflow.com/a/34020023) can achieve the same behaviour.
By using the int as a key, any duplicates will overwrite each other. And the two-value return can be used to determine if an element is not present in the map.


### Part 1 TDD

Testing the set workaround, if I have 1-3 and 3-5 I should end up just 1,2,3,4,5 - no duplicates.  
But in the case of a gap, 1-3 5-7 then I should not see number 6.
The input data needs a test to ensure we capture the ranges, as well as the IDs, and handle that space separating the two groups.  
First time using reflect deepEqual, to compare a nested slice of slices


#### Part 1 Impl

When it came to the input parsing, it turns out the blank row is filtered out by the strings.Fields function - nice.
Piecing together the individual functions used for the tests was straight forward. I suspect Part 2 will want to do more with the list of IDs which are in the set, but for now it's enough to count them.
When I tried to run it with the challenge input it was still running after 2 minutes!  
I know I'm going to have to redesign this, but adding some print statements to see where we get up to.  
It's an issue with expanding the range between two long numbers. Or that's the first issue anyway.
Creating a second input file, and removing most of the ranges reveals the rest of the program finishes immediately. That doesn't mean that once the map is full it won't take ages to lookup the IDs but we'll come back to that.

##### Rethinking the range expand

Clearly expanding the range so that I have every number in the range in a list is not going to be an option. I'm going to have to operate on the basis of if X is greater than or equal to startRange and lessThan or equal to endRange.  
So that means treating the ranges as a list of strings.  
Which means my set idea goes out the window. NOOOOOOO!  
I then need to update the ParseInput function to return the ranges as a list of strings rather than the nested list of ints.
Had to rethink the duplicates where a single item is in multiple of the ranges. Breaking the loop was enough for the unit test, let's see how the real input handles it.
Nailed it!

## Part 2

What! If I could expand the ranges, this would already have been done.
I need to work with the left and right integers of the range, so creating a function to extract them.  
First time sorting a slice!
Ok trying to handle the overlap is tricky.
And now to count the numbers in the range. I've totally butchered CountOfFreshIngredients, it should not have the second return type, this is how API interfaces drift over time.

## Take away

This was the first time encoutering execution timeouts. Not a real timeout, just me bored of waiting and killing it.  
The result was throwing away a nice little idea of using sets, but sometimes we have to say goodbye.  
On that note this was how I learned Golang does not have sets, and that maps can be used to replicate the behaviour.  
Comparing nested slices of slices requires more than the basic slices.Equals. Enter the reflect.DeepEqual which I subsequently took out as part of aboritng the set idea.  
We got to import a sort library.
Finally I've fallen into the classic trap of updating inputs/outputs for a function without respecting the functions name, just because I couldn't be bothered to break it apart - slaps wrist.
