package piscine

func ListSize(l *List) int {
	iterator := l.Head
	count := 0
	for iterator != nil {
		count++
		iterator = iterator.Next
	}
	return count
}
