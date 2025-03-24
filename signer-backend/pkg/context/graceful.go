package context

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"sync"
	"time"
)

type gracefulHandleContextKey struct{}

var GracefulHandleContextKey = &gracefulHandleContextKey{}

type GracefulHandle struct {
	logger            *slog.Logger
	timeout           time.Duration
	done              chan struct{}
	shutdownNotifiers []chan struct{}
	wg                *sync.WaitGroup
}

func NewGracefulHandle(logger *slog.Logger, timeout time.Duration) *GracefulHandle {
	return &GracefulHandle{
		logger:            logger,
		timeout:           timeout,
		done:              make(chan struct{}),
		shutdownNotifiers: make([]chan struct{}, 0),
		wg:                &sync.WaitGroup{},
	}
}

func (h *GracefulHandle) Done(
	cancellableCtxCreator func(ctx context.Context) (context.Context, context.CancelFunc),
	ctx context.Context,
) <-chan struct{} {
	ctx, cancel := cancellableCtxCreator(ctx)

	go func() {
		<-ctx.Done()
		cancel()

		for _, shutdownNotifier := range h.shutdownNotifiers {
			h.logger.Info("notifying graceful shutdown")
			shutdownNotifier <- struct{}{}
			close(shutdownNotifier)
		}

		wgDoneChan := make(chan struct{})

		go func() {
			h.wg.Wait()
			wgDoneChan <- struct{}{}
		}()

		h.logger.Info("waiting for graceful shutdown or all graceful contexts completed")
		select {
		case <-wgDoneChan:
			h.logger.Info("all graceful contexts completed")
		case <-time.After(h.timeout):
			h.logger.Info("graceful shutdown timeout")
		}

		h.done <- struct{}{}
		close(h.done)
	}()

	return h.done
}

func (h *GracefulHandle) WithGraceful(ctx context.Context) (context.Context, context.CancelFunc) {
	notifier := make(chan struct{})
	consumerCtx, consumerCancel := context.WithCancel(ctx)
	notifierCtx, notifierCancel := context.WithCancel(consumerCtx)
	h.shutdownNotifiers = append(h.shutdownNotifiers, notifier)
	h.wg.Add(1)

	// when graceful shutdown signal is received, emit a cancel signal to consumer context
	go func() {
		<-notifier
		notifierCancel()
	}()

	// when consumer completed normally, or consumer context is cancelled due to graceful shutdown
	go func() {
		<-consumerCtx.Done()
		h.wg.Done()
		idx := slices.Index(h.shutdownNotifiers, notifier)
		h.shutdownNotifiers = append(h.shutdownNotifiers[:idx], h.shutdownNotifiers[idx+1:]...)
	}()

	return notifierCtx, consumerCancel
}

type GracefulHandleContext context.Context

func WithGracefulHandle(ctx context.Context, handle *GracefulHandle) GracefulHandleContext {
	return context.WithValue(ctx, GracefulHandleContextKey, handle)
}

func GracefulHandleFromContext(ctx GracefulHandleContext) (*GracefulHandle, error) {
	handle, ok := ctx.Value(GracefulHandleContextKey).(*GracefulHandle)
	if !ok {
		return nil, fmt.Errorf("graceful handle not found in context, please use WithGracefulHandle to create a graceful handle context")
	}
	return handle, nil
}
