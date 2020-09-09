package main

import (
	"fmt"
	"time"
	"math/rand"
	bmapi "git.fisica.unipg.it/bondmachine/bmapiusbuart.git"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	if acc, err := bmapi.AcceleratorInit("/dev/ttyUSB1") ; err == nil {

		min := 0
		max := 127
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 1000)
			inv := rand.Intn(max - min) + min
			nerr := acc.BMr2o(0, uint8(inv))
			if nerr != nil {
				fmt.Println("Error in send value")
			} else  {
				fmt.Println("Send value: ", inv)
			}
			time.Sleep(time.Millisecond * 1000)
			value, _ := acc.BMi2r(0)
			fmt.Println("Recv value: ", value)
			if value == uint8(inv+1) {
				fmt.Println("")
			} else {
				fmt.Println("Some problem accurred")
			}
		}
	}
}
