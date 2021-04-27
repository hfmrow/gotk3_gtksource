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
// #include "source_since_3_16.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_source_style_scheme_chooser_button_get_type()), marshalSourceStyleSchemeChooserButton},
		{glib.Type(C.gtk_source_style_scheme_chooser_widget_get_type()), marshalSourceStyleSchemeChooserWidget},
		{glib.Type(C.gtk_source_background_pattern_type_get_type()), marshalSourceBackgroundPatternType},
	}

	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceStyleSchemeChooserButton"] = wrapSourceStyleSchemeChooserButton
	gtk.WrapMap["GtkSourceStyleSchemeChooserWidget"] = wrapSourceStyleSchemeChooserWidget
}

/*
 * GtkSourceStyleSchemeChooser (full)
 * Interface implemented by widgets for choosing style schemes
 */

// SourceStyleSchemeChooser is a representation of GTK's GtkSourceStyleSchemeChooser.
type SourceStyleSchemeChooser struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceStyleSchemeChooser.
func (v *SourceStyleSchemeChooser) native() *C.GtkSourceStyleSchemeChooser {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyleSchemeChooser(p)
}

func marshalSourceStyleSchemeChooser(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleSchemeChooser(obj), nil
}

func wrapSourceStyleSchemeChooser(obj *glib.Object) *SourceStyleSchemeChooser {
	return &SourceStyleSchemeChooser{obj}
}

// GetStyleScheme is a wrapper around gtk_source_style_scheme_chooser_get_style_scheme().
func (v *SourceStyleSchemeChooser) GetStyleScheme() (*SourceStyleScheme, error) {
	c := C.gtk_source_style_scheme_chooser_get_style_scheme(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceStyleScheme(glib.Take(unsafe.Pointer(c))), nil
}

// SetStyleScheme is a wrapper around gtk_source_style_scheme_chooser_set_style_scheme().
func (v *SourceStyleSchemeChooser) SetStyleScheme(scheme *SourceStyleScheme) {
	C.gtk_source_style_scheme_chooser_set_style_scheme(v.native(), scheme.native())
}

/*
 * GtkSourceStyleSchemeChooserWidget (full)
 * A widget for choosing style schemes
 */

// SourceStyleSchemeChooserWidget is a representation of GTK's GtkSourceStyleSchemeChooserWidget.
type SourceStyleSchemeChooserWidget struct {
	gtk.Bin

	SourceStyleSchemeChooser
}

// native returns a pointer to the underlying GtkSourceStyleSchemeChooserWidget.
func (v *SourceStyleSchemeChooserWidget) native() *C.GtkSourceStyleSchemeChooserWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyleSchemeChooserWidget(p)
}

func marshalSourceStyleSchemeChooserWidget(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleSchemeChooserWidget(obj), nil
}

func wrapSourceStyleSchemeChooserWidget(obj *glib.Object) *SourceStyleSchemeChooserWidget {
	c := wrapSourceStyleSchemeChooser(obj)

	return &SourceStyleSchemeChooserWidget{gtk.Bin{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}, *c}
}

// SourceStyleSchemeChooserWidgetNew is a wrapper around gtk_source_style_scheme_chooser_widget_new().
func SourceStyleSchemeChooserWidgetNew() (*SourceStyleSchemeChooserWidget, error) {
	c := C.gtk_source_style_scheme_chooser_widget_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := (glib.Take(unsafe.Pointer(c)))
	return wrapSourceStyleSchemeChooserWidget(obj), nil
}

/*
 * GtkSourceStyleSchemeChooserButton (full)
 * A button to launch a style scheme selection dialog
 */

// SourceStyleSchemeChooserButton is a representation of GTK's GtkSourceStyleSchemeChooserButton.
type SourceStyleSchemeChooserButton struct {
	gtk.Button
	SourceStyleSchemeChooser
}

// native returns a pointer to the underlying GtkSourceStyleSchemeChooserButton.
func (v *SourceStyleSchemeChooserButton) native() *C.GtkSourceStyleSchemeChooserButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyleSchemeChooserButton(p)
}

