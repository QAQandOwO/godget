package enum

import (
	"errors"
	"testing"
)

type level string

var (
	infoLevel  = New[level]("info", WithValue[level]("INFO"))
	warnLevel  = New[level]("warn", WithNumber(2))
	errorLevel = New[level]("error", WithIgnoreCase(true))
	debugLevel = New[level]("debug", WithNumber(5))
)

func TestGetEnumByName(t *testing.T) {
	tests := []struct {
		name    string
		want    Enum[level]
		existed bool
	}{
		0: {
			name:    "info",
			want:    infoLevel,
			existed: true,
		},
		1: {
			name:    "warn",
			want:    warnLevel,
			existed: true,
		},
		2: {
			name:    "error",
			want:    errorLevel,
			existed: true,
		},
		3: {
			name:    "unknown",
			want:    Enum[level]{},
			existed: false,
		},
		4: {
			name:    "INFO",
			want:    Enum[level]{},
			existed: false,
		},
		5: {
			name:    "ERROR",
			want:    errorLevel,
			existed: true,
		},
	}

	for i, test := range tests {
		got, existed := GetEnumByName[level](test.name)
		if existed != test.existed {
			if !existed {
				t.Errorf("[%d]: got: not existed, want: %v", i, test.want)
			} else {
				t.Errorf("[%d]: got: %v, want: not existed", i, test.want)
			}
		} else if got != test.want {
			t.Errorf("[%d]: got: %#v, want: %#v", i, got, test.want)
		}
	}
}

func TestGetEnums(t *testing.T) {
	t.Run("existed type", func(t *testing.T) {
		want := map[Enum[level]]struct{}{infoLevel: {}, errorLevel: {}, warnLevel: {}, debugLevel: {}}

		if got, existed := GetEnums[level](); !existed {
			t.Errorf("got: not existed, want: %v", want)
		} else if len(got) != len(want) {
			t.Errorf("got: %v, want: %v", got, want)
		} else {
			for _, enum := range got {
				if _, ok := want[enum]; !ok {
					t.Errorf("got: %v, want: %v", got, want)
					break
				}
			}
		}
	})

	t.Run("not existed type", func(t *testing.T) {
		if got, existed := GetEnums[struct{}](); existed {
			t.Errorf("got: %#v, want: not existed", got)
		}
	})
}

func TestGetEnumNames(t *testing.T) {
	t.Run("existed type", func(t *testing.T) {
		want := map[string]struct{}{"info": {}, "error": {}, "warn": {}, "debug": {}}

		if got, existed := GetEnumNames[level](); !existed {
			t.Errorf("got: not existed, want: %v", want)
		} else if len(got) != len(want) {
			t.Errorf("got: %v, want: %v", got, want)
		} else {
			for _, name := range got {
				if _, ok := want[name]; !ok {
					t.Errorf("got: %v, want: %v", got, want)
					break
				}
			}
		}
	})

	t.Run("not existed type", func(t *testing.T) {
		if got, existed := GetEnumNames[struct{}](); existed {
			t.Errorf("got: %#v, want: not existed", got)
		}
	})
}

