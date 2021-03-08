package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

//BEGIN_SOLUTION

// DisjointSetImpl satisfies the DisjointSet interface.
// A node is a representative if it points to itself in `parent` map.
// An undefined key in 'parent' should be taken as pointing to itself.
// The size of a class is stored in the representative's `size` map.
type DisjointSetImpl struct {
	parent map[int]int
	size   map[int]int
}

// FindSet returns representative of the class that s belongs to.
func (ds *DisjointSetImpl) FindSet(s int) int {
	// Compares path for all non-representative nodes visited.
	if p, ok := ds.parent[s]; !ok {
		ds.parent[s] = s
		ds.size[s] = 1
		return s
	} else if s != p {
		r := ds.FindSet(p)
		ds.parent[s] = r
		return r
	} else {
		return s
	}
}

// UnionSet merges the classes represented by s and t, using Union by size
// and returns the new class representative.
func (ds *DisjointSetImpl) UnionSet(s, t int) int {
	s, t = ds.FindSet(s), ds.FindSet(t)
	if s == t {
		return s
	}
	sizeS, sizeT := ds.size[s], ds.size[t]
	if sizeS < sizeT {
		s, t = t, s
	}
	ds.parent[t] = s
	ds.size[s] = sizeS + sizeT
	return s
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	return &DisjointSetImpl{
		parent: make(map[int]int),
		size:   make(map[int]int),
	}
}
