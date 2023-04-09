package lessons1

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

func printExampleMap() {
	example := map[string]uint8 {"ax0": 00000001, "cx0": 00000010, "dx0": 00000011, "ex0": 00000100}
	iterators := []string {"ax0", "cx0", "dx0", "ex0"}

	for i := 0; i < len(example); i++ {
		println(example[iterators[i]])
	}
}






