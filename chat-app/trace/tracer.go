package trace

import (
  "io"
  "fmt"
)

type Tracer interface {
  Trace(...interface{})
}

func New(w io.Writer) Tracer {
  return &tracer{out:w}
}

type tracer struct {
  out io.Writer
}

func (t *tracer) Trace(in ...interface{}) {
  fmt.Fprint(t.out, in...)
  fmt.Fprintln(t.out)

}

func Off() Tracer {
  return &nilTracer{}
}

type nilTracer struct {}

func (tracer *nilTracer) Trace(...interface{}) {

}
