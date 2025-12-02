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
