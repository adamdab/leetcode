int maxDifference(char* s) {
    int min_even = 100, max_odd = 0, counter[26];
    int index = 0;
    while (s[index] != '\0') {
        counter[s[index] - 'a']++;
        index++;
    }
    for (int i = 0; i < 26; i++) {
        if (counter[i] % 2 == 0) {
            if (counter[i] != 0)
                min_even = min_even > counter[i] ? counter[i] : min_even;
        } else
            max_odd = max_odd > counter[i] ? max_odd : counter[i];
    }

    return max_odd - min_even;
}