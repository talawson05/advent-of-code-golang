# Day 07

## Part 1

Looks like we're drawing in the console.
Does this count as ascii art?

### Part 1 TDD

Given the task is to count the number of splits, and splits occur for every ^ character, could we just count the characters and avoid drawing this out?
Created a simple test using strings.Count then -1.
Submitted the answer for the real input; turns out the answer is incorrect.
Ah, reading the input more closely merges occur where two splits cause outputs onto the same location, and there are some splitters which don't receive a beam to split.

### Part 1 Impl

Given we have a grid with a fixed width, we can create a list with numbers for each column.
As we progress down the grid, we can togle the positions as we encounter a splitter; and we can merge where they meet.
And where a splitter occurs with no input, do nothing

## Part 2

## Take aways

Reading the input clearly, and noticing the small details is vitally important to avoid going down the wrong path. Luckily the simple count trial and error was quick and painless.
