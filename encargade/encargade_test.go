package encargade_test

import (
	"testing"

	"github.com/nickrisaro/freezer-bot/encargade"
	"github.com/nickrisaro/freezer-bot/freezer"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const conexiónALaBase = "file::memory:?cache=shared"

type EncargadeTestSuite struct {
	suite.Suite
	encargadeConBaseDeDatos *encargade.Encargade
	miFreezer               *freezer.Freezer
}

func (suite *EncargadeTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(conexiónALaBase), &gorm.Config{})
	suite.NoError(err, "Debería conectarse a la base de datos")
	suite.NotNil(db, "La base no debería ser nula")
	err = db.AutoMigrate(&freezer.Freezer{}, &freezer.Producto{})
	suite.NoError(err, "Debería ejecutar las migraciones")

	suite.encargadeConBaseDeDatos = encargade.NewEncargadeConBaseDeDatos(db)

	suite.NotNil(suite.encargadeConBaseDeDatos, "Le encargade no debería ser nil")

	suite.miFreezer = freezer.NewFreezer(1, "Un Freezer")
	result := db.Create(suite.miFreezer)
	suite.NoError(result.Error, "Debería haber creado el freezer")
	suite.NotNil(suite.miFreezer.ID, "El freezer debería tener ID")
}

func (suite *EncargadeTestSuite) TestSiNoHayNadaEnElFreezerLeEncargadeSabeQueEstáVacío() {
	encargade := encargade.NewEncargade(suite.miFreezer)

	suite.Equal("El freezer está vacío", encargade.QueCosasHayEnElFreezer(), "Esperaba que el freezer esté vacío")
}

func (suite *EncargadeTestSuite) TestSiLeDigoALeEncargadeQueAgregueUnaPizzaLaAgrega() {
	encargade := encargade.NewEncargade(suite.miFreezer)

	err := encargade.Meter("Pizza, 1, unidad")

	suite.Nil(err, "No esperaba un error")

	miProducto := freezer.NewProducto("Pizza", 1.0, freezer.Unidad)
	suite.Equal([]*freezer.Producto{miProducto}, suite.miFreezer.Productos, "Esperaba que haya una pizza")
}

func (suite *EncargadeTestSuite) TestSiHayUnaPizzaEnElFreezerLeEncargadeMeLoDice() {
	encargade := encargade.NewEncargade(suite.miFreezer)
	encargade.Meter("Pizza, 1, unidad")

	suite.Equal("El freezer tiene:\n\n- Pizza: 1.00 unidad(es)\n", encargade.QueCosasHayEnElFreezer(), "Esperaba que le encargade me diga que hay una pizza")
}

func TestEncargadeTestSuite(t *testing.T) {
	suite.Run(t, new(EncargadeTestSuite))
}
