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

// Orthotope represents an orthotope in N = len(Lengths) dimentions with side lengths
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

// BuildRandom places a bridge at an unoccupied location.
func (o *Orthotope) BuildRandom() error {

	if len(o.nonBridges) == 0 {
		return fmt.Errorf("no more unocuppied space to build: %w", ErrInternalState)
	}

	// Select random unoccupied location
	var nb string
	for nb = range o.nonBridges {
		break
	}

	_, ok := o.bridges[nb]
	if ok {
		return fmt.Errorf("location %v in built locations: %w", nb, ErrInternalState)
	}

	delete(o.nonBridges, nb)
	o.bridges[nb] = true

	return nil
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

// BridgeComplete returns true if there is an orthoginally connected path
// from 0 to o.Lengths[0]-1 along the 1st dimention.
func (o *Orthotope) BridgeComplete() bool {

	// TODO

	return false
}

func (o *Orthotope) inBound(locs ...int) bool {

	for i, loc := range locs {
		length := o.Lengths[i]
		if loc < 0 || loc > length {
			return false
		}
	}

	return true
}

func key(locs ...int) string {

	var locKey string
	for _, loc := range locs {
		locKey += fmt.Sprintf("%d-", loc)
	}

	locKey = locKey[:len(locKey)-1]
	return locKey
}

func locations(key string) ([]int, error) {

	locStings := strings.Split(key, "-")

	var locs []int
	for _, s := range locStings {
		loc, err := strconv.Atoi(s)
		if err != nil {
			return locs, fmt.Errorf("invalid key format %q, must be int", s)
		}
		locs = append(locs, loc)
	}

	return locs, nil
}
