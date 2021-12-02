## First Pass - Part 1

Did the problem in Go, since that's what I'm most familiar with. It was fairly straightforward - iterate through the inputs and count the number of times that the input increased from the previous one. I was one over the correct answer initially, having 1293 instead of the correct 1292. This is because `previous` is 0 and triggers our condition to increment `count` on the first pass, even though there is no previous instance. Solved this with an else conditional and a continue to not evaluate our condition on the first element.

SOLUTION: 1292

## First Pass - Part 2
Mostly the same as the first part, but with the added complexity of checking array boundaries so that I didn't access any out of bounds indices.

SOLUTION: 1262

### TODO
Implement this in Rust.