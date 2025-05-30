//Integer Formatting
// %d	Decimal integer	fmt.Printf("%d", 10) → 10
// %b	Binary representation	fmt.Printf("%b", 10) → 1010
// %o	Octal representation	fmt.Printf("%o", 10) → 12
// %x	Lowercase hexadecimal	fmt.Printf("%x", 255) → ff
// %X	Uppercase hexadecimal	fmt.Printf("%X", 255) → FF
//
//Float Formatting
// %f	Decimal point, no exponent	fmt.Printf("%.2f", 10.5) → 10.50
// %e	Scientific notation	fmt.Printf("%e", 1000000.0) → 1.000000e+06
// %E	Scientific notation	fmt.Printf("%E", 1000000.0) → 1.000000E+06
// %g	%e for large exponents, %f otherwise	fmt.Printf("%g", 10.5) → 10.5
// %G	%E for large exponents, %f otherwise	fmt.Printf("%G", 10.5) → 10.5
//
//String Formatting
// %s	String	fmt.Printf("%s", "Hello") → Hello
// %q	Quoted string	fmt.Printf("%q", "Hello") → "Hello"
// %v	Default format	fmt.Printf("%v", "Hello") → Hello
// %#v	Go-syntax representation	fmt.Printf("%#v", "Hello") → "Hello"
//
//Boolean Formatting
// %t	Boolean (true/false)	fmt.Printf("%t", true) → true
//
//Character Formatting
// %c	Character	fmt.Printf("%c", 65) → A
//
//Pointer Formatting
// %p	Pointer address	fmt.Printf("%p", &var) → 0xc000010200
//
//Width and Precision
// %f	Default width	fmt.Printf("%f", 10.5) → 10.500000
// %9f	Width 9, default precision	fmt.Printf("%9f", 10.5) → 10.500000
// %.2f	Default width, precision 2	fmt.Printf("%.2f", 10.5) → 10.50
// %9.2f	Width 9, precision 2	fmt.Printf("%9.2f", 10.5) →     10.50
// %9.f	Width 9, precision 0	fmt.Printf("%9.f", 10.5) →        10
//
//Special Formatting
// %T	Type of the value	fmt.Printf("%T", 10) → int
// %%	Literal percent	fmt.Printf("%%") → %
//
//Padding and Alignment
// %-9s	Left-justified	fmt.Printf("%-9s", "Hello") → Hello
// %9s	Right-justified	fmt.Printf("%9s", "Hello") →     Hello
// %09d	Zero padding	fmt.Printf("%09d", 10) → 000000010
// %+d	Always show sign	fmt.Printf("%+d", 10) → +10
// % d	Space for positive numbers	fmt.Printf("% d", 10) →  10
//
//Complex Number Formatting
// %v	Default format	fmt.Printf("%v", complex(1, 2)) → (1+2i)
// %f	Decimal point	fmt.Printf("%f", complex(1.5, -2.5)) → (1.500000-2.500000i)
//
//Other Useful Placeholders
// %U	Unicode	fmt.Printf("%U", 'A') → U+0041
// %#x	Hex with 0x prefix	fmt.Printf("%#x", 255) → 0xff

package main

import (
	"fmt"
)

func main() {
	fmt.Println("PlaceHolders")
}
