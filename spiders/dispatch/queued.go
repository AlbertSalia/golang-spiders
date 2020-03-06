package dispatch

import "golang-spiders/spiders/engine"

type QueuedDispatch struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedDispatch) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueuedDispatch) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedDispatch) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedDispatch) Ru() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeWorker chan engine.Request
			var activeRequest engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]

			}
		}
	}()
}
