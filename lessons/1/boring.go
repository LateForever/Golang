package lessons

// In GO you can call init function in packges instead of main function

func init() {

	pizzas := map[int8]string {1: "Pepperoni", 2: "Hawaiian", 3: "Meat Lovers", 4: "Veggie Lovers", 5: "Cheese Lovers"};

	for i := 0; i < len(pizzas); i++ {
		println(pizzas[int8(i + 1)]);
	}

}
