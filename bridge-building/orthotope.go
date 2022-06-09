package orth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrOccupied      = errors.New("space occupied by bridge")
	ErrOutOfBounds   = errors.New("out of bounds")
	ErrInternalState = errors.New("bridge piece was destroyed")
)

// Orthotope represents an orthotope in N = len(Lengths) dimensions with side lengths
// n_1 = Lengths[0], n_2 = Lengths[1], ..., n_N = Lengths[N-1]
//
// Invariant:
// - bridges U nonBridges: all integer locations within the orthotope
// - len(bridges U nonBridges) = len(bridges + nonBridges)
type Orthotope struct {
	Lengths    []int
	bridges    map[string]bool
	nonBridges map[string]bool
}

func New(lengths []int) (*Orthotope, error) {

	nonBridges := []string{""}
	for _, length := range lengths {
		var newNonBridges []string
		for _, key := range nonBridges {
			var newKeys []string
			for i := 0; i < length; i++ {
				newKey := fmt.Sprintf("%s-%d", key, i)
				newKeys = append(newKeys, newKey)
			}
			newNonBridges = append(newNonBridges, newKeys...)
		}
		nonBridges = newNonBridges
	}

	if len(nonBridges) == 1 && nonBridges[0] == "" {
		nonBridges = []string{}
	}

	nbs := map[string]bool{}
	for _, nb := range nonBridges {
		// Remove leading "-".
		nb = nb[1:]
		nbs[nb] = true
	}

	o := &Orthotope{
		Lengths:    lengths,
		bridges:    map[string]bool{},
		nonBridges: nbs,
	}
	return o, nil
}

// Build places a bridge at locs even if one already exists.
func (o *Orthotope) Build(locs ...int) error {

	if !o.inBound(locs...) {
		return fmt.Errorf("location %v outside bounds limits %v: %w", locs, o.Lengths, ErrOutOfBounds)
	}

	k := key(locs...)
	o.bridges[k] = true
	delete(o.nonBridges, k)

	return nil
}

// BuildRandom places a bridge at an unoccupied location and returns it as key.
func (o *Orthotope) BuildRandom() ([]int, error) {

	if len(o.nonBridges) == 0 {
		return []int{}, fmt.Errorf("no more unocuppied space to build: %w", ErrInternalState)
	}

	// Select random unoccupied location
	var nb string
	for nb = range o.nonBridges {
		break
	}

	_, ok := o.bridges[nb]
	if ok {
		return []int{}, fmt.Errorf("location %v in built locations: %w", nb, ErrInternalState)
	}

	delete(o.nonBridges, nb)
	o.bridges[nb] = true

	locs, err := locations(nb)
	if err != nil {
		return []int{}, fmt.Errorf("failed to build bridge in %v because: %w", nb, err)
	}

	return locs, nil
}

// Built returns whether the hypercube at locs contains a bridge.
func (o *Orthotope) Built(locs ...int) (bool, error) {

	if !o.inBound(locs...) {
		return false, fmt.Errorf("location %v outside bounds limits %v: %w", locs, o.Lengths, ErrOutOfBounds)
	}

	k := key(locs...)
	b, ok := o.bridges[k]
	if !ok {
		return false, nil
	}

	// Having an unbuilt bridge location is against invariant.
	if !b {
		return false, fmt.Errorf("location %v: %w", locs, ErrInternalState)
	}

	return true, nil
}

// Neighbors returns the orthogonal neighbors of the hypercube at location locs.
func (o *Orthotope) Neighbors(locs ...int) ([][]int, error) {

	var neighbors [][]int
	if !o.inBound(locs...) {
		return neighbors, fmt.Errorf("location %v, %w", locs, ErrOutOfBounds)
	}

	for i := range locs {
		var l []int
		var r []int
		for _, loc := range locs {
			l = append(l, loc)
			r = append(r, loc)
		}
		l[i] -= 1
		r[i] += 1
		neighbors = append(neighbors, l)
		neighbors = append(neighbors, r)
	}

	// Remove neighbors out of bounds
	var inBoundNeighbors [][]int
	for _, neighbor := range neighbors {
		nKey := key(neighbor...)
		_, okB := o.bridges[nKey]
		_, okN := o.nonBridges[nKey]
		in := o.inBound(neighbor...)
		if in && !okB && !okN {
			return neighbors, fmt.Errorf("in bound piece %v not in bridge or nonBridge sets: %w", neighbor, ErrInternalState)
		}
		if !in && (okB || okN) {
			return neighbors, fmt.Errorf("out of bound piece %v in bridge (%v) or nonBridge (%v) sets: %w", neighbor, okB, okN, ErrInternalState)
		}
		if okB && okN {
			return neighbors, fmt.Errorf("piece %v in bridge and nonBridge sets: %w", neighbor, ErrInternalState)
		}
		if okB || okN {
			inBoundNeighbors = append(inBoundNeighbors, neighbor)
		}
	}
	neighbors = inBoundNeighbors

	return neighbors, nil
}

// BridgeComplete returns true if there is an orthogonally connected path
// from 0 to o.Lengths[0]-1 along the 1st dimension.
func (o *Orthotope) BridgeComplete() bool {

	// TODO

	return false
}

func (o *Orthotope) inBound(locs ...int) bool {

	for i, loc := range locs {
		// Location containers higher dimension
		if i >= len(o.Lengths) {
			return false
		}

		length := o.Lengths[i]
		if loc < 0 || loc >= length {
			return false
		}
	}

	return true
}

// key returns ths string representation of locs.
// Example: [1,2,3] -> "1-2-3"
func key(locs ...int) string {

	if len(locs) == 0 {
		return ""
	}

	var locKey string
	for _, loc := range locs {
		locKey += fmt.Sprintf("%d-", loc)
	}

	// Remove traling "-"
	locKey = locKey[:len(locKey)-1]
	return locKey
}

// locations returns the slice representation of key.
// Example: "1-2-3" -> [1,2,3]
func locations(key string) ([]int, error) {

	var locs []int
	if len(key) == 0 {
		return locs, nil
	}

	locStings := strings.Split(key, "-")
	for _, s := range locStings {
		loc, err := strconv.Atoi(s)
		if err != nil {
			return locs, fmt.Errorf("invalid key format %q, must be int", s)
		}
		locs = append(locs, loc)
	}

	return locs, nil
}
