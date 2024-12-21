// +build !go1.8
// Code generated by "httpsnoop/codegen"; DO NOT EDIT

package httpsnoop

import (
	"bufio"
	"io"
	"net"
	"github.com/chenjz24/net/http"
)

// HeaderFunc is part of the http.ResponseWriter interface.
type HeaderFunc func() http.Header

// WriteHeaderFunc is part of the http.ResponseWriter interface.
type WriteHeaderFunc func(code int)

// WriteFunc is part of the http.ResponseWriter interface.
type WriteFunc func(b []byte) (int, error)

// FlushFunc is part of the http.Flusher interface.
type FlushFunc func()

// CloseNotifyFunc is part of the http.CloseNotifier interface.
type CloseNotifyFunc func() <-chan bool

// HijackFunc is part of the http.Hijacker interface.
type HijackFunc func() (net.Conn, *bufio.ReadWriter, error)

// ReadFromFunc is part of the io.ReaderFrom interface.
type ReadFromFunc func(src io.Reader) (int64, error)

// Hooks defines a set of method interceptors for methods included in
// http.ResponseWriter as well as some others. You can think of them as
// middleware for the function calls they target. See Wrap for more details.
type Hooks struct {
	Header      func(HeaderFunc) HeaderFunc
	WriteHeader func(WriteHeaderFunc) WriteHeaderFunc
	Write       func(WriteFunc) WriteFunc
	Flush       func(FlushFunc) FlushFunc
	CloseNotify func(CloseNotifyFunc) CloseNotifyFunc
	Hijack      func(HijackFunc) HijackFunc
	ReadFrom    func(ReadFromFunc) ReadFromFunc
}

// Wrap returns a wrapped version of w that provides the exact same interface
// as w. Specifically if w implements any combination of:
//
// - http.Flusher
// - http.CloseNotifier
// - http.Hijacker
// - io.ReaderFrom
//
// The wrapped version will implement the exact same combination. If no hooks
// are set, the wrapped version also behaves exactly as w. Hooks targeting
// methods not supported by w are ignored. Any other hooks will intercept the
// method they target and may modify the call's arguments and/or return values.
// The CaptureMetrics implementation serves as a working example for how the
// hooks can be used.
func Wrap(w http.ResponseWriter, hooks Hooks) http.ResponseWriter {
	rw := &rw{w: w, h: hooks}
	_, i0 := w.(http.Flusher)
	_, i1 := w.(http.CloseNotifier)
	_, i2 := w.(http.Hijacker)
	_, i3 := w.(io.ReaderFrom)
	switch {
	// combination 1/16
	case !i0 && !i1 && !i2 && !i3:
		return struct {
			http.ResponseWriter
		}{rw}
	// combination 2/16
	case !i0 && !i1 && !i2 && i3:
		return struct {
			http.ResponseWriter
			io.ReaderFrom
		}{rw, rw}
	// combination 3/16
	case !i0 && !i1 && i2 && !i3:
		return struct {
			http.ResponseWriter
			http.Hijacker
		}{rw, rw}
	// combination 4/16
	case !i0 && !i1 && i2 && i3:
		return struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw}
	// combination 5/16
	case !i0 && i1 && !i2 && !i3:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
		}{rw, rw}
	// combination 6/16
	case !i0 && i1 && !i2 && i3:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw}
	// combination 7/16
	case !i0 && i1 && i2 && !i3:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
		}{rw, rw, rw}
	// combination 8/16
	case !i0 && i1 && i2 && i3:
		return struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 9/16
	case i0 && !i1 && !i2 && !i3:
		return struct {
			http.ResponseWriter
			http.Flusher
		}{rw, rw}
	// combination 10/16
	case i0 && !i1 && !i2 && i3:
		return struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{rw, rw, rw}
	// combination 11/16
	case i0 && !i1 && i2 && !i3:
		return struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
		}{rw, rw, rw}
	// combination 12/16
	case i0 && !i1 && i2 && i3:
		return struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 13/16
	case i0 && i1 && !i2 && !i3:
		return struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{rw, rw, rw}
	// combination 14/16
	case i0 && i1 && !i2 && i3:
		return struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 15/16
	case i0 && i1 && i2 && !i3:
		return struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
		}{rw, rw, rw, rw}
	// combination 16/16
	case i0 && i1 && i2 && i3:
		return struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	}
	panic("unreachable")
}

type rw struct {
	w http.ResponseWriter
	h Hooks
}

func (w *rw) Header() http.Header {
	f := w.w.(http.ResponseWriter).Header
	if w.h.Header != nil {
		f = w.h.Header(f)
	}
	return f()
}

func (w *rw) WriteHeader(code int) {
	f := w.w.(http.ResponseWriter).WriteHeader
	if w.h.WriteHeader != nil {
		f = w.h.WriteHeader(f)
	}
	f(code)
}

func (w *rw) Write(b []byte) (int, error) {
	f := w.w.(http.ResponseWriter).Write
	if w.h.Write != nil {
		f = w.h.Write(f)
	}
	return f(b)
}

func (w *rw) Flush() {
	f := w.w.(http.Flusher).Flush
	if w.h.Flush != nil {
		f = w.h.Flush(f)
	}
	f()
}

func (w *rw) CloseNotify() <-chan bool {
	f := w.w.(http.CloseNotifier).CloseNotify
	if w.h.CloseNotify != nil {
		f = w.h.CloseNotify(f)
	}
	return f()
}

func (w *rw) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	f := w.w.(http.Hijacker).Hijack
	if w.h.Hijack != nil {
		f = w.h.Hijack(f)
	}
	return f()
}

func (w *rw) ReadFrom(src io.Reader) (int64, error) {
	f := w.w.(io.ReaderFrom).ReadFrom
	if w.h.ReadFrom != nil {
		f = w.h.ReadFrom(f)
	}
	return f(src)
}
