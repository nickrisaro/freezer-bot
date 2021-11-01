package encargade_test

import (
	"testing"

	"github.com/nickrisaro/freezer-bot/encargade"
	"github.com/nickrisaro/freezer-bot/freezer"
	"github.com/stretchr/testify/assert"
)

func TestSiNoHayNadaEnElFreezerLeEncargadeSabeQueEstáVacío(t *testing.T) {
	miFreezer := freezer.NewFreezerEfímero()
	encargade := encargade.NewEncargade(miFreezer)

	assert.Equal(t, "El freezer está vacío", encargade.QueCosasHayEnElFreezer(), "Esperaba que el freezer esté vacío")
}

func TestSiLeDigoALeEncargadeQueAgregueUnaPizzaLaAgrega(t *testing.T) {
	miFreezer := freezer.NewFreezerEfímero()
	encargade := encargade.NewEncargade(miFreezer)

	err := encargade.Meter("Pizza, 1, unidad")

	assert.Nil(t, err, "No esperaba un error")

	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)
	assert.Equal(t, []*freezer.Producto{miProducto}, miFreezer.Productos(), "Esperaba que haya una pizza")
}
