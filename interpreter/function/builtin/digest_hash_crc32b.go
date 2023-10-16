// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"encoding/binary"
	"fmt"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Digest_hash_crc32b_Name = "digest.hash_crc32b"

var Digest_hash_crc32b_ArgumentTypes = []value.Type{value.StringType}

func Digest_hash_crc32b_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Digest_hash_crc32b_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Digest_hash_crc32b_ArgumentTypes[i] {
			return errors.TypeMismatch(Digest_hash_crc32b_Name, i+1, Digest_hash_crc32b_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of digest.hash_crc32b
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-crc32b/
func Digest_hash_crc32b(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Digest_hash_crc32b_Validate(args); err != nil {
		return value.Null, err
	}

	input := value.Unwrap[*value.String](args[0])

	// https://github.com/whik/crc-lib-c/blob/master/crcLib.c#L527
	var crc uint32 = 0xffffffff
	for _, c := range []byte(input.Value) {
		crc = crc ^ (uint32)(c)
		for i := 0; i < 8; i++ {
			if crc&0x1 != 0 {
				crc = (crc >> 1) ^ 0xEDB88320
			} else {
				crc = (crc >> 1)
			}
		}
	}
	crc = 0xffffffff ^ crc
	buf := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(buf, crc)

	return &value.String{
		Value: fmt.Sprintf("%08x", buf),
	}, nil
}