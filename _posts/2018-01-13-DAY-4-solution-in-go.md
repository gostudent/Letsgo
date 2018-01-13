---
published: true
---

###Beautiful SubSequence

###Question:

Nowadays Babul is solving problems on sub-sequence. He is struck with a problem in which he has to find the longest sub-sequence in an array A of size N such that for all (i,j) where i!=j either A[i] divides A[j] or vice versa. If no such sub-sequence exists then print -1. Help him to accomplish this task.

###Constraints :

- 1<=T<=100
- 2<=N<=1000
- 1<=A[i]<=1000

###Sample Input:

The First line contains T no. of test cases.
Each Test case is of two lines.
The First line contains N size of the array.
Next line contains N-Space separated integers, denoting elements of an array.

### Sample Output:

For each T print the size of the longest sub-sequence satisfying the above criteria.

- Sample TestCase
- Input :
- 2
- 5
- 5 3 1 4 7
- 6
- 2 4 6 1 3 11

###Output :

- 2
- 3

###Explanation:

- First Test Case :
- Longest Sub Sequence are {5,1} , {4,1}, {3,1} etc. so size is 2.

- Second Test Case :
- Longest Sub Sequence are {1, 2, 6}, {1, 3, 6} so size is 3.