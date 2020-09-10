package bmapiusbuart

//import "fmt"

func (ba *BMAPI) decoder() {
	for {
		select {
		case b := <-ba.recvChan:
			switch {
			case (b & cmdMASK) == cmdNEWVAL:
				reg := (b & ^cmdMASK)
				//fmt.Println("Command: NEWVAL - Reg:", reg)
				// BMAPI Customize
				value := uint8(0)
				for i := 0; i < regsizeB; i++ {
					value = (value << 8) + <-ba.recvChan
				}
				switch reg {
				// BMAPI Customize
				case uint8(0):
					ba.o0Mutex.Lock()
					ba.o0 = value
					ba.o0Mutex.Unlock()
				}
			case (b & cmdMASK) == cmdDVALIDH:
			case (b & cmdMASK) == cmdDVALIDL:
			case (b & cmdMASK) == cmdDRECVH:
			case (b & cmdMASK) == cmdDRECVL:
			default:
				continue
				//fmt.Println(b)
			}
		}
	}
}
