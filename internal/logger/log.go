package logger

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"
)

type Logger struct {
	out io.Writer
	mu  sync.Mutex
}

func New(out io.Writer) *Logger {
	return &Logger{
		out: out,
	}
}

type logData struct {
	Time    string `json:"time"`
	Message string `json:"message"`
	Method  string `json:"method"`
	Path    string `json:"path"`
}

func (l *Logger) LogInfo(r *http.Request, message string) (int, error) {
	data := logData{
		Time:    time.Now().UTC().Format(time.RFC3339),
		Message: message,
		Method:  r.Method,
		Path:    r.URL.Path,
	}

	line, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, '\n'))
}
