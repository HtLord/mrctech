package mrctech

func Fibonacci(i int) int{
	if i == 0 {
		return 0
	}else if i == 1 || i == 2{
		return 1
	}
	return Fibonacci(i-1)+Fibonacci(i-2)
}