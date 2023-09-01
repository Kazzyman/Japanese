package main

// **do-this**
/*
In the first of the following statements, cyclicArrayHits is a variable that can store a CyclicArrayHits instance.
The statement does not initialized that var to refer to a specific CyclicArrayHits object/instance.
The statement only declares allocated memory of type CyclicArrayHits, named cyclicArrayHits

i.e., it declares a variable called cyclicArrayHits of type CyclicArrayHits. The variable is declared but is
not initialized: it will initially have the zero-value of/for the CyclicArrayHits struct type.

i.e., it creates a memory allocation for a future instance of CyclicArrayHits to be called cyclicArrayHits without yet
initializing cyclicArrayHits as an actual instance, yet.

var cyclicArrayHits CyclicArrayHits // An empty, uninitialized, var of the CyclicArrayHits type.
... is, therefore, equivalent to:
cyclicArrayHits := CyclicArrayHits{} // Explicitly, an empty, uninitialized, var of the CyclicArrayHits type.
.
.
*/

// The universal hits array:
//
// Declare a memory allocation for storing the most-recent 300 data-points; a cyclic array of structures
// ... this will be used universally among the activities 1-4 to log both Right! and Oops! hits
var cyclicArrayHits CyclicArrayHits

// Create a structure which will populate the cyclic array, and then two methods, one for each of the two data fields
type CyclicArrayHits struct { // Declare a new type, in this case a struct, to be known as a CyclicArrayHits type.
	jchar       [300]string // a 300-element array of string vars, as the field 'jchar' (to hold the various japanese chars as strings)
	RightOrOops [300]string // a 300-element array of string vars, as the field 'RightOrOops' (to hold Right! or Oops! strings)
	index       int         // a single var of type int, as the field 'index'
}

// Create a method for adding elements to the foregoing memory allocation such that it implements the behavior of a cyclic array
func (ca *CyclicArrayHits) InsertRightOrOops(RightOrOops string) { // Please refer also to the lengthy discussion below: //
	ca.RightOrOops[ca.index] = RightOrOops          // Assign the passed string value 'RightOrOops' to the 'data' array at position 'ca.index'
	ca.index = (ca.index + 1) % len(ca.RightOrOops) // Increment 'index' (an integer indexing element) such that it loops-back to ...
	// ... the first position of the 'RightOrOops' array -- having determined its length using: len(ca.RightOrOops).
}

// Create another method for adding elements to the foregoing file of cards such that it implements the behavior of a cyclic array
func (ca *CyclicArrayHits) InsertChar(jchar string) { // Please refer also to the lengthy discussion below: //
	ca.jchar[ca.index] = jchar                // Assign the passed string 'value' to the 'data' array at position 'ca.index'
	ca.index = (ca.index + 1) % len(ca.jchar) // Increment 'index' (an integer indexing element) such that it loops-back to ...
	// ... the first position of the 'jchar' array -- having determined its length using: len(ca.jchar).
}

/*
.
.
*/
// Analogous documentation for the next three statements has already been given above,
// ... while the finer-details of the syntax follow below
var cyclicArrayOfTheJcharsGottenWrong CyclicArrayOfTheJcharsGottenWrong

type CyclicArrayOfTheJcharsGottenWrong struct {
	jchar [300]string
	index int
}

func (ca *CyclicArrayOfTheJcharsGottenWrong) InsertCharsWrong(Jchar string) {
	ca.jchar[ca.index] = Jchar
	ca.index = (ca.index + 1) % len(ca.jchar)
}

//
//
/*
The foregoing three functions all define 'methods' named 'Insert...' and, in each case, they associate said method with a
'CyclicArray...' type

        - - - - - - - - - - - - - - - - - - - - - - - -

In the various method functions defined -- way-up-above-thar: ...

(ca *CyclicArray...) is the method receiver; or method receiver declaration.
Here 'ca' is a local variable: and contains a pointer to a type (in this case a user-defined struct).

Each of the forgoing method functions explicitly create a local var 'ca' as a pointer to an instance of CyclicArray'...'
'ca', thereafter, functions as an instance of the object CyclicArray'...' in statements such as:

ca.index = (some value or expression)

... It (ca) specifies that the 'Insert' method can be called on an instance or value of the CyclicArray'...' type:
and, can operate on that instance or value;
the 'Insert' method is, then, said to be associated with a pointer to a CyclicArray'...' object.
... By use of a pointer (instead of a value), the 'Insert...' method can modify the internal state of the
CyclicArray'...' object that it is called on: in one case, the 'jchar' and the 'index' fields of the user-defined
struct within said object.

Insert...(value string) is the method's signature. It merely says that the 'Insert...' method takes a parameter of type string.

Inside the method body, 'ca' is a local variable that holds the receiver: a pointer to a CyclicArray object.
... This allows access or modification to the fields of the CyclicArray object; in one case: 'jchar' and 'index'.

The line:
ca.jchar[ca.index] = Jchar
... assigns the 'value' parameter Jchar to the position in the 'jchar' array located at 'index'.

The line "ca.index = (ca.index + 1) % len(ca.jchar)" updates the index to the next position in a cyclic manner.
... It ensures that the index loops back to 0 when it reaches the end of the array.

In summary, the (ca *CyclicArray...) method receiver declaration defines a method named 'Insert...' that operates on
... a CyclicArray... object. The method updates the array at the current index and cycles the index to the next position,
... (in such a way as to loop back to first position) effectively implementing the behavior of a cyclic array.
*/
