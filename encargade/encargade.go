package encargade

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/nickrisaro/freezer-bot/freezer"
	"gorm.io/gorm"
)

type Encargade struct {
	miBaseDeDatos *gorm.DB
}

func NewEncargade(baseDeDatos *gorm.DB) *Encargade {
	return &Encargade{miBaseDeDatos: baseDeDatos}
}

func (e *Encargade) NuevoFreezer(identificador int64, nombre string) error {
	freezerParaLaDB := freezer.NewFreezer(identificador, nombre)

	resultado := e.miBaseDeDatos.Create(freezerParaLaDB)
	if resultado.Error != nil && strings.Contains(resultado.Error.Error(), "UNIQUE") {
		return errors.New("ya existe ese freezer")
	}
	return resultado.Error
}

func (e *Encargade) QueCosasHayEnEsteFreezer(identificador int64) string {
	productos := make([]freezer.Producto, 0)
	resultado := e.miBaseDeDatos.Table("productos").
		Select("productos.*").
		Joins("left join freezers on freezers.id = productos.freezer_id").
		Where("freezers.identificador = ?", identificador).
		Where("productos.cantidad > 0").
		Scan(&productos)

	if resultado.Error != nil {
		return "No pude encontrar el freezer que me pedís"
	}

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

	partes := strings.Split(producto, ",")

	if len(partes) < 3 {
		return errors.New("formato de producto incorrecto")
	}

	cantidad, err := strconv.ParseFloat(strings.TrimSpace(partes[1]), 64)
	if err != nil {
		return err
	}

	if cantidad <= 0 {
		return fmt.Errorf("cantidadNegativa")
	}

	elProducto := freezer.NewProducto(strings.TrimSpace(partes[0]), cantidad, stringAunidadDeMedida(strings.TrimSpace(partes[2])))

	freezerDeLaDB := freezer.Freezer{Identificador: identificador}
	resultado := e.miBaseDeDatos.Where(&freezerDeLaDB).Preload("Productos").First(&freezerDeLaDB)
	if resultado.Error != nil {
		return resultado.Error
	}

	elProducto = freezerDeLaDB.Agregar(elProducto)

	resultado = e.miBaseDeDatos.Save(elProducto)
	return resultado.Error
}

func (e *Encargade) SacarDelFreezer(identificador int64, producto string) error {

	partes := strings.Split(producto, ",")

	if len(partes) < 2 {
		return errors.New("formato de producto incorrecto")
	}

	cantidad, err := strconv.ParseFloat(strings.TrimSpace(partes[1]), 64)
	if err != nil {
		return err
	}

	if cantidad <= 0 {
		return fmt.Errorf("cantidadNegativa")
	}

	freezerDeLaDB := freezer.Freezer{Identificador: identificador}
	resultado := e.miBaseDeDatos.Where(&freezerDeLaDB).Preload("Productos").First(&freezerDeLaDB)
	if resultado.Error != nil {
		return resultado.Error
	}

	productoActualizado := freezerDeLaDB.Quitar(strings.TrimSpace(partes[0]), cantidad)
	if productoActualizado == nil {
		return errors.New("noExisteProducto")
	}

	resultado = e.miBaseDeDatos.Save(productoActualizado)

	return resultado.Error
}

func stringAunidadDeMedida(unidadDeMedida string) freezer.Medida {
	switch strings.ToUpper(unidadDeMedida) {
	case "UNIDAD", "UNIDADES", "U":
		return freezer.Unidad
	case "KILO", "KILOS", "K":
		return freezer.Kilo
	case "GRAMO", "GRAMOS", "G":
		return freezer.Gramo
	case "LITRO", "LITROS", "L":
		return freezer.Litro
	case "MILILITRO", "MILILITROS", "M":
		return freezer.Mililitro
	}
	return freezer.Otra
}
