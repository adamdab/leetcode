### Task
Find numbers for ginev `n>0` and `limit>0` such that:
- `i+j+k=n`
- `0 ≤ i,j,k ≤ min(n,limit)`

### Exploration
Let's assuem that `0≤i≤min(n,limit)` is given,
then we can express `k` as `k=n-i-j`,
from constraints we have `0≤n-i-j≤min(n,limit)`🡢 `n-i-j≤n` and `n-i-j≤limit`
it it obvious that `n-i-j≤0` so we are left with constraint `n-i-j≤limit`.
From constraint `n-i-j≤limit` 
🡢 `n-i-j-limit≤0` 
🡢 `n-i-limit ≤ j`,
from base constraints we get `max(0, n-i-limit)≤j`.
From base constraints we get `i+j+k=0` and `j<min(n, limit)` 
🡢 `i+j≤n` and `j≤min(n,limit)`
🡢 `j≤n-i` and `j≤min(n,limit)` 
🡢 `j≤min(n-i, limit)`.

### Conclusion
From above we have constraints:
- `0≤i≤min(n,limit)`
- `max(0, n-i-limit)≤j≤min(limit, n-i)`
- `k=n-i-j`

From above we see that for given `i` so taht `0≤i≤min(n,limit)`,
there are `max(min(limit, n-i)-max(o,n-i-limit)+1, 0)` numbers that fulfill constraints.


