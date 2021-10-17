package constvar

func Bool(c bool) func() bool                   { return func() bool { return c } }
func String(c string) func() string             { return func() string { return c } }
func Int(c int) func() int                      { return func() int { return c } }
func Int8(c int8) func() int8                   { return func() int8 { return c } }
func Int16(c int16) func() int16                { return func() int16 { return c } }
func Int32(c int32) func() int32                { return func() int32 { return c } }
func Int64(c int64) func() int64                { return func() int64 { return c } }
func Uint(c uint) func() uint                   { return func() uint { return c } }
func Uint8(c uint8) func() uint8                { return func() uint8 { return c } }
func Uint16(c uint16) func() uint16             { return func() uint16 { return c } }
func Uint32(c uint32) func() uint32             { return func() uint32 { return c } }
func Uint64(c uint64) func() uint64             { return func() uint64 { return c } }
func Uintptr(c uintptr) func() uintptr          { return func() uintptr { return c } }
func Float32(c float32) func() float32          { return func() float32 { return c } }
func Float64(c float64) func() float64          { return func() float64 { return c } }
func Complex64(c complex64) func() complex64    { return func() complex64 { return c } }
func Complex128(c complex128) func() complex128 { return func() complex128 { return c } }
func Byte(c byte) func() byte                   { return func() byte { return c } }
func Rune(c rune) func() rune                   { return func() rune { return c } }
func Error(c error) func() error                { return func() error { return c } }
