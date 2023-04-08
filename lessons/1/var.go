package lessons

// Every go source file belongs to a package
// The above package declaration must be the first line of code in your Go source file.
// All the functions, types, and variables defined in your Go source file become part of the declared package.

const (
	i = 10
	pi = 3.14
	prefix = "GO_"
)

var (
	I int
	Pi float32
	Prefix string
)

func createArray() {
	example := [10]int{};
	memory := [256]uint16{0b0001010100101001, 0b1010101000101001, 0b1110010010001011};

	println(memory);
	println(example);
}

// @SLICE 
// In many situations, the array type is not a good choice -for instance when we don't know how long the array will be when
// we define it. Thus, we need a "dynamic array". This is called slice in Go.

var fslice []int;

func createSlice() {
	slice := []byte{}

	println(slice);
}

// @MAP 
// map behaves like a dictionary in Python. Use the form map[keyType]valueType to define it.

var numbers map[string] int;

func setMap(key string, value int) {
	numbers[key] = value;
}


