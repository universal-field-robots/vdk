package mjpeg

import "github.com/universal-field-robots/vdk/av"

type CodecData struct {
}

func (d CodecData) Type() av.CodecType {
	return av.MJPEG
}
