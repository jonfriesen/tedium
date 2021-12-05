# Tedium - Multi-tenant worker pool pacakge
_This is an experiment. I'm still learning. Never use this._

Tedium is a simple multi-tenant bounded resource worker pool. It uses [gammazero/workerpool](https://github.com/gammazero/workerpool) as the core worker. It allows tenants to be created with bounded concurrency, so for example, a tenant can have the ability to run 3 operations concurrently. If enough workers are open on the core worker pool all 3 will start concurrently otherwise they will run as workers become available. If a tenant worker has more work than workers, work will queue until a worker is available and trickle down to the core worker.

## Usage

```go
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

/**
* Example output:
* done odd 1
* done even 2
* done even 0
* done even 4
* done odd 3
* done odd 7
* done odd 5
* done odd 11
* done even 8
* done odd 9
* done even 12
* done even 14
* done even 10
* done odd 15
* done odd 13
* done even 6
*/

```