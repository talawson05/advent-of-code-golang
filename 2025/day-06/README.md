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

## Take aways