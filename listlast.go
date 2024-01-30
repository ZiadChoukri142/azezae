package piscine

func ListLast(l *List) interface{} {
	iterator := l.Head
	var lastData interface{}
	for iterator != nil {
		lastData = iterator.Data
		iterator = iterator.Next
	}
	return lastData
}
