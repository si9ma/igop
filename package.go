package gossa

import (
	"go/constant"
	"reflect"
)

var (
	registerPkgs = make(map[string]*Package)
)

// LookupPackage lookup register pkgs
func LookupPackage(name string) (pkg *Package, ok bool) {
	pkg, ok = registerPkgs[name]
	return
}

// RegisterPackage register pkg
func RegisterPackage(pkg *Package) {
	if p, ok := registerPkgs[pkg.Path]; ok {
		p.merge(pkg)
		return
	}
	registerPkgs[pkg.Path] = pkg
	//	externPackages[pkg.Path] = true
}

type TypedConst struct {
	Typ   reflect.Type
	Value constant.Value
}

type UntypedConst struct {
	Typ   string
	Value constant.Value
}

type NamedType struct {
	Typ        reflect.Type
	Methods    string
	PtrMethods string
}

type Package struct {
	Name          string
	Path          string
	Interfaces    map[string]reflect.Type
	NamedTypes    map[string]NamedType
	AliasTypes    map[string]reflect.Type
	Vars          map[string]reflect.Value
	Funcs         map[string]reflect.Value
	TypedConsts   map[string]TypedConst
	UntypedConsts map[string]UntypedConst
	Deps          map[string]string        // path : name
	methods       map[string]reflect.Value // methods cached
}

// merge same package
func (p *Package) merge(same *Package) {
	for k, v := range same.Interfaces {
		p.Interfaces[k] = v
	}
	for k, v := range same.NamedTypes {
		p.NamedTypes[k] = v
	}
	for k, v := range same.Vars {
		p.Vars[k] = v
	}
	for k, v := range same.Funcs {
		p.Funcs[k] = v
	}
	for k, v := range same.UntypedConsts {
		p.UntypedConsts[k] = v
	}
}

var (
	externValues = make(map[string]reflect.Value)
)

// register external function for no function body
func RegisterExternal(key string, i interface{}) {
	externValues[key] = reflect.ValueOf(i)
}
