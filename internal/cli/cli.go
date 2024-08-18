package cli

import "io"

var NewMarshaler = newEncoder

func (e *encoder) Marshal(v any, w io.Writer) error { return e.marshal(v, w) }
