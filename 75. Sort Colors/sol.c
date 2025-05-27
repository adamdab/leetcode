void sortColors(int* nums, int numsSize) {
    int colorCounter[3];
    for(int i=0; i<numsSize; i++) colorCounter[*(nums+i)]++;
    int i=0;
    for(int color=0; color<3; color++) {
        while(colorCounter[color]--!=0) {
            *(nums+i)=color;
            i++;
        }
    }
}