int findLucky(int* arr, int arrSize) {
    int lucky = -1;
    int temp[500];
    for (int i=0; i<arrSize; i++) {
        temp[arr[i]-1]++;
    }
    for (int i = 0; i < 500; i++) {
        if (temp[i] == i + 1)
            lucky = i + 1;
    }
    return lucky;
}
