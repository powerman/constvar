package constvar

type Bool bool

func (c Bool) Val() bool { return bool(c) }

type String string

func (c String) Val() string { return string(c) }

type Int int

func (c Int) Val() int { return int(c) }

type Int8 int8

func (c Int8) Val() int8 { return int8(c) }

type Int16 int16

func (c Int16) Val() int16 { return int16(c) }

type Int32 int32

func (c Int32) Val() int32 { return int32(c) }

type Int64 int64

func (c Int64) Val() int64 { return int64(c) }

type Uint uint

func (c Uint) Val() uint { return uint(c) }

type Uint8 uint8

func (c Uint8) Val() uint8 { return uint8(c) }

type Uint16 uint16

func (c Uint16) Val() uint16 { return uint16(c) }

type Uint32 uint32

func (c Uint32) Val() uint32 { return uint32(c) }

type Uint64 uint64

func (c Uint64) Val() uint64 { return uint64(c) }

type Uintptr uintptr

func (c Uintptr) Val() uintptr { return uintptr(c) }

type Float32 float32

func (c Float32) Val() float32 { return float32(c) }

type Float64 float64

func (c Float64) Val() float64 { return float64(c) }

type Complex64 complex64

func (c Complex64) Val() complex64 { return complex64(c) }

type Complex128 complex128

func (c Complex128) Val() complex128 { return complex128(c) }

type Byte = Uint8

type Rune = Int32

type constError struct {
	error
}

func Error(err error) interface{ Val() error } { return constError{error: err} }

func (c constError) Val() error { return c.error }
