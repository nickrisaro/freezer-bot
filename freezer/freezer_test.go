package freezer_test

import (
	"testing"

	"github.com/nickrisaro/freezer-bot/freezer"
	"github.com/stretchr/testify/assert"
)

func TestInicialmenteElFreezerEstáVacío(t *testing.T) {
	miFreezer := freezer.NewFreezerEfímero()

	assert.Equal(t, make([]interface{}, 0), miFreezer.Productos(), "Esperaba que el freezer esté vacío")
}
