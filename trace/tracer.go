package trace

// tracerはコード内での出来事を記録できるオブジェクトを表すインターフェイスです

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}
type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// OffはTraceメソッドの読み出しを無視するTracerを返します。
func Off() Tracer {
	return &nilTracer{}
}
