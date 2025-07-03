# Simulation
Let's observe **Example 1** example from *description.md*
> Input: k = 5
> Output: "b"

| step | word |
| --- | --- |
| 0 | a |
| 1 | ab |
| 2 | abbc |
| 3 | abbcbccd |

From last step we have enough letters generated to give the ansewer 'b'.

# Observation 1
First we can observe that we can instead of operating on chars we can move to ints and use a single array to store all adjustments
| step | array |
| --- | --- |
| 0 | [0]   |
| 1 | [0, 1] |
| 2 | [0, 1, 1, 2] |
| 3 | [0, 1, 1, 2, 1, 2, 2, 3] |

# Observation 2
To calculate next values in array we can use previously calculated values
| step | start new | end new | array |
| --- | --- | --- | --- | 
| 0 | - | - | [0] |
| 1 | 1 | 1| [0,1] |
| 2 | 2 | 3 | [0,1,1,2] |
| 3 | 4 | 7 | [0,1,1,2,1,2,2,3] |
| 4 | 8 | 15 | [...] |

We can observe that for step $s$ we have to update numbers $nums[i]:i\in[2^{s-1}; 2^s-1]$ based on $nums[i]:i\in[0; 2^{s-1}-1]$

# Observation 3
Let's put this sequence and devide it by $log_2$ parts
num:  [ 0, | 1, | 1, 2, | 1, 2, 2, 3, | ... ]
indx: [ 0,|  1, | 2, 3, | 4, 5, 6, 7, | ...]
eg. **k=7**
deps: [ **0**,|  1, | **2**, 3, | 4, 5, **6**, 7, | ...]
eg. **k=8**
deps: [ **0**,|  **1**, | 2, **3**, | 4, 5, 6, **7**, | ...]

We can observe that depends on number of 'hops' between this $log_2$ segments

# Observation 4
Number of 'hops' in $log_2$ segments for index $i$ is equal to sum of bits in $base_2(i)%26$ representation of $i$

# Conclusion
To get $k-th$ letter in word we need to add to *a* offset of sum of bits in $base_2(k-1)$
