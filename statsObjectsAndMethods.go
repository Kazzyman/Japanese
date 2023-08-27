package main

// **do-this**
/*
.
.
//
var cyclicArrayUserNeedsWorkOn CyclicArrayUserNeedsWorkOn

type CyclicArrayUserNeedsWorkOn struct {
	useThisChar [80]string
	index       int
}

func (ca *CyclicArrayUserNeedsWorkOn) InsertMissedPrompt(UseThisChar string) {
	ca.useThisChar[ca.index] = UseThisChar
	ca.index = (ca.index + 1) % len(ca.useThisChar)
}
.
.
.
.
.
.
*/

/*

Next will come the following statement: var cyclicArrayHits CyclicArrayHits

Here, cyclicArrayHits is a variable that can store a CyclicArrayHits instance.
It is not initialized to refer to a specific CyclicArrayHits object/instance, yet.
It currently only declares allocated memory of type CyclicArrayHits, named cyclicArrayHits

i.e., it declares a variable called cyclicArrayHits of type CyclicArrayHits
The variable is declared but not initialized, so it will have the zero
value for the CyclicArrayHits struct type

... creates memory allocation for a future instance of CyclicArrayHits called
cyclicArrayHits without yet initializing cyclicArrayHits as an actual instance, yet.

... equivalent to:
cyclicArrayHits := CyclicArrayHits{}
.
*/
/*
.
.
.
*/
// Universal hits array
var cyclicArrayHits CyclicArrayHits

type CyclicArrayHits struct {
	jchar       [300]string // a 300 element array of string vars, as the field 'jchar' (various '...'ana chars)
	RightOrOops [300]string // a 300 element array of string vars, as the field 'RightOrOops' (? right or Oops ?)
	index       int         // a field in the struct called 'index' (a single var of type int)
}

func (ca *CyclicArrayHits) InsertRightOrOops(RightOrOops string) { // Please refer also to the lengthy discussion below: //
	ca.RightOrOops[ca.index] = RightOrOops          // Assign the passed string 'value' to the 'data' array at position 'ca.index'
	ca.index = (ca.index + 1) % len(ca.RightOrOops) // Increment 'index' (an integer element) such that it loops-back ...
	// ... to the first position of the 'jchar' array -- having determined its length using: len(ca.jchar).
}
func (ca *CyclicArrayHits) InsertChar(jchar string) { // Please refer also to the lengthy discussion below: //
	ca.jchar[ca.index] = jchar // Assign the passed string 'value' to the 'data' array at position 'ca.index'
	ca.index = (ca.index + 1) % len(ca.jchar)
}

/*
.
.
.
*/
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
The foregoing functions all define 'methods' named 'Insert...' and, in each case, they associate said method with a
'CyclicArray...' type: which happens to be a 'struct' of the form:

type CyclicArray... struct {
	jchar  [300]string // a 300 element array of string vars, as the field 'jchar'
	index int         // a field in the struct called 'index' (a single var of type int)
}

        - - - - - - - - - - - - - - - - - - - - - - - -

In the various method functions defined -- way-up-above-thar: ...

(ca *CyclicArray) is the method receiver; or method receiver declaration.
Here 'ca' is a local variable: and contains a pointer.

Each of the forgoing method functions explicitly create a local var 'ca' as a pointer to an instance of CyclicArray'...'
'ca', thereafter, functions as an instance of the object CyclicArray'...' in statements such as:

ca.index = (some value or expression)

... It (ca) specifies that the 'Insert' method can be called on an instance or value of the CyclicArray'...' type:
and, can operate on that instance or value;
the 'Insert' method is, then, said to be associated with a pointer to a CyclicArray'...' object.
... Thereby, the 'Insert' method can modify the internal state of the CyclicArray'...' object that it is called on:
in this case, the 'jchar' and the 'index' fields of the aforementioned struct.

Insert(value string) is the method's signature. It merely says that the 'Insert' method takes a parameter of type string.

Inside the method body, 'ca' is a local variable that holds the receiver: a pointer to a CyclicArray object.
... This allows access or modification to the fields of the CyclicArray object; in this case: 'jchar' and 'index'.

The line:
ca.jchar[ca.index] = value
... assigns the 'value' parameter to the position in the 'jchar' array located at 'index'.

The line "ca.index = (ca.index + 1) % len(ca.jchar)" updates the index to the next position in a cyclic manner.
... It ensures that the index loops back to 0 when it reaches the end of the array.

In summary, the (ca *CyclicArray) method receiver declaration defines a method named 'Insert' that operates on
... a CyclicArray object. The method updates the array at the current index and cycles the index to the next position,
... (in such a way as to loop back to first position) effectively implementing the behavior of a cyclic array.
*/
