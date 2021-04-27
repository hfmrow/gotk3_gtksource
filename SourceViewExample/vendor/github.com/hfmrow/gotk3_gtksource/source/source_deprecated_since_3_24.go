// +build linux
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14
// +build gtksourceview_3_18 gtksourceview_3_20 gtksourceview_3_22 gtksourceview_deprecated

// Filename version number based on gtk_sourceview

// Limit > gtk_3_14 because libgtksourceview-3.0 start at gtk >= 3.16
// Latest gtksourceview_3_18 start with gtk_3_16
// Latest gtksourceview_4_0 start with gtk_3_22

// Same copyright and license as the rest of the files in this project

package source

// #include <gtk/gtk.h>
// #include <gtksourceview/gtksource.h>
// #include "source_since_3_24.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"

	"github.com/gotk3/gotk3/glib"
)

/*
 * SourceCompletionItem
 */

// SourceCompletionItemNew is a wrapper around gtk_source_completion_item_new().
func SourceCompletionItemNew(label, text string, icon *gdk.Pixbuf, info string) (*SourceCompletionItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	cstr1 := C.CString(text)
	defer C.free(unsafe.Pointer(cstr1))
	cstr2 := C.CString(info)
	defer C.free(unsafe.Pointer(cstr2))

	c := C.gtk_source_completion_item_new(
		(*C.gchar)(cstr),
		(*C.gchar)(cstr1),
		(*C.GdkPixbuf)(unsafe.Pointer(icon.Native())),
		(*C.gchar)(cstr2))

	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceCompletionItem(glib.Take(unsafe.Pointer(c))), nil
}

// SourceCompletionItemNewWithMarkup is a wrapper around gtk_source_completion_item_new_with_markup().
func SourceCompletionItemNewWithMarkup(markup, text string, icon *gdk.Pixbuf, info string) (*SourceCompletionItem, error) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	cstr1 := C.CString(text)
	defer C.free(unsafe.Pointer(cstr1))
	cstr2 := C.CString(info)
	defer C.free(unsafe.Pointer(cstr2))

	c := C.gtk_source_completion_item_new_with_markup(
		(*C.gchar)(cstr),
		(*C.gchar)(cstr1),
		(*C.GdkPixbuf)(unsafe.Pointer(icon.Native())),
		(*C.gchar)(cstr2))

	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceCompletionItem(glib.Take(unsafe.Pointer(c))), nil
}
