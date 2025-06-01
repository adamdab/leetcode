int max(int a, int b) {
    return a > b ? a : b;
}

int min(int a, int b) {
    return a < b ? a : b;
}

long long distributeCandies(int n, int limit) {
    long long res = 0;
    for(int i=0; i<=min(n, limit); i++) {
        res+=max(min(limit, n-i) - max(0,n-i-limit)+1, 0);
    }
    return res;
}