package igop

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/tools/go/types/typeutil"
)

var (
	xtypeTypeNames = make(map[string]*types.Basic)
)

var (
	typesDummyStruct    = types.NewStruct(nil, nil)
	typesDummySig       = types.NewSignature(nil, nil, nil, false)
	typesError          = types.Universe.Lookup("error").Type()
	typesEmptyInterface = types.NewInterfaceType(nil, nil)
)

var (
	tyEmptyInterface = reflect.TypeOf((*interface{})(nil)).Elem()
	tyErrorInterface = reflect.TypeOf((*error)(nil)).Elem()
)

func init() {
	for i := types.Invalid; i <= types.UntypedNil; i++ {
		typ := types.Typ[i]
		xtypeTypeNames[typ.String()] = typ
	}
}

type TypesLoader struct {
	packages  map[string]*types.Package
	installed map[string]*Package
	rcache    map[reflect.Type]types.Type
	tcache    *typeutil.Map
	curpkg    *Package
	mode      Mode
}

// NewTypesLoader install package and readonly
func NewTypesLoader(mode Mode) Loader {
	r := &TypesLoader{
		packages:  make(map[string]*types.Package),
		installed: make(map[string]*Package),
		rcache:    make(map[reflect.Type]types.Type),
		tcache:    &typeutil.Map{},
		mode:      mode,
	}
	r.packages["unsafe"] = types.Unsafe
	r.rcache[tyErrorInterface] = typesError
	r.rcache[tyEmptyInterface] = typesEmptyInterface
	return r
}

func (r *TypesLoader) Installed(path string) (pkg *Package, ok bool) {
	pkg, ok = r.installed[path]
	return
}

func (r *TypesLoader) Packages() (pkgs []*types.Package) {
	for _, pkg := range r.packages {
		pkgs = append(pkgs, pkg)
	}
	return
}

func (r *TypesLoader) LookupPackage(pkgpath string) (*types.Package, bool) {
	pkg, ok := r.packages[pkgpath]
	return pkg, ok
}

func (r *TypesLoader) LookupReflect(typ types.Type) (reflect.Type, bool) {
	if rt := r.tcache.At(typ); rt != nil {
		return rt.(reflect.Type), true
	}
	return nil, false
}

func (r *TypesLoader) LookupTypes(typ reflect.Type) (types.Type, bool) {
	t, ok := r.rcache[typ]
	return t, ok
}

func (r *TypesLoader) Import(path string) (*types.Package, error) {
	if p, ok := r.packages[path]; ok {
		return p, nil
	}
	pkg, ok := registerPkgs[path]
	if !ok {
		return nil, fmt.Errorf("Not found package %v", path)
	}
	p := types.NewPackage(pkg.Path, pkg.Name)
	r.packages[path] = p
	var list []*types.Package
	for dep, _ := range pkg.Deps {
		p, err := r.Import(dep)
		if err == nil {
			list = append(list, p)
		}
	}
	if err := r.installPackage(pkg); err != nil {
		return nil, err
	}
	p.SetImports(list)
	p.MarkComplete()
	return p, nil
}

func (r *TypesLoader) installPackage(pkg *Package) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		r.curpkg = nil
	}()
	r.curpkg = pkg
	r.installed[pkg.Path] = pkg
	p, ok := r.packages[pkg.Path]
	if !ok {
		p = types.NewPackage(pkg.Path, pkg.Name)
		r.packages[pkg.Path] = p
	}
	for name, typ := range pkg.Interfaces {
		r.InsertInterface(p, name, typ)
	}
	for name, typ := range pkg.NamedTypes {
		if typ.Typ.Kind() == reflect.Struct {
			r.InsertNamedType(p, name, typ)
		}
	}
	for name, typ := range pkg.NamedTypes {
		if typ.Typ.Kind() != reflect.Struct {
			r.InsertNamedType(p, name, typ)
		}
	}
	for name, typ := range pkg.AliasTypes {
		r.InsertAlias(p, name, typ)
	}
	for name, fn := range pkg.Funcs {
		r.InsertFunc(p, name, fn)
	}
	for name, v := range pkg.Vars {
		r.InsertVar(p, name, v.Elem())
	}
	for name, c := range pkg.TypedConsts {
		r.InsertTypedConst(p, name, c)
	}
	for name, c := range pkg.UntypedConsts {
		r.InsertUntypedConst(p, name, c)
	}
	return
}

func (r *TypesLoader) InsertInterface(p *types.Package, name string, rt reflect.Type) {
	r.ToType(rt)
}

func (r *TypesLoader) InsertNamedType(p *types.Package, name string, t NamedType) {
	r.ToType(t.Typ)
}

func (r *TypesLoader) InsertAlias(p *types.Package, name string, rt reflect.Type) {
	typ := r.ToType(rt)
	p.Scope().Insert(types.NewTypeName(token.NoPos, p, name, typ))
}

func (r *TypesLoader) InsertFunc(p *types.Package, name string, v reflect.Value) {
	typ := r.ToType(v.Type())
	p.Scope().Insert(types.NewFunc(token.NoPos, p, name, typ.(*types.Signature)))
}

func (r *TypesLoader) InsertVar(p *types.Package, name string, v reflect.Value) {
	typ := r.ToType(v.Type())
	p.Scope().Insert(types.NewVar(token.NoPos, p, name, typ))
}

func (r *TypesLoader) InsertConst(p *types.Package, name string, typ types.Type, c constant.Value) {
	p.Scope().Insert(types.NewConst(token.NoPos, p, name, typ, c))
}

func splitPath(path string) (pkg string, name string, ok bool) {
	pos := strings.LastIndex(path, ".")
	if pos == -1 {
		return path, "", false
	}
	return path[:pos], path[pos+1:], true
}

func (r *TypesLoader) parserNamed(path string) (*types.Package, string) {
	if pkg, name, ok := splitPath(path); ok {
		if p := r.GetPackage(pkg); p != nil {
			return p, name
		}
	}
	panic(fmt.Errorf("parse path failed: %v", path))
}

func (r *TypesLoader) LookupType(typ string) types.Type {
	if t, ok := xtypeTypeNames[typ]; ok {
		return t
	}
	p, name := r.parserNamed(typ)
	return p.Scope().Lookup(name).Type()
}

func (r *TypesLoader) InsertTypedConst(p *types.Package, name string, v TypedConst) {
	typ := r.ToType(v.Typ)
	r.InsertConst(p, name, typ, v.Value)
}

func (r *TypesLoader) InsertUntypedConst(p *types.Package, name string, v UntypedConst) {
	var typ types.Type
	if t, ok := xtypeTypeNames[v.Typ]; ok {
		typ = t
	} else {
		typ = r.LookupType(v.Typ)
	}
	r.InsertConst(p, name, typ, v.Value)
}

func (r *TypesLoader) GetPackage(pkg string) *types.Package {
	if pkg == "" {
		return nil
	}
	if p, ok := r.packages[pkg]; ok {
		return p
	}
	var name string
	if r.curpkg != nil {
		name = r.curpkg.Deps[pkg]
	}
	if name == "" {
		pkgs := strings.Split(pkg, "/")
		name = pkgs[len(pkgs)-1]
	}
	p := types.NewPackage(pkg, name)
	r.packages[pkg] = p
	return p
}

func toTypeChanDir(dir reflect.ChanDir) types.ChanDir {
	switch dir {
	case reflect.RecvDir:
		return types.RecvOnly
	case reflect.SendDir:
		return types.SendOnly
	case reflect.BothDir:
		return types.SendRecv
	}
	panic("unreachable")
}

func (r *TypesLoader) Insert(v reflect.Value) {
	typ := r.ToType(v.Type())
	if v.Kind() == reflect.Func {
		name := runtime.FuncForPC(v.Pointer()).Name()
		names := strings.Split(name, ".")
		pkg := r.GetPackage(names[0])
		pkg.Scope().Insert(types.NewFunc(token.NoPos, pkg, names[1], typ.(*types.Signature)))
	}
}

func (r *TypesLoader) toFunc(pkg *types.Package, recv *types.Var, inoff int, rt reflect.Type) *types.Signature {
	numIn := rt.NumIn()
	numOut := rt.NumOut()
	in := make([]*types.Var, numIn-inoff, numIn-inoff)
	out := make([]*types.Var, numOut, numOut)
	for i := inoff; i < numIn; i++ {
		it := r.ToType(rt.In(i))
		in[i-inoff] = types.NewVar(token.NoPos, pkg, "", it)
	}
	for i := 0; i < numOut; i++ {
		it := r.ToType(rt.Out(i))
		out[i] = types.NewVar(token.NoPos, pkg, "", it)
	}
	return types.NewSignature(recv, types.NewTuple(in...), types.NewTuple(out...), rt.IsVariadic())
}

func (r *TypesLoader) ToType(rt reflect.Type) types.Type {
	if t, ok := r.rcache[rt]; ok {
		return t
	}
	var typ types.Type
	var fields []*types.Var
	var imethods []*types.Func
	kind := rt.Kind()
	switch kind {
	case reflect.Invalid:
		typ = types.Typ[types.Invalid]
	case reflect.Bool:
		typ = types.Typ[types.Bool]
	case reflect.Int:
		typ = types.Typ[types.Int]
	case reflect.Int8:
		typ = types.Typ[types.Int8]
	case reflect.Int16:
		typ = types.Typ[types.Int16]
	case reflect.Int32:
		typ = types.Typ[types.Int32]
	case reflect.Int64:
		typ = types.Typ[types.Int64]
	case reflect.Uint:
		typ = types.Typ[types.Uint]
	case reflect.Uint8:
		typ = types.Typ[types.Uint8]
	case reflect.Uint16:
		typ = types.Typ[types.Uint16]
	case reflect.Uint32:
		typ = types.Typ[types.Uint32]
	case reflect.Uint64:
		typ = types.Typ[types.Uint64]
	case reflect.Uintptr:
		typ = types.Typ[types.Uintptr]
	case reflect.Float32:
		typ = types.Typ[types.Float32]
	case reflect.Float64:
		typ = types.Typ[types.Float64]
	case reflect.Complex64:
		typ = types.Typ[types.Complex64]
	case reflect.Complex128:
		typ = types.Typ[types.Complex128]
	case reflect.Array:
		elem := r.ToType(rt.Elem())
		typ = types.NewArray(elem, int64(rt.Len()))
	case reflect.Chan:
		elem := r.ToType(rt.Elem())
		dir := toTypeChanDir(rt.ChanDir())
		typ = types.NewChan(dir, elem)
	case reflect.Func:
		pkg := r.GetPackage(r.curpkg.Path)
		typ = r.toFunc(pkg, nil, 0, rt)
	case reflect.Interface:
		n := rt.NumMethod()
		imethods = make([]*types.Func, n, n)
		pkg := r.GetPackage(rt.PkgPath())
		for i := 0; i < n; i++ {
			im := rt.Method(i)
			sig := typesDummySig
			imethods[i] = types.NewFunc(token.NoPos, pkg, im.Name, sig)
		}
		typ = types.NewInterfaceType(imethods, nil)
	case reflect.Map:
		key := r.ToType(rt.Key())
		elem := r.ToType(rt.Elem())
		typ = types.NewMap(key, elem)
	case reflect.Ptr:
		elem := r.ToType(rt.Elem())
		typ = types.NewPointer(elem)
	case reflect.Slice:
		elem := r.ToType(rt.Elem())
		typ = types.NewSlice(elem)
	case reflect.String:
		typ = types.Typ[types.String]
	case reflect.Struct:
		n := rt.NumField()
		fields = make([]*types.Var, n, n)
		tags := make([]string, n, n)
		pkg := r.GetPackage(rt.PkgPath())
		for i := 0; i < n; i++ {
			f := rt.Field(i)
			ft := types.Typ[types.UnsafePointer] //r.ToType(f.Type)
			fields[i] = types.NewVar(token.NoPos, pkg, f.Name, ft)
			tags[i] = string(f.Tag)
		}
		typ = types.NewStruct(fields, tags)
	case reflect.UnsafePointer:
		typ = types.Typ[types.UnsafePointer]
	default:
		panic("unreachable")
	}
	var named *types.Named
	if path := rt.PkgPath(); path != "" {
		pkg := r.GetPackage(path)
		obj := types.NewTypeName(token.NoPos, pkg, rt.Name(), nil)
		named = types.NewNamed(obj, typ, nil)
		typ = named
		pkg.Scope().Insert(obj)
	}
	r.rcache[rt] = typ
	r.tcache.Set(typ, rt)
	if kind == reflect.Struct {
		n := rt.NumField()
		pkg := r.GetPackage(rt.PkgPath())
		for i := 0; i < n; i++ {
			f := rt.Field(i)
			ft := r.ToType(f.Type)
			fields[i] = types.NewField(token.NoPos, pkg, f.Name, ft, f.Anonymous)
		}
	} else if kind == reflect.Interface {
		n := rt.NumMethod()
		pkg := named.Obj().Pkg()
		recv := types.NewVar(token.NoPos, pkg, "", typ)
		for i := 0; i < n; i++ {
			im := rt.Method(i)
			sig := r.toFunc(pkg, recv, 0, im.Type)
			imethods[i] = types.NewFunc(token.NoPos, pkg, im.Name, sig)
		}
		typ.Underlying().(*types.Interface).Complete()
	}
	if named != nil {
		if kind != reflect.Interface {
			var filter func(name string, ptr bool) bool
			pkg := named.Obj().Pkg()
			if p, ok := r.installed[pkg.Path()]; ok {
				if t, ok := p.NamedTypes[named.Obj().Name()]; ok {
					m := make(map[string]bool)
					pm := make(map[string]bool)
					for _, v := range strings.Split(t.Methods, ",") {
						m[v] = true
					}
					for _, v := range strings.Split(t.PtrMethods, ",") {
						pm[v] = true
					}
					filter = func(name string, ptr bool) bool {
						if ptr {
							return pm[name]
						}
						return m[name]
					}
				}
			}

			prt := reflect.PtrTo(rt)
			ptyp := r.ToType(prt)
			precv := types.NewVar(token.NoPos, pkg, "", ptyp)

			skip := make(map[string]bool)
			for _, im := range AllMethod(prt, r.mode&DisableUnexportMethods == 0) {
				if filter != nil && !filter(im.Name, true) {
					continue
				}
				var sig *types.Signature
				if im.Type != nil {
					sig = r.toFunc(pkg, precv, 1, im.Type)
				} else {
					sig = typesDummySig
				}
				skip[im.Name] = true
				named.AddMethod(types.NewFunc(token.NoPos, pkg, im.Name, sig))
			}
			recv := types.NewVar(token.NoPos, pkg, "", typ)
			for _, im := range AllMethod(rt, r.mode&DisableUnexportMethods == 0) {
				if skip[im.Name] {
					continue
				}
				if filter != nil && !filter(im.Name, false) {
					continue
				}
				var sig *types.Signature
				if im.Type != nil {
					sig = r.toFunc(pkg, recv, 1, im.Type)
				} else {
					sig = typesDummySig
				}
				named.AddMethod(types.NewFunc(token.NoPos, pkg, im.Name, sig))
			}
		}
	}
	return typ
}
