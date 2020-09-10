package bmapiusbuart

import (
	"context"
	"sync"
)

const (
	// BMAPI Customize
	regsizeB = 1
)

type BMAPI struct {
	// BMAPI Customize
	i0       uint8
	i0Mutex  sync.RWMutex
	o0       uint8
	o0Send   chan uint8
	o0Mutex  sync.RWMutex
	recvChan <-chan uint8
	sendChan chan<- uint8
	cancel   context.CancelFunc
}

func BMAPIInit(device string, tr func(context.Context, string) (chan<- uint8, <-chan uint8)) (*BMAPI, error) {
	result := new(BMAPI)

	ctx, cancel := context.WithCancel(context.Background())

	result.cancel = cancel
	result.sendChan, result.recvChan = tr(ctx, device)
	// TODO SendChan

	result.o0Send = make(chan uint8)

	go result.decoder()
	go result.encoder()

	return result, nil
}

func AcceleratorInit(device string) (*BMAPI, error) {
	return BMAPIInit(device, startTransceiver)
}
