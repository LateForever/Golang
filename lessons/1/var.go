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