func TestGetEnumCount(t *testing.T) {
	t.Run("existed type", func(t *testing.T) {
		want := 4

		if got := GetEnumCount[level](); got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("not existed type", func(t *testing.T) {
		if got := GetEnumCount[struct{}](); got != 0 {
			t.Errorf("got: %v, want: not existed", got)
		}
	})
}

func TestEnum_IsValid(t *testing.T) {
	tests := []struct {
		enum   Enum[level]
		wanted bool
	}{
		0: {
			enum:   infoLevel,
			wanted: true,
		},
		1: {
			enum:   warnLevel,
			wanted: true,
		},
		2: {
			enum:   errorLevel,
			wanted: true,
		},
		3: {
			enum:   debugLevel,
			wanted: true,
		},
		4: {
			enum:   Enum[level]{},
			wanted: false,
		},
	}
	for i, test := range tests {
		if got := test.enum.IsValid(); got != test.wanted {
			t.Errorf("[%d]: got %v, want %v", i, got, test.wanted)
		}
	}
}

func TestEnum_Name(t *testing.T) {
	tests := []struct {
		enum   Enum[level]
		wanted string
	}{
		0: {
			enum:   infoLevel,
			wanted: "info",
		},
		1: {
			enum:   warnLevel,
			wanted: "warn",
		},
		2: {
			enum:   errorLevel,
			wanted: "error",
		},
		3: {
			enum:   debugLevel,
			wanted: "debug",
		},
		4: {
			enum:   Enum[level]{},
			wanted: "",
		},
	}

	for i, test := range tests {
		if got := test.enum.Name(); got != test.wanted {
			t.Errorf("[%d]: got %v, want %v", i, got, test.wanted)
		}
	}
}

func TestEnum_Value(t *testing.T) {
	tests := []struct {
		enum   Enum[level]
		wanted level
	}{
		0: {
			enum:   infoLevel,
			wanted: "INFO",
		},
		1: {
			enum:   warnLevel,
			wanted: "",
		},
		2: {
			enum:   errorLevel,
			wanted: "",
		},
		3: {
			enum:   debugLevel,
			wanted: "",
		},
		4: {
			enum:   Enum[level]{},
			wanted: "",
		},
	}

	for i, test := range tests {
		if got := test.enum.Value(); got != test.wanted {
			t.Errorf("[%d]: got %v, want %v", i, got, test.wanted)
		}
	}
}

func TestEnum_Equal(t *testing.T) {
	testLevel := infoLevel
	tests := []struct {
		enum   Enum[level]
		other  Enum[level]
		wanted bool
	}{
		0: {
			enum:   infoLevel,
			other:  infoLevel,
			wanted: true,
		},
		1: {
			enum:   infoLevel,
			other:  testLevel,
			wanted: true,
		},
		2: {
			enum:   infoLevel,
			other:  errorLevel,
			wanted: true,
		},
		3: {
			enum:   infoLevel,
			other:  warnLevel,
			wanted: false,
		},
		4: {
			enum:   infoLevel,
			other:  Enum[level]{},
			wanted: false,
		},
		5: {
			enum:   Enum[level]{},
			other:  Enum[level]{},
			wanted: true,
		},
	}

	for i, test := range tests {
		if got := test.enum.Equal(test.other); got != test.wanted {
			t.Errorf("[%d]: got %v, want %v", i, got, test.wanted)
		}
	}
}

func TestEnum_Compare(t *testing.T) {
	testLevel := warnLevel
	tests := []struct {
		enum   Enum[level]
		other  Enum[level]
		wanted int
	}{
		0: {
			enum:   warnLevel,
			other:  warnLevel,
			wanted: 0,
		},
		1: {
			enum:   warnLevel,
			other:  testLevel,
			wanted: 0,
		},
		2: {
			enum:   warnLevel,
			other:  debugLevel,
			wanted: -1,
		},
		3: {
			enum:   warnLevel,
			other:  infoLevel,
			wanted: 1,
		},
		4: {
			enum:   warnLevel,
			other:  Enum[level]{},
			wanted: 1,
		},
		5: {
			enum:   Enum[level]{},
			other:  warnLevel,
			wanted: -1,
		},
		6: {
			enum:   Enum[level]{},
			other:  Enum[level]{},
			wanted: 0,
		},
	}

	for i, test := range tests {
		if got := test.enum.Compare(test.other); got != test.wanted {
			t.Errorf("[%d]: got %v, want %v", i, got, test.wanted)
		}
	}
}

func TestEnum_MarshalText(t *testing.T) {
	tests := []struct {
		enum   Enum[level]
		wanted string
		hasErr bool
	}{
		0: {
			enum:   infoLevel,
			wanted: "info",
		},
		1: {
			enum:   warnLevel,
			wanted: "warn",
		},
		2: {
			enum:   errorLevel,
			wanted: "error",
		},
		3: {
			enum:   debugLevel,
			wanted: "debug",
		},
		4: {
			enum:   Enum[level]{},
			hasErr: true,
		},
	}

	for i, test := range tests {
		got, err := test.enum.MarshalText()
		if !test.hasErr {
			if err != nil {
				t.Errorf("[%d]: got error %v, want no error", i, err)
			} else if text := string(got); text != test.wanted {
				t.Errorf("[%d]: got %v, want %v", i, text, test.wanted)
			}
		} else {
			wantErr := new(InvalidError)
			if !errors.As(err, &wantErr) {
				t.Errorf("[%d]: got %v, want %v", i, err, wantErr)
			}
		}
	}
}

func TestEnum_UnmarshalText(t *testing.T) {
	tests := []struct {
		text   string
		wanted Enum[level]
		hasErr bool
	}{
		0: {
			text:   "info",
			wanted: infoLevel,
		},
		1: {
			text:   "error",
			wanted: errorLevel,
		},
		2: {
			text:   "ERROR",
			wanted: errorLevel,
		},
		3: {
			text:   "invalid",
			hasErr: true,
		},
	}

	for i, test := range tests {
		var e Enum[level]
		err := e.UnmarshalText([]byte(test.text))
		if !test.hasErr {
			if err != nil {
				t.Errorf("[%d]: got error %v, want no error", i, err)
			} else if !e.Equal(test.wanted) {
				t.Errorf("[%d]: got %v, want %v", i, e, test.wanted)
			}
		} else {
			wantErr := new(NameNotExistedError)
			if !errors.As(err, &wantErr) {
				t.Errorf("[%d]: got %v, want %v", i, err, wantErr)
			}
		}
	}
}

func TestEnum_MarshalJSON(t *testing.T) {
	tests := []struct {
		enum   Enum[level]
		wanted string
		hasErr bool
	}{
		0: {
			enum:   infoLevel,
			wanted: `"info"`,
		},
		1: {
			enum:   warnLevel,
			wanted: `"warn"`,
		},
		2: {
			enum:   errorLevel,
			wanted: `"error"`,
		},
		3: {
			enum:   debugLevel,
			wanted: `"debug"`,
		},
		4: {
			enum:   Enum[level]{},
			hasErr: true,
		},
	}

	for i, test := range tests {
		got, err := test.enum.MarshalJSON()
		if !test.hasErr {
			if err != nil {
				t.Errorf("[%d]: got error %v, want no error", i, err)
			} else if text := string(got); text != test.wanted {
				t.Errorf("[%d]: got %v, want %v", i, text, test.wanted)
			}
		} else {
			wantErr := new(InvalidError)
			if !errors.As(err, &wantErr) {
				t.Errorf("[%d]: got %v, want %v", i, err, wantErr)
			}
		}
	}
}

func TestEnum_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		text   string
		wanted Enum[level]
		hasErr bool
	}{
		0: {
			text:   `"info"`,
			wanted: infoLevel,
		},
		1: {
			text:   `"error"`,
			wanted: errorLevel,
		},
		2: {
			text:   `"ERROR"`,
			wanted: errorLevel,
		},
		3: {
			text:   `"invalid"`,
			hasErr: true,
		},
	}
	for i, test := range tests {
		var e Enum[level]
		err := e.UnmarshalJSON([]byte(test.text))
		if !test.hasErr {
			if err != nil {
				t.Errorf("[%d]: got error %v, want no error", i, err)
			} else if !e.Equal(test.wanted) {
				t.Errorf("[%d]: got %v, want %v", i, e, test.wanted)
			}
		} else {
			wantErr := new(NameNotExistedError)
			if !errors.As(err, &wantErr) {
				t.Errorf("[%d]: got %v, want %v", i, err, wantErr)
			}
		}
	}
}
