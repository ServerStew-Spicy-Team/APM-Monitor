package tools

import (
	"github.com/spf13/viper"
	"reflect"
	"strconv"
	"time"
)

func StructToMap(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		fieldName := objType.Field(i).Name
		fieldVal := objVal.Field(i).Interface()
		data[fieldName] = fieldVal
	}
	return data
}

func NewTimeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

func ScientificToNumber(scientific float64) string {
	return strconv.FormatFloat(scientific, 'f', -1, 64)
}

func GetIP() string {
	return viper.GetString("ip")
}
