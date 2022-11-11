package utility

// the recursive swap fucntion to swap the head and tail elements of the given array
func Swap(a *[5]int, head int, tail int) {
	if head >= tail {
		return
	}
	var temp = a[tail]
	a[tail] = a[head]
	a[head] = temp

	head++ //increment head
	tail-- //increment tail
	Swap(a, head, tail)
}
