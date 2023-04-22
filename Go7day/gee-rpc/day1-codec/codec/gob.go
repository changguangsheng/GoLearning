package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	//conn 是由构建函数传入，通常是通过 TCP 或者 Unix 建立 socket 时得到的链接实例
	conn io.ReadWriteCloser
	//buf 是为了防止阻塞而创建的带缓冲的 Writer，一般这么做能提升性能。
	buf *bufio.Writer
	//dec 和 enc 对应 gob 的 Decoder 和 Encoder，
	dec *gob.Decoder
	enc *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}
	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	return nil
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}
