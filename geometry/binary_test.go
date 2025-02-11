package geometry

import (
	"encoding"
	"encoding/hex"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/twpayne/go-geos"
)

var (
	_ encoding.BinaryMarshaler   = &Geometry{}
	_ encoding.BinaryUnmarshaler = &Geometry{}
)

func TestBinary(t *testing.T) {
	for i, tc := range []struct {
		geom      *Geometry
		binaryStr string
	}{
		{
			geom:      NewGeometry(geos.NewPoint([]float64{1, 2})),
			binaryStr: "0101000000000000000000f03f0000000000000040",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actualBinary, err := tc.geom.MarshalBinary()
			require.NoError(t, err)
			assert.Equal(t, tc.binaryStr, hex.EncodeToString(actualBinary))

			var geom Geometry
			binary, err := hex.DecodeString(tc.binaryStr)
			require.NoError(t, err)
			require.NoError(t, geom.UnmarshalBinary(binary))
			assert.True(t, tc.geom.Equals(geom.Geom))
		})
	}
}
