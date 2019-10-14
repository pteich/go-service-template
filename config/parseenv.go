package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ParseEnv enriches value of a given struct from values of env
func ParseEnv(c interface{}) {

	// get env vars by use of reflection and struct tags
	valueRef := reflect.ValueOf(c)
	confType := valueRef.Elem().Type()

	for i := 0; i < confType.NumField(); i++ {

		field := confType.Field(i)
		tag := field.Tag.Get("env")

		envValue, found := os.LookupEnv(tag)
		if found {
			// set the env value by type, currently only string and bool implemented
			switch field.Type.Name() {
			case "string":
				valueRef.Elem().FieldByName(field.Name).SetString(envValue)
			case "bool":
				valueRef.Elem().FieldByName(field.Name).SetBool(false)
				if strings.ToLower(envValue) == "true" {
					valueRef.Elem().FieldByName(field.Name).SetBool(true)
				}
			case "int":
				value, err := strconv.ParseInt(envValue, 0, 64)
				if err == nil {
					valueRef.Elem().FieldByName(field.Name).SetInt(value)
				}
			default:
				panic(fmt.Sprintf("config env type %s not implemented", field.Type.Name()))
			}
		}

	}

}
