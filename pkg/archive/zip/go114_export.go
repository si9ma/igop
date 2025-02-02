// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package zip

import (
	q "archive/zip"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "zip",
		Path: "archive/zip",
		Deps: map[string]string{
			"bufio":           "bufio",
			"compress/flate":  "flate",
			"encoding/binary": "binary",
			"errors":          "errors",
			"hash":            "hash",
			"hash/crc32":      "crc32",
			"io":              "io",
			"io/ioutil":       "ioutil",
			"os":              "os",
			"path":            "path",
			"strings":         "strings",
			"sync":            "sync",
			"time":            "time",
			"unicode/utf8":    "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Compressor":   {reflect.TypeOf((*q.Compressor)(nil)).Elem(), "", ""},
			"Decompressor": {reflect.TypeOf((*q.Decompressor)(nil)).Elem(), "", ""},
			"File":         {reflect.TypeOf((*q.File)(nil)).Elem(), "", "DataOffset,Open,findBodyOffset,hasDataDescriptor"},
			"FileHeader":   {reflect.TypeOf((*q.FileHeader)(nil)).Elem(), "", "FileInfo,ModTime,Mode,SetModTime,SetMode,isZip64"},
			"ReadCloser":   {reflect.TypeOf((*q.ReadCloser)(nil)).Elem(), "", "Close"},
			"Reader":       {reflect.TypeOf((*q.Reader)(nil)).Elem(), "", "RegisterDecompressor,decompressor,init"},
			"Writer":       {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Close,Create,CreateHeader,Flush,RegisterCompressor,SetComment,SetOffset,compressor"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrAlgorithm": reflect.ValueOf(&q.ErrAlgorithm),
			"ErrChecksum":  reflect.ValueOf(&q.ErrChecksum),
			"ErrFormat":    reflect.ValueOf(&q.ErrFormat),
		},
		Funcs: map[string]reflect.Value{
			"FileInfoHeader":       reflect.ValueOf(q.FileInfoHeader),
			"NewReader":            reflect.ValueOf(q.NewReader),
			"NewWriter":            reflect.ValueOf(q.NewWriter),
			"OpenReader":           reflect.ValueOf(q.OpenReader),
			"RegisterCompressor":   reflect.ValueOf(q.RegisterCompressor),
			"RegisterDecompressor": reflect.ValueOf(q.RegisterDecompressor),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Deflate": {reflect.TypeOf(q.Deflate), constant.MakeInt64(int64(q.Deflate))},
			"Store":   {reflect.TypeOf(q.Store), constant.MakeInt64(int64(q.Store))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
