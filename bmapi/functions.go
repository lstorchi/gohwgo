package bmapiusbuart

func (ba *BMAPI) BMi2r(register uint8) (uint8, error) {
	switch register {
	case uint8(0):
		ba.o0Mutex.RLock()
		result := ba.o0
		ba.o0Mutex.RUnlock()
		return result, nil
	}
	return 0, nil
}
func (*BMAPI) BMi2rw(register uint8) (uint8, error) {
	return 0, nil
}
func (ba *BMAPI) BMr2o(register uint8, value uint8) error {
	switch register {
	case uint8(0):
		ba.o0Send <- value
	}
	return nil
}
func (*BMAPI) BMr2ow(register uint8, value uint8) error {
	return nil
}
func (*BMAPI) BMr2owa(register uint8, value uint8) error {
	return nil
}
