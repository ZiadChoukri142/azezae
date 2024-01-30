package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for x := '9'; x >= '0'; x-- {
		for y := '9'; y >= '0'; y-- {
			for i := '9'; i >= '0'; i-- {
				for j := '9'; j >= '0'; j-- {
					if (x > i) || (x == i && y > j) {
						z01.PrintRune(x)
						z01.PrintRune(y)
						z01.PrintRune(' ')
						z01.PrintRune(i)
						z01.PrintRune(j)

						if x != '0' || y != '1' || i != '0' || j != '0' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
