// Package itkacl provides an interface for go applications to interface with the itkacl library.
// At the start of your application use the ItkaclInit function to create a context used globaly for all ITKACL queries.
// Make sure to allways free the context using ItkaclFree at the end of your program.
// Once a context has been created, use the ItkaclCheck method to make ITKACL queries.
package itkacl

// #cgo LDFLAGS: -litkacl
// #include <itkacl.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"unsafe"
)

var ctx *C.struct_itkacl_ctx

// ItkaclInit creates a ITKACL context. This is needed to make ITKACL queries.
func ItkaclInit() error {
	errorMsg := make([]C.char, 1024)
	ctx = C.itkacl_create_ctx(&errorMsg[0], 1024)
	if ctx == nil {
		return errors.New(C.GoString(&errorMsg[0]))
	}
	return nil
}

// ItkaclFree frees the memory allocated for the ITKACL context.
func ItkaclFree() {
	C.itkacl_free_ctx(ctx)
}

// ItkaclCheck makes an ITKACL query.
func ItkaclCheck(realm string, user string) (*bool, error) {
	_realm := C.CString(realm)
	_user := C.CString(user)
	defer C.free(unsafe.Pointer(_realm))
	defer C.free(unsafe.Pointer(_user))

	errorMsg := make([]C.char, 1024)
	result := C.itkacl_check_with_ctx(ctx, _realm, _user, &errorMsg[0], 1024)

	if result == -1 {
		return nil, errors.New(C.GoString(&errorMsg[0]))
	}

	boolResult := result == 0
	return &boolResult, nil
}
