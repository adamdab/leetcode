int compare(char* word, int s1, int e1, int s2, int e2) {
    if (s1 == e1 && s2 != e2)
        return -1;
    if (s1 != e1 && s2 == e2)
        return 1;
    int minLength = (e2 - s2) < (e1 - s1) ? (e2 - s2) : (e1 - s1);
    for (int i = 0; i < minLength; i++) {
        if (word[s1 + i] > word[s2 + i])
            return 1;
        if (word[s2 + i] > word[s1 + i])
            return -1;
    }
    return (e1 - s1) > (e2 - s2) ? 1 : -1;
}

char* answerString(char* word, int numFriends) {
    int N = 0;
    while (word[N] != '\0') {
        N++;
    };

    if (numFriends == 1)
        return word;

    int maxSubstringSize = N - numFriends;
    int maxStart = 0, maxEnd = 0;

    for (int currentStart = 0; currentStart < N; currentStart++) {
        int currentEnd = currentStart + maxSubstringSize + 1 < N
                             ? currentStart + maxSubstringSize + 1
                             : N;
        if (compare(word, maxStart, maxEnd, currentStart, currentEnd) < 0) {
            maxStart = currentStart;
            maxEnd = currentEnd;
        }
    }

    if (maxEnd < N)
        word[maxEnd] = '\0';
    return word + maxStart;
}