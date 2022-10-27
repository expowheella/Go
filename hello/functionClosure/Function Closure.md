### Closures
Functions declared inside of functions are closures.

Create any function, for ex.:

    func myFunc() func(s string) string {
        result := func(s string) string {
            str := "Hi, " + s + "!"
            return str
        }
        return result
    }

Here:
* myFunc() - is a name of a new func
* func (s string) - means that the input for myFunc() is another function with a string as an argument.
* string - means that the output of the function myFunc() will be a string

---

    func main() {
        callFunc := myFunc()

        fmt.Print(callFunc("Bulat"))
    }
the output:

    Hi, Bulat!