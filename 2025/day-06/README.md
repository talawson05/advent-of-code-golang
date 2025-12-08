# Day 06

## Part 1

Looks like another nested loop but instead of counting columns then rows, we're going rows, then columns.

### Part 1 TDD

Parsing the input first seems like a sensible idea.  
Then deal with the addition/multiplication after.  
Attempting to use DeepEqual again, let's hope this time is stays.
Splitting out the operator from the rest of the equation, to use later, should be helpful so will return as a separate return value.

### Part 1 Impl

I seem to be having a problem creating a 2d slice.
Where I'm expecting row 0 to contain a slice of many values, instead I'm getting each value appended to a new row so that everything is in column 0.  
I think the issue is the splitting on the input string using fields rather than new line characters.  
Turns out counting column then row is mentally trickier than originally thought.

## Part 2

It looks like another rip out and start again.
Input is essentially reversed. 

### Part 2 TDD

Need to parse that string from right to left, top to bottom.  
And because the spacing/indents are important, read it from a file.


### Part 2 Impl

Reversing a slice is easy enough with slices.Reverse, which manipulates the slice so no returning.
But during the attempt, I'm going to read the input line by line instead of a whole string.
Which means reading from file to take advantage of the line scanner; taking the reverse slice out.  
Once the input was parsed, turns out it was the regular way so can traverse it like row x column as normal.
This meant creating a second DoCalc function as it's going through it in a different order.

## Take aways

Parsing the input is one thing, but normalising would have been a good idea.
By normalising it I mean converting the column x row layout from part 1 into the standard row x column.
The benefit of doing it this way would mean a single DoCalc function.
