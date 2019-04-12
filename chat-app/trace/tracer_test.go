package trace

import (
  "testing"
  "bytes"
)

func TestNew(t *testing.T) {
  var buf bytes.Buffer
  tracer := New(&buf)
  if tracer == nil {
    t.Error("should not be nil")
  } else {
    tracer.Trace("Hello World")
    if buf.String() != "Hello World\n" {
      t.Errorf("Should be equal to Hello World and not as %s", buf.String())
    }
  }
}

func TestOff(t *testing.T) {
  var silentTracer Tracer = Off()
  silentTracer.Trace("something")
}
