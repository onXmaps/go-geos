package geometry

import (
	"testing"

	"github.com/stretchr/testify/require"

	geos "github.com/twpayne/go-geos"
)

func mustNewGeometryFromWKT(t *testing.T, wkt string) *Geometry {
	t.Helper()
	geom, err := geos.NewGeomFromWKT(wkt)
	require.NoError(t, err)
	return &Geometry{Geom: geom}
}
