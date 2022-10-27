Create a structure named Person:

    type Person struct {
        name, surname string
    }


Create a method - fullName() for Person structure which returns string

    func (receiver Person) fullName() string {
        return (receiver.name + " " + receiver.surname)
    }

Create an instance of structure (class in Python)
    
    func main() {    
        human := Person{"Bulat", "Iakhin"}
        
call method fullName() on human

        fmt.Println(human.fullName())
    }

Output:

    Bulat Iakhin