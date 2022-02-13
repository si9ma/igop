// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package debug

import (
	q "runtime/debug"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "debug",
		Path: "runtime/debug",
		Deps: map[string]string{
			"os":      "os",
			"runtime": "runtime",
			"sort":    "sort",
			"strings": "strings",
			"time":    "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"BuildInfo": {reflect.TypeOf((*q.BuildInfo)(nil)).Elem(), "", ""},
			"GCStats":   {reflect.TypeOf((*q.GCStats)(nil)).Elem(), "", ""},
			"Module":    {reflect.TypeOf((*q.Module)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"FreeOSMemory":    reflect.ValueOf(q.FreeOSMemory),
			"PrintStack":      reflect.ValueOf(q.PrintStack),
			"ReadBuildInfo":   reflect.ValueOf(q.ReadBuildInfo),
			"ReadGCStats":     reflect.ValueOf(q.ReadGCStats),
			"SetGCPercent":    reflect.ValueOf(q.SetGCPercent),
			"SetMaxStack":     reflect.ValueOf(q.SetMaxStack),
			"SetMaxThreads":   reflect.ValueOf(q.SetMaxThreads),
			"SetPanicOnFault": reflect.ValueOf(q.SetPanicOnFault),
			"SetTraceback":    reflect.ValueOf(q.SetTraceback),
			"Stack":           reflect.ValueOf(q.Stack),
			"WriteHeapDump":   reflect.ValueOf(q.WriteHeapDump),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}