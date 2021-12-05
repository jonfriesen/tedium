//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"sync"

	"github.com/jonfriesen/tedium"
)

const demoSize int = 16

func main() {
	w := tedium.NewTenantWorker(1024, 1024)

	wg := sync.WaitGroup{}

	wg.Add(demoSize)
	for i := 0; i < demoSize; i++ {
		group := "odd"
		if i%2 == 0 {
			group = "even"
		}

		i := i
		w.AddWork(group, func() {
			defer wg.Done()
			fmt.Println("done", group, i)
		})
	}

	wg.Wait()
}
