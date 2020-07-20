package pointers

import "fmt"

// resources:
// - https://tour.golang.org/moretypes/1

// P undocummented.
type P struct {
	A int
}

// Truth: there is no such as thing as "pass by reference" in golang. Everything in golang is a value.
// reference type: a variable contains a reference to another value.
// In contrasts, in c++ a reference is: "A reference is a type of C++ variable that acts as an alias to another object or value."

// #include <stdio.h>

// int main() {
//         int a = 10;
//         int &b = a;
//         int &c = b;

//         printf("%p %p %p\n", &a, &b, &c); // 0x7ffe114f0b14 0x7ffe114f0b14 0x7ffe114f0b14
//         return 0;
// }

// It is not possible to create a Go program where two variables share the same storage location in memory.
// It is possible to create two variables whose contents point to the same storage location,
// but that is not the same thing as two variables who share the same storage location.

// PMem prints the memory address of *r
func PMem(r *P) {
	// r is a copy of the memory address of the argument of type *P. In this example, r is a "temporary"
	// variable created for the scope of this function that takes the value of "a" (in main.go).
	// There is no such thing as reference variables in go. Having the following:
	// a := 42
	// b := &a
	// yields:
	// fmt.Println(&a) // 0x00000012
	// fmt.Println(&b) // 0x00000020
	// b is a variable that points to the memory address of a. So B and A point to the same position in memory.
	// this implies that modifying B, it modifies A as well.
	// Understanding this premise and having a look at this example we see the same behavior but instead of a and b we have r (arg)
	fmt.Printf("%p\n", &r)
	fmt.Printf("%p\n", r)

	// that translates that modifying r should modify the caller:
	r.A = 21
}
