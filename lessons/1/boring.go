package lessons

// @STRUCT
// Structs are a collection of fields. They're useful for grouping data together to form records.

type Person struct {
	id int
	name string
	surname string
	age int
	favouritePizza string
}

// In GO you can call init function in packges instead of main function

func init() {

	pizzas := map[int8]string {1: "Pepperoni", 2: "Hawaiian", 3: "Meat Lovers", 4: "Veggie Lovers", 5: "Cheese Lovers"};

	for i := 0; i < len(pizzas); i++ {
		println(pizzas[int8(i + 1)]);
	}
}
