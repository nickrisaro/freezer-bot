package encargade

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nickrisaro/freezer-bot/freezer"
	"gorm.io/gorm"
)

type Encargade struct {
	miBaseDeDatos *gorm.DB
	miFreezer     *freezer.Freezer
}

func NewEncargadeConBaseDeDatos(baseDeDatos *gorm.DB) *Encargade {
	return &Encargade{miBaseDeDatos: baseDeDatos}
}

func NewEncargade(miFreezer *freezer.Freezer) *Encargade {
	return &Encargade{miFreezer: miFreezer}
}

func (e *Encargade) QueCosasHayEnEsteFreezer(identificador int64) string {
	freezerDeLaDB := freezer.Freezer{Identificador: identificador}
	resultado := e.miBaseDeDatos.Where(&freezerDeLaDB).Preload("Productos").First(&freezerDeLaDB)
	if resultado.Error != nil {
		return "No pude encontrar el freezer que me pedís"
	}

	productos := freezerDeLaDB.Productos

	if len(productos) == 0 {
		return "El freezer está vacío"
	}

	inventario := "El freezer tiene:\n\n"

	for _, producto := range productos {
		inventario += fmt.Sprintf("- %s\n", producto.String())
	}

	return inventario
}

func (e *Encargade) QueCosasHayEnElFreezer() string {
	productos := e.miFreezer.Productos

	if len(productos) == 0 {
		return "El freezer está vacío"
	}

	inventario := "El freezer tiene:\n\n"

	for _, producto := range productos {
		inventario += fmt.Sprintf("- %s\n", producto.String())
	}

	return inventario
}

func (e *Encargade) MeterEnFreezer(identificador int64, producto string) error {
	freezerDeLaDB := freezer.Freezer{Identificador: identificador}
	resultado := e.miBaseDeDatos.Where(&freezerDeLaDB).First(&freezerDeLaDB)
	if resultado.Error != nil {
		return resultado.Error
	}

	partes := strings.Split(producto, ",")

	cantidad, err := strconv.ParseFloat(strings.TrimSpace(partes[1]), 64)
	if err != nil {
		return err
	}

	elProducto := freezer.NewProducto(strings.TrimSpace(partes[0]), cantidad, stringAunidadDeMedida(strings.TrimSpace(partes[2])))
	freezerDeLaDB.Agregar(elProducto)
	if err != nil {
		return err
	}

	resultado = e.miBaseDeDatos.Save(freezerDeLaDB)
	if resultado.Error != nil {
		return resultado.Error
	}

	return nil
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
