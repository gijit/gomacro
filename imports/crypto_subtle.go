// this file was generated by gomacro command: import "crypto/subtle"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/subtle"
)

func init() {
	Binds["crypto/subtle"] = map[string]Value{
		"ConstantTimeByteEq":	ValueOf(subtle.ConstantTimeByteEq),
		"ConstantTimeCompare":	ValueOf(subtle.ConstantTimeCompare),
		"ConstantTimeCopy":	ValueOf(subtle.ConstantTimeCopy),
		"ConstantTimeEq":	ValueOf(subtle.ConstantTimeEq),
		"ConstantTimeLessOrEq":	ValueOf(subtle.ConstantTimeLessOrEq),
		"ConstantTimeSelect":	ValueOf(subtle.ConstantTimeSelect),
	}
	Types["crypto/subtle"] = map[string]Type{
	}
	Proxies["crypto/subtle"] = map[string]Type{
	}
}