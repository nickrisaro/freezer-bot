package freezer

type Freezer interface {
	Productos() []interface{}
}

type FreezerEfímero struct {

}

func NewFreezerEfímero() Freezer {
	return new(FreezerEfímero)
}

func (f *FreezerEfímero) Productos() []interface{} {
	return make([]interface{}, 0)
}