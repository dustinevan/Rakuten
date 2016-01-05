package main

import (
	"fmt"
	"time"
)



func minuteQ(toMinuteQ, toFiveMinQ, toMainQ, quit chan int) {
	fmt.Println("Entering minuteQ")
	q := make([]int, 0, 4)
	for {
		select {
		case ball := <-toMinuteQ:
			if len(q) < 4 {
				fmt.Println("minuteQ received: ", ball)
				q = append(q, ball)
			} else {
				toFiveMinQ <- ball
				for i := len(q)-2; i > -1; i-- {//backwards
					toMainQ <- q[i]
				}
				q = q[0:0]
			}
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fiveMinQ(toFiveMinQ, toHourQ, toMainQ, quit chan int) {
	fmt.Println("Entering five minuteQ")
	q := make([]int, 0, 11)
	for {
		select {
		case ball := <-toFiveMinQ:
			if len(q) < 11 {
				fmt.Println("fiveminuteQ received: ", ball)
				q = append(q, ball)
			} else {
				toHourQ <- ball
				for i := len(q)-2; i > -1; i-- {//backwards
					toMainQ <- q[i]
				}
				q = q[0:0]
			}
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func hourQ(toHourQ, toMainQ, quit chan int) {
	fmt.Println("Entering hourQ")
	q := make([]int, 0, 11)
	for {
		select {
		case ball := <-toHourQ:
			if len(q) < 11 {
				fmt.Println("hourQ received: ", ball)
				q = append(q, ball)
			} else {
				for i := len(q)-2; i > -1; i-- {//backwards
					toMainQ <- q[i]
				}
				q = q[0:0]

			}
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func mainQ(numberOfBalls int, toMainQ, toMinuteQ chan int) int{
	fmt.Println("Entering mainQ")
	q := createBalls(numberOfBalls)
	fmt.Println(q)
	count, front, back := 0, 0, numberOfBalls-1
	start := true

	for len(q) != numberOfBalls || start{
		select {
		case ball := <-toMainQ:
			q[back] = ball
			back = (back+1)%numberOfBalls//the mod will loop this around
			fmt.Println("Main Queue Received: ", ball)
		default:
			toMinuteQ <- q[front]
			front = (front+1)%numberOfBalls
			count++
			fmt.Println(count)
		}
	}
	return count
}

func createBalls(numberOfBalls int) []int{
	fmt.Println("creating balls")
	a := make([]int, numberOfBalls)
	for i := 0; i < numberOfBalls; i++ {
		a[i] = i+1 //starting the count at 1
	}
	return a

}



func main() {
	start := time.Now().UnixNano()
	fmt.Println("start")
	toMainQ := make(chan int, 10)
	toHourQ := make(chan int, 10)
	toFiveMinQ := make(chan int, 10)
	toMinuteQ := make(chan int, 10)
	quit := make(chan int, 10)

	fmt.Println("channels are done")

	go minuteQ(toMinuteQ, toFiveMinQ, toMainQ, quit)
	go fiveMinQ(toFiveMinQ, toHourQ, toMainQ, quit)
	go hourQ(toHourQ, toMainQ, quit)

	minutes := mainQ(27, toMainQ, toMinuteQ)

	time.Sleep(1 * time.Millisecond)
	quit <- 1
	quit <- 1
	quit <- 1
	fmt.Println(minutes)
	stop := time.Now().UnixNano()
	fmt.Println(stop - start)
}

