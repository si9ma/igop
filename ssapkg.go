package gossa

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
)

func BuildPackage(loader Loader, importer types.Importer, fset *token.FileSet, pkg *types.Package, files []*ast.File, mode ssa.BuilderMode) (*ssa.Package, *types.Info, error) {
	if fset == nil {
		panic("no token.FileSet")
	}
	if pkg.Path() == "" {
		panic("package has no import path")
	}

	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Scopes:     make(map[ast.Node]*types.Scope),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}

	tc := &types.Config{
		Importer: NewImporter(loader, importer),
	}
	if err := types.NewChecker(tc, fset, pkg, info).Files(files); err != nil {
		return nil, nil, err
	}
	prog := ssa.NewProgram(fset, mode)

	// Create SSA packages for all imports.
	// Order is not significant.
	created := make(map[*types.Package]bool)
	var createAll func(pkgs []*types.Package)
	createAll = func(pkgs []*types.Package) {
		for _, p := range pkgs {
			if !created[p] {
				created[p] = true
				prog.CreatePackage(p, nil, nil, true)
				createAll(p.Imports())
			}
		}
	}
	createAll(pkg.Imports())

	// create depends
	for _, p := range loader.Packages() {
		// indirect incomplete package
		if !p.Complete() {
			p.MarkComplete()
		}
		if !created[p] {
			prog.CreatePackage(p, nil, nil, true)
		}
	}

	// Create and build the primary package.
	ssapkg := prog.CreatePackage(pkg, files, info, false)
	ssapkg.Build()
	return ssapkg, info, nil
}