func marshalSourceStyleSchemeChooserButton(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleSchemeChooserButton(obj), nil
}

func wrapSourceStyleSchemeChooserButton(obj *glib.Object) *SourceStyleSchemeChooserButton {
	c := wrapSourceStyleSchemeChooser(obj)
	a := &gtk.Actionable{obj}
	return &SourceStyleSchemeChooserButton{gtk.Button{gtk.Bin{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}, a}, *c}
}

// SourceStyleSchemeChooserButtonNew is a wrapper around gtk_source_style_scheme_chooser_button_new().
func SourceStyleSchemeChooserButtonNew() (*SourceStyleSchemeChooserButton, error) {
	c := C.gtk_source_style_scheme_chooser_button_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := (glib.Take(unsafe.Pointer(c)))
	return wrapSourceStyleSchemeChooserButton(obj), nil
}

/*
 * SourceView
 */

// SourceBackgroundPatternType is a representation of GTK's GtkSourceBackgroundPatternType.
type SourceBackgroundPatternType int

const (
	SOURCE_BACKGROUND_PATTERN_TYPE_NONE SourceBackgroundPatternType = C.GTK_SOURCE_BACKGROUND_PATTERN_TYPE_NONE
	SOURCE_BACKGROUND_PATTERN_TYPE_GRID                             = C.GTK_SOURCE_BACKGROUND_PATTERN_TYPE_GRID
)

func marshalSourceBackgroundPatternType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceBackgroundPatternType(c), nil
}

// GetVersion same as GTK_SOURCE_CHECK_VERSION() macro. It is a personal
// implementation to check which version will be deployed in prod.
func SourceGetVersion() (int, int, int, bool) {
	major := C.gtk_source_get_major_version()
	minor := C.gtk_source_get_minor_version()
	micro := C.gtk_source_get_micro_version()
	return int(major),
		int(minor),
		int(micro),
		gobool(C.gtk_source_check_version(
			major,
			minor,
			micro))
}

// IndentLines is a wrapper around gtk_source_view_indent_lines().
func (v *SourceView) IndentLines(start, end *gtk.TextIter) {
	C.gtk_source_view_indent_lines(v.native(), nativeTextIter(start), nativeTextIter(end))
}

// UnIndentLines is a wrapper around gtk_source_view_unindent_lines().
func (v *SourceView) UnIndentLines(start, end *gtk.TextIter) {
	C.gtk_source_view_unindent_lines(v.native(), nativeTextIter(start), nativeTextIter(end))
}

// SetBackgroundPattern is a wrapper around gtk_source_view_set_background_pattern().
func (v *SourceView) SetBackgroundPattern(bp SourceBackgroundPatternType) {
	C.gtk_source_view_set_background_pattern(v.native(), C.GtkSourceBackgroundPatternType(bp))
}

// GetBackgroundPattern is a wrapper around gtk_source_view_get_background_pattern().
func (v *SourceView) GetBackgroundPattern() SourceBackgroundPatternType {
	return SourceBackgroundPatternType(C.gtk_source_view_get_background_pattern(v.native()))
}

/*
 * SourceBuffer
 */

// JoinLines is a wrapper around gtk_source_buffer_join_lines().
func (v *SourceBuffer) JoinLines(caseType SourceChangeCaseType, start, end *gtk.TextIter) {
	C.gtk_source_buffer_join_lines(v.native(), nativeTextIter(start), nativeTextIter(end))
}

/*
 * SourceSearchContext
 */

// GetMatchStyle is a wrapper around gtk_source_search_context_get_match_style().
func (v *SourceSearchContext) GetMatchStyle() (*SourceStyle, error) {
	c := C.gtk_source_search_context_get_match_style(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceStyle(glib.Take(unsafe.Pointer(c))), nil
}

// SetMatchStyle is a wrapper around gtk_source_search_context_set_match_style().
func (v *SourceSearchContext) SetMatchStyle(style *SourceStyle) {
	C.gtk_source_search_context_set_match_style(v.native(),
		style.native())
}
