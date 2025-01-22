package car_test

import (
	"gold/car"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCar_DefaultValues(t *testing.T) {
	car := car.NewCar()
	assert.Equal(t, "Unknown", car.Brand, "Brand should be default value")
	assert.Equal(t, "Unknown", car.Color, "Color should be default value")
	assert.Equal(t, "Unknown", car.Model, "Model should be default value")
	assert.Equal(t, 0, car.Price, "Price should be default value")
	assert.Equal(t, 0, car.Year, "Year should be default value")
}

func TestNewCar_WithBrand(t *testing.T) {
	car := car.NewCar(car.WithBrand("Toyota"))
	assert.Equal(t, "Toyota", car.Brand, "Brand should be 'Toyota'")
}

func TestNewCar_WithModel(t *testing.T) {
	car := car.NewCar(car.WithModel("Corolla"))
	assert.Equal(t, "Corolla", car.Model, "Model should be 'Corolla'")
}

func TestNewCar_WithYear(t *testing.T) {
	car := car.NewCar(car.WithYear(2022))
	assert.Equal(t, 2022, car.Year, "Year should be 2022")
}

func TestNewCar_WithPrice(t *testing.T) {
	car := car.NewCar(car.WithPrice(25000))
	assert.Equal(t, 25000, car.Price, "Price should be 25000")
}

func TestNewCar_WithColor(t *testing.T) {
	car := car.NewCar(car.WithColor("Red"))
	assert.Equal(t, "Red", car.Color, "Color should be 'Red'")
}

func TestNewCar_WithMultipleOptions(t *testing.T) {
	car := car.NewCar(
		car.WithBrand("Toyota"),
		car.WithModel("Corolla"),
		car.WithYear(2022),
		car.WithPrice(25000),
		car.WithColor("Red"),
	)
	assert.Equal(t, "Toyota", car.Brand, "Brand should be 'Toyota'")
	assert.Equal(t, "Corolla", car.Model, "Model should be 'Corolla'")
	assert.Equal(t, 2022, car.Year, "Year should be 2022")
	assert.Equal(t, 25000, car.Price, "Price should be 25000")
	assert.Equal(t, "Red", car.Color, "Color should be 'Red'")
}
