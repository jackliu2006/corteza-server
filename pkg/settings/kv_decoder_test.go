package settings

import (
	"testing"

	"github.com/jmoiron/sqlx/types"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	type (
		subdst struct {
			S string `kv:"s"`
			B bool   `kv:"b"`

			Bar struct {
				Foo string `kv:"foo"`
			} `kv:"bar"`
		}

		withHandler struct{}

		dst struct {
			S string `kv:"s"`
			B bool   `kv:"b"`
			N int    `kv:"n"`

			NoKV string

			WH withHandler

			Ptr *string

			Sub subdst `kv:"sub"`

			Map map[string]string `kv:"sub.map"`
			S2I map[string]int    `kv:"sub.s2i"`
		}
	)

	var (
		ptr = "point-me"

		aux = dst{}
		kv  = KV{
			"s":           types.JSONText(`"string"`),
			"b":           types.JSONText("true"),
			"n":           types.JSONText("42"),
			"sub.s":       types.JSONText(`"string"`),
			"sub.b":       types.JSONText("true"),
			"sub.bar":     nil,
			"sub.bar.foo": types.JSONText(`"foobar"`),

			"noKV": types.JSONText(`"NO-KV-!"`),
			"ptr":  types.JSONText(`"point-me"`),

			"sub.map.foo": types.JSONText(`"foo"`),
			"sub.map.bar": types.JSONText(`"bar"`),
			"sub.map.baz": types.JSONText(`"baz"`),

			"sub.s2i.one": types.JSONText(`1`),
			"sub.s2i.two": types.JSONText(`2`),
		}

		eq = dst{
			S: "string",
			B: true,
			N: 42,

			NoKV: "NO-KV-!",
			Ptr:  &ptr,

			Sub: subdst{
				S: "string",
				B: true,
			},

			Map: map[string]string{
				"foo": "foo",
				"bar": "bar",
				"baz": "baz",
			},

			S2I: map[string]int{
				"one": 1,
				"two": 2,
			},
		}
	)

	// setting this externaly (embedded structs)
	eq.Sub.Bar.Foo = "foobar"

	require.NoError(t, Decode(kv, &aux))
	require.Equal(t, eq, aux)
}
