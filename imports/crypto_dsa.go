// this file was generated by gomacro command: import "crypto/dsa"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/dsa"
)

func init() {
	Binds["crypto/dsa"] = map[string]Value{
		"ErrInvalidPublicKey":	ValueOf(&dsa.ErrInvalidPublicKey).Elem(),
		"GenerateKey":	ValueOf(dsa.GenerateKey),
		"GenerateParameters":	ValueOf(dsa.GenerateParameters),
		"L1024N160":	ValueOf(dsa.L1024N160),
		"L2048N224":	ValueOf(dsa.L2048N224),
		"L2048N256":	ValueOf(dsa.L2048N256),
		"L3072N256":	ValueOf(dsa.L3072N256),
		"Sign":	ValueOf(dsa.Sign),
		"Verify":	ValueOf(dsa.Verify),
	}
	Types["crypto/dsa"] = map[string]Type{
		"ParameterSizes":	TypeOf((*dsa.ParameterSizes)(nil)).Elem(),
		"Parameters":	TypeOf((*dsa.Parameters)(nil)).Elem(),
		"PrivateKey":	TypeOf((*dsa.PrivateKey)(nil)).Elem(),
		"PublicKey":	TypeOf((*dsa.PublicKey)(nil)).Elem(),
	}
	Proxies["crypto/dsa"] = map[string]Type{
	}
}