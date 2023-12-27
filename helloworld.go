package main

func main() {
	for i := 0; i <= 100; i = i + 1 {
		if i%2 == 0 {
			println("%d-偶数", i)
		} else {
			println("%d-奇数", i)
		}
	}
}
