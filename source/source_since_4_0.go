// +build linux
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20
// +build !gtksourceview_3_18,!gtksourceview_3_20,!gtksourceview_3_22,!gtksourceview_3_24

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
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

/*
 * GtkSource initialization
 */

// SourceInit is a wrapper around gtk_source_init().
func SourceInit() {
	C.gtk_source_init()

}

// SourceFinalize is a wrapper around gtk_source_finalize().
func SourceFinalize() {
	C.gtk_source_finalize()
}

/*
 * SourceCompletion
 */

// Start is a wrapper around gtk_source_completion_start().
// since GtkSourceView 4, gtk_source_completion_show() has been renamed to
// gtk_source_completion_start().
// A List must be manually freed by either calling Free() or FreeFull()
func (v *SourceCompletion) Start(providers *glib.List, context *SourceCompletionContext) bool {
	return gobool(
		C.gtk_source_completion_start(v.native(), toCGList(providers), context.native()))
}

/*
 * SourceCompletionItem
 */

// SourceCompletionItemNew is a wrapper around gtk_source_completion_item_new().
func SourceCompletionItemNew() (*SourceCompletionItem, error) {
	c := C.gtk_source_completion_item_new()
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapSourceCompletionItem(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * SourceSearchContext
 */

// Forward is a wrapper around gtk_source_search_context_forward().
func (v *SourceSearchContext) Forward(iter *gtk.TextIter) (*gtk.TextIter, *gtk.TextIter, bool, bool) {

	start, end := new(gtk.TextIter), new(gtk.TextIter)
	var hasWrappedAround C.gboolean

	c := C.gtk_source_search_context_forward(
		v.native(),
		nativeTextIter(iter),
		nativeTextIter(start),
		nativeTextIter(end),
		&hasWrappedAround)

	return start, end, gobool(hasWrappedAround), gobool(c)
}

// ForwardFinish is a wrapper around gtk_source_search_context_forward_finish().
func (v *SourceSearchContext) ForwardFinish(result *glib.AsyncResult) (*gtk.TextIter, *gtk.TextIter, bool, bool, error) {

	var (
		start            = new(gtk.TextIter)
		end              = new(gtk.TextIter)
		hasWrappedAround C.gboolean
		cerr             *C.GError = nil
		err              error
	)

	c := C.gtk_source_search_context_forward_finish(
		v.native(),
		(*C.GAsyncResult)(unsafe.Pointer(result.Native())),
		nativeTextIter(start),
		nativeTextIter(end),
		&hasWrappedAround,
		&cerr)

	if cerr != nil {
		defer C.g_error_free(cerr)
		err = errors.New(goString(cerr.message))
	}

	return start, end, gobool(hasWrappedAround), gobool(c), err
}

// Backward is a wrapper around gtk_source_search_context_backward().
func (v *SourceSearchContext) Backward(iter *gtk.TextIter) (*gtk.TextIter, *gtk.TextIter, bool, bool) {

	start, end := new(gtk.TextIter), new(gtk.TextIter)
	var hasWrappedAround C.gboolean

	c := C.gtk_source_search_context_backward(
		v.native(),
		nativeTextIter(iter),
		nativeTextIter(start),
		nativeTextIter(end),
		&hasWrappedAround)

	return start, end, gobool(hasWrappedAround), gobool(c)
}

// BackwardFinish is a wrapper around gtk_source_search_context_backward_finish().
func (v *SourceSearchContext) BackwardFinish(result *glib.AsyncResult) (*gtk.TextIter, *gtk.TextIter, bool, bool, error) {

	var (
		start            = new(gtk.TextIter)
		end              = new(gtk.TextIter)
		hasWrappedAround C.gboolean
		cerr             *C.GError = nil
		err              error
	)

	c := C.gtk_source_search_context_backward_finish(
		v.native(),
		(*C.GAsyncResult)(unsafe.Pointer(result.Native())),
		nativeTextIter(start),
		nativeTextIter(end),
		&hasWrappedAround,
		&cerr)

	if cerr != nil {
		defer C.g_error_free(cerr)
		err = errors.New(goString(cerr.message))
	}

	return start, end, gobool(hasWrappedAround), gobool(c), err
}

// Replace is a wrapper around gtk_source_search_context_replace().
func (v *SourceSearchContext) Replace(start, end *gtk.TextIter, replace string) (bool, error) {
	var err *C.GError = nil
	cstr := C.CString(replace)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_source_search_context_replace(
		v.native(),
		nativeTextIter(start), nativeTextIter(end),
		(*C.gchar)(cstr),
		C.gint(len(replace)),
		&err)

	if err != nil {
		defer C.g_error_free(err)
		return gobool(c), errors.New(goString(err.message))
	}
	return gobool(c), nil
}
