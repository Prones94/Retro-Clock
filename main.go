package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

/* Step #1 Create Digits */
func main(){
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	digits := [...]placeholder {
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	/* Step #2 Print Clock */
	/* Step #3 Animate Clock */
	screen.Clear()
	for { // this creates an infinite loop to always recreate the clock
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		
		fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", hour, min, sec)
		clock := [...]placeholder { // [8][5] --> [8] placeholder
			digits[hour/10], digits[hour%10],
			colon, //separator array
			digits[min/10], digits[min%10],
			colon, //separator array
			digits[sec/10], digits[sec%10],
		}
		for line := range clock[0] {
			for index, digit := range clock {
				next := clock [index][line]
				if digit == colon && sec%2 == 0{
					next = "  "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()

		time.Sleep(time.Second) // function that takes in time.Duration value and stops world for given duration
	}
}