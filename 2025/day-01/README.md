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

Let's also throw in a test for an unexpected direction of Up

## Implementing the dial rotation

Now that the tests are in place, lets try implementing the logic.  
I'm sure there is a clever way to simple add/subtract the number of clicks and using modulo to handle returning to 0; but since I'm visualising a circlar dial I'm going with a loop.  

Gone with a switch statement to cover the direction edge case. Not 100% happy it's being evaluated every loop iteration, but even moving it will mean I need to do an IF to tell whether we're counting up or down.

Found one of my expected results was off, the large right turn needed updating.
