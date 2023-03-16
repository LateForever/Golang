package pseudo

type Employee struct {
	Id   string
	Name string
	Age  int
}

var employee = Employee{"0", "Bob", 20}
var employee2 = Employee{Name: "Mike", Age: 30}

// Constructor function for Employee
func createEmployee(id string, name string, age int) Employee {
	employee3 := new(Employee)

	employee3.Id = id
	employee3.Name = name
	employee3.Age = age

	return *employee3
}

var employee4 = &Employee{"3", "Jack", 33}

// Functions as methods in classes

type EmployeeMethods interface {
	GetId() string
	SetId(id string)
	GetName() string
	SetName(name string)
	GetAge() int
	SetAge(age int)
}

// @Getter for Name
func (e *Employee) GetName() string {
	return e.Name
}

// @Setter for Name
func (e *Employee) SetName(name string) {
	e.Name = name
}

// @Getter for Age
func (e *Employee) GetAge() int {
	return e.Age
}

// @Setter for Age
func (e *Employee) SetAge(age int) {
	e.Age = age
}
