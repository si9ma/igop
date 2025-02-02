// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package bzip2

import (
	q "compress/bzip2"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "bzip2",
		Path: "compress/bzip2",
		Deps: map[string]string{
			"bufio": "bufio",
			"io":    "io",
			"sort":  "sort",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"StructuralError": {reflect.TypeOf((*q.StructuralError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
