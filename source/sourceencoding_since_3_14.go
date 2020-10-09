// +build linux
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

// Filename version number based on gtk_sourceview

// Limit > gtk_3_14 because libgtksourceview-3.0 start at gtk >= 3.16
// Latest gtksourceview_3_18 start with gtk_3_16
// Latest gtksourceview_4_0 start with gtk_3_22

// Same copyright and license as the rest of the files in this project

package source

// #include <gtk/gtk.h>
// #include <gtksourceview/gtksource.h>
// #include "sourceencoding_since_3_14.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_source_encoding_get_type()), marshalSourceEncoding},
	}

	glib.RegisterGValueMarshalers(tm)

	// gtk.WrapMap["GtkSourceEncoding"] = wrapSourceEncoding
}

/*
 * GtkSourceEncoding
 */

type SourceEncoding C.GtkSourceEncoding

// native returns a pointer to the underlying GtkSourceEncoding.
func (v *SourceEncoding) native() *C.GtkSourceEncoding {
	if v == nil {
		return nil
	}
	p := unsafe.Pointer(v)
	return C.toGtkSourceEncoding(p)
}

func marshalSourceEncoding(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*SourceEncoding)(unsafe.Pointer(c)), nil
}

// SourceEncodingGetUtf8 is a wrapper around gtk_source_encoding_get_utf8().
func SourceEncodingGetUtf8() (*SourceEncoding, error) {
	c := C.gtk_source_encoding_get_utf8()
	if c == nil {
		return nil, nilPtrErr
	}

	e := (*SourceEncoding)(unsafe.Pointer(c))
	return e, nil
}

// SourceEncodingGetCurrent is a wrapper around gtk_source_encoding_get_current().
func SourceEncodingGetCurrent() (*SourceEncoding, error) {
	c := C.gtk_source_encoding_get_current()
	if c == nil {
		return nil, nilPtrErr
	}

	e := (*SourceEncoding)(unsafe.Pointer(c))
	return e, nil
}

// SourceEncodingGetFromCharset is a wrapper around gtk_source_encoding_get_from_charset().
func SourceEncodingGetFromCharset(charset string) (*SourceEncoding, error) {
	cstr := C.CString(charset)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_source_encoding_get_from_charset((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	e := (*SourceEncoding)(unsafe.Pointer(c))
	return e, nil
}

// ToString is a wrapper around gtk_source_encoding_to_string().
func (v *SourceEncoding) ToString() string {
	return toGoStringFree(C.gtk_source_encoding_to_string(v.native()))
}

// GetName is a wrapper around gtk_source_encoding_get_name().
func (v *SourceEncoding) GetName() string {
	return goString(C.gtk_source_encoding_get_name(v.native()))
}

// GetCharset is a wrapper around gtk_source_encoding_get_charset().
func (v *SourceEncoding) GetCharset() string {
	return goString(C.gtk_source_encoding_get_charset(v.native()))
}

// GetAll is a wrapper around gtk_source_encoding_get_all().
// the returned list must be Free with (*glib.SList).Free().
func SourceEncodingGetAll() *glib.SList {
	c := (*C.GSList)(C.gtk_source_encoding_get_all())

	all := glib.WrapSList(uintptr(unsafe.Pointer(c)))
	if all == nil {
		return nil
	}
	return all
}

// Free is a wrapper around gtk_source_encoding_free().
func (v *SourceEncoding) Free() {
	C.gtk_source_encoding_free(v.native())
}

// Copy is a wrapper around gtk_source_encoding_copy().
func (v *SourceEncoding) Copy() (*SourceEncoding, error) {
	c := C.gtk_source_encoding_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := (*SourceEncoding)(unsafe.Pointer(c))
	return e, nil
}
