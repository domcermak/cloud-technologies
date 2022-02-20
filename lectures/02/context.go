package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func main() {
	//context1()
	//context2()
	//context3()
	//context4()
}

func context1() {
	ctx := context.Background()
	_ = context.TODO()

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		cancel()
	}()

	<-ctx.Done()

	fmt.Println(ctx.Err())
	// ctx vs chan?
}

func context2() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	fmt.Println(ctx.Deadline())
	// cancel must be always called
	defer cancel()
	_, _ = http.NewRequestWithContext(ctx, "GET", "http://localhost", http.NoBody)
}

func context3() {
	ctx := context.WithValue(context.Background(), "k", "v")
	fmt.Println(ctx.Value("k"))
}

func context4() {
	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		time.Sleep(3 * time.Second)
		return errors.New("error")
	})

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Second):
				fmt.Println("tick")
			}
		}
	})

	fmt.Println("result", eg.Wait())
}
