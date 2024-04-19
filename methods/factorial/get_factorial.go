package factorial

func Get_factorial(n int, ch chan int){
	reult := 1
	for i := 1; i <= n; i++{
        reult = reult * i
    }
	ch <- reult
}