package freezer

import "fmt"

// Freezer guarda productos y los mantiene refrigerados
// Se pueden poner y sacar cosas en él y puedo fijarme que hay adentro
type Freezer interface {
	// Productos me dice que cosas tengo freezadas
	Productos() []*Producto
	// Agregar me deja meter algo en el freezer
	Agregar(*Producto)
	// Quitar saca algo del freezer
	Quitar(string)
}

// FreezerEfimero guarda las cosas pero por un ratito no más
type FreezerEfímero struct {
	productos []*Producto
}

// NewFreezerEfímero construye un freezer efímero
func NewFreezerEfímero() Freezer {
	productos := make([]*Producto, 0)
	return &FreezerEfímero{productos: productos}
}

func (f *FreezerEfímero) Productos() []*Producto {
	return f.productos
}

func (f *FreezerEfímero) Agregar(producto *Producto) {
	f.productos = append(f.productos, producto)
}

func (f *FreezerEfímero) Quitar(nombreProducto string) {

	index := -1

	for i, producto := range f.productos {
		if producto.nombre == nombreProducto {
			index = i
			break
		}
	}

	if index != -1 {
		nuevosProductos := make([]*Producto, 0)
		nuevosProductos = append(nuevosProductos, f.productos[:index]...)

		f.productos = append(nuevosProductos, f.productos[index+1:]...)
	}
}

// Medida unidad de medida de los productos que guardo en el Freezer
type Medida int

const (
	Otra Medida = iota
	Unidad
	Kilo
	Gramo
	Litro
	Mililitro
)

func (s Medida) String() string {
	switch s {
	case Unidad:
		return "unidad(es)"
	case Kilo:
		return "kilo(s)"
	case Gramo:
		return "gramo(s)"
	case Litro:
		return "litro(s)"
	case Mililitro:
		return "mililitro(s)"
	}
	return ""
}

// Producto algo que se puede freezar
type Producto struct {
	nombre         string
	cantidad       float64
	unidadDeMedida Medida
}

// NewProducto construye un nuevo producto para guardarlo en el freezer
func NewProducto(nombre string, cantidad float64, unidadDeMedida Medida) *Producto {
	return &Producto{nombre: nombre, cantidad: cantidad, unidadDeMedida: unidadDeMedida}
}

func (p *Producto) String() string {
	return fmt.Sprintf("%s: %0.2f %s", p.nombre, p.cantidad, p.unidadDeMedida)
}
