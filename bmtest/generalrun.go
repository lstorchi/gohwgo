package main

import (
	"fmt"
	"time"
	bmapi "git.fisica.unipg.it/bondmachine/bmapiusbuart.git"
)

func main() {

	if acc, err := bmapi.AcceleratorInit("/dev/ttyUSB1") ; err == nil {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 100)
			fmt.Println("Sending", i)
		        acc.BMr2o(0, uint8(i))
		        time.Sleep(time.Millisecond * 100)
		        icheck, _ := acc.BMi2r(0)
			fmt.Println("Receiving", uint8(icheck))
		}

	}
}
