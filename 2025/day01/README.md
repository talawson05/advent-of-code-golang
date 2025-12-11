# Day 01

It's day 1 of Advent of Code and we have a safe cracking challenge.  
Picture a large metal safe, with a big dial on the front.  
We're going to be turning that dial left, and right, and it's going to spin all the way around. The dial has numbers 0-99. If we start on position 0, and move Right 5 clicks then we'll land on 5; if we then rotate Left for 10 clicks, we're on 95.  

## TDD

First things first, lets write some tests for this.

| Test case | Starting position | Direction of rotation | Number of clicks to rotate | Expected end position |  
| -------- | ------- | ------- | ------- | ------- |
| Simple right turn | 10 | Right | 2 | 12 |  
| Simple left turn | 10 | Left | 2 | 8 |  
| Completing the circle right | 99 | Right | 1 | 0 |  
| Completing the circle left | 0 | Left |1 | 99|  
| Going nowhere right| 42 | Right | 0 | 42 |  
| Going nowhere left | 42 | Left | 0 | 42 |  
| Large right turn| 10 | Right| 120| 30 |  
| Large left turn | 10 | Left| 145| 65 |  

Let's also throw in a test for an unexpected direction of Up.  
Tests located in [main_test.go](./main_test.go)

## Implementing the dial rotation

Now that the tests are in place, lets try implementing the logic.  
I'm sure there is a clever way to simple add/subtract the number of clicks and using modulo to handle returning to 0; but since I'm visualising a circlar dial I'm going with a loop.  

Gone with a switch statement to cover the direction edge case. Not 100% happy it's being evaluated every loop iteration, but even moving it will mean I need to do an IF to tell whether we're counting up or down.

Found one of my expected results was off, the large right turn needed updating.

## Example input

Heading back to the AoC task description they give us an example list of 10 steps; so let's implement some file reading.  
Created the [example_input.txt](./example_input.txt) to hold the values.  
Wrote test 'TestExampleInput' which just compared the file content to a slice, to make sure I was happy with my approach.  
Next added the step-by-step rotation and validation. Stopped after 5 steps as it was clear the approach was working.  

## The twist

It's not enough to rotate the dial, we need to keep a count of how many times we land on 0. That's land, not pass.  
A simple if == 0 counter should suffice.

## Running

Put the full input into new file [input.txt](./input.txt) and implemented the main.  
Only to bump into package path problems.  
Ended up creating the [cmd main](./cmd/main.go) to act as the trigger. Still figuring the packages/classes out

### GOTCHA

The number of 0s was coming out incorrect. Turns out I was occluding the current dial position variable, which lead to overwriting the count. Changing `currentDialPosition := DoRotation` to `currentDialPosition = returnPosition` solved it.  
Onto part 2.

## Part 2

Now we count every time we touch the number 0, passing and stopping.  
My zero check is outside of the rotation, so a refactor is in order.  
The ability to have multiple return values coming out of a function is great;
and the ability to ignore the new value with `_` keeps the existing tests working.  
The addition of optional arguments helped keep the interface too, but at the cost of extra logic inside the function.
Added the same step by step test, plus tests for when a single rotation will loop over the zero multiple times.  
Success!
