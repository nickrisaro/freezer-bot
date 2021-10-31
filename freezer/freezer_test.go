package freezer_test

import (
	"testing"

	"github.com/nickrisaro/freezer-bot/freezer"
	"github.com/stretchr/testify/assert"
)

func TestInicialmenteElFreezerEstáVacío(t *testing.T) {
	miFreezer := freezer.NewFreezerEfímero()

	assert.Equal(t, []*freezer.Producto{}, miFreezer.Productos(), "Esperaba que el freezer esté vacío")
}

func TestPuedoAgregarUnaUnidadDeUnProducto(t *testing.T) {
	miFreezer := freezer.NewFreezerEfímero()
	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)

	miFreezer.Agregar(miProducto)

	assert.Equal(t, []*freezer.Producto{miProducto}, miFreezer.Productos(), "Esperaba que el freezer tenga una pizza")
}
