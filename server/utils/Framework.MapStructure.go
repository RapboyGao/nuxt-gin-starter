package utils

import (
	"reflect"
	"strconv"
	"time"

	"github.com/golang-module/carbon"
	"github.com/mitchellh/mapstructure"
)

func ToTimeHookFunc() mapstructure.DecodeHookFunc {
	var float64Example float64 = 1
	var float32Example float32 = 1
	// var int64Example int64 = 1

	return func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
		if to == reflect.TypeOf(time.Time{}) {
			switch from.Kind() {
			case reflect.String:
				parsed := carbon.Parse(data.(string))
				return parsed.ToStdTime(), nil
			case reflect.Float64:
				return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
			case reflect.Float32:
				return time.Unix(0, int64(data.(float32))*int64(time.Millisecond)), nil
			case reflect.Int64:
				return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
			default:
				return data, nil
			}
		} else if to == reflect.TypeOf(float64Example) {
			return strconv.ParseFloat(data.(string), 64)
		} else if to == reflect.TypeOf(float32Example) {
			return strconv.ParseFloat(data.(string), 32)
		} else {
			return data, nil
		}
	}
}

func Decode(input interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:   nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(ToTimeHookFunc()),
		Result:     result,
	})
	if err != nil {
		return err
	}
	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}
