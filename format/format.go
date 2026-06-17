package format

import (
	"github.com/universal-field-robots/vdk/av/avutil"
	"github.com/universal-field-robots/vdk/format/aac"
	"github.com/universal-field-robots/vdk/format/flv"
	"github.com/universal-field-robots/vdk/format/mp4"
	"github.com/universal-field-robots/vdk/format/rtmp"
	"github.com/universal-field-robots/vdk/format/rtsp"
	"github.com/universal-field-robots/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
