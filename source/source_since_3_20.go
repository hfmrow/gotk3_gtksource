// +build linux
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14
// +build !gtksourceview_3_18

// Filename version number based on gtk_sourceview

// Limit > gtk_3_14 because libgtksourceview-3.0 start at gtk >= 3.16
// Latest gtksourceview_3_18 start with gtk_3_16
// Latest gtksourceview_4_0 start with gtk_3_22

// Same copyright and license as the rest of the files in this project

package source

// #include <gtk/gtk.h>
// #include <glib.h>
// #include <glib-object.h>
// #include <gtksourceview/gtksource.h>
// #include "source_since_3_20.go.h"
// #include "source.go.h"
// #include "source_facility.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_source_tag_get_type()), marshalSourceTag},
	}

	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceTag"] = wrapSourceTag
}

/*
 * GtkSourceTag (full)
 */

// SourceTag is a representation of GTK's GtkSourceTag.
// Subclass of GtkTextTag
// A tag that can be applied to text in a GtkSourceBuffer
type SourceTag struct {
	gtk.TextTag
}

// native returns a pointer to the underlying GtkTextBuffer.
func (v *SourceTag) native() *C.GtkSourceTag {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceTag(p)
}

func marshalSourceTag(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))

	obj := glib.Take(unsafe.Pointer(c))

	return &gtk.TextTag{obj}, nil
}

func wrapSourceTag(obj *glib.Object) *SourceTag {
	return &SourceTag{gtk.TextTag{obj}}
}

// SourceTagNew is a wrapper around gtk_source_tag_new().
func SourceTagNew(tagName string) (*gtk.TextTag, error) {
	cstr := C.CString(tagName)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_source_tag_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}

	return toTextTag(c), nil
}

/*
 * SourceBuffer
 */

// SourceBufferCreateSourceTag is a wrapper around gtk_source_buffer_create_source_tag().
// nameWvalue := map[string]interface{}{}
func (v *SourceBuffer) CreateSourceTag(tagName string, nameWvalue map[string]interface{}) (*gtk.TextTag, error) {
	var cstr *C.gchar = nil

	if len(tagName) != 0 {
		cstr = C.CString(tagName)
		defer C.free(unsafe.Pointer(cstr))
	}
	// 20200927 -got an error: unexpected type: ... with the function designed for:
	// c := C.gtk_source_buffer_create_source_tag((*C.gchar)(cstr), nil, nil)
	// therefore, after a long test using different ways without resolving
	// this state, I use another tag-new function instead to create the tag.
	// And add it to SourceBuffer's TagTable like original fonction do.
	c := C.gtk_source_tag_new((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	e := toTextTag(c)
	if nameWvalue != nil { // Filling the tag with the properties if they are there
		for name, value := range nameWvalue {
			e.SetProperty(name, value)
		}
	}
	// Add tag to buffer like original function ...
	if tt, err := v.GetTagTable(); err == nil {
		tt.Add(e)
	} else {
		return nil, err
	}

	return e, nil
}
