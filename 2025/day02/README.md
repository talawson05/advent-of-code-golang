# Day 02

## Part 1

Started with copying my day01 package setup just to get something running.
There will be a way to have the tests executed without the cmd subfolder I just haven't spotted it yet.

Todays challenge is a pattern matching one, so my first thought is to use regex.
However on reading the examples, specifically `95-115 has one invalid ID, 99` meaning 111 is a valid number even though it's all the same number. So it looks like the trick is to convert to a string and compare the two halves.

### TDD

For the Invalid/Valid pattern matching here is what I think should cover it:

|Input value|True/False|
|---|---|
| 99 | true |
| 111 | false |
| 0 | false |
| 1122 | false |
| 1212 | true |
| 1111 | true |

I then need to test the range expansion correctly so that I parse the input properly:

|Input range| Expected Values|
|---|---|
|11-13|11, 12, 13|
|99-101|99, 100, 101|
|1-2| 1, 2|
|5-5| 5 |
| 42 | Invalid |
| 42- | Invalid |
| -42 | Invalid |

### Implementation

Rune scape baby!
Using runes to get the characters in the string so I can compare the first half to the second half.
And reused the `slices.Equal` I picked up yesterday

#### Unexpected error

As I was implementing the ExpandRange function to cover the test cases where a number is missing e.g. `42-` I encountered a problem where splitting on the dash resulted in a slice with a length of two but the second element was empty and was triggering a `panic: runtime error: invalid memory address or nil pointer dereference`. As this was my own test case rather than the AoC challenge I'm moving on and will come back to this later.  
And now I've learned how to skip tests in go.

##### Workaround

So I know why I got the slice of 2 items but one doesn't exist. In a string `42-` which splits on the dash I get the 42 and the righthand side is null. I would have thought it be nil if it didn't truncate the slice. I still don't know how to catch this, but did find a workaround: `strings.LastIndexByte(inputRange, '-')` and then validating the position of the dash relative to the length of the string.

#### GOTCHA

When looping over the expanded range of IDs I want each element to be the ID e.g.

```go
for id := range expandedRange {
    // do something with id
}
```

But what happened was I was given the index of the loop e,g 0, 1, 2 instead of the actual ID value.
Updated to below to ignore the index:

```go
for _, id := range expandedRange {}
```

## Part 2

Ah ha! I knew the simple string pattern was going to come back; suddenly 111 is now invalid.  
I'm going to pause here to watch [the Regex chapter on the course I've been following](https://youtu.be/XCE0psygwj8?si=UrSNAEkTO6cSpy6R).  
_insert image of neo downloading kungfu from The Matrix film_  
aaaand we're back. Turns out Go doesn't support something called backreferencing so regex is out.
Let's update those tests so that:

|Input value|True/False|
|---|---|
| 99 | true |
| 111 | true |
| 101 | false |
|123123123|true|
|1212121212|true|
|1111111|true|

### Part 2 Implementation

#### Attempt 1 unsuccessful

First attempt at continuing to split the number down the middle worked for all the same character like 111 by offsetting the middle so that the first and second half overlapped by a character e.g. 11 == 11.  
But that ran into issues with the 123 repeating 3 times.

#### Attempt 2

Trying to loop over the string to see if we can build up a matching string by taking the index of characters and multiply by the length e.g. for 111 take 1, repeat it 3 times to 111 will match; but keep looping where no match is found e.g. for 123123123 we need to take the first 3 characters.  
Must remember to truncate the built up string so we don't overshoot.

## Takeaway

Got there in the end. Looping solves everything right?  
Must configure the debugger in VSCode so I can step through instead of relying on printing.
