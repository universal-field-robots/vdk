package nvr

import "github.com/universal-field-robots/vdk/av"

type Stream struct {
	codec av.CodecData
	idx   int
}
