package bytebufferpool_test

import (
	"fmt"

	"github.com/evan361425/bytebufferpool"
)

func ExampleByteBuffer() {
	bb := bytebufferpool.Get()

	_, _ = bb.WriteString("first line\n")
	_, _ = bb.Write([]byte("second line\n"))
	bb.B = append(bb.B, "third line\n"...)

	fmt.Printf("bytebuffer contents=%q", bb.B)

	// It is safe to release byte buffer now, since it is
	// no longer used.
	bytebufferpool.Put(bb)
}
