package view

import (
	"sync"
)

// LogHandler ...
type LogHandler interface {
	HandleLog(string)
}

// LogWriter ...
type LogWriter struct {
	sync.Mutex
	logs     []string
	index    int
	handlers map[LogHandler]struct{}
}

// NewLogWriter ...
func NewLogWriter(buf int) *LogWriter {
	return &LogWriter{
		logs:     make([]string, buf),
		index:    0,
		handlers: make(map[LogHandler]struct{}),
	}
}

// RegisterHandler ...
func (l *LogWriter) RegisterHandler(lh LogHandler) {
	l.Lock()
	defer l.Unlock()

	if _, ok := l.handlers[lh]; ok {
		return
	}

	l.handlers[lh] = struct{}{}

	if l.logs[l.index] != "" {
		for i := l.index; i < len(l.logs); i++ {
			lh.HandleLog(l.logs[i])
		}
	}
	for i := 0; i < l.index; i++ {
		lh.HandleLog(l.logs[i])
	}
}

// DeregisterHandler ...
func (l *LogWriter) DeregisterHandler(lh LogHandler) {
	l.Lock()
	defer l.Unlock()
	delete(l.handlers, lh)
}

// Write ...
func (l *LogWriter) Write(p []byte) (n int, err error) {
	l.Lock()
	defer l.Unlock()

	n = len(p)
	if p[n-1] == '\n' {
		p = p[:n-1]
	}

	l.logs[l.index] = string(p)
	l.index = (l.index + 1) % len(l.logs)

	for lh, _ := range l.handlers {
		lh.HandleLog(string(p))
	}
	return
}
