package model

type Registry interface {
	Register(p Provider) error
	Provider(name string) (Provider, bool)
	List() []Provider
}
