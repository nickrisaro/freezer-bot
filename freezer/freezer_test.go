package freezer_test

import (
	"testing"

	"github.com/nickrisaro/freezer-bot/freezer"
	"github.com/stretchr/testify/assert"
)

func TestElFreezerTieneIdentificadorYNombre(t *testing.T) {
	var identificador int64 = 1
	nombre := "El freezer de Nick"
	miFreezer := freezer.NewFreezer(identificador, nombre)

	assert.Equal(t, identificador, miFreezer.Identificador, "El identificador no es el esperado")
	assert.Equal(t, nombre, miFreezer.Nombre, "El nombre no es el esperado")
}

func TestInicialmenteElFreezerEstáVacío(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")

	assert.Equal(t, []*freezer.Producto{}, miFreezer.Productos, "Esperaba que el freezer esté vacío")
}

func TestPuedoAgregarUnaUnidadDeUnProducto(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)

	miFreezer.Agregar(miProducto)

	assert.Equal(t, []*freezer.Producto{miProducto}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
}

func TestSiUnProductoYaEstáEnElFreezerAlAgregarloDeVueltaSeActualizaLaCantidad(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miFreezer.Agregar(miProducto)

	assert.Equal(t, []*freezer.Producto{freezer.NewProducto("Pizza", 2.0, freezer.Unidad)}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
}

func TestSiUnProductoYaEstáEnElFreezerAlAgregarloEnMinúsculasSeActualizaLaCantidad(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miProductoEnMinúsuclas := freezer.NewProducto("pizza", 1.0, freezer.Unidad)
	miFreezer.Agregar(miProductoEnMinúsuclas)

	assert.Equal(t, []*freezer.Producto{freezer.NewProducto("Pizza", 2.0, freezer.Unidad)}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
}

func TestSePuedeSacarUnProductoDelFreezer(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miFreezer.Quitar("Pizza", 1.0)

	assert.Equal(t, []*freezer.Producto{}, miFreezer.Productos, "Esperaba que el freezer esté vacío")
}

func TestSiHayMásDeUnaUnidadAlQuitarSeReduceLaCantidad(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 2.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miFreezer.Quitar("Pizza", 1.0)

	assert.Equal(t, []*freezer.Producto{freezer.NewProducto("Pizza", 1.0, freezer.Unidad)}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
}

func TestSiSacoMásDeLoQueHayEliminaElProducto(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 2.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miFreezer.Quitar("Pizza", 3.0)

	assert.Equal(t, []*freezer.Producto{}, miFreezer.Productos, "Esperaba que el freezer esté vacío")
}

func TestSiQuitoUnProductoEnMinúsculasSeReduceLaCantidad(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 2.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miFreezer.Quitar("pizza", 1.0)

	assert.Equal(t, []*freezer.Producto{freezer.NewProducto("Pizza", 1.0, freezer.Unidad)}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
}
