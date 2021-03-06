package drreflect

import (
	"reflect"
)

type reflectStructTypeHandler struct {
	*reflectTypeHandler
}

func NewReflectStructTypeHandlerFromValue(v interface{}) StructTypeHandler {
	return NewReflectStructTypeHandler(reflect.TypeOf(v))
}

func NewReflectStructTypeHandler(t reflect.Type) *reflectStructTypeHandler {
	if t.Kind() == reflect.Ptr {
		if t.Elem().Kind() != reflect.Struct {
			panic("provided type was a pointer but not to a struct: " + t.String())
		}
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("provided type wasn't a struct or a pointer to a struct: " + t.String())
	}

	sth := reflectStructTypeHandler{
		NewReflectTypeHandler(t),
	}
	return &sth
}

func (r *reflectStructTypeHandler) GetFieldValue(input interface{}, fieldName string) interface{} {
	t := reflect.TypeOf(input)
	if r.t != t && r.ptr != t {
		panic("unexpected value type " + t.String() + " for accessor of type " + r.t.String())
	}
	v := reflect.ValueOf(input)
	if r.ptr == t {
		v = v.Elem()
	}
	fieldValue := v.FieldByName(fieldName)
	return fieldValue.Interface()
}
