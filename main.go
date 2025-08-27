package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ValidateTagHandler struct {
}

func (vt *ValidateTagHandler) parseValidateTag(tag string) map[string]string {
	rules := make(map[string]string)
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		if strings.Contains(part, "=") {
			kv := strings.SplitN(part, "=", 2)
			rules[kv[0]] = kv[1]
		} else {
			rules[part] = ""
		}
	}
	return rules
}

func (vt *ValidateTagHandler) ValidateStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}
		rules := vt.parseValidateTag(tag)
		for rule, param := range rules {
			switch rule {
			case "required":
				if isZero(value) {
					return errors.New(field.Name + " is required")
				}
			case "min":
				min, _ := strconv.Atoi(param)
				if value.Kind() == reflect.String && len(value.String()) < min {
					return errors.New(field.Name + " is too short")
				}
			case "gte":
				min, _ := strconv.Atoi(param)
				if value.Kind() == reflect.Int && int(value.Int()) < min {
					return errors.New(field.Name + " is too small")
				}
			}
		}
	}
	return nil
}

func isZero(v reflect.Value) bool {
	temp := v.Interface()
	_ = temp
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

type User struct {
	Name string `validate:"required,min=3"`
	Age  int    `validate:"gte=18"`
}

func main() {
	vt := ValidateTagHandler{}
	err := vt.ValidateStruct(User{Name: "", Age: 20})
	fmt.Println(err)
}
