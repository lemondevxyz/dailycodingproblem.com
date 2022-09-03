Good morning! Here's your coding interview problem for today.

This problem was asked by Palantir.

Given a number represented by a list of digits, find the next greater permutation of a number, in terms of lexicographic ordering. If there is not greater permutation possible, return the permutation with the lowest value/ordering.

For example, the list [1,2,3] should return [1,3,2]. The list [1,3,2] should return [2,1,3]. The list [3,2,1] should return [1,2,3].

Can you perform the operation without allocating extra memory (disregarding the input memory)?

---

It should be like a cursor, target one number and move it across.

```
; (1 2 3 4)
((1 2 3 4) (1 2 4 3) (1 3 2 4) (1 3 4 2) (1 4 2 3) (1 4 3 2)
(2 1 3 4) (2 1 4 3) (2 3 1 4) (2 3 4 1) (2 4 1 3) (2 4 3 1)
(3 1 2 4) (3 1 4 2) (3 2 1 4) (3 2 4 1) (3 4 1 2) (3 4 2 1)
(4 1 2 3) (4 1 3 2) (4 2 1 3) (4 2 3 1) (4 3 1 2) (4 3 2 1))
```