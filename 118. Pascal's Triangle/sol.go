package main

func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    res[0] = []int{1}
    for i:=1; i<numRows; i++ {
        res[i] = append(res[i], 1)
        prev := res[i-1]
        for j:=1; j<len(prev); j++ {
            res[i] = append(res[i], prev[j]+prev[j-1])
        }
        res[i] = append(res[i], 1)
    }
    return res
}