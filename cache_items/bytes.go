package cache_items

type BytesData struct {
	key  []byte
	data []byte
}

func NewByteData(key []byte, data []byte) *BytesData {
	return &BytesData{
		key:  key,
		data: data,
	}
}

func (bd *BytesData) GetKey() string {
	return string(bd.key)
}

func (bd *BytesData) String() string {
	return bd.GetKey()
}

func (bd *BytesData) GetKeyBytes() []byte {
	return bd.key
}

func (bd *BytesData) GetData() []byte {
	return bd.data
}
