// This file is automatically generated. DO NOT EDIT.

package geos

// #include "geos.h"
import "C"

// Area returns g's area.
func (g *Geom) Area() float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	var area float64
	if C.GEOSArea_r(g.context.handle, g.geom, (*C.double)(&area)) == 0 {
		panic(g.context.err)
	}
	return area
}

func (g *Geom) Buffer(width float64, quadsegs int) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSBuffer_r(g.context.handle, g.geom, C.double(width), C.int(quadsegs)), nil)
}

// BufferWithStyle returns a buffer using the provided style parameters.
func (g *Geom) BufferWithStyle(width float64, quadsegs int, endCapStyle BufCapStyle, joinStyle BufJoinStyle, mitreLimit float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSBufferWithStyle_r(g.context.handle, g.geom, C.double(width), C.int(quadsegs), C.int(endCapStyle), C.int(joinStyle), C.double(mitreLimit)), nil)
}

// Clone returns a clone of g.
func (g *Geom) Clone() *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSGeom_clone_r(g.context.handle, g.geom), nil)
}

func (g *Geom) ConcaveHull(ratio float64, allowHoles uint) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSConcaveHull_r(g.context.handle, g.geom, C.double(ratio), C.unsigned(allowHoles)), nil)
}

// ConvexHull returns g's convex hull.
func (g *Geom) ConvexHull() *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSConvexHull_r(g.context.handle, g.geom), nil)
}

// Contains returns true if g contains other.
func (g *Geom) Contains(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSContains_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// CoveredBy returns true if g is covered by other.
func (g *Geom) CoveredBy(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSCoveredBy_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// Covers returns true if g covers other.
func (g *Geom) Covers(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSCovers_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// Crosses returns true if g crosses other.
func (g *Geom) Crosses(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSCrosses_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

func (g *Geom) Densify(tolerance float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSDensify_r(g.context.handle, g.geom, C.double(tolerance)), nil)
}

// Difference returns the difference between g and other.
func (g *Geom) Difference(other *Geom) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	return g.context.newGeom(C.GEOSDifference_r(g.context.handle, g.geom, other.geom), nil)
}

// DifferencePrec returns the difference between g and other.
func (g *Geom) DifferencePrec(other *Geom, gridSize float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	return g.context.newGeom(C.GEOSDifferencePrec_r(g.context.handle, g.geom, other.geom, C.double(gridSize)), nil)
}

// Disjoint returns true if g is disjoint from other.
func (g *Geom) Disjoint(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSDisjoint_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// Distance returns the distance between the closes points on g and other.
func (g *Geom) Distance(other *Geom) float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	var distance float64
	if C.GEOSDistance_r(g.context.handle, g.geom, other.geom, (*C.double)(&distance)) == 0 {
		panic(g.context.err)
	}
	return distance
}

// DistanceIndexed returns the distance between g and other, using the indexed facet distance.
func (g *Geom) DistanceIndexed(other *Geom) float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	var distanceIndexed float64
	if C.GEOSDistanceIndexed_r(g.context.handle, g.geom, other.geom, (*C.double)(&distanceIndexed)) == 0 {
		panic(g.context.err)
	}
	return distanceIndexed
}

// DistanceWithin returns whether the distance between g and other is within the given dist.
func (g *Geom) DistanceWithin(other *Geom, dist float64) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSDistanceWithin_r(g.context.handle, g.geom, other.geom, C.double(dist)) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// Envelope returns the envelope of g.
func (g *Geom) Envelope() *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSEnvelope_r(g.context.handle, g.geom), nil)
}

// Equals returns true if g equals other.
func (g *Geom) Equals(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSEquals_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// EqualsExact returns true if g equals other exactly.
func (g *Geom) EqualsExact(other *Geom, tolerance float64) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSEqualsExact_r(g.context.handle, g.geom, other.geom, C.double(tolerance)) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// FrechetDistance returns the Fréchet distance between g and other.
func (g *Geom) FrechetDistance(other *Geom) float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	var frechetDistance float64
	if C.GEOSFrechetDistance_r(g.context.handle, g.geom, other.geom, (*C.double)(&frechetDistance)) == 0 {
		panic(g.context.err)
	}
	return frechetDistance
}

// FrechetDistanceDensify returns the Fréchet distance between g and other.
func (g *Geom) FrechetDistanceDensify(other *Geom, densifyFrac float64) float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	var frechetDistanceDensify float64
	if C.GEOSFrechetDistanceDensify_r(g.context.handle, g.geom, other.geom, C.double(densifyFrac), (*C.double)(&frechetDistanceDensify)) == 0 {
		panic(g.context.err)
	}
	return frechetDistanceDensify
}

// HausdorffDistance returns the Hausdorff distance between g and other.
func (g *Geom) HausdorffDistance(other *Geom) float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	var hausdorffDistance float64
	if C.GEOSHausdorffDistance_r(g.context.handle, g.geom, other.geom, (*C.double)(&hausdorffDistance)) == 0 {
		panic(g.context.err)
	}
	return hausdorffDistance
}

// HausdorffDistanceDensify returns the Hausdorff distance between g and other.
func (g *Geom) HausdorffDistanceDensify(other *Geom, densifyFrac float64) float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	var hausdorffDistanceDensify float64
	if C.GEOSHausdorffDistanceDensify_r(g.context.handle, g.geom, other.geom, C.double(densifyFrac), (*C.double)(&hausdorffDistanceDensify)) == 0 {
		panic(g.context.err)
	}
	return hausdorffDistanceDensify
}

