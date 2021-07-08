package error_learn

import (
	"fmt"
	"io"
)

type Status struct {
	Code int
	Reason string
}

type Header struct {
	Key, Value string
}

type errWrite struct {
	io.Writer
	err error
}

func (e *errWrite) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}

	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

/*
struct的匿名成员？
*/

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWrite{w, nil}
	_, _ = fmt.Fprintf(ew, "HTTP/1.1 %d %s \r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)

	return ew.err
}
