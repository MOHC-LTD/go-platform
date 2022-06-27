package pool

import (
	"context"
	"sync"
)

// Pool groups together goroutines into a place where they can be gracefully shutdown
type Pool interface {
	// Go starts a new goroutine in the pool
	Go(func())
	// Shutdown waits for all running goroutines to finish before returning
	Shutdown(context.Context) error
}

// pool is an implementation of Pool
type pool struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

// New creates a new Pool
func New() Pool {
	ctx, cancel := context.WithCancel(context.Background())
	return pool{
		ctx:    ctx,
		cancel: cancel,
		wg:     &sync.WaitGroup{},
	}
}

// Go starts a new goroutine in the pool
func (p pool) Go(fn func()) {
	select {
	case <-p.ctx.Done():
		return
	default:
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()

			fn()
		}()
	}
}

// Shutdown waits for all running goroutines to finish before returning
func (p pool) Shutdown(ctx context.Context) error {
	p.cancel()

	// Channel to alert when wait group has Wait()'d
	wgChan := make(chan struct{})
	go func() {
		p.wg.Wait()
		wgChan <- struct{}{}
	}()

	// Wait until jobs are 0 or context cancels
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-wgChan:
	}

	return nil
}
