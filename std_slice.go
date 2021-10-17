package constvar

type Bools []bool

func (c Bools) Val() []bool { return []bool(c) }

type Strings []string

func (c Strings) Val() []string { return []string(c) }

type Ints []int

func (c Ints) Val() []int { return []int(c) }

type Int8s []int8

func (c Int8s) Val() []int8 { return []int8(c) }

type Int16s []int16

func (c Int16s) Val() []int16 { return []int16(c) }

type Int32s []int32

func (c Int32s) Val() []int32 { return []int32(c) }

type Int64s []int64

func (c Int64s) Val() []int64 { return []int64(c) }

type Uints []uint

func (c Uints) Val() []uint { return []uint(c) }

type Uint8s []uint8

func (c Uint8s) Val() []uint8 { return []uint8(c) }

type Uint16s []uint16

func (c Uint16s) Val() []uint16 { return []uint16(c) }

type Uint32s []uint32

func (c Uint32s) Val() []uint32 { return []uint32(c) }

type Uint64s []uint64

func (c Uint64s) Val() []uint64 { return []uint64(c) }

type Uintptrs []uintptr

func (c Uintptrs) Val() []uintptr { return []uintptr(c) }

type Float32s []float32

func (c Float32s) Val() []float32 { return []float32(c) }

type Float64s []float64

func (c Float64s) Val() []float64 { return []float64(c) }

type Complex64s []complex64

func (c Complex64s) Val() []complex64 { return []complex64(c) }

type Complex128s []complex128

func (c Complex128s) Val() []complex128 { return []complex128(c) }

type Bytes = Uint8s

type Runes = Int32s

type Errors []error

func (c Errors) Val() []error { return []error(c) }
