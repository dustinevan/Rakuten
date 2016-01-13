package main
import (
	"fmt"
	"time"
)
const qsize int = 123

/*
these constants are used to test whether we can just add all the balls in the respective Qs
and increment the back all at once, or if we have to circle the array and increment one at a time.
*/
const qBackMin int = qsize - 4
const qBack5Min int = qsize - 11
const qBackHour int = qsize - 12
const qSizeMinusOne int = qsize - 1

func main() {

	start := time.Now()

	//MainQ
	var mainQ [qsize]int
	var mainQfront, mainQback int = 0, 0
	for i := 0; i < qsize; i++ {
		mainQ[i] = i+1;
	}

	//MinuteQ
	var minuteQ [5]int
	minuteQback := 0

	//FiveMinQ
	var fiveMinQ [12]int
	fiveMinQback := 0

	//HourQ
	var hourQ [12]int
	hourQback := 0

	done := false
	minute := 0
	for !done {
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
			done = true;
		}
		notInOrder:
	}
	fmt.Printf("\nThis took %v\n\n%v\n", time.Since(start), (minute)/1440)
}
