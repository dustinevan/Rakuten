package main
import (
	"fmt"
	"time"
)



//initialize the front and back as 0!!
type CircularMainQueue struct {
	q []int
	front int
	back int
}

//All the Queues combined with a counter
type BallClock struct {
	mainQueue CircularMainQueue
	minuteQ []int
	fiveMinuteQ []int
	hourQ []int
	minuteCount int
}

//Send a ball from the mainQ to the minuteQ
func (clock *BallClock) increment() {
	clock.pushMinuteQ(clock.mainQueue.q[clock.mainQueue.front])
	clock.mainQueue.front = (clock.mainQueue.front+1)%(len(clock.mainQueue.q))//circular
	clock.minuteCount++
}

//The order checking is more involved due to the circular array
func (mainQ *CircularMainQueue) isInOrder() bool {
	if(mainQ.front != mainQ.back){
		return false
	}

	indexOffset := 0
	for i := 0; i < len(mainQ.q)-1; i++ {
		if mainQ.q[i] == 1 {
			indexOffset = i
			break
		}
	}

	for i := 0; i < len(mainQ.q)-1; i++ {
		index1 := (indexOffset + i)%(len(mainQ.q))
		index2 := (index1+1)%(len(mainQ.q))
		if (mainQ.q[index2] - mainQ.q[index1]) != 1{
			return false
		}
	}
	return true
}

/**
These Queues model the operation of the rows in the ball clock
 */
func (clock *BallClock) pushMainQ(input []int) {
	for i := len(input)-1; i > -1; i-- {
		clock.mainQueue.q[clock.mainQueue.back] = input[i];
		clock.mainQueue.back = (clock.mainQueue.back+1)%(len(clock.mainQueue.q))//this keeps it circular
	}
	return
}

func (clock *BallClock) pushMinuteQ(ball int) {
	if(len(clock.minuteQ) < 4){
		clock.minuteQ = append(clock.minuteQ, ball)
	} else {
		clock.pushMainQ(clock.minuteQ)
		clock.pushFiveMinuteQ(ball)
		clock.minuteQ = clock.minuteQ[0:0]
	}
}

func (clock *BallClock) pushFiveMinuteQ(ball int) {
	if(len(clock.fiveMinuteQ) < 11){
		clock.fiveMinuteQ = append(clock.fiveMinuteQ, ball)
	} else {
		clock.pushMainQ(clock.fiveMinuteQ)
		clock.pushHourQ(ball)
		clock.fiveMinuteQ = clock.fiveMinuteQ[0:0]
	}
}

func (clock *BallClock) pushHourQ(ball int) {
	if(len(clock.hourQ) == 11){
		clock.pushMainQ(clock.hourQ)
		clock.pushMainQ([]int{ball})
		clock.hourQ = clock.hourQ[0:0]
	} else {
		clock.hourQ = append(clock.hourQ, ball)
	}
}

/**
Stringer for Ball Clock
 */
func (clock BallClock) String() string {
	s := fmt.Sprintf("{Min:%v, FiveMin:%v, Hour:%v, Main:%v}", clock.minuteQ, clock.fiveMinuteQ, clock.hourQ, clock.mainQueue);
	s += fmt.Sprintf("; MinutesPassed = %v", clock.minuteCount )
	return s
}

/**
This Stringer is needed so the circular array prints a normal slice, scaled to how many
valid elements it contains. The circular array has old element between the back and front
pointers
 */
func (Q CircularMainQueue) String() string {
	a := make([]int ,0)
	//The Queue is full if front == back
	if(Q.front == Q.back){
		return fmt.Sprintf("%v", Q.q)
	}
	front := Q.front
	back := Q.back
	for front != back {
		a = append(a, Q.q[front])
		front = (front+1)%len(Q.q)
	}
	return fmt.Sprintf("%v", a)
}


//Contructors
func makeMainQueue(numberOfBalls int) *CircularMainQueue{
	q := make([]int, numberOfBalls)
	for i := 1; i < numberOfBalls+1; i++ {
		q[i-1] = i
	}
	return &CircularMainQueue{q, 0, 0}
}

func makeBallClock(numberOfBalls int) *BallClock{
	mainQueue := makeMainQueue(numberOfBalls)
	minuteQ := make([]int, 0, 11)
	fiveMinuteQ := make([]int, 0, 11)
	hourQ := make([]int, 0, 11)
	clock := &BallClock{*mainQueue, minuteQ, fiveMinuteQ, hourQ, 0}

	return clock
}

//Higher level runner functions
func incrementTo(numberOfBalls, numberOfMinutes int){
	start := time.Now()
	clock := makeBallClock(numberOfBalls)
	for clock.minuteCount != numberOfMinutes {
		clock.increment()
	}
	fmt.Println(clock)
	fmt.Printf("This took %v\n\n", time.Since(start))

}

func findRepeat(numberOfBalls int){
	start := time.Now()
	clock := makeBallClock(numberOfBalls)

	clock.increment()
	for !clock.mainQueue.isInOrder() {
		clock.increment()
		if(clock.minuteCount%1440 == 0 ){
			//fmt.Println(clock)
		}
	}
	fmt.Printf("Number of Balls: %v  Number of Days: %v ", numberOfBalls, clock.minuteCount/1440);
	fmt.Printf("\nThis took %v\n\n", time.Since(start))
}

func testSuite(){
	//incrementTo(27, 1440)
	//incrementTo(45, 1000000)
	//incrementTo(90, 1000000)
	//incrementTo(127, 1000000)

	findRepeat(30)
	findRepeat(45)
	findRepeat(60)
	findRepeat(90)
	findRepeat(123)

}

func main() {
	testSuite()
}

