package vag

type Storage interface {
	Load(any) error
	Save(any) error
}
