#include <stdio.h>
#include <stdbool.h>

void setZeroes(int** matrix, int matrixSize, int* matrixColSize) {

    int colSize = matrixColSize[0];
    bool zeroHeader = false;

    for (int i = 0; i < matrixSize; i++) {
        for (int j = 0; j < colSize; j++) {
            int val = *(*(matrix + i) + j);
            if (val == 0) {
                if (i == 0) {
                    zeroHeader = true;
                } else {
                    // 1st col
                    *(*(matrix + i) + 0) = 0;
                    // 1st row
                    *(*(matrix + 0) + j) = 0;
                }
            }
        }
    }

    for (int i = matrixSize - 1; i > 0; i--) {
        for (int j = colSize - 1; j >= 0; j--) {
            int c = *(*(matrix + i) + 0);
            int r = *(*(matrix + 0) + j);
            if (c == 0 || r == 0) {
                *(*(matrix + i) + j) = 0;
            }
        }
    }

    if (zeroHeader) {
        for (int j = colSize - 1; j >= 0; j--) {
            *(*(matrix + 0) + j) = 0;
        }
    }
}