// +build !linux

package neighbor

// getNeighbors returns ErrNotImplemented on non-Linux platforms.
func getNeighbors(family Family) ([]*Neighbor, error) {
	return nil, ErrNotImplemented
}
