# Day 10

## Part 1

3 types on input per line, that'll be fun.
There's a note saying default values are off, that'll be important.
Ah, toggling switches on/off... great. There's going to be a lot of iterating isn't there.
Joltage can be safely ignored; yeah right. We'll see that in part 2.
Finally determine the fewest total toggles.
Sum the lot for the final answer.
And with 185 lines in the input, I'm not manually calculating it any time soon.

### Part 1 TDD

Test the input parsing first.
Test the toggling, where different buttons control one or more light, and it's state is maintained.
Create tests for the example input.

## Part 2

"it's time to worry about the joltage requirements" surprise surprise.  
We need to increment a couter, for each light, every time a connected button is pressed.
It was at this point I stopped trying to figure it out, and look up existing solutions.  
Found an example of someone calculating part 2 separately to part 1 so I didn't need to start again. They even labelled the approach as [Gaussian Elimination](https://en.wikipedia.org/wiki/Gaussian_elimination); woosh.
The smaller math functions like isolate, substitue, and eval I took as written; but I wanted to understand the main logic in the joltage calculation better.

## Take Away

First time using reciever arguments in a function; helps keep the calling code consise compared to passing in and reassigning the return value.
The Go tour tells me they call this a 'method'.  
Importing the math library, things got serious.

Day 10 was a hard day; the math wasn't mathing.
Whilst my formal education was some time ago, I'm not sure it covered this kind of area.  
A good thing about doing this after everyone else has finished, is looking how others have solved it.  
A lot of love for something called [z3](https://en.wikipedia.org/wiki/Z3_Theorem_Prover) which looks to me like an equation solving library. Something about [satisfiability modulo theories (SMT)](https://en.wikipedia.org/wiki/Satisfiability_modulo_theories).  
A couple of people using a [bit mask](https://en.wikipedia.org/wiki/Mask_(computing)).  
As an academic exercise this was a fun opportunity to learn something new.
