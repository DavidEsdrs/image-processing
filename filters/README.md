# filters

This package has the actual execution of all the filters, from cropping to blurring.

This package follows the `command` design pattern, i.e, each filter (command)
contains the computation logic in the Execute method for the filter that 
furthermore will be invoked by the `processor`.

Each filter implements the following interface (that is located in processor.go):
```go
type Command interface {
	Execute(*[][]color.Color) error
}
```