package draw

func (list *List) BeginCommand() {
	list.Commands = append(list.Commands, Command{
		Clip:    list.CurrentClip,
		Texture: list.CurrentTexture,
	})
	list.CurrentCommand = &list.Commands[len(list.Commands)-1]
}

func (list *List) Primitive_Reserve(index_count, vertex_count int) {
	cmd := list.CurrentCommand
	cmd.Count += Index(index_count)
}

func (list *List) Primitive_Rect(r *Rectangle, color Color) {
	a, b, c, d := r.Corners()

	base := Index(len(list.Vertices))
	list.Indicies = append(list.Indicies,
		base+0, base+1, base+2,
		base+0, base+2, base+3,
	)
	list.Vertices = append(list.Vertices,
		Vertex{a, noUV, color},
		Vertex{b, noUV, color},
		Vertex{c, noUV, color},
		Vertex{d, noUV, color},
	)
}

func (list *List) Primitive_RectUV(r *Rectangle, uv *Rectangle, color Color) {
	a, b, c, d := r.Corners()
	uv_a, uv_b, uv_c, uv_d := uv.Corners()

	base := Index(len(list.Vertices))
	list.Indicies = append(list.Indicies,
		base+0, base+1, base+2,
		base+0, base+2, base+3,
	)
	list.Vertices = append(list.Vertices,
		Vertex{a, uv_a, color},
		Vertex{b, uv_b, color},
		Vertex{c, uv_c, color},
		Vertex{d, uv_d, color},
	)
}

func (list *List) Primitive_Quad(a, b, c, d Vector, color Color) {
	base := Index(len(list.Vertices))
	list.Indicies = append(list.Indicies,
		base+0, base+1, base+2,
		base+0, base+2, base+3,
	)
	list.Vertices = append(list.Vertices,
		Vertex{a, noUV, color},
		Vertex{b, noUV, color},
		Vertex{c, noUV, color},
		Vertex{d, noUV, color},
	)
}

func (list *List) Primitive_Tri(a, b, c Vector, color Color) {
	base := Index(len(list.Vertices))
	list.Indicies = append(list.Indicies,
		base+0, base+1, base+2,
	)
	list.Vertices = append(list.Vertices,
		Vertex{a, noUV, color},
		Vertex{b, noUV, color},
		Vertex{c, noUV, color},
	)
}

func (list *List) Primitive_QuadUV(q *[4]Vector, uv *Rectangle, color Color) {
	a, b, c, d := q[0], q[1], q[2], q[3]
	uv_a, uv_b, uv_c, uv_d := uv.Corners()

	base := Index(len(list.Vertices))
	list.Indicies = append(list.Indicies,
		base+0, base+1, base+2,
		base+0, base+2, base+3,
	)
	list.Vertices = append(list.Vertices,
		Vertex{a, uv_a, color},
		Vertex{b, uv_b, color},
		Vertex{c, uv_c, color},
		Vertex{d, uv_d, color},
	)
}

func (list *List) AddLine(points []Vector, thickness float32, color Color) {
	if len(points) < 2 || color.Transparent() || thickness == 0 {
		return
	}

	startIndexCount := len(list.Indicies)

	R := abs(thickness / 2.0)

	a := points[0]
	var x1, x2, xn Vector

	// draw each segment, where
	// a1-------^---------b1
	// |        | abn      |
	// a - - - - - - - - - b
	// |                   |
	// a2-----------------b2
	// x1, x2, xn are the previous segments end corners and normal
	for i, b := range points[1:] {
		// segment normal
		abn := SegmentNormal(a, b).ScaleTo(R)
		// segment corners
		a1, a2 := a.Add(abn), a.Sub(abn)
		b1, b2 := b.Add(abn), b.Sub(abn)

		if i > 0 && R > 1.5 {
			// draw segment chamfer
			d := xn.Rotate().Dot(abn)
			if d < 0 {
				list.Primitive_Tri(x1, a1, a, color)
			} else if d > 0 {
				list.Primitive_Tri(x2, a2, a, color)

			}
		}
		// draw block segment
		list.Primitive_Quad(a1, b1, b2, a2, color)

		a = b
		x1, x2, xn = b1, b2, abn
	}

	list.CurrentCommand.Count += Index(len(list.Indicies) - startIndexCount)
}