// Interpolate returns a point distance d from the start of g, which must be a linestring.
func (g *Geom) Interpolate(d float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newGeom(C.GEOSInterpolate_r(g.context.handle, g.geom, C.double(d)), nil)
}

// InterpolateNormalized returns the point that is at proportion from the start.
func (g *Geom) InterpolateNormalized(proportion float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newGeom(C.GEOSInterpolateNormalized_r(g.context.handle, g.geom, C.double(proportion)), nil)
}

func (g *Geom) Intersection(other *Geom) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	return g.context.newGeom(C.GEOSIntersection_r(g.context.handle, g.geom, other.geom), nil)
}

func (g *Geom) IntersectionPrec(other *Geom, gridSize float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	return g.context.newGeom(C.GEOSIntersectionPrec_r(g.context.handle, g.geom, other.geom, C.double(gridSize)), nil)
}

// Intersects returns true if g intersects other.
func (g *Geom) Intersects(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSIntersects_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// IsClosed returns true if g is closed.
func (g *Geom) IsClosed() bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	switch C.GEOSisClosed_r(g.context.handle, g.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// IsEmpty returns true if g is empty.
func (g *Geom) IsEmpty() bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	switch C.GEOSisEmpty_r(g.context.handle, g.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// IsRing returns true if g is a ring.
func (g *Geom) IsRing() bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	switch C.GEOSisRing_r(g.context.handle, g.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// IsSimple returns true if g is simple.
func (g *Geom) IsSimple() bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	switch C.GEOSisSimple_r(g.context.handle, g.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// IsValid returns true if g is valid.
func (g *Geom) IsValid() bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	switch C.GEOSisValid_r(g.context.handle, g.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// Length returns g's length.
func (g *Geom) Length() float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	var length float64
	if C.GEOSLength_r(g.context.handle, g.geom, (*C.double)(&length)) == 0 {
		panic(g.context.err)
	}
	return length
}

// MakeValid repairs an invalid geometry, returning a valid output.
func (g *Geom) MakeValid() *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSMakeValid_r(g.context.handle, g.geom), nil)
}

func (g *Geom) MaximumInscribedCircle(tolerance float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSMaximumInscribedCircle_r(g.context.handle, g.geom, C.double(tolerance)), nil)
}

func (g *Geom) MinimumRotatedRectangle() *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSMinimumRotatedRectangle_r(g.context.handle, g.geom), nil)
}

// MinimumWidth returns a linestring geometry which represents the minimum diameter of g.
func (g *Geom) MinimumWidth() *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSMinimumWidth_r(g.context.handle, g.geom), nil)
}

func (g *Geom) OffsetCurve(width float64, quadsegs int, joinStyle BufJoinStyle, mitreLimit float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSOffsetCurve_r(g.context.handle, g.geom, C.double(width), C.int(quadsegs), C.int(joinStyle), C.double(mitreLimit)), nil)
}

// Overlaps returns true if g overlaps other.
func (g *Geom) Overlaps(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSOverlaps_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

func (g *Geom) SetPrecision(gridSize float64, flags PrecisionRule) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSGeom_setPrecision_r(g.context.handle, g.geom, C.double(gridSize), C.int(flags)), nil)
}

// Simplify returns a simplified geometry.
func (g *Geom) Simplify(tolerance float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSSimplify_r(g.context.handle, g.geom, C.double(tolerance)), nil)
}

func (g *Geom) SymDifference(other *Geom) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	return g.context.newGeom(C.GEOSSymDifference_r(g.context.handle, g.geom, other.geom), nil)
}

// TopologyPreserveSimplify returns a simplified geometry preserving topology.
func (g *Geom) TopologyPreserveSimplify(tolerance float64) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSTopologyPreserveSimplify_r(g.context.handle, g.geom, C.double(tolerance)), nil)
}

// Touches returns true if g touches other.
func (g *Geom) Touches(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSTouches_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// UnaryUnion returns the union of all components of a single geometry.
func (g *Geom) UnaryUnion() *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	return g.context.newNonNilGeom(C.GEOSUnaryUnion_r(g.context.handle, g.geom), nil)
}

// Union returns the union of g and other.
func (g *Geom) Union(other *Geom) *Geom {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	return g.context.newGeom(C.GEOSUnion_r(g.context.handle, g.geom, other.geom), nil)
}

// Within returns true if g is within other.
func (g *Geom) Within(other *Geom) bool {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	if other.context != g.context {
		other.context.Lock()
		defer other.context.Unlock()
	}
	switch C.GEOSWithin_r(g.context.handle, g.geom, other.geom) {
	case 0:
		return false
	case 1:
		return true
	default:
		panic(g.context.err)
	}
}

// X returns g's X coordinate.
func (g *Geom) X() float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	var x float64
	if C.GEOSGeomGetX_r(g.context.handle, g.geom, (*C.double)(&x)) == 0 {
		panic(g.context.err)
	}
	return x
}

// Y returns g's Y coordinate.
func (g *Geom) Y() float64 {
	g.mustNotBeDestroyed()
	g.context.Lock()
	defer g.context.Unlock()
	var y float64
	if C.GEOSGeomGetY_r(g.context.handle, g.geom, (*C.double)(&y)) == 0 {
		panic(g.context.err)
	}
	return y
}
