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

// @Return created person
func createPerson(id int, name string, surname string, age int, favouritePizza string) Person {
	return Person{id, name, surname, age, favouritePizza};
}

// @Return fav pizza of person
func getFavouritePizza(person Person) string {
	return person.favouritePizza;
}

func init() {

	pizzas := map[int8]string {1: "Pepperoni", 2: "Hawaiian", 3: "Meat Lovers", 4: "Veggie Lovers", 5: "Cheese Lovers"};

	for i := 0; i < len(pizzas); i++ {
		println(pizzas[int8(i + 1)]);
	}

	// clients 
	var Alberto Person;
	var Marcello Person;
	var Juan Person;

	Alberto = createPerson(1, "Alberto", "Garcia", 21, pizzas[1]);
	Marcello = createPerson(2, "Marcello", "Garcia", 21, pizzas[2]);
	Juan = createPerson(3, "Juan", "Garcia", 21, pizzas[3]);

	clients := [] Person {Alberto, Marcello, Juan};

	// Print favourite pizza of each client
	for i := 0; i < len(clients); i++ {
		println(clients[i].favouritePizza);
	}
}
