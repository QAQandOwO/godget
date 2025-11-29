package fieldenum

import "math"

var builtinValues = map[string]any{
	"i":   1i,
	"e":   math.E,
	"pi":  math.Pi,
	"Pi":  math.Pi,
	"phi": math.Phi,
	"Phi": math.Phi,
	"nan": math.NaN(),
	"NaN": math.NaN(),
	"inf": math.Inf(1),
	"Inf": math.Inf(1),
}
