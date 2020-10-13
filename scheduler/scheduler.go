package scheduler

type Scheduler interface {
	Start(done chan struct{})
}
