package encargade

import (
	"strconv"
	"strings"

	"github.com/nickrisaro/freezer-bot/freezer"
)

type Encargade struct {
	miFreezer freezer.Freezer
}

func NewEncargade(miFreezer freezer.Freezer) *Encargade {
	return &Encargade{miFreezer: miFreezer}
}

func (e *Encargade) QueCosasHayEnElFreezer() string {
	return "El freezer está vacío"
}

func (e *Encargade) Meter(producto string) error {

	partes := strings.Split(producto, ",")

	cantidad, err := strconv.ParseFloat(strings.TrimSpace(partes[1]), 64)
	if err != nil {
		return err
	}

	elProducto := freezer.NewProducto(strings.TrimSpace(partes[0]), cantidad, stringAunidadDeMedida(strings.TrimSpace(partes[2])))
	e.miFreezer.Agregar(elProducto)

	return nil
}

func stringAunidadDeMedida(unidadDeMedida string) freezer.Medida {
	switch strings.ToUpper(unidadDeMedida) {
	case "UNIDAD":
		return freezer.Unidad
	case "KILO":
		return freezer.Kilo
	case "GRAMO":
		return freezer.Gramo
	case "LITRO":
		return freezer.Litro
	case "MILILITRO":
		return freezer.Mililitro
	}
	return freezer.Otra
}