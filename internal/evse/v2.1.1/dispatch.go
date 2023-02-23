package evse

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

type Job struct {
	EvseResolver *EvseResolver
	Credential   db.Credential
	Location     db.Location
	Uid          string
	Dto          dto.EvseDto
}

func (j Job) Process() {
	ctx := context.Background()
	evse := j.EvseResolver.ReplaceEvse(ctx, j.Credential, j.Location, j.Uid, &j.Dto)

	if evse != nil {
		updateLocationLastUpdatedParams := param.NewUpdateLocationLastUpdatedParams(j.Location)

		if j.Dto.Capabilities != nil || j.Dto.Status != nil {
			if locationAvailabilityParams, err := j.EvseResolver.updateLocationAvailability(ctx, j.Location); err == nil {
				updateLocationLastUpdatedParams = locationAvailabilityParams
			}
		}

		updateLocationLastUpdatedParams.LastUpdated = evse.LastUpdated

		err := j.EvseResolver.Repository.UpdateLocationLastUpdated(ctx, updateLocationLastUpdatedParams)

		if err != nil {
			metrics.RecordError("OCPI112", "Error updating evse", err)
			log.Printf("OCPI112: Params=%#v", updateLocationLastUpdatedParams)
		}
	}
}

type Worker struct {
	id         int
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(id int, workerPool chan chan Job) Worker {
	return Worker{
		id:         id,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

func (w Worker) Start() {
	go func() {
		for {
			// Register the current worker into the worker queue
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// We have received a work request
				job.Process()
			case <-w.quit:
				// We have received a signal to stop
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	JobQueue   chan Job
	WorkerPool chan chan Job
	maxWorkers int
	workers    []*Worker
	queueSize  int
	tick       int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	queue := make(chan Job)
	pool := make(chan chan Job, maxWorkers)

	return &Dispatcher{
		JobQueue:   queue,
		WorkerPool: pool,
		maxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Start() {
	log.Printf("Starting Evse Dispatch service")

	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()

		d.workers = append(d.workers, &worker)
	}

	go d.dispatch()
}

func (d *Dispatcher) QueueJob(job Job) {
	d.queueSize++
	d.JobQueue <- job
}

func (d *Dispatcher) Stop() {
	log.Printf("Shutting down Evse Dispatch service")

	for _, worker := range d.workers {
		worker.Stop()
	}
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			// A job request has been received
			go func(job Job) {
				// Try to obtain a worker job channel that is available.
				// This will block until a worker is idle
				jobChannel := <-d.WorkerPool

				if d.tick > 1000 {
					log.Printf("Evse Dispatch: %d", d.queueSize)
					d.tick = 0
				}

				d.tick++
				d.queueSize--

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
