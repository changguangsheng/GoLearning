package codec

import "io"

type Header struct {
	//服务名和方法名，通常与 Go 语言中的结构体和方法相映射。
	ServiceMethod string //format "Service.Method"
	//请求的序号，也可以认为是某个请求的 ID，用来区分不同的请求。
	Seq uint64 // sequence number chosen by client
	//错误信息，客户端置为空，服务端如果如果发生错误，将错误信息置于 Error 中。
	Error error
}

// 抽象出 Codec 的构造函数，客户端和服务端可以通过 Codec 的 Type 得到构造函数，从而创建 Codec 实例。
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodeFunc func(reader io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodeFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodeFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
