package rect

//Rect is a Rectangle
type Rect struct {
	X float64
	Y float64
	W float64
	H float64
}

//GetLeft returns the left x coordinate of the rect
func (r *Rect) GetLeft() float64 {
	if r.W > 0 {
		return r.X
	}
	return r.X + r.W

}

//GetRight returns the right x coordinate of the rect
func (r *Rect) GetRight() float64 {
	if r.W > 0 {
		return r.X + r.W
	}
	return r.X

}

//GetTop returns the top y coordinate of the rect
func (r *Rect) GetTop() float64 {
	if r.H > 0 {
		return r.Y
	}
	return r.Y + r.H

}

//GetBottom returns the bottom y coordinate of the rect
func (r *Rect) GetBottom() float64 {
	if r.H > 0 {
		return r.Y + r.H
	}

	return r.Y

}
