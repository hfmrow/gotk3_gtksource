// +build linux
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20
// +build gtksourceview_3_18 gtksourceview_3_20 gtksourceview_3_22 gtksourceview_3_24 gtksourceview_deprecated

// Filename version number based on gtk_sourceview

// Limit > gtk_3_14 because libgtksourceview-3.0 start at gtk >= 3.16
// Latest gtksourceview_3_18 start with gtk_3_16
// Latest gtksourceview_4_0 start with gtk_3_22

// GtkSourceView 3.24 is the latest stable GtkSourceView 3 version.

// Same copyright and license as the rest of the files in this project

package source

// #include <gtk/gtk.h>
// #include <gtksourceview/gtksource.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * SourceCompletion
 */

// Start is a wrapper around gtk_source_completion_start().
// since GtkSourceView 3.99.7, gtk_source_completion_show() has been renamed to
// gtk_source_completion_start().
// A List must be manually freed by either calling Free() or FreeFull()
func (v *SourceCompletion) Show(providers *glib.List, context *SourceCompletionContext) bool {
	return gobool(
		C.gtk_source_completion_show(v.native(), toCGList(providers), context.native()))
}

/*
 * SourceCompletionItem
 */

// SourceCompletionItemNew2 is a wrapper around gtk_source_completion_item_new2().
func SourceCompletionItemNew2() (*SourceCompletionItem, error) {
	c := C.gtk_source_completion_item_new2()

	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceCompletionItem(glib.Take(unsafe.Pointer(c))), nil
}
