// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package suffixarray

import (
	q "index/suffixarray"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "suffixarray",
		Path: "index/suffixarray",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding/binary": "binary",
			"errors":          "errors",
			"io":              "io",
			"math":            "math",
			"regexp":          "regexp",
			"sort":            "sort",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Index": {reflect.TypeOf((*q.Index)(nil)).Elem(), "", "Bytes,FindAllIndex,Lookup,Read,Write,at,lookupAll"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
