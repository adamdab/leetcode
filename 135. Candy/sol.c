int candy(int *ratings, int ratingsSize)
{
    int counter = 1;

    int increments = 0;
    int decrements = 0;
    int max = 0;

    for (int i = 1; i < ratingsSize; i++)
    {
        if (*(ratings + i - 1) < *(ratings + i))
        {
            increments++;
            decrements = 0;
            max = increments + 1;
            counter += max;
        }
        else if (*(ratings + i - 1) == *(ratings + i))
        {
            increments = 0;
            decrements = 0;
            max = 0;
            counter++;
        }
        else
        {
            decrements++;
            increments = 0;
            counter += decrements + (max > decrements ? 0 : 1);
        }
    }

    return counter;
}