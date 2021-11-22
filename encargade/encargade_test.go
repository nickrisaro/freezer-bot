package encargade_test

import (
	"math/rand"
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
	db        *gorm.DB
	encargade *encargade.Encargade
	miFreezer *freezer.Freezer
}

func (suite *EncargadeTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(conexiónALaBase), &gorm.Config{})
	suite.NoError(err, "Debería conectarse a la base de datos")
	suite.NotNil(db, "La base no debería ser nula")
	suite.db = db

	err = suite.db.AutoMigrate(&freezer.Freezer{}, &freezer.Producto{})
	suite.NoError(err, "Debería ejecutar las migraciones")

	suite.encargade = encargade.NewEncargade(suite.db)

	suite.NotNil(suite.encargade, "Le encargade no debería ser nil")

	suite.miFreezer = freezer.NewFreezer(int64(rand.Int()), "Un Freezer")
	result := suite.db.Create(suite.miFreezer)
	suite.NoError(result.Error, "Debería haber creado el freezer")
	suite.NotNil(suite.miFreezer.ID, "El freezer debería tener ID")
}

func (suite *EncargadeTestSuite) TestSiNoHayNadaEnElFreezerLeEncargadeSabeQueEstáVacío() {
	suite.Equal("El freezer está vacío", suite.encargade.QueCosasHayEnEsteFreezer(suite.miFreezer.Identificador), "Esperaba que el freezer esté vacío")
}

func (suite *EncargadeTestSuite) TestSiLeDigoALeEncargadeQueAgregueUnaPizzaLaAgrega() {
	err := suite.encargade.MeterEnFreezer(suite.miFreezer.Identificador, "Pizza, 1, unidad")

	suite.Nil(err, "No esperaba un error")

	freezerDeLaDB := freezer.Freezer{}
	resultado := suite.db.Preload("Productos").First(&freezerDeLaDB, suite.miFreezer.ID)
	suite.NoError(resultado.Error, "Debería haber encontrado el freezer")
	suite.NotEmpty(freezerDeLaDB.Productos, "Debería tener productos")

	miProducto := freezerDeLaDB.Productos[0]
	suite.Equal("Pizza", miProducto.Nombre, "Esperaba una pizza")
	suite.Equal(1.0, miProducto.Cantidad, "Esperaba una unidad")
	suite.Equal(freezer.Unidad, miProducto.UnidadDeMedida, "Esperaba unidad como unidad de medida")
}

func (suite *EncargadeTestSuite) TestSiHayUnaPizzaEnElFreezerLeEncargadeMeLoDice() {
	err := suite.encargade.MeterEnFreezer(suite.miFreezer.Identificador, "Pizza, 1, unidad")

	suite.Nil(err, "No esperaba un error")

	suite.Equal("El freezer tiene:\n\n- Pizza: 1.00 unidad(es)\n", suite.encargade.QueCosasHayEnEsteFreezer(suite.miFreezer.Identificador), "Esperaba que le encargade me diga que hay una pizza")
}

func TestEncargadeTestSuite(t *testing.T) {
	suite.Run(t, new(EncargadeTestSuite))
}
