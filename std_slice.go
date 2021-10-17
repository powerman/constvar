package constvar

func Bools(c []bool) func() []bool                   { return func() []bool { return c } }
func Strings(c []string) func() []string             { return func() []string { return c } }
func Ints(c []int) func() []int                      { return func() []int { return c } }
func Int8s(c []int8) func() []int8                   { return func() []int8 { return c } }
func Int16s(c []int16) func() []int16                { return func() []int16 { return c } }
func Int32s(c []int32) func() []int32                { return func() []int32 { return c } }
func Int64s(c []int64) func() []int64                { return func() []int64 { return c } }
func Uints(c []uint) func() []uint                   { return func() []uint { return c } }
func Uint8s(c []uint8) func() []uint8                { return func() []uint8 { return c } }
func Uint16s(c []uint16) func() []uint16             { return func() []uint16 { return c } }
func Uint32s(c []uint32) func() []uint32             { return func() []uint32 { return c } }
func Uint64s(c []uint64) func() []uint64             { return func() []uint64 { return c } }
func Uintptrs(c []uintptr) func() []uintptr          { return func() []uintptr { return c } }
func Float32s(c []float32) func() []float32          { return func() []float32 { return c } }
func Float64s(c []float64) func() []float64          { return func() []float64 { return c } }
func Complex64s(c []complex64) func() []complex64    { return func() []complex64 { return c } }
func Complex128s(c []complex128) func() []complex128 { return func() []complex128 { return c } }
func Bytes(c []byte) func() []byte                   { return func() []byte { return c } }
func Runes(c []rune) func() []rune                   { return func() []rune { return c } }
func Errors(c []error) func() []error                { return func() []error { return c } }
