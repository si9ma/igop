// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package pkix

import (
	q "crypto/x509/pkix"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "pkix",
		Path: "crypto/x509/pkix",
		Deps: map[string]string{
			"encoding/asn1": "asn1",
			"encoding/hex":  "hex",
			"fmt":           "fmt",
			"math/big":      "big",
			"time":          "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"AlgorithmIdentifier":          {reflect.TypeOf((*q.AlgorithmIdentifier)(nil)).Elem(), "", ""},
			"AttributeTypeAndValue":        {reflect.TypeOf((*q.AttributeTypeAndValue)(nil)).Elem(), "", ""},
			"AttributeTypeAndValueSET":     {reflect.TypeOf((*q.AttributeTypeAndValueSET)(nil)).Elem(), "", ""},
			"CertificateList":              {reflect.TypeOf((*q.CertificateList)(nil)).Elem(), "", "HasExpired"},
			"Extension":                    {reflect.TypeOf((*q.Extension)(nil)).Elem(), "", ""},
			"Name":                         {reflect.TypeOf((*q.Name)(nil)).Elem(), "String,ToRDNSequence,appendRDNs", "FillFromRDNSequence"},
			"RDNSequence":                  {reflect.TypeOf((*q.RDNSequence)(nil)).Elem(), "String", ""},
			"RelativeDistinguishedNameSET": {reflect.TypeOf((*q.RelativeDistinguishedNameSET)(nil)).Elem(), "", ""},
			"RevokedCertificate":           {reflect.TypeOf((*q.RevokedCertificate)(nil)).Elem(), "", ""},
			"TBSCertificateList":           {reflect.TypeOf((*q.TBSCertificateList)(nil)).Elem(), "", ""},
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
