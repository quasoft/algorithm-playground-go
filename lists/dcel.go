package goalgorithms

type Vertex struct {
	X, Y int
	Edge *HalfEdge
}

type Face struct {
	Edge *HalfEdge
}

type HalfEdge struct {
	Target *Vertex
	Face   *Face
	Twin   *HalfEdge
	Next   *HalfEdge
	Prev   *HalfEdge
}

type DCEL struct {
	Vertices []Vertex
	Faces    []Face
	Edges    []HalfEdge
}
