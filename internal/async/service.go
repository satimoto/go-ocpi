package async

import "log"

type AsyncResult struct {
	String string
	Bool bool
}

type AsyncService struct {
	channels map[string]chan AsyncResult
}

func NewService() *AsyncService {
	return &AsyncService{
		channels: make(map[string]chan AsyncResult),
	}
}

func (r *AsyncService) Add(key string) <-chan AsyncResult {
	log.Printf("Async: Add key=%v", key)
	r.channels[key] = make(chan AsyncResult)
	return r.channels[key]
}

func (r *AsyncService) Remove(key string) {
	log.Printf("Async: Remove key=%v", key)
	if _, ok := r.channels[key]; ok {
		close(r.channels[key])
		delete(r.channels, key)
	}
}

func (r *AsyncService) Set(key string, result AsyncResult) bool {
	if _, ok := r.channels[key]; ok {
		log.Printf("Async: Set key=%v", key)
		r.channels[key] <-result
		return true
	}

	log.Printf("Async: Invalid key=%v", key)
	return false
}