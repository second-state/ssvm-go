// +build image

package wasmedge

/*
#cgo linux LDFLAGS: -lwasmedge-image_c -ljpeg -lpng

#include <wasmedge-image.h>
*/
import "C"

func NewImageImportObject() *ImportObject {
	self := &ImportObject{
		_inner: C.WasmEdge_Image_ImportObjectCreate(),
	}
	if self._inner == nil {
		return nil
	}
	return self
}
