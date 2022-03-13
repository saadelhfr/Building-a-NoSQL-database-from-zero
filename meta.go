package LibraDB

import "encoding/binary"

const (
	metaPageNum = 0
)

// meta is the meta page of the db
type meta struct {
	freelistPage pgnum
}

func newEmptyMeta() *meta {
	return &meta{}
}

func (m *meta) serialize() []byte {
	buf := make([]byte, metaSize)
	pos := 0

	binary.LittleEndian.PutUint64(buf[pos:], uint64(m.freelistPage))
	pos += pageNumSize

	return buf
}

func (m *meta) deserialize(buf []byte) {
	pos := 0
	m.freelistPage = pgnum(binary.LittleEndian.Uint64(buf[pos:]))
	pos += pageNumSize
}
