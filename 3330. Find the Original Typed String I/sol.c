int possibleStringCount(char* word) {
    int res = 1, counter = 0, i = 1;
    char prev = word[0];
    while (word[i] != '\0') {
        if (prev == word[i])
            counter++;
        else {
            res += counter;
            counter = 0;
            prev = word[i];
        }
        i++;
    }
    res += counter;
    return res;
}