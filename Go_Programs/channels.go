package main

import "fmt"

func sum(s []int, mychannel chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	mychannel <- sum // send sum to channel
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mychannel := make(chan int)

	go sum(s[:len(s)/2], mychannel)
	go sum(s[len(s)/2:], mychannel)

	x := <-mychannel // recieve from mychannel
	y := <-mychannel // recieve from mychannel

	fmt.Println("Sum computed in first goroutine: ", x)
	fmt.Println("Sum computed in second goroutine: ", y)
	fmt.Println("Total sum: ", x+y)
}
