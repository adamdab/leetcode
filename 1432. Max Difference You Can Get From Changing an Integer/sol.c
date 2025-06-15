#include <stdbool.h>

int iterator(int num) {
    int iter = 10;
    while (num / iter != 0)
        iter *= 10;
    return iter / 10;
}
int max(int num, int iter) {
    int res = 0, chNum = 10;
    while (iter != 0) {
        int n = num / iter;
        num %= iter;
        if (n != 9 && chNum == 10) {
            chNum = n;
        }
        n = n == chNum ? 9 : n;
        res += n * iter;
        iter /= 10;
    }
    return res;
}

int min(int num, int iter) {
    int res = 0, chNum = 10, ch = 0;
    bool firstPass = true;
    while (iter != 0) {
        int n = num / iter;
        num %= iter;
        if (n > 1 && chNum == 10) {
            chNum = n;
            ch = firstPass ? 1 : 0;
        }
        n = n == chNum ? ch : n;
        res += n * iter;
        iter /= 10;
        firstPass = false;
    }
    return res;
}

int maxDiff(int num) {
    int iter = iterator(num);
    return max(num, iter) - min(num, iter);
}