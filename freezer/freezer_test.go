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
	miFreezer.ID = 1
	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)

	productoAgregado := miFreezer.Agregar(miProducto)

	assert.Equal(t, []*freezer.Producto{miProducto}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
	assert.Equal(t, "Pizza", productoAgregado.Nombre, "El producto no tiene el nombre esperado")
	assert.Equal(t, miFreezer.ID, productoAgregado.FreezerID, "El producto no tiene el nombre esperado")
}

func TestSiUnProductoYaEstáEnElFreezerAlAgregarloDeVueltaSeActualizaLaCantidad(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miFreezer.ID = 1
	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	productoAgregado := miFreezer.Agregar(miProducto)

	productoEsperado := freezer.NewProducto("Pizza", 2.0, freezer.Unidad)
	productoEsperado.FreezerID = miFreezer.ID

	assert.Equal(t, []*freezer.Producto{productoEsperado}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
	assert.Equal(t, productoEsperado, productoAgregado, "El producto agregado no es el esperado")
}

func TestSiUnProductoYaEstáEnElFreezerAlAgregarloEnMinúsculasSeActualizaLaCantidad(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miFreezer.ID = 1
	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miProductoEnMinúsuclas := freezer.NewProducto("pizza", 1.0, freezer.Unidad)
	productoAgregado := miFreezer.Agregar(miProductoEnMinúsuclas)

	productoEsperado := freezer.NewProducto("Pizza", 2.0, freezer.Unidad)
	productoEsperado.FreezerID = miFreezer.ID

	assert.Equal(t, []*freezer.Producto{productoEsperado}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
	assert.Equal(t, productoEsperado, productoAgregado, "El producto agregado no es el esperado")
}

func TestSePuedeSacarUnProductoDelFreezer(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miPizza := miFreezer.Quitar("Pizza", 1.0)

	assert.Equal(t, []*freezer.Producto{}, miFreezer.Productos, "Esperaba que el freezer esté vacío")
	assert.Equal(t, freezer.NewProducto("Pizza", 0.0, freezer.Unidad), miPizza, "Esperaba una pizza")
}

func TestSiHayMásDeUnaUnidadAlQuitarSeReduceLaCantidad(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 2.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miPizza := miFreezer.Quitar("Pizza", 1.0)

	assert.Equal(t, []*freezer.Producto{freezer.NewProducto("Pizza", 1.0, freezer.Unidad)}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
	assert.Equal(t, freezer.NewProducto("Pizza", 1.0, freezer.Unidad), miPizza, "Esperaba una pizza")
}

func TestSiSacoMásDeLoQueHayEliminaElProducto(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 2.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miFreezer.Quitar("Pizza", 3.0)

	assert.Equal(t, []*freezer.Producto{}, miFreezer.Productos, "Esperaba que el freezer esté vacío")
}

func TestSiSacoMásDeLoQueHayElProductoRetornadoTieneCeroEnLaCantidad(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 2.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	productoQuitado := miFreezer.Quitar("Pizza", 3.0)

	assert.Equal(t, float64(0), productoQuitado.Cantidad, "Esperaba que el producto esté en cero")
}

func TestSiQuitoUnProductoEnMinúsculasSeReduceLaCantidad(t *testing.T) {
	miFreezer := freezer.NewFreezer(1, "Un Freezer")
	miProducto := freezer.NewProducto("Pizza", 2.0, freezer.Unidad)
	miFreezer.Agregar(miProducto)

	miFreezer.Quitar("pizza", 1.0)

	assert.Equal(t, []*freezer.Producto{freezer.NewProducto("Pizza", 1.0, freezer.Unidad)}, miFreezer.Productos, "Esperaba que el freezer tenga una pizza")
}
