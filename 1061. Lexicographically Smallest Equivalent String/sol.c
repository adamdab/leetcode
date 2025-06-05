void swap(char from, char to, char* chars) {
    int fromIndex = from - 'a';
    int toIndex = to - 'a';
    char currentMin = chars[fromIndex];
    char mappedTo = chars[toIndex];
    if (mappedTo < currentMin) {
        for (int i = 0; i < 26; i++) {
            if (chars[i] == currentMin) {
                printf("Changing mapping of %c from %c to %c\n", i + 'a',
                       chars[i], mappedTo);
                chars[i] = mappedTo;
            }
        }
    }
}

char* smallestEquivalentString(char* s1, char* s2, char* baseStr) {
    char chars[26] = {'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'};

    int index = 0;
    while (s1[index] != '\0') {
        swap(s1[index], s2[index], chars);
        swap(s2[index], s1[index], chars);
        index++;
    }

    index = 0;
    while (baseStr[index] != '\0') {
        int encoding = baseStr[index] - 'a';
        baseStr[index] = chars[encoding];
        index++;
    }

    return baseStr;
}