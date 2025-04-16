package main 

import (
	"context"
	"fmt"	
	"os"
	"os/signal"
	"sync"
	"time"
)


func main() {
	fmt.Println("I'm mining bitcoins")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt )
	defer stop()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done() 
		tk := time.NewTicker(time.Second)

		for {
			select {
			case <- tk.C:
				fmt.Print("...")
			case <- ctx.Done():
				stop()
				fmt.Println("")
				return 
			}
		}
	}()

	wg.Wait()
	fmt.Println("bye bye")
}
