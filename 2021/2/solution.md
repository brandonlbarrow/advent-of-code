# Day 2 Thoughts

This was a really interesting puzzle. The biggest thing I got out of it was learning how to use fmt.Sscan to read space delimited strings into variables, which was very useful when iterating over the inputs to get both the action (string) and delta (int).
Originally I tried to use a map[string]int to hold my change and delta, but it proved a bit cumbersome with multiple loops, and I much prefer the fmt.Sscan solution. I am really enjoying the ocean theme and modeling submarine traversal.

Solution to part 1: 1813801
Solution to part 2: 1960569556
