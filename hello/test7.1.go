package main

import "fmt"

func main() {
	a := 1
	fmt.Println(&a)
	b := a
	fmt.Println(&b)
	c := ([3]int{1, 2, 3})
	fmt.Printf("%v", c[:][1:2])
	println()
	fmt.Println(len(extends_len([]int{1, 2, 3}, 2)))

	fmt.Printf("偶数为:%v", Filter([]int{1, 2, 3, 4, 5}, even))
	println()
	fmt.Printf("添加后为%v", Insert_slice([]string{"1", "2", "3", "4"}, []string{"s", "a"}, 2))
	fmt.Printf("删除后为%v", remove_slice([]string{"a", "v", "s", "a"}, 2, 4))

	fmt.Println(split_string("ashdalsdkaklsd", 6))

	fmt.Printf("反转的字符串为:%s", reverse_string("sjahhjjsdjakal     kjdhA"))
	fmt.Printf("不同字符%v", uniq([]string{"a", "b", "a", "a", "a", "c", "d", "e", "f", "g"}))
	println()
	fmt.Printf("排序后的结果为:%v", bubble_sort([]int{5, 2, 4, 6, 2, 8}))
	println()
	fmt.Printf("map转换结果为：%v", mapfunc([]int{1, 4, 5}, func(x int) int { return x * 10 }))
}

func extends_len(slice []int, factor int) (slice1 []int) {
	if factor >= 1 {
		slice_T := make([]int, len(slice)*factor)
		copy(slice_T, slice)
		slice1 = slice_T
		return
	}
	slice1 = slice
	return
}

func Filter(slice []int, fn func(int) bool) []int {
	var p []int
	for _, v := range slice {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func even(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func Insert_slice(toSlice []string, fromSlice []string, index int) []string {
	var result []string
	result = make([]string, len(toSlice)+len(fromSlice))
	at := copy(result, toSlice[:index])
	at += copy(result[at:], fromSlice)
	copy(result[at:], toSlice[index:])
	return result
}

func remove_slice(slice []string, start int, end int) []string {
	// var result []string
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	at += copy(result[at:], slice[end:])
	return result
}

func split_string(s string, index int) (string, string) {
	return s[:index], s[index:]
}

func reverse_string(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func uniq(s []string) []string {
	arru := make([]string, 0)
	tem := "0"
	for i := 0; i < len(s); i++ {
		if tem != s[i] {
			arru = append(arru, s[i])
			tem = s[i]
		}
	}
	return arru
}

func bubble_sort(x []int) []int {
	for i := 1; i < len(x); i++ {
		for j := 0; j < len(x)-i; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	return x
}

func mapfunc(x []int, fn func(int) int) []int {
	arru := make([]int, len(x))
	for i, val := range x {
		arru[i] = fn(val)
	}
	return arru
}
