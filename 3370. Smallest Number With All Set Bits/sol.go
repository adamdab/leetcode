package main

func smallestNumber(n int) int {
    return nextOfPower2(n)-1
}

func nextOfPower2(n int) int {
    k := 1
    for k <= n {
        k = k<<1
    }
    return k
}
