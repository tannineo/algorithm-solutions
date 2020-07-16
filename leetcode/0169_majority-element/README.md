# 0168_majority-element

https://leetcode.com/problems/majority-element/

## Solutions

### Brute Force

Nested loops:

1. select the number to test
2. loop the whole array to test if its count is bigger than `n/2`

Complexity:

- time: `O(n^2)`
- space: `O(1)`

### Hashmap

The intuitive solution: count all the elements using a map and get the answer with biggest count.

Complexity:

- time: `O(n)`
- space: `O(n)`

### Randomization

Nested loops:

1. randonly select the number to test
   - we have `>50%` chance to find the right answer
2. loop the whole array to test if its count is bigger than `n/2`

Complexity:

- time: `O(infinity)` (with bad luck it can run without end)
- space: `O(1)`

### Bit Voting

Using one loop, count the 32 bit `bit 1`, 32 is for the `int` bit size

Complexity:

- time: `O(n)`
- space: `O(1)`

### Boyer-Moore Voting

https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm

Complexity:

- time: `O(n)`
- space: `O(1)`

### Full Sort

Use quicksort, we don't have to maintain the elements relative order, so use unstable sort.

The medium after sort must be the answer, since it occurs more than `n/2` times.

Complexity:

- time: `O(nlogn)` (sorting)
- space: `O(1)`

### nth_element (partial sort)

In `C++` std library we have a method `nth_element`, giving the element that must be `n_th` element after sorting.

See: https://stackoverflow.com/questions/11068429/nth-element-implementations-complexities

We don't have it in `js`, so we implement it by our own.

TODO

Complexity:

- time: `O(n)`
- space: `O(1)`

### Divide and Conquer

Deivde the array into 2 parts.

The majority must be one of the 2 parts' majority. To compare which one is the answer to the 'current whole array' we need to return the majority elements and their counts to compare

- time: `O(nlogn)` we do iterations inside each calls
- space: `O(logn)` divide and conquer call stack

## Refference

Huahua's Blog
https://zxi.mytechroad.com/blog/divide-and-conquer/leetcode-169-majority-element/
