package freezer

import (
	"fmt"

	"gorm.io/gorm"
)

// Freezer guarda productos y los mantiene refrigerados
// Se pueden poner y sacar cosas en él y puedo fijarme que hay adentro
type Freezer interface {
	// Agregar me deja meter algo en el freezer
	Agregar(*Producto)
	// Quitar saca algo del freezer
	Quitar(string)
}

// FreezerEfimero guarda las cosas pero por un ratito no más
type FreezerEfímero struct {
	gorm.Model
	Identificador int64
	Nombre        string
	Productos     []*Producto
}

// NewFreezerEfímero construye un freezer efímero
func NewFreezerEfímero(identificador int64, nombre string) *FreezerEfímero {
	productos := make([]*Producto, 0)
	return &FreezerEfímero{Identificador: identificador, Nombre: nombre, Productos: productos}
}

func (f *FreezerEfímero) Agregar(producto *Producto) {
	f.Productos = append(f.Productos, producto)
}

func (f *FreezerEfímero) Quitar(nombreProducto string) {

	index := -1

	for i, producto := range f.Productos {
		if producto.Nombre == nombreProducto {
			index = i
			break
		}
	}

	if index != -1 {
		nuevosProductos := make([]*Producto, 0)
		nuevosProductos = append(nuevosProductos, f.Productos[:index]...)

		f.Productos = append(nuevosProductos, f.Productos[index+1:]...)
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
	gorm.Model
	Nombre         string
	Cantidad       float64
	UnidadDeMedida Medida
}

// NewProducto construye un nuevo producto para guardarlo en el freezer
func NewProducto(nombre string, cantidad float64, unidadDeMedida Medida) *Producto {
	return &Producto{Nombre: nombre, Cantidad: cantidad, UnidadDeMedida: unidadDeMedida}
}

func (p *Producto) String() string {
	return fmt.Sprintf("%s: %0.2f %s", p.Nombre, p.Cantidad, p.UnidadDeMedida)
}
