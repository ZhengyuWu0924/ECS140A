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

// TODO: implement a type that satisfies the DisjointSet interface.
type Set struct {
	mValueParent map[int]int
	mValueRank   map[int]int
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	s := Set{make(map[int]int), make(map[int]int)}
	return s
}

func (s Set) FindSet(a int) int {
	if elem, ok := s.mValueParent[a]; ok == true {
		if elem != a {
			s.mValueParent[a] = s.FindSet(elem)
			return s.mValueParent[a]
		}
		return a
	}
	s.mValueParent[a] = a
	s.mValueRank[a] = 0
	return a
}

func (s Set) UnionSet(a int, b int) int {
	a = s.FindSet(a)
	b = s.FindSet(b)
	if a == b {
		return a
	}
	if s.mValueRank[a] < s.mValueRank[b] {
		a, b = b, a
	}
	s.mValueParent[b] = a
	if s.mValueRank[a] == s.mValueRank[b] {
		s.mValueRank[a]++
	}
	return a
}
