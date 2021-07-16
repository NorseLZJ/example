package pb

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"

	"reflect"

	_ "github.com/davyxu/cellnet/codec/gogopb"
	"github.com/davyxu/cellnet/util"
)

func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{Codec: codec.MustGetCodec("gogopb"), Type: reflect.TypeOf((*Req)(nil)).Elem(), ID: int(util.StringHash("wsgate.Req"))})
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{Codec: codec.MustGetCodec("gogopb"), Type: reflect.TypeOf((*Ack)(nil)).Elem(), ID: int(util.StringHash("wsgate.Ack"))})
}
