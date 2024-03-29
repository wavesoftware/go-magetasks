package metadata

import (
	"fmt"
	"reflect"
)

type marker struct{}

func importPath(variable string) string {
	m := marker{}
	p := findPackageForType(m)
	return fmt.Sprintf("%s.%s", p, variable)
}

func findPackageForType(a interface{}) string {
	return reflect.TypeOf(a).PkgPath()
}