func (list *List) AddClosedLine(points []Vector, thickness float32, color Color) {
	if len(points) < 2 || color.Transparent() || thickness == 0 {
		return
	}
	if len(points) < 3 {
		list.AddLine(points, thickness, color)
		return
	}

	startIndexCount := len(list.Indicies)

	R := abs(thickness / 2.0)
	a := points[len(points)-1]
	xn := SegmentNormal(points[len(points)-2], a).ScaleTo(R)
	x1, x2 := a.Add(xn), a.Sub(xn)

	// draw each segment, where
	// a1-------^---------b1
	// |        | abn      |
	// a - - - - - - - - - b
	// |                   |
	// a2-----------------b2
	// x1, x2, xn are the previous segments end corners and normal
	for _, b := range points {
		// segment normal
		abn := SegmentNormal(a, b).ScaleTo(R)
		// segment corners
		a1, a2 := a.Add(abn), a.Sub(abn)
		b1, b2 := b.Add(abn), b.Sub(abn)

		// draw segment chamfer
		if R > 1.5 {
			d := xn.Rotate().Dot(abn)
			if d < 0 {
				list.Primitive_Tri(x1, a1, a, color)
			} else if d > 0 {
				list.Primitive_Tri(x2, a2, a, color)
			}
		}

		// draw block segment
		list.Primitive_Quad(a1, b1, b2, a2, color)

		a = b
		x1, x2, xn = b1, b2, abn
	}

	list.CurrentCommand.Count += Index(len(list.Indicies) - startIndexCount)
}

func (list *List) AddRectFill(r *Rectangle, color Color) {
	if color.Transparent() {
		return
	}

	list.Primitive_Reserve(6, 4)
	list.Primitive_Rect(r, color)
}

const segmentsPerArc = 24

func (list *List) AddArc(center Vector, R float32, start, sweep float32, color Color) {
	if color.Transparent() || R == 0 {
		return
	}
	R = abs(R)
	startIndexCount := len(list.Indicies)

	// N := sweep * R gives one segment per pixel
	N := Index(clamp(abs(sweep)*R/(2*pi), 3, segmentsPerArc))

	theta := sweep / float32(N)
	rotc, rots := cos(theta), sin(theta)

	dx, dy := cos(start)*R, sin(start)*R

	// add center point to the vertex buffer
	base := Index(len(list.Vertices))
	list.Vertices = append(list.Vertices, Vertex{center, noUV, color})
	// add the first point the vertex buffer
	p := Vector{center.X + dx, center.Y + dy}
	list.Vertices = append(list.Vertices, Vertex{p, noUV, color})
	// loop over rest of the points
	for i := Index(0); i < N; i++ {
		dx, dy = dx*rotc-dy*rots, dx*rots+dy*rotc
		p = Vector{center.X + dx, center.Y + dy}
		list.Vertices = append(list.Vertices, Vertex{p, noUV, color})
		list.Indicies = append(list.Indicies, base, base+i+1, base+i+2)
	}

	list.CurrentCommand.Count += Index(len(list.Indicies) - startIndexCount)
}

func (list *List) AddCircle(center Vector, R float32, color Color) {
	if color.Transparent() || R == 0 {
		return
	}
	R = abs(R)
	startIndexCount := len(list.Indicies)

	// N := 2 * PI * R gives one segment per pixel
	N := Index(clamp(R, 3, segmentsPerArc))

	theta := 2 * pi / float32(N)
	rotc, rots := cos(theta), sin(theta)

	dx, dy := R, float32(0)

	// add center point to the vertex buffer
	base := Index(len(list.Vertices))
	list.Vertices = append(list.Vertices, Vertex{center, noUV, color})
	// add the first point the vertex buffer
	p := Vector{center.X + dx, center.Y + dy}
	list.Vertices = append(list.Vertices, Vertex{p, noUV, color})

	// loop over rest of the points
	for i := Index(0); i < N; i++ {
		dx, dy = dx*rotc-dy*rots, dx*rots+dy*rotc
		p = Vector{center.X + dx, center.Y + dy}
		list.Vertices = append(list.Vertices, Vertex{p, noUV, color})
		list.Indicies = append(list.Indicies, base, base+i+1, base+i+2)
	}

	list.CurrentCommand.Count += Index(len(list.Indicies) - startIndexCount)
}
