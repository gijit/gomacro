// this file was generated by gomacro command: import "compress/bzip2"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"compress/bzip2"
)

func init() {
	Binds["compress/bzip2"] = map[string]Value{
		"NewReader":	ValueOf(bzip2.NewReader),
	}
	Types["compress/bzip2"] = map[string]Type{
		"StructuralError":	TypeOf((*bzip2.StructuralError)(nil)).Elem(),
	}
	Proxies["compress/bzip2"] = map[string]Type{
	}
}