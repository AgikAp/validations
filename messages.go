package validations

var constantMessage map[string]string = map[string]string{
	"max":   "max value is %v characters",
	"min":   "min value is %v characters",
	"email": "invalid format email",
	"image": "file is not image",
	"gt":    "numbers cannot greater than %v",
	"lt":    "numbers cannot less than %v",
}

var errorValidations = "error validation field"
