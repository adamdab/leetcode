int differenceOfSums(int n, int m) {
    int q=n/m;
    return (n*(n+1)/2) - q*(q+1)*m;
}