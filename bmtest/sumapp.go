package main

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"math/rand"
	bmapi "git.fisica.unipg.it/bondmachine/bmapiusbuart.git"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	args := os.Args

	numofloop := 10
	if len(args) == 2 {
		rnumofloop, err := strconv.Atoi(args[1])
		numofloop = rnumofloop
		if  err != nil {
			fmt.Println("Error in atoi conversion ", err)
			return 
		}
	}


	if acc, err := bmapi.AcceleratorInit("/dev/ttyUSB1") ; err == nil {

		value, _ := acc.BMi2r(0)
		fmt.Println("Read initial value: ", value)
		fmt.Println("")

		min := 0
		max := 127
		for i := 0; i < numofloop; i++ {
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

		time.Sleep(time.Millisecond * 1000)

		fmt.Println("")
		value, _ = acc.BMi2r(0)
		fmt.Println("Read Final value: ", value)
	}
}
