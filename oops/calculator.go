package main
//struct + constructor + methods 
type Calculator struct {
	Radius float64}
//constructor
func NewCalculator(radius float64)*Calculator{
	return &Calculator{Radius:radius}}

//comes reference to struct first then function name and if any additional params required
//u can do now &Calculator.Perimeter
func (c *Calculator)Perimeter() float64{
	return 2*3.14*c.Radius}
func (c *Calculator)  Area() float64{
	return 3.14*c.Radius*c.Radius}



// interface + struct + methods

type Shape interface {
	Area() float64
	Perimeter() float64}


type Circle struct {
	Radius float64	}

func (c *Circle)Perimeter() float64{
	return 2*3.14*c.Radius}
func (c *Circle)  Area() float64{
	return 3.14*c.Radius*c.Radius}


type Rectangle struct {Radius float64
	Height float64}



func (c *Rectangle)Perimeter() float64{
	return 2*c.Radius+2*c.Height}

func (c *Rectangle)  Area() float64{
	return c.Radius*c.Height}