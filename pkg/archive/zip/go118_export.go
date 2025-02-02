// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

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
			"io/fs":           "fs",
			"os":              "os",
			"path":            "path",
			"sort":            "sort",
			"strings":         "strings",
			"sync":            "sync",
			"time":            "time",
			"unicode/utf8":    "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Compressor":   {reflect.TypeOf((*q.Compressor)(nil)).Elem(), "", ""},
			"Decompressor": {reflect.TypeOf((*q.Decompressor)(nil)).Elem(), "", ""},
			"File":         {reflect.TypeOf((*q.File)(nil)).Elem(), "", "DataOffset,Open,OpenRaw,findBodyOffset"},
			"FileHeader":   {reflect.TypeOf((*q.FileHeader)(nil)).Elem(), "", "FileInfo,ModTime,Mode,SetModTime,SetMode,hasDataDescriptor,isZip64"},
			"ReadCloser":   {reflect.TypeOf((*q.ReadCloser)(nil)).Elem(), "", "Close"},
			"Reader":       {reflect.TypeOf((*q.Reader)(nil)).Elem(), "", "Open,RegisterDecompressor,decompressor,init,initFileList,openLookup,openReadDir"},
			"Writer":       {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Close,Copy,Create,CreateHeader,CreateRaw,Flush,RegisterCompressor,SetComment,SetOffset,compressor,prepare"},
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
