package tedium

import (
	"github.com/gammazero/workerpool"
)

type TenantWorker struct {
	core            *workerpool.WorkerPool
	tenantQueueSize int
	tenantQueues    map[string]*workerpool.WorkerPool
}

func NewTenantWorker(coreQueueSize int, tenantQueueSize int) *TenantWorker {
	return &TenantWorker{
		core:            workerpool.New(coreQueueSize),
		tenantQueueSize: tenantQueueSize,
		tenantQueues:    make(map[string]*workerpool.WorkerPool),
	}
}

func (t *TenantWorker) AddWork(id string, fn func()) {
	tw, ok := t.tenantQueues[id]
	if !ok {
		tw = workerpool.New(t.tenantQueueSize)
		t.tenantQueues[id] = tw
	}

	// create channel to hold the tenant worker up until
	// the core worker finishes processing the function
	hold := make(chan struct{})

	tw.Submit(func() {
		defer close(hold)
		func() {
			// after completion execute channel
			defer func() {
				hold <- struct{}{}
			}()

			// run our workload
			t.core.Submit(fn)
		}()

		// pause until function returns
		<-hold
	})
}
