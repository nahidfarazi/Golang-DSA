package main

import "fmt"

func main() {
	var n, q int
	fmt.Printf("array and query size: ")
	fmt.Scanf("%d %d", &n, &q)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Println(a)
	fmt.Println(q)
	var pre = make([]int, n)
	pre[0] = a[0]
	for i := 1; i < n; i++ {
		pre[i] = a[i] + pre[i-1]
	}
	fmt.Println(pre)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scanf("%d %d", &l, &r)
		l--
		r--
		var sum int
		if l == 0 {
			sum = pre[r]
		} else {
			sum = pre[r] - pre[l-1]
		}
		fmt.Println(sum)

	}
}

//01798335376
