// Package constvar provides safer constant values for global variables.
//
// Some Go types can't be used for constant values, thus we've to use
// global variables instead. But variables may be occasionally modified,
// which usually result in hard-to-spot bugs. This package helps to work
// around.
//
//
// Usage
//
//   // Replace this:
//   var appStart = time.Now()
//   main() {
//       fmt.Println(appStart)
//   }
//
//   // with this:
//   var appStart = constvar.Time(time.Now())
//   main() {
//       fmt.Println(appStart())
//   }
//
//
// Safety
//
// Of course, no one stops you from re-defining global variable, but it's
// much more easier to notice (on code review) usage of constvar package
// outside of defining global vars than usual variable assignment.
//
//
// Side-effects
//
// It's now possible to set to nil global variables defined using
// constvar. This shouldn't be a real issue because this will result in
// run-time panic on accessing it value and thus won't go unnoticed.
//
// Code is more self-documented now. :)
//
//
// Hints
//
// ★ Configuration for golangci-lint
//
//   issues:
//     exclude-rules:
//       - linters: gochecknoglobals
//         source: "constvar[.]"
package constvar
