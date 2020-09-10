package bmapiusbuart

//import "fmt"

func (ba *BMAPI) encoder() {
	for {
		select {
		case o0Data := <-ba.o0Send:
			reg := uint8(0)
			cmd := cmdNEWVAL & reg
			ba.sendChan <- cmd
			//fmt.Println("Command: NEWVAL sent - Reg:", reg)
			// BMAPI Customize
			mask8 := uint8(255)
			for i := 0; i < regsizeB; i++ {
				value := uint8((o0Data << (8 * i)) & mask8)
				ba.sendChan <- value
				//fmt.Println("value:", value)
			}
		}
	}
}
