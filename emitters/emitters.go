package emitters

type callback func(args ...interface{})

type eventEmitter struct {
	messages map[string][]callback
}

// Emitter is an interface containing the two methods for interacting with events
type Emitter interface {
	Emit(eventName string, args ...interface{})
	On(eventName string, cb callback)
}

// New returns an Emitter with an internal map of registered events
func New() Emitter {
	return &eventEmitter{make(map[string][]callback)}
}

func (r *eventEmitter) Emit(eventName string, args ...interface{}) {
	if len(r.messages[eventName]) > 0 {
		for _, cb := range r.messages[eventName] {
			cb(args...)
		}
	}
}

func (r *eventEmitter) On(eventName string, cb callback) {
	r.messages[eventName] = append(r.messages[eventName], cb)
}

// m := make() make a map of Events that emit will be searched when emit('event name') is called
