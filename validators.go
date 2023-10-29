package validations

import (
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type validators struct {
	tag      string
	messages map[string]string
}

func NewValidatorsAndCustomMessage(messages map[string]string) *validators {
	for key, val := range messages {
		constantMessage[key] = val
	}

	return &validators{messages: constantMessage}
}

func NewValidators() *validators {
	return &validators{messages: constantMessage}
}

func (v *validators) SetFieldNameTag(tag string) {
	v.tag = tag
}

func (v *validators) Valid(object interface{}) (msgError *cerros, err error) {
	var resp = make(map[string]string)
	v.rekursifNestedStruct(object, &resp, msgError, &err)
	if 0 < len(resp) {
		err = errors.New(errorValidations)
		msgError = NewCerros(resp)
	}
	return
}

func (v *validators) rekursifNestedStruct(object interface{}, resp *map[string]string, msgError *cerros, err *error) {
	objValue := reflect.ValueOf(object)
	objType := objValue.Type()
	for i := 0; i < objValue.NumField(); i++ {
		fieldType := objType.Field(i)
		fieldName := fieldType.Name
		if v.tag != "" && fieldType.Tag.Get(v.tag) != "" {
			fieldName = fieldType.Tag.Get(v.tag)
		}

		tag := fieldType.Tag.Get("valid")
		value := objValue.Field(i)

		if value.Kind() == reflect.Struct {
			v.rekursifNestedStruct(value.Interface(), resp, msgError, err)
		} else {
			mainLogicValidation(tag, fieldName, value, resp, err)
		}
	}
}

func mainLogicValidation(tag string, fieldName string, value reflect.Value, resp *map[string]string, err *error) {
	if tag != "" {
		// ambil satu yang pertama
		rule := strings.Split(tag, ";")
		for _, rul := range rule {
			var req string
			// ambil requirement jika ada
			requirement := strings.Split(rul, ":")
			if len(requirement) > 1 {
				rul = strings.TrimSpace(requirement[0])
				req = strings.TrimSpace(requirement[1])
			}

			// validasi
			*err = validator(fieldName, value, rul, req, resp)
		}
	}
}

func validator(fieldName string, value reflect.Value, rule string, req string, messages *map[string]string) (err error) {
	if rule != "" {
		switch rule {
		case "max":
			err = validMax(fieldName, value, req, messages)
			break
		case "min":
			err = validMin(fieldName, value, req, messages)
			break
		case "gt":
			err = validIntegerGreaterThan(fieldName, value, req, messages)
			break
		case "lt":
			err = validIntegerLessThan(fieldName, value, req, messages)
			break
		case "email":
			err = validRegexp(fieldName, value, messages, emailPattern, constantMessage["email"])
			// case "image":
			// 	_ = validImage(fieldName, value, messages)
		}
	}

	return
}

func validMax(fieldName string, value reflect.Value, max string, messages *map[string]string) (err error) {
	val := reflectValueToString(value)
	maxNum, err := strconv.Atoi(max)
	if len(val) > maxNum {
		(*messages)[fieldName] = fmt.Sprintf(constantMessage["max"], maxNum)
		return
	}
	return
}

func validMin(fieldName string, value reflect.Value, min string, messages *map[string]string) (err error) {
	val := reflectValueToString(value)
	minNum, err := strconv.Atoi(min)
	if len(val) < minNum {
		(*messages)[fieldName] = fmt.Sprintf(constantMessage["min"], minNum)
		return
	}
	return
}

func validIntegerGreaterThan(fieldName string, value reflect.Value, requirement string, messages *map[string]string) (err error) {
	num, err := strconv.Atoi(reflectValueToString(value))
	numRequirement, err := strconv.Atoi(requirement)
	if err != nil || num > numRequirement {
		(*messages)[fieldName] = fmt.Sprintf(constantMessage["gt"], requirement)
		return
	}
	return
}

func validIntegerLessThan(fieldName string, value reflect.Value, requirement string, messages *map[string]string) (err error) {
	num, err := strconv.Atoi(reflectValueToString(value))
	numRequirement, err := strconv.Atoi(requirement)
	if err != nil || num < numRequirement {
		(*messages)[fieldName] = fmt.Sprintf(constantMessage["lt"], requirement)
		return
	}
	return
}

func validRegexp(fieldName string, value reflect.Value, messages *map[string]string, regExp string, msgConstant string) (err error) {
	val := reflectValueToString(value)
	valid, err := regexp.MatchString(regExp, val)
	if !valid {
		(*messages)[fieldName] = fmt.Sprintf(msgConstant)
		return
	}
	return
}

func validImage(fieldName string, value reflect.Value, messages *map[string]string) (err error) {
	extensionList := strings.Split(imageExtension, ",")
	valid := false
	for _, extension := range extensionList {
		err = validationImageByType(value.Interface().(multipart.File), extension, &valid)
		if valid {
			break
		}
	}

	if !valid {
		(*messages)[fieldName] = fmt.Sprintf(constantMessage["image"])
	}

	return
}

func validationImageByType(file multipart.File, requirement string, valid *bool) (err error) {
	switch requirement {
	case "png":
		err = validationIsImageValid(file, pngImageSignature, valid)
		break
	case "jpg":
		err = validationIsImageValid(file, jpegImageSignature, valid)
		break
	case "webp":
		err = validationIsImageValid(file, webPImageSignature, valid)
		break
	case "tiff":
		err = validationIsImageValid(file, tiffImageSignature, valid)
		break
	case "svg":
		err = validationIsImageValid(file, svgImageSignature, valid)
		break
	}

	return
}

func validationIsImageValid(file multipart.File, signature []byte, valid *bool) (err error) {
	_, err = file.Read(signature)
	if err == nil {
		*valid = false
	}

	return
}

func reflectValueToString(value reflect.Value) string {
	return fmt.Sprintf("%v", value.Interface())
}

// func (v *validators) ValidByTagJson(object interface{})
