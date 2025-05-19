#include <stdio.h>
#include <stdbool.h>

bool checkIfTriangle(int a, int b, int c) {
    return a + b > c && a + c > b && b + c > a;
} 

bool checkIfEquilateral(int a, int b, int c) {
    return a == b && b == c;
}

bool checkIfIsosceles(int a, int b, int c) {
    // we assuem that it is not equilateral
    return a == b || b == c || a == c;
}

char* triangleType(int* nums, int numsSize) {
    char* cases[4] = { "equilateral", "isosceles", "scalene", "none" };
    int a = nums[0], b = nums[1], c = nums[2];

    if(checkIfEquilateral(a,b,c)) return cases[0];
    if(!checkIfTriangle(a,b,c)) return cases[3];
    if(checkIfIsosceles(a,b,c)) return cases[1];

    return cases[2];
}


