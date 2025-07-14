int getDecimalValue(struct ListNode* head) {
    int res = 0;
    while (head) {
        res = (res | head->val);
        head = head->next;
        if (head)
            res = res << 1;
    }
    return res;
}