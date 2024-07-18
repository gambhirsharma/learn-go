/*
NOTE:

- String are sequence of UTF-8 characters enclosed in double-quotes
- Numeric are of 8, 16, 32, and 64-bit variants for both signed (int) and unsigned (uint) integers.
- A byte is an alias for uint8.
- A rune is an alias for int32. 
- Floats (or floating-point numbers) are either float32 or float64. 
- Complex numbers are also supported and can be represented as complex128 or complex64.

*/

package main

import "fmt"

func main() {
        const a int32 = 12         // 32-bit integer
	/* User specified types */
        const b float32 = 20.5      // 32-bit float
        var c complex128 = 1 + 4i  // 128-bit complex number
        var d uint16 = 14          // 16-bit unsigned integer

        /* Default types */
        n := 42              // int
        pi := 3.14           // float64
        x, y := true, false  // bool
        z := "Go is awesome" // string

        fmt.Printf("user-specified types:\n %T %T %T %T\n", a, b, c, d)
        fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, x, y, z)
}
