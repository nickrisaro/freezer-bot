package freezer

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// Freezer guarda productos y los mantiene refrigerados
// Se pueden poner y sacar cosas en él y puedo fijarme que hay adentro
type Freezer struct {
	gorm.Model
	Identificador int64 `gorm:"unique"`
	Nombre        string
	Productos     []*Producto
}

// NewFreezer construye un freezer efímero
func NewFreezer(identificador int64, nombre string) *Freezer {
	productos := make([]*Producto, 0)
	return &Freezer{Identificador: identificador, Nombre: nombre, Productos: productos}
}

// Agregar agrega un producto al freezer.
// Si el producto ya estaba en el freezer se suman las cantidades
// Asume que las cantidades son en la misma unidad
func (f *Freezer) Agregar(producto *Producto) *Producto {
	index := -1

	for i, productoActual := range f.Productos {
		if strings.EqualFold(productoActual.Nombre, producto.Nombre) {
			index = i
			break
		}
	}

	var productoAActualizar *Producto = producto
	if index == -1 {
		f.Productos = append(f.Productos, producto)
	} else {
		productoAActualizar = f.Productos[index]
		productoAActualizar.Cantidad += producto.Cantidad
	}
	productoAActualizar.FreezerID = f.ID

	return productoAActualizar
}

// Quitar remueve cantidad unidades del producto identificado por nombreProducto
// Si luego de remover esas unidades no quedan más unidades de ese producto en el freezer lo elimina del freezer
func (f *Freezer) Quitar(nombreProducto string, cantidad float64) *Producto {

	index := -1

	for i, producto := range f.Productos {
		if strings.EqualFold(producto.Nombre, nombreProducto) {
			index = i
			break
		}
	}

	var productoAActualizar *Producto = nil
	if index != -1 {
		productoAActualizar = f.Productos[index]
		productoAActualizar.Cantidad -= cantidad

		if productoAActualizar.Cantidad <= 0.0 {
			nuevosProductos := make([]*Producto, 0)
			nuevosProductos = append(nuevosProductos, f.Productos[:index]...)

			f.Productos = append(nuevosProductos, f.Productos[index+1:]...)
		}
	}

	return productoAActualizar
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
	FreezerID      uint
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
