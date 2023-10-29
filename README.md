# Validations
## Instalations
Use go get for install this pakage
[go get github.com/AgikAp/validations](https://github.com/AgikAp/validations)

## Usage
Use tag 'valid' for usage the validator
```go
package main

import (
	"log"

	"github.com/AgikAp/validations"
)

type Data struct {
	Name       string `valid:"max:20;min:4" json:"name"`
	Email      string `valid:"max:5;min:4;email"`
	NestedData NestedData
}

type NestedData struct {
	Data string `valid:"min:8"`
}

/*
if not declaration of custom message
validator will usage default message for response error
*/
var customMessages = map[string]string{"max": "maksimum value string %v character"}

func main() {
	// Initialization
	validation := validations.NewValidatorsAndCustomMessage(customMessages)
	validation.SetFieldNameTag("json")

	// Validation
	data := Data{Name: "Jhon Doe", Email: "jhon@mail.com", NestedData: NestedData{Data: "op"}}
	msg, err := validation.Valid(data)
	if err != nil {
		log.Print(err)               // error
		log.Print(msg.GetMessages()) // map [string]string
	}
}

```

## Available To Use Validation
### String
max   : use for validation length of string

min   : use for validation length of string

email : use for validation string match with pattern email

### Integer
gt    : use for validation value int not greater than requirement

lt    : use for validation value int not less than requirement
