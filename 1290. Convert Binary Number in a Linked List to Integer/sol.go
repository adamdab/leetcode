func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = (res | head.Val)
		head = head.Next
		if head != nil {
			res = res << 1
		}
	}
	return res
}