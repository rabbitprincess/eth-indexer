package db

// DocType is an interface for structs to be used as database documents
type DocType interface {
	GetID() string
	SetID(string)
}

// BaseEsType implements DocType and contains the document's id
type BaseEsType struct {
	Id string `json:"-" db:"id"`
}

// GetID returns the document's id
func (m *BaseEsType) GetID() string {
	return m.Id
}

// SetID sets the document's id
func (m *BaseEsType) SetID(id string) {
	m.Id = id
}

var EsSchema map[string]string
