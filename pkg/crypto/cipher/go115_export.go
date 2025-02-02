// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package cipher

import (
	q "crypto/cipher"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cipher",
		Path: "crypto/cipher",
		Deps: map[string]string{
			"crypto/internal/subtle": "subtle",
			"crypto/subtle":          "subtle",
			"encoding/binary":        "binary",
			"errors":                 "errors",
			"io":                     "io",
		},
		Interfaces: map[string]reflect.Type{
			"AEAD":      reflect.TypeOf((*q.AEAD)(nil)).Elem(),
			"Block":     reflect.TypeOf((*q.Block)(nil)).Elem(),
			"BlockMode": reflect.TypeOf((*q.BlockMode)(nil)).Elem(),
			"Stream":    reflect.TypeOf((*q.Stream)(nil)).Elem(),
		},
		NamedTypes: map[string]igop.NamedType{
			"StreamReader": {reflect.TypeOf((*q.StreamReader)(nil)).Elem(), "Read", ""},
			"StreamWriter": {reflect.TypeOf((*q.StreamWriter)(nil)).Elem(), "Close,Write", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCBCDecrypter":     reflect.ValueOf(q.NewCBCDecrypter),
			"NewCBCEncrypter":     reflect.ValueOf(q.NewCBCEncrypter),
			"NewCFBDecrypter":     reflect.ValueOf(q.NewCFBDecrypter),
			"NewCFBEncrypter":     reflect.ValueOf(q.NewCFBEncrypter),
			"NewCTR":              reflect.ValueOf(q.NewCTR),
			"NewGCM":              reflect.ValueOf(q.NewGCM),
			"NewGCMWithNonceSize": reflect.ValueOf(q.NewGCMWithNonceSize),
			"NewGCMWithTagSize":   reflect.ValueOf(q.NewGCMWithTagSize),
			"NewOFB":              reflect.ValueOf(q.NewOFB),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
