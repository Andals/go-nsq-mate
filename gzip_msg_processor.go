package nsqmate

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

type GzipMessageProcessor struct {
}

func (this *GzipMessageProcessor) Process(msg []byte) []byte {
	b := bytes.NewBuffer([]byte{})
	w := gzip.NewWriter(b)
	w.Write(msg)
	w.Flush()
	w.Close()

	return b.Bytes()
}

func (this *GzipMessageProcessor) MultiProcess(msgs [][]byte) [][]byte {
	bs := make([][]byte, len(msgs))
	for i, msg := range msgs {
		bs[i] = this.Process(msg)
	}

	return bs
}

func (this *GzipMessageProcessor) Restore(msg []byte) ([]byte, error) {
	b := bytes.NewReader(msg)
	r, e := gzip.NewReader(b)
	if e != nil {
		return nil, e
	}

	body, e := ioutil.ReadAll(r)
	if e != nil && len(body) == 0 {
		return nil, e
	}

	return body, nil
}
