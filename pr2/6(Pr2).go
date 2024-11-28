package Pr2

func Fib(arr []int, le int) []int {
	if le == 2 {
		return arr
	}
	arr = append(arr, arr[len(arr)-2]+arr[len(arr)-1])
	return Fib(arr, le-1)
}
