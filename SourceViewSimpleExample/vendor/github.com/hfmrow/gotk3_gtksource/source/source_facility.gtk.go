package source

// #include <glib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// #include <gio/gio.h>
// #include <gtksourceview/gtksource.h>
// #include "source_facility.gtk.go.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

/*
 * Gotk3 duplicated function / var
 */

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	return b != C.FALSE
}

func cGSList(clist *glib.SList) *C.GSList {
	if clist == nil {
		return nil
	}
	return (*C.GSList)(unsafe.Pointer(clist.Native()))
}

func free(str ...interface{}) {
	for _, s := range str {
		switch x := s.(type) {
		case *C.char:
			C.free(unsafe.Pointer(x))
		case []*C.char:
			for _, cp := range x {
				C.free(unsafe.Pointer(cp))
			}
			/*
				case C.gpointer:
					C.g_free(C.gpointer(c))
			*/
		default:
			fmt.Printf("utils.go free(): Unknown type: %T\n", x)
		}

	}
}

func goString(cstr *C.gchar) string {
	return C.GoString((*C.char)(cstr))
}

// same implementation as package glib
func toGoStringArray(c **C.gchar) []string {
	var strs []string
	originalc := c
	defer C.g_strfreev(originalc)

	for *c != nil {
		strs = append(strs, C.GoString((*C.char)(*c)))
		c = C.next_gcharptr(c)
	}

	return strs
}

var nilPtrErr = errors.New("cgo returned unexpected nil pointer")

/*
 * GdkRectangle
 */

func nativeGdkRectangle(rect gdk.Rectangle) *C.GdkRectangle {
	// Note: Here we can't use rect.GdkRectangle because it would return
	// C type prefixed with gdk package. A ways how to resolve this Go
	// issue with same C structs in different Go packages is documented
	// here https://github.com/golang/go/issues/13467 .
	// This is the easiest way how to resolve the problem.
	return &C.GdkRectangle{
		x:      C.int(rect.GetX()),
		y:      C.int(rect.GetY()),
		width:  C.int(rect.GetWidth()),
		height: C.int(rect.GetHeight()),
	}
}
