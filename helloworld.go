package main

/*
*

	func main() {
		t := time.Now().UnixNano()
		rand.Seed(t)
		s := rand.Intn(6)
		fmt.Println(s)

		switch s + 1 {
		case 1:
			fmt.Println("凶")
		case 2, 3:
			fmt.Println("吉")
		case 4, 5:
			fmt.Println("中吉")
		case 6:
			fmt.Println("大吉")
		}
	}

	func main() {
		var f float64 = 10
		var n int = int(f)
		println(n)
	}

	func main() {
		var sum int
		sum = 5 + 6 + 3
		avg := float64(sum / 3)
		println(avg)
	}

	func main() {
		var a, b, c bool
		a = true
		b = true
		c = false
		if a && b || !c {
			println("true")
		} else {
			println("false")
		}
	}

	func main() {
		n1 := 19
		n2 := 86
		n3 := 1
		n4 := 12

		sum := n1 + n2 + n3 + n4
		println(sum)
	}

*/

func main() {
	ns := []int{19, 86, 1, 12}

	for i := 0; i <= 4; i++ {
		println(ns[i])
	}
}
