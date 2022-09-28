package edit

type RotateTransform struct {
	Angle int `json:"angle"`
}

type SkewTransform struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type FlipTransform struct {
	Horizontal bool `json:"horizontal"`
	Vertical   bool `json:"vertical"`
}
type Transform struct {
	Rotate *RotateTransform `json:"rotate"`
	Skew   *SkewTransform   `json:"skew"`
	Flip   *FlipTransform   `json:"flip"`
}

func NewTransform() *Transform {
	return new(Transform)
}

func (t *Transform) SetFlip(flip *FlipTransform) *Transform {
	t.Flip = flip
	return t
}

func (t *Transform) SetRotate(rotate *RotateTransform) *Transform {
	t.Rotate = rotate
	return t
}

func (t *Transform) SetSkew(skew *SkewTransform) *Transform {
	t.Skew = skew
	return t
}

// rotate

func NewRotateTransform() *RotateTransform {
	return new(RotateTransform)
}

func (r *RotateTransform) SetAngle(angle int) *RotateTransform {
	r.Angle = angle
	return r
}

// skew

func NewSkewTransform() *SkewTransform {
	return new(SkewTransform)
}

func (s *SkewTransform) SetX(x float32) *SkewTransform {
	s.X = x
	return s
}

func (s *SkewTransform) SetY(y float32) *SkewTransform {
	s.Y = y
	return s
}

// flip

func NewFlipTransform() *FlipTransform {
	return new(FlipTransform)
}

func (f *FlipTransform) SetHorizontal(h bool) *FlipTransform {
	f.Horizontal = h
	return f
}

func (f *FlipTransform) SetVertical(v bool) *FlipTransform {
	f.Vertical = v
	return f
}
