# GO
## Installation
...

## Syntax
### Package Visability
Small letter for variabels and functions makes the func/var package private, whilst capital letter makes it public and accessible outside the package.

```go
    // Package private
    var special string

    // Public, visable outside the package
    var Special string


    // "Package Private" Only visable from package main
    func test() {

    }

    // "Public" Visable outside the main Package
    func Test() {

    }

```

### IF and Switches

```go
	if 1 == 1 {
		log.Println("1 = 1")
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	switch cat {
	case "cat":
		log.Println("Meow!")
	case "dog":
		log.Println("Bark!")
	default:
		log.Println("*crickets*")

	}
```

### Loops
```go

	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "cat", "rihno"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	animalMap := make(map[string]string)
	animalMap["dog"] = "FIDO"
	animalMap["cat"] = "FIDOx2"

	for animalType, animal := range animalMap {
		log.Println(animalType, animal)
	}


```