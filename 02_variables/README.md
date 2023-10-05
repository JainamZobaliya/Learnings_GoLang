# Variables & Types in GoLang

## Variables
The `var` statement declares a list of variables; as in function argument lists, the type is last.<br />
A `var` statement can be at package or function level.<br />
A `var` declaration can include initializers, one per variable.<br />
If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

### Short variable declarations
1. Inside a function, the `:=` short assignment statement can be used in place of a var declaration with implicit type.
2. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

## Types
Go's basic types are:
1. `bool`
2. `string`
3. `int`  `int8`  `int16`  `int32`  `int64` <br /> `uint` `uint8` `uint16` `uint32` `uint64` `uintptr`
4. `byte` // alias for uint8
5. `rune` // alias for int32 // represents a Unicode code point
6. `float32` `float64`
7. `complex64` `complex128`


## Default Values
Variables declared without an explicit initial value are given their zero value.

The zero value is:
1. `0` for __numeric types__,
2. `false` for the __boolean type__, and
3. `""` (the empty string) for __strings__.