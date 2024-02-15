package responsesnoop

import (
	"net/http"
	"sync"

	"github.com/felixge/httpsnoop"
)

type Snoop struct {
	writer  http.ResponseWriter
	status  int
	written bool
}

var snoopPool = sync.Pool{
	New: func() any {
		return &Snoop{}
	},
}

func NewSnooper(w http.ResponseWriter) *Snoop {
	s := snoopPool.Get().(*Snoop)
	s.status = http.StatusOK
	s.written = false
	s.writer = httpsnoop.Wrap(w, httpsnoop.Hooks{
		Write: func(next httpsnoop.WriteFunc) httpsnoop.WriteFunc {
			return func(b []byte) (int, error) {
				if !s.written {
					s.written = true
				}
				return next(b)
			}
		},
		WriteHeader: func(next httpsnoop.WriteHeaderFunc) httpsnoop.WriteHeaderFunc {
			return func(statusCode int) {
				if !s.written {
					s.written = true
					s.status = statusCode
				}
				next(statusCode)
			}
		},
	})
	return s
}

func (s *Snoop) GetWriter() http.ResponseWriter {
	return s.writer
}

func (s *Snoop) GetStatus() int {
	return s.status
}

func (s *Snoop) Release() {
	s.writer = nil
	s.status = 0
	s.written = false
	snoopPool.Put(s)
}
