package main
import (
	"fmt"
)



//initialize the front and back as 0!!
type MainQueue struct {
	q []int
	front int
	back int
	minuteCount int
}

func (mainQ *MainQueue) toMain(input []int) {

	for i := len(input)-1; i > -1; i-- {
		mainQ.q[mainQ.back] = input[i];
		mainQ.back = (mainQ.back+1)%(len(mainQ.q))//this keeps it circular
	}
	return
}

func (mainQ *MainQueue) increment(toMinute func(int)) bool{

	if(mainQ.isInOrder() && mainQ.minuteCount > 0){
		return true
	}
	toMinute(mainQ.q[mainQ.front])
	mainQ.front = (mainQ.front+1)%(len(mainQ.q))//circular
	mainQ.minuteCount++
	return false
}

func (mainQ *MainQueue) isInOrder() bool {
	if(mainQ.front != mainQ.back || mainQ.minuteCount == 0){
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


type BallClock struct {
	mainQueue MainQueue
	minuteQ []int
	fiveMinuteQ []int
	hourQ []int
	toMinute func(int) //this is kind of weird, but I need to pass it to the incrementor...there must be a better way
}

func (clock *BallClock) getToMinuteQ(toFiveMinute func(int), toMain func([]int)) func(int){

	return func(ball int){
		if(len(clock.minuteQ) < 4){
			clock.minuteQ = append(clock.minuteQ, ball)
		} else {
			toMain(clock.minuteQ)
			toFiveMinute(ball)
			clock.minuteQ = clock.minuteQ[0:0]
		}
	}
}

func (clock *BallClock) getToFiveMinuteQ(toHour func(int), toMain func([]int)) func(int){
	//q := make([]int, 0, 11)

	return func(ball int){
		if(len(clock.fiveMinuteQ) < 11){
			clock.fiveMinuteQ = append(clock.fiveMinuteQ, ball)
		} else {
			toMain(clock.fiveMinuteQ)
			toHour(ball)
			clock.fiveMinuteQ = clock.fiveMinuteQ[0:0]
		}
	}
}

func (clock *BallClock) getToHourQ(toMain func([]int)) func(int){

	return func(ball int){
		if(len(clock.hourQ) == 11){
			toMain([]int{ball})
			toMain(clock.hourQ)
			clock.hourQ = clock.hourQ[0:0]
		} else {
			clock.hourQ = append(clock.hourQ, ball)
		}
	}
}

func (clock BallClock) String() string {
	s := fmt.Sprintf("{Min:%v, FiveMin:%v, Hour:%v, Main:%v}", clock.minuteQ, clock.fiveMinuteQ, clock.hourQ, clock.mainQueue.q);
	s += fmt.Sprintf("; MinutesPassed = %v", clock.mainQueue.minuteCount )
	return s
}


//Contruction Methods
func makeMainQueue(numberOfBalls int) *MainQueue{
	q := make([]int, numberOfBalls)
	for i := 1; i < numberOfBalls+1; i++ {
		q[i-1] = i
	}
	return &MainQueue{q, 0, 0, 0}
}

func makeBallClock(numberOfBalls int) *BallClock{
	mainQueue := makeMainQueue(numberOfBalls)
	minuteQ := make([]int, 0, 11)
	fiveMinuteQ := make([]int, 0, 11)
	hourQ := make([]int, 0, 11)
	clock := &BallClock{*mainQueue, minuteQ, fiveMinuteQ, hourQ, func(int){}}

	//wire up the functions
	toMain := clock.mainQueue.toMain
	toHour := clock.getToHourQ(toMain)
	toFiveMinute := clock.getToFiveMinuteQ(toHour, toMain)
	clock.toMinute = clock.getToMinuteQ(toFiveMinute, toMain)

	return clock
}

func incrementTo(numberOfBalls, numberOfMinutes int){
	clock := makeBallClock(numberOfBalls)
	for clock.mainQueue.minuteCount != numberOfMinutes {
		clock.mainQueue.increment(clock.toMinute)
	}
	fmt.Println(clock)
}

func findRepeat(numberOfBalls int){
	clock := makeBallClock(numberOfBalls)
	for !clock.mainQueue.isInOrder() {
		clock.mainQueue.increment(clock.toMinute)
	}
	fmt.Printf("Number of Balls:%v  Number of Days:%v", numberOfBalls, clock.mainQueue.minuteCount/1440);
}

func main() {
	fmt.Println("Start")
	numberOfBalls := 127
	minuteMode := true
	numberOfMinutes := 2348723

	if(minuteMode){
		incrementTo(numberOfBalls, numberOfMinutes)
	} else {
		findRepeat(numberOfBalls)
	}
}

