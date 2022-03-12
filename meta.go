package LibraDB

import "encoding/binary"

// meta is the meta page of the db
type meta struct {
	// The database has a root collection that holds all the collections in the database. It is called root and the
	// root property of meta holds page number containing the root of collections collection. The keys are the
	// collections names and the values are the page number of the root of each collection. Then, once the collection
	// and the root page are located, a search inside a collection can be made.
	root         pgnum
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
