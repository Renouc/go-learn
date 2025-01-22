package car

// 选项模式 创建对象

type Car struct {
	Brand string
	Color string
	Model string
	Price int
	Year  int
}

type Option func(*Car)

func WithBrand(brand string) Option {
	return func(c *Car) {
		c.Brand = brand
	}
}

func WithModel(model string) Option {
	return func(c *Car) {
		c.Model = model
	}
}

func WithYear(year int) Option {
	return func(c *Car) {
		c.Year = year
	}
}

func WithPrice(price int) Option {
	return func(c *Car) {
		c.Price = price
	}
}

func WithColor(color string) Option {
	return func(c *Car) {
		c.Color = color
	}
}

func NewCar(opts ...Option) *Car {
	car := &Car{
		Brand: "Unknown",
		Color: "Unknown",
		Model: "Unknown",
		Price: 0,
		Year:  0,
	}
	for _, opt := range opts {
		opt(car)
	}
	return car
}
