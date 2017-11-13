package main

import "fmt"

/*
Java code:

public class Person {

	private String name;

	public String getName() {
		return this.name;
	}

}

public class Saiyan {

	private Person person;

	public String getName() {
		return this.person.getName();
	}

}
*/

type Person struct {
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}

type Sayan struct {
	*Person
	Power int
}

func main() {
	goku := Sayan{
		Person: &Person{"Goku"},
		Power: 1000,
	}

	fmt.Println(goku.GetName())
}