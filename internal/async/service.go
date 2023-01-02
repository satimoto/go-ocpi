package async

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
	r.channels[key] = make(chan AsyncResult)
	return r.channels[key]
}

func (r *AsyncService) Remove(key string) {
	if _, ok := r.channels[key]; ok {
		close(r.channels[key])
		delete(r.channels, key)
	}
}

func (r *AsyncService) Set(key string, result AsyncResult) bool {
	if _, ok := r.channels[key]; ok {
		r.channels[key] <-result
		return true
	}

	return false
}