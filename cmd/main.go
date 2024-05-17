package main

import "fmt"

func main() {
	var len int
	fmt.Println("Введите длину массива : ")
	fmt.Scan(&len)
	arr := make([]int, len)
	fmt.Println("Введите данные массива : ")
	for i := 0; i < len; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)
	b := Find(arr)
	fmt.Println(b)
}

func Find(a []int) (b []int) {
	var k, r, l, d int
	d = a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > k {
			k = a[i]
			r = i
		}
		if a[i] < d {
			d = a[i]
		}

	}
	for g := 2; g <= d; g++ {
		if a[r]%g == 0 {
			for j := 0; j < len(a); j++ {
				if a[j]%g != 0 {
					l++
					break
				}
			}
			if l == 0 {
				b = append(b, g)
			}
			l = 0
		}
	}

	return b
}
