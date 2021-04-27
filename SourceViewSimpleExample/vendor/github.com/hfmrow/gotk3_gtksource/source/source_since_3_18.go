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
// #include "source_since_3_18.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_source_map_get_type()), marshalSourceMap},
		{glib.Type(C.gtk_source_sort_flags_get_type()), marshalSourceSortFlags},
	}

	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceMap"] = wrapSourceMap
}

// SourceSortFlags is a representation of GTK's GtkSourceSortFlags.
type SourceSortFlags int

const (
	SOURCE_SORT_FLAGS_NONE              SourceSortFlags = C.GTK_SOURCE_SORT_FLAGS_NONE
	SOURCE_SORT_FLAGS_CASE_SENSITIVE    SourceSortFlags = C.GTK_SOURCE_SORT_FLAGS_CASE_SENSITIVE
	SOURCE_SORT_FLAGS_REVERSE_ORDER     SourceSortFlags = C.GTK_SOURCE_SORT_FLAGS_REVERSE_ORDER
	SOURCE_SORT_FLAGS_REMOVE_DUPLICATES SourceSortFlags = C.GTK_SOURCE_SORT_FLAGS_REMOVE_DUPLICATES
)

func marshalSourceSortFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceSortFlags(c), nil
}

/*
 * SourceView
 */

// SetSmartBackspace is a wrapper around gtk_source_view_set_smart_backspace().
func (v *SourceView) SetSmartBackspace(enable bool) {
	C.gtk_source_view_set_smart_backspace(v.native(), gbool(enable))
}

// GetSmartBackspace is a wrapper around gtk_source_view_get_smart_backspace().
func (v *SourceView) GetSmartBackspace() bool {
	return gobool(C.gtk_source_view_get_smart_backspace(v.native()))
}

/*
 * GtkSourceMap (full)
 */

// SourceMap is a representation of GTK's GtkSourceMap.
// Widget that displays a map for a specific GtkSourceView
type SourceMap struct {
	SourceView
}

// native returns a pointer to the underlying GtkSourceStyleSchemeManager.
func (v *SourceMap) native() *C.GtkSourceMap {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceMap(p)
}

func marshalSourceMap(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceMap(obj), nil
}

func wrapSourceMap(obj *glib.Object) *SourceMap {
	return &SourceMap{SourceView{gtk.TextView{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}}}
}

// SourceMapNew is a wrapper around gtk_source_map_new().
// Original documentation figure out return a GtkWidget
// but here, to be able to use SetView and GetView methods,
// we get a SourceMap object instead. To convert,
// use SourceMap.ToWidget()
func SourceMapNew() (*SourceMap, error) {
	c := C.gtk_source_map_new()
	if c == nil {
		return nil, nilPtrErr
	}
	e := wrapSourceMap(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SetView is a wrapper around gtk_source_map_set_view ().
func (v *SourceMap) SetView(view *SourceView) {
	C.gtk_source_map_set_view(v.native(), view.native())
}

// GetView is a wrapper around gtk_source_map_get_view ().
func (v *SourceMap) GetView() (*SourceView, error) {
	c := C.gtk_source_map_get_view(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceView(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * SourceBuffer
 */

// JoinLines is a wrapper around gtk_source_buffer_sort_lines().
func (v *SourceBuffer) SortLines(caseType SourceChangeCaseType, start, end *gtk.TextIter, flags SourceSortFlags, column int) {
	C.gtk_source_buffer_sort_lines(v.native(), nativeTextIter(start), nativeTextIter(end),
		C.GtkSourceSortFlags(flags), C.gint(column))
}

/*
 * SourceEncoding
 */

// GetDefaultCandidates is a wrapper around gtk_source_encoding_get_default_candidates().
// the returned list must be Free with (*glib.SList).Free().
func SourceEncodingGetDefaultCandidates() *glib.SList {
	c := (*C.GSList)(C.gtk_source_encoding_get_default_candidates())

	candidates := glib.WrapSList(uintptr(unsafe.Pointer(c)))
	if candidates == nil {
		return nil // empty.
	}
	return candidates
}

/*
 * SourceCompletionProposal
 */

// GetIcon is a wrapper around gtk_source_completion_proposal_get_icon().
func (v *SourceCompletionProposal) GetIcon() (*gdk.Pixbuf, error) {
	return toPixbuf(C.gtk_source_completion_proposal_get_icon(v.native()))
}

// GetIconName is a wrapper around gtk_source_completion_proposal_get_icon_name().
func (v *SourceCompletionProposal) GetIconName() string {
	return toGoStringFree(C.gtk_source_completion_proposal_get_icon_name(v.native()))
}
