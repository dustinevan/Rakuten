package main
import (
	"fmt"
	"time"
	"os"
	"strconv"
)


func main() {

	start := time.Now()
	qsize := 27
	var err error;
	minutesToRun := -1
	args := os.Args[1:]
	switch len(args) {
	case 0:
		fmt.Println("Usage:\n\tfasterballclock.go numberOfBalls [minutes to run]\n\n")
		os.Exit(1)
	case 1:
		if qsize, err = strconv.Atoi(args[0]); err != nil {
			fmt.Println("Usage:\n\tfasterballclock.go numberOfBalls [minutes to run]\n\n")
			os.Exit(1)
		}
	case 2:
		if qsize, err = strconv.Atoi(args[0]); err != nil {
			fmt.Println("Usage:\n\tfasterballclock.go numberOfBalls [minutes to run]\n\n")
			os.Exit(1)
		}
		if minutesToRun, err = strconv.Atoi(args[1]); err != nil {
			fmt.Println("Usage:\n\tfasterballclock.go numberOfBalls [minutes to run]\n\n")
			os.Exit(1)
		}
	}

	qBackMin := qsize - 4
	qBack5Min := qsize - 11
	qBackHour := qsize - 12
	qSizeMinusOne := qsize - 1

	//MainQ
	mainQ := make([]int, qsize, qsize)
	mainQfront, mainQback := 0, 0
	for i := 0; i < qsize; i++ {
		mainQ[i] = i+1;
	}

	//MinuteQ
	minuteQ := make([]int, 5, 5)
	minuteQback := 0

	//FiveMinQ
	fiveMinQ := make([]int, 12, 12)
	fiveMinQback := 0

	//HourQ
	hourQ := make([]int, 12, 12)
	hourQback := 0

	done := false
	minute := 0
	for !done && (minutesToRun != minute) {
		minute++

		//move ball out of mainQ
		minuteQ[minuteQback] = mainQ[mainQfront]
		if mainQfront == qSizeMinusOne {
			mainQfront = 0
		} else {
			mainQfront++
		}

		//Did we just fill the minuteQ? if yes do more stuff
		if minuteQback == 4 {
			fiveMinQ[fiveMinQback] = minuteQ[4]
			if mainQback < qBackMin {
				//mainQ[mainQback:mainQback+3] = []int{minuteQ[3], minuteQ[2], minuteQ[1], minuteQ[0]}
				mainQ[mainQback] = minuteQ[3]
				mainQ[mainQback+1] = minuteQ[2]
				mainQ[mainQback+2] = minuteQ[1]
				mainQ[mainQback+3] = minuteQ[0]
				mainQback+=4
			} else {
				mainQ[mainQback] = minuteQ[3]
				if mainQback == qSizeMinusOne {
					mainQback = 0
				} else {
					mainQback++
				}
				mainQ[mainQback] = minuteQ[2]
				if mainQback == qSizeMinusOne {
					mainQback = 0
				} else {
					mainQback++
				}
				mainQ[mainQback] = minuteQ[1]
				if mainQback == qSizeMinusOne {
					mainQback = 0
				} else {
					mainQback++
				}
				mainQ[mainQback] = minuteQ[0]
				if mainQback == qSizeMinusOne {
					mainQback = 0
				} else {
					mainQback++
				}
			}
			if fiveMinQback == 11 {
				hourQ[hourQback] = fiveMinQ[11]
				if mainQback < qBack5Min {
					mainQ[mainQback] = fiveMinQ[10]
					mainQ[mainQback+1] = fiveMinQ[9]
					mainQ[mainQback+2] = fiveMinQ[8]
					mainQ[mainQback+3] = fiveMinQ[7]
					mainQ[mainQback+4] = fiveMinQ[6]
					mainQ[mainQback+5] = fiveMinQ[5]
					mainQ[mainQback+6] = fiveMinQ[4]
					mainQ[mainQback+7] = fiveMinQ[3]
					mainQ[mainQback+8] = fiveMinQ[2]
					mainQ[mainQback+9] = fiveMinQ[1]
					mainQ[mainQback+10] = fiveMinQ[0]
					mainQback+=11
				} else {
					mainQ[mainQback] = fiveMinQ[10]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[9]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[8]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[7]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[6]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[5]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[4]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[3]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[2]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[1]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
					mainQ[mainQback] = fiveMinQ[0]
					if mainQback == qSizeMinusOne {
						mainQback = 0
					} else {
						mainQback++
					}
				}
				if hourQback == 11 {
					if mainQback < qBackHour {
						mainQ[mainQback] = hourQ[10]
						mainQ[mainQback + 1] = hourQ[9]
						mainQ[mainQback + 2] = hourQ[8]
						mainQ[mainQback + 3] = hourQ[7]
						mainQ[mainQback + 4] = hourQ[6]
						mainQ[mainQback + 5] = hourQ[5]
						mainQ[mainQback + 6] = hourQ[4]
						mainQ[mainQback + 7] = hourQ[3]
						mainQ[mainQback + 8] = hourQ[2]
						mainQ[mainQback + 9] = hourQ[1]
						mainQ[mainQback + 10] = hourQ[0]
						mainQ[mainQback+ 11] = hourQ[11]
						mainQback += 12
					} else {
						mainQ[mainQback] = hourQ[10]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[9]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[8]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[7]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[6]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[5]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[4]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[3]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[2]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[1]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[0]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
						mainQ[mainQback] = hourQ[11]
						if mainQback == qSizeMinusOne {
							mainQback = 0
						} else {
							mainQback++
						}
					}
					hourQback = 0
				} else {
					hourQback++
				}
				fiveMinQback = 0
			} else {
				fiveMinQback++
			}
			minuteQback = 0
		} else {
			minuteQback++
		}

		if mainQfront == mainQback {
			for i := 0; i < qsize-1; i++ {
				if( mainQ[i] > mainQ[i+1] ){
					if(i == mainQback){
						continue
					} else {
						goto notInOrder
					}
				}
			}
			minute++
			done = true;
		}
		notInOrder:
	}
	fmt.Printf("\nThis took %v\n\n%v\n", time.Since(start), (minute)/1440)
	fmt.Printf("{Min:%v, FiveMin:%v, Hour:%v, Main:%v}", minuteQ, fiveMinQ, hourQ, mainQ)
}
