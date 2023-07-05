package abi

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

func EncodeFunctionSignature(funcName string) []byte {
	return crypto.Keccak256([]byte(funcName))[:4]
}

func AbiDecode(parameters string, data []byte) []interface{} {

	args := make(abi.Arguments, 0)

	parametersArr := strings.Split(parameters, ",")
	for _, p := range parametersArr {
		arg := abi.Argument{}
		var err error
		arg.Type, err = abi.NewType(p, "", nil)
		if err != nil {
			return []interface{}{}
		}
		args = append(args, arg)
	}

	result, err := args.Unpack(data)
	if err != nil {
		return []interface{}{}
	} else {
		return result
	}
}

func AbiEncode(parameters string, vars ...interface{}) []byte {

	args := make(abi.Arguments, 0)

	parametersArr := strings.Split(parameters, ",")
	for _, p := range parametersArr {
		arg := abi.Argument{}
		var err error
		arg.Type, err = abi.NewType(p, "", nil)
		if err != nil {
			return []byte{}
		}
		args = append(args, arg)
	}

	result, err := args.Pack(vars...)
	if err != nil {
		return []byte{}
	} else {
		return result
	}
}

func getArgs(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
		if j >= 0 && j > i+1 {
			return s[i+1 : j]
		}
	}
	return ""
}

func ToAbi(signature string, params ...interface{}) []byte {
	methodSig := EncodeFunctionSignature(signature)
	args := getArgs(signature)
	if len(args) == 0 {
		return methodSig
	}
	inputCode := AbiEncode(args, params...)
	if len(inputCode) == 0 {
		//panic(err)
		return []byte{}
	}
	code := append(methodSig, inputCode...)
	return code
}
