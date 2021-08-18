package main

import "fmt"

func main() {
	a := [5]int{2, 3, 3, 3, 2}

	n := len(a)

	for i, j := 0, n-1; ; i, j = i+1, j+1 {
		n := len(a)

		if n%2 != 0 && i == j {
			break

		}

		if n%2 == 0 && i == n/2 {
			break

		}

		if a[i] != a[j] {
			fmt.Println("not a palindrome")
			break
		}

		if a[i] == a[j] {
			fmt.Println("palindrome")
			break
		}

	}

}
