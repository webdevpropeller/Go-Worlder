package presenter

import (
	"fmt"
	"go_worlder_system/config"
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	"reflect"
)

var (
	URL = config.NewURL()
)

func convertOption(options []model.Option) []outputdata.Option {
	oOptions := []outputdata.Option{}
	for _, industry := range options {
		oOption := outputdata.Option{
			Value: industry.ID,
			Text:  industry.Name,
		}
		oOptions = append(oOptions, oOption)
	}
	return oOptions
}

func convert(src interface{}, dst interface{}) error {
	fv := reflect.ValueOf(src)

	ft := fv.Type()
	if fv.Kind() == reflect.Ptr {
		ft = ft.Elem()
		fv = fv.Elem()
	}

	tv := reflect.ValueOf(dst)
	if tv.Kind() != reflect.Ptr {
		return fmt.Errorf("[Error] non-pointer: %v", dst)
	}

	num := ft.NumField()
	for i := 0; i < num; i++ {
		field := ft.Field(i)

		if !field.Anonymous {
			name := field.Name
			srcField := fv.FieldByName(name)
			dstField := tv.Elem().FieldByName(name)

			if srcField.IsValid() && dstField.IsValid() {
				if srcField.Type() == dstField.Type() {
					dstField.Set(srcField)
				}
			}
		}
	}

	return nil
}
