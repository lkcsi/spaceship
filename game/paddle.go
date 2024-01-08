package game

type Paddle struct {
	Pos    Vector
	Width  int
	Height int
}

func NewPaddle(x, y float64, width, height int) *Paddle {
	return &Paddle{Vector{x, y}, width, height}
}
