package game

type Paddle struct {
	Pos    Vector
	Width  int
	Height int
}

func (paddle *Paddle) NewPaddle(pos Vector, width, height int) *Paddle {
	return &Paddle{pos, width, height}
}
