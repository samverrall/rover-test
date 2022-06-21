package plateau

// Plateau represents the grid for vehicle.
type Plateau struct {
	// Unexported fields so that the plateau cannot be modified during execution.
	width  int
	height int
}

// New initialises a new Plateau.
func New(width, height int) *Plateau {
	return &Plateau{
		width:  width,
		height: height,
	}
}

// GetWidth returns the set width for the plateau.
func (p *Plateau) GetWidth() int {
	return p.width
}

// GetHeight returns the set width for the plateau.
func (p *Plateau) GetHeight() int {
	return p.height
}
