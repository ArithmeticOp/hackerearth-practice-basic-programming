package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	var n []int

	for i := 0; i < t; i++ {
		var sn int
		fmt.Scanf("%d\n", &sn)
		n = append(n, sn)
	}

	for _, v := range n {
		solve(v)
	}
}

func solve(sn int) {
	var face int
	var in int
	face, in = facing(sn)
	fmt.Println(face, findSeat(in))
}

func facing(sn int) (nsn, in int) {
	var mod12 = sn % 12
	switch mod12 {
	case 0:
		{
			nsn = sn - 11
			in = 12
			break
		}
	case 11:
		{
			nsn = sn - 9
			in = 11
			break
		}
	case 10:
		{
			nsn = sn - 7
			in = 10
			break
		}
	case 9:
		{
			nsn = sn - 5
			in = 9
			break
		}
	case 8:
		{
			nsn = sn - 3
			in = 8
			break
		}
	case 7:
		{
			nsn = sn - 1
			in = 7
			break
		}
	case 6:
		{
			nsn = sn + 1
			in = 6
			break
		}
	case 5:
		{
			nsn = sn + 3
			in = 5
			break
		}
	case 4:
		{
			nsn = sn + 5
			in = 4
			break
		}
	case 3:
		{
			nsn = sn + 7
			in = 3
			break
		}
	case 2:
		{
			nsn = sn + 9
			in = 2
			break
		}
	case 1:
		{
			nsn = sn + 11
			in = 1
			break
		}
	}
	return
}

func findSeat(sn int) (s string) {
	if sn%12 == 0 || sn%12 == 1 || sn%12 == 7 || sn%12 == 6 {
		s = "WS"
	} else if sn%12 == 5 || sn%12 == 8 || sn%12 == 2 || sn%12 == 11 {
		s = "MS"
	} else if sn%12 == 4 || sn%12 == 9 || sn%12 == 3 || sn%12 == 10 {
		s = "AS"
	}
	return
}
