# Greetings in GO
This package helps you to get nice greetings

## Instalations
run this command to install the package
```
go get -u github.com/David29-ra/greetings_go
```

## Example

```
package main

import (
	"fmt"
	"log"

	"github.com/David29-ra/greetings_go"
)

func main() {
	log.SetPrefix("greetings: ")

	message, err := greetings.Hello("Monita")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
```