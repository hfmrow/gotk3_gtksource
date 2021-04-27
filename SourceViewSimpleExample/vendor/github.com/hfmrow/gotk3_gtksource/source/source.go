// +build linux
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

// Limit > gtk_3_14 because libgtksourceview-3.0 start at gtk >= 3.16
// Latest gtksourceview_3_18 start with gtk_3_16
// Latest gtksourceview_4_0 start with gtk_3_22

// Same copyright and license as the rest of the files in this project

/*
	require:
			libgtksourceview-4-dev > gtk_3_18
		or	libgtksourceview-3.0-dev with: // #cgo pkg-config: gtksourceview-3.0 for < gtk_3_20
*/

package source

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gtk/gtk.h>
// #include <gtksourceview/gtksource.h>
// #include "source.go.h"
// #include "source_facility.go.h"
// #include "source_facility.gtk.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_source_buffer_get_type()), marshalSourceBuffer},
		{glib.Type(C.gtk_source_view_get_type()), marshalSourceView},
		{glib.Type(C.gtk_source_style_scheme_manager_get_type()), marshalSourceStyleSchemeManager},
		{glib.Type(C.gtk_source_style_scheme_get_type()), marshalSourceStyleScheme},
		{glib.Type(C.gtk_source_style_get_type()), marshalSourceStyle},
		{glib.Type(C.gtk_source_language_manager_get_type()), marshalSourceLanguageManager},
		{glib.Type(C.gtk_source_language_get_type()), marshalSourceLanguage},
		{glib.Type(C.gtk_source_mark_attributes_get_type()), marshalSourceMarkAttributes},
		{glib.Type(C.gtk_source_mark_get_type()), marshalSourceMark},
		{glib.Type(C.gtk_source_completion_get_type()), marshalSourceCompletion},
		{glib.Type(C.gtk_source_completion_context_get_type()), marshalSourceCompletionContext},
		{glib.Type(C.gtk_source_completion_proposal_get_type()), marshalSourceCompletionProposal},
		{glib.Type(C.gtk_source_completion_info_get_type()), marshalSourceCompletionInfo},
		{glib.Type(C.gtk_source_completion_words_get_type()), marshalSourceCompletionWords},
		{glib.Type(C.gtk_source_completion_item_get_type()), marshalSourceCompletionItem},
		{glib.Type(C.gtk_source_gutter_get_type()), marshalSourceGutter},
		{glib.Type(C.gtk_source_gutter_renderer_get_type()), marshalSourceGutterRenderer},
		{glib.Type(C.gtk_source_gutter_renderer_text_get_type()), marshalSourceGutterRendererText},
		{glib.Type(C.gtk_source_gutter_renderer_pixbuf_get_type()), marshalSourceGutterRendererPixbuf},
		{glib.Type(C.gtk_source_undo_manager_get_type()), marshalSourceUndoManager},
		{glib.Type(C.gtk_source_smart_home_end_type_get_type()), marshalSourceSmartHomeEndType},
		{glib.Type(C.gtk_source_view_gutter_position_get_type()), marshalSourceViewGutterPosition},
		{glib.Type(C.gtk_source_completion_error_get_type()), marshalSourceCompletionError},
		{glib.Type(C.gtk_source_completion_activation_get_type()), marshalSourceCompletionActivation},
		// {glib.Type(C.gtk_source_gutter_renderer_state_get_type()), marshalSourceGutterRendererState},
		{glib.Type(C.gtk_source_gutter_renderer_alignment_mode_get_type()), marshalSourceGutterRendererAlignmentMode},
		{glib.Type(C.gtk_source_bracket_match_type_get_type()), marshalSourceBracketMatchType},
	}

	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceBuffer"] = wrapSourceBuffer
	gtk.WrapMap["GtkSourceView"] = wrapSourceView
	gtk.WrapMap["GtkSourceStyleSchemeManager"] = wrapSourceStyleSchemeManager
	gtk.WrapMap["GtkSourceStyleScheme"] = wrapSourceStyleScheme
	gtk.WrapMap["GtkSourceStyle"] = wrapSourceStyle
	gtk.WrapMap["GtkSourceLanguageManager"] = wrapSourceLanguageManager
	gtk.WrapMap["GtkSourceLanguage"] = wrapSourceLanguage
	gtk.WrapMap["GtkSourceMarkAttributes"] = wrapSourceMarkAttributes
	gtk.WrapMap["GtkSourceMark"] = wrapSourceMark
	gtk.WrapMap["GtkSourceCompletion"] = wrapSourceCompletion
	gtk.WrapMap["GtkSourceCompletionContext"] = wrapSourceCompletionContext
	gtk.WrapMap["GtkSourceCompletionProposal"] = wrapSourceCompletionProposal
	gtk.WrapMap["GtkSourceCompletionInfo"] = wrapSourceCompletionInfo
	gtk.WrapMap["GtkSourceCompletionWords"] = wrapSourceCompletionWords
	gtk.WrapMap["GtkSourceCompletionItem"] = wrapSourceCompletionItem
	gtk.WrapMap["GtkSourceGutter"] = wrapSourceGutter
	gtk.WrapMap["GtkSourceGutterRenderer"] = wrapSourceGutterRenderer
	gtk.WrapMap["GtkSourceGutterRendererText"] = wrapSourceGutterRendererText
	gtk.WrapMap["GtkSourceGutterRendererPixbuf"] = wrapSourceGutterRendererPixbuf
	gtk.WrapMap["GtkSourceUndoManager"] = wrapSourceUndoManager
}

// SourceSmartHomeEndType is a representation of GTK's GtkSourceSmartHomeEndType.
type SourceSmartHomeEndType int

const (
	SOURCE_SMART_HOME_END_DISABLED SourceSmartHomeEndType = C.GTK_SOURCE_SMART_HOME_END_DISABLED
	SOURCE_SMART_HOME_END_BEFORE   SourceSmartHomeEndType = C.GTK_SOURCE_SMART_HOME_END_BEFORE
	SOURCE_SMART_HOME_END_AFTER    SourceSmartHomeEndType = C.GTK_SOURCE_SMART_HOME_END_AFTER
	SOURCE_SMART_HOME_END_ALWAYS   SourceSmartHomeEndType = C.GTK_SOURCE_SMART_HOME_END_ALWAYS
)

func marshalSourceSmartHomeEndType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceSmartHomeEndType(c), nil
}

// SourceViewGutterPosition is a representation of GTK's GtkSourceViewGutterPosition.
type SourceViewGutterPosition int

const (
	SOURCE_VIEW_GUTTER_POSITION_LINES SourceViewGutterPosition = C.GTK_SOURCE_VIEW_GUTTER_POSITION_LINES
	SOURCE_VIEW_GUTTER_POSITION_MARKS SourceViewGutterPosition = C.GTK_SOURCE_VIEW_GUTTER_POSITION_MARKS
)

func marshalSourceViewGutterPosition(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceViewGutterPosition(c), nil
}

// SourceCompletionError is a representation of GTK's GtkSourceCompletionError.
type SourceCompletionError int

const (
	SOURCE_COMPLETION_ERROR_ALREADY_BOUND SourceCompletionError = C.GTK_SOURCE_COMPLETION_ERROR_ALREADY_BOUND
	SOURCE_COMPLETION_ERROR_NOT_BOUND     SourceCompletionError = C.GTK_SOURCE_COMPLETION_ERROR_NOT_BOUND
)

func marshalSourceCompletionError(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceCompletionError(c), nil
}

// SourceCompletionActivation is a representation of GTK's GtkSourceCompletionActivation.
type SourceCompletionActivation int

const (
	SOURCE_COMPLETION_ACTIVATION_NONE           SourceCompletionActivation = C.GTK_SOURCE_COMPLETION_ACTIVATION_NONE
	SOURCE_COMPLETION_ACTIVATION_INTERACTIVE    SourceCompletionActivation = C.GTK_SOURCE_COMPLETION_ACTIVATION_INTERACTIVE
	SOURCE_COMPLETION_ACTIVATION_USER_REQUESTED SourceCompletionActivation = C.GTK_SOURCE_COMPLETION_ACTIVATION_USER_REQUESTED
)

func marshalSourceCompletionActivation(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceCompletionActivation(c), nil
}

// SourceGutterRendererState is a representation of GTK's GtkSourceGutterRendererState.
type SourceGutterRendererState int

const (
	SOURCE_GUTTER_RENDERER_STATE_NORMAL   SourceGutterRendererState = C.GTK_SOURCE_GUTTER_RENDERER_STATE_NORMAL
	SOURCE_GUTTER_RENDERER_STATE_CURSOR   SourceGutterRendererState = C.GTK_SOURCE_GUTTER_RENDERER_STATE_CURSOR
	SOURCE_GUTTER_RENDERER_STATE_PRELIT   SourceGutterRendererState = C.GTK_SOURCE_GUTTER_RENDERER_STATE_PRELIT
	SOURCE_GUTTER_RENDERER_STATE_SELECTED SourceGutterRendererState = C.GTK_SOURCE_GUTTER_RENDERER_STATE_SELECTED
)

// NOTE When declared to 'glib.TypeMarshaler', i have some GTK error when a signal try
// To convert this value ...
// func marshalSourceGutterRendererState(p uintptr) (interface{}, error) {
// 	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
// 	return SourceGutterRendererState(c), nil
// }

// SourceGutterRendererAlignmentMode is a representation of GTK's GtkSourceGutterRendererAlignmentMode.
type SourceGutterRendererAlignmentMode int

const (
	SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_CELL  SourceGutterRendererAlignmentMode = C.GTK_SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_CELL
	SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_FIRST SourceGutterRendererAlignmentMode = C.GTK_SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_FIRST
	SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_LAST  SourceGutterRendererAlignmentMode = C.GTK_SOURCE_GUTTER_RENDERER_ALIGNMENT_MODE_LAST
)

func marshalSourceGutterRendererAlignmentMode(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceGutterRendererAlignmentMode(c), nil
}

// SourceBracketMatchType is a representation of GTK's GtkSourceBracketMatchType.
type SourceBracketMatchType int

const (
	SOURCE_BRACKET_MATCH_NONE         SourceBracketMatchType = C.GTK_SOURCE_BRACKET_MATCH_NONE
	SOURCE_BRACKET_MATCH_OUT_OF_RANGE SourceBracketMatchType = C.GTK_SOURCE_BRACKET_MATCH_OUT_OF_RANGE
	SOURCE_BRACKET_MATCH_NOT_FOUND    SourceBracketMatchType = C.GTK_SOURCE_BRACKET_MATCH_NOT_FOUND
	SOURCE_BRACKET_MATCH_FOUND        SourceBracketMatchType = C.GTK_SOURCE_BRACKET_MATCH_FOUND
)

func marshalSourceBracketMatchType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceBracketMatchType(c), nil
}

/*
 * GtkSourceBuffer (full)
 */

// SourceBuffer is a representation of GTK's GtkSourceBuffer.
// Subclass of GtkTextBuffer
type SourceBuffer struct {
	gtk.TextBuffer
}

func (v *SourceBuffer) toGtkTextBuffer() *C.GtkTextBuffer {
	if v == nil {
		return nil
	}
	return C.toGtkTextBuffer(unsafe.Pointer(v.GObject))
}

// native returns a pointer to the underlying GtkSourceBuffer.
func (v *SourceBuffer) native() *C.GtkSourceBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceBuffer(p)
}

func marshalSourceBuffer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceBuffer(obj), nil
}

func wrapSourceBuffer(obj *glib.Object) *SourceBuffer {
	return &SourceBuffer{gtk.TextBuffer{obj}}
}

// ToTextBuffer returns a *TextBuffer
// Some derived methods accept only a TextBuffer source as pointer
func (v *SourceBuffer) ToTextBuffer() *gtk.TextBuffer {
	if v == nil {
		return nil
	}
	return &gtk.TextBuffer{v.Object}
}

// Serialize overriding gtk_text_buffer_serialize() to use with GtkSourceBuffer
func (v *SourceBuffer) Serialize(content *SourceBuffer, format gdk.Atom, start, end *gtk.TextIter) string {
	var length = new(C.gsize)
	ptr := C.gtk_text_buffer_serialize(
		v.toGtkTextBuffer(),
		content.toGtkTextBuffer(),
		C.GdkAtom(unsafe.Pointer(format)),
		nativeTextIter(start),
		nativeTextIter(end),
		length)
	return C.GoStringN((*C.char)(unsafe.Pointer(ptr)), (C.int)(*length))
}

// Deserialize overriding gtk_text_buffer_deserialize() to use with GtkSourceBuffer
func (v *SourceBuffer) Deserialize(content *SourceBuffer, format gdk.Atom, iter *gtk.TextIter, data []byte) (ok bool, err error) {
	var length = (C.gsize)(len(data))
	var cerr *C.GError = nil
	cbool := C.gtk_text_buffer_deserialize(v.toGtkTextBuffer(), content.toGtkTextBuffer(), C.GdkAtom(unsafe.Pointer(format)),
		nativeTextIter(iter), (*C.guint8)(unsafe.Pointer(&data[0])), length, &cerr)
	if !gobool(cbool) {
		defer C.g_error_free(cerr)
		return false, errors.New(goString(cerr.message))
	}
	return gobool(cbool), nil
}

// SourceBufferNew is a wrapper around gtk_source_buffer_new().
func SourceBufferNew(table *gtk.TextTagTable) (*SourceBuffer, error) {
	c := C.gtk_source_buffer_new(nativeTextTagTable(table))
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceBuffer(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SourceBufferNew is a wrapper around gtk_source_buffer_new_with_language().
func SourceBufferNewWithLanguage(language *SourceLanguage) (*SourceBuffer, error) {
	c := C.gtk_source_buffer_new_with_language(language.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceBuffer(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SetLanguage is a wrapper around gtk_source_buffer_set_language().
func (v *SourceBuffer) SetLanguage(language *SourceLanguage) {
	C.gtk_source_buffer_set_language(v.native(), language.native())
}

// GetLanguage is a wrapper around gtk_source_buffer_get_language().
func (v *SourceBuffer) GetLanguage() (*SourceLanguage, error) {
	c := C.gtk_source_buffer_get_language(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceLanguage(glib.Take(unsafe.Pointer(c))), nil
}

// SetStyleSheme is a wrapper around gtk_source_buffer_set_style_scheme().
func (v *SourceBuffer) SetStyleSheme(scheme *SourceStyleScheme) {
	C.gtk_source_buffer_set_style_scheme(v.native(), scheme.native())
}

// GetStyleSheme is a wrapper around gtk_source_buffer_get_style_scheme().
func (v *SourceBuffer) GetStyleSheme() (*SourceStyleScheme, error) {
	c := C.gtk_source_buffer_get_style_scheme(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceStyleScheme(glib.Take(unsafe.Pointer(c))), nil
}

// SetHighlightSyntax is a wrapper around gtk_source_buffer_set_highlight_syntax().
func (v *SourceBuffer) SetHighlightSyntax(highlight bool) {
	C.gtk_source_buffer_set_highlight_syntax(v.native(), gbool(highlight))
}

// GetHighlightSyntax is a wrapper around gtk_source_buffer_get_highlight_syntax().
func (v *SourceBuffer) GetHighlightSyntax() bool {
	return gobool(C.gtk_source_buffer_get_highlight_syntax(v.native()))
}

// SetHighlightMatchingBrackets is a wrapper around gtk_source_buffer_set_highlight_matching_brackets().
func (v *SourceBuffer) SetHighlightMatchingBrackets(highlight bool) {
	C.gtk_source_buffer_set_highlight_matching_brackets(v.native(), gbool(highlight))
}

// GetHighlightMatchingBrackets is a wrapper around gtk_source_buffer_get_highlight_matching_brackets().
func (v *SourceBuffer) GetHighlightMatchingBrackets() bool {
	return gobool(C.gtk_source_buffer_get_highlight_matching_brackets(v.native()))
}

// EnsureHightlight is a wrapper around gtk_source_buffer_ensure_highlight().
func (v *SourceBuffer) EnsureHightlight(start, end *gtk.TextIter) {
	C.gtk_source_buffer_ensure_highlight(v.native(), nativeTextIter(start), nativeTextIter(end))
}

// Undo is a wrapper around gtk_source_buffer_undo().
func (v *SourceBuffer) Undo() {
	C.gtk_source_buffer_undo(v.native())
}

// Redo is a wrapper around gtk_source_buffer_redo().
func (v *SourceBuffer) Redo() {
	C.gtk_source_buffer_redo(v.native())
}

// CanUndo is a wrapper around gtk_source_buffer_can_undo().
func (v *SourceBuffer) CanUndo() bool {
	return gobool(C.gtk_source_buffer_can_undo(v.native()))
}

// CanRedo is a wrapper around gtk_source_buffer_can_redo().
func (v *SourceBuffer) CanRedo() bool {
	return gobool(C.gtk_source_buffer_can_redo(v.native()))
}

// BeginNotUndoableAction is a wrapper around gtk_source_buffer_begin_not_undoable_action().
func (v *SourceBuffer) BeginNotUndoableAction() {
	C.gtk_source_buffer_begin_not_undoable_action(v.native())
}

// EndNotUndoableAction is a wrapper around gtk_source_buffer_end_not_undoable_action().
func (v *SourceBuffer) EndNotUndoableAction() {
	C.gtk_source_buffer_end_not_undoable_action(v.native())
}

// GetMaxUndoLevels is a wrapper around gtk_source_buffer_get_max_undo_levels().
func (v *SourceBuffer) GetMaxUndoLevels() int {
	return int(C.gtk_source_buffer_get_max_undo_levels(v.native()))
}

// SetMaxUndoLevels is a wrapper around gtk_source_buffer_set_max_undo_levels().
func (v *SourceBuffer) SetMaxUndoLevels(maxLevels int) {
	C.gtk_source_buffer_set_max_undo_levels(v.native(), C.gint(maxLevels))
}

// GetUndoManager is a wrapper around gtk_source_buffer_get_undo_manager().
func (v *SourceBuffer) GetUndoManager() (*SourceUndoManager, error) {
	c := C.gtk_source_buffer_get_undo_manager(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceUndoManager(glib.Take(unsafe.Pointer(c))), nil
}

// SetUndoManager is a wrapper around gtk_source_buffer_set_undo_manager().
func (v *SourceBuffer) SetUndoManager(manager *SourceUndoManager) {
	C.gtk_source_buffer_set_undo_manager(v.native(), manager.native())
}

// HasContextClass is a wrapper around gtk_source_buffer_iter_has_context_class().
func (v *SourceBuffer) HasContextClass(iter *gtk.TextIter, contextClass string) bool {
	cstr := C.CString(contextClass)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(
		C.gtk_source_buffer_iter_has_context_class(
			v.native(),
			nativeTextIter(iter),
			(*C.gchar)(cstr)))
}

// GetContextClassAtIter is a wrapper around gtk_source_buffer_get_context_classes_at_iter().
func (v *SourceBuffer) GetContextClassAtIter(iter *gtk.TextIter) []string {
	return toGoStringArray(C.gtk_source_buffer_get_context_classes_at_iter(v.native(), nativeTextIter(iter)))
}

// IterForwardToContextClassToggle is a wrapper around gtk_source_buffer_iter_forward_to_context_class_toggle().
func (v *SourceBuffer) IterForwardToContextClassToggle(iter *gtk.TextIter, contextClass string) bool {
	cstr := C.CString(contextClass)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(
		C.gtk_source_buffer_iter_forward_to_context_class_toggle(
			v.native(),
			nativeTextIter(iter),
			(*C.gchar)(cstr)))
}

// IterBackwardToContextClassToggle is a wrapper around gtk_source_buffer_iter_backward_to_context_class_toggle().
func (v *SourceBuffer) IterBackwardToContextClassToggle(iter *gtk.TextIter, contextClass string) bool {
	cstr := C.CString(contextClass)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(
		C.gtk_source_buffer_iter_backward_to_context_class_toggle(
			v.native(),
			nativeTextIter(iter),
			(*C.gchar)(cstr)))
}

// CreateSourceMark is a wrapper around gtk_source_buffer_create_source_mark().
func (v *SourceBuffer) CreateSourceMark(name, category string, where *gtk.TextIter) (*SourceMark, error) {
	var cstr *C.gchar = nil
	if len(name) != 0 {
		cstr = C.CString(name)
		defer C.free(unsafe.Pointer(cstr))
	}
	cstr1 := C.CString(category)
	defer C.free(unsafe.Pointer(cstr1))

	c := C.gtk_source_buffer_create_source_mark(
		v.native(),
		(*C.gchar)(cstr),
		(*C.gchar)(cstr1),
		nativeTextIter(where))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceMark(glib.Take(unsafe.Pointer(c))), nil
}

// ForwardIterToSourceMark is a wrapper around gtk_source_buffer_forward_iter_to_source_mark().
func (v *SourceBuffer) ForwardIterToSourceMark(iter *gtk.TextIter, category string) bool {
	var cstr *C.gchar = nil
	if len(category) != 0 {
		cstr := C.CString(category)
		defer C.free(unsafe.Pointer(cstr))
	}
	return gobool(
		C.gtk_source_buffer_forward_iter_to_source_mark(
			v.native(),
			nativeTextIter(iter),
			(*C.gchar)(cstr)))
}

// BackwardIterToSourceMark is a wrapper around gtk_source_buffer_backward_iter_to_source_mark().
func (v *SourceBuffer) BackwardIterToSourceMark(iter *gtk.TextIter, category string) bool {
	var cstr *C.gchar = nil
	if len(category) != 0 {
		cstr := C.CString(category)
		defer C.free(unsafe.Pointer(cstr))
	}

	return gobool(
		C.gtk_source_buffer_backward_iter_to_source_mark(
			v.native(),
			nativeTextIter(iter),
			(*C.gchar)(cstr)))
}

// GetSourceMarksAtLine is a wrapper around gtk_source_buffer_get_source_marks_at_line().
func (v *SourceBuffer) GetSourceMarksAtLine(line int, category string) *glib.List {
	var cstr *C.gchar = nil
	if len(category) != 0 {
		cstr := C.CString(category)
		defer C.free(unsafe.Pointer(cstr))
	}
	clist := C.gtk_source_buffer_get_source_marks_at_line(
		v.native(),
		C.gint(line),
		(*C.gchar)(cstr))

	if clist == nil {
		return nil
	}
	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapSourceMark(glib.Take(ptr))
	})
	runtime.SetFinalizer(glist, func(l *glib.List) {
		l.Free()
	})

	return glist
}

// GetSourceMarksAtIter is a wrapper around gtk_source_buffer_get_source_marks_at_iter().
func (v *SourceBuffer) GetSourceMarksAtIter(iter *gtk.TextIter, category string) *glib.List {
	var cstr *C.gchar = nil
	if len(category) != 0 {
		cstr := C.CString(category)
		defer C.free(unsafe.Pointer(cstr))
	}
	clist := C.gtk_source_buffer_get_source_marks_at_iter(
		v.native(),
		nativeTextIter(iter),
		(*C.gchar)(cstr))

	if clist == nil {
		return nil
	}
	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapSourceMark(glib.Take(ptr))
	})
	runtime.SetFinalizer(glist, func(l *glib.List) {
		l.Free()
	})

	return glist
}

// RemoveSourceMarks is a wrapper around gtk_source_buffer_remove_source_marks().
func (v *SourceBuffer) RemoveSourceMarks(start, end *gtk.TextIter, category string) {
	var cstr *C.gchar = nil
	if len(category) != 0 {
		cstr := C.CString(category)
		defer C.free(unsafe.Pointer(cstr))
	}
	C.gtk_source_buffer_remove_source_marks(
		v.native(), nativeTextIter(start), nativeTextIter(end), (*C.gchar)(cstr))
}

/*
 * GtkSourceView (full)
 */

// SourceView is a representation of GTK's GtkSourceView
// Subclass of GtkTextView
type SourceView struct {
	gtk.TextView
}

func (v *SourceView) toGtkTextView() *C.GtkTextView {
	if v == nil {
		return nil
	}
	return C.toGtkTextView(unsafe.Pointer(v.GObject))
}

// native returns a pointer to the underlying GtkSourceView.
func (v *SourceView) native() *C.GtkSourceView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceView(p)
}

func marshalSourceView(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceView(obj), nil
}

func wrapSourceView(obj *glib.Object) *SourceView {
	return &SourceView{gtk.TextView{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}}
}

// ToTextView returns a *TextView
// Some derived methods accept only a TextView source as pointer
func (v *SourceView) ToTextView() *gtk.TextView {
	if v == nil {
		return nil
	}
	return &gtk.TextView{gtk.Container{gtk.Widget{glib.InitiallyUnowned{v.Object}}}}
}

// SourceViewNew is a wrapper around gtk_source_view_new ().
func SourceViewNew() (*SourceView, error) {
	c := C.gtk_source_view_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceView(glib.Take(unsafe.Pointer(c))), nil
}

// SourceViewNewWithBuffer is a wrapper around gtk_source_view_new_with_buffer().
func SourceViewNewWithBuffer(buf *SourceBuffer) (*SourceView, error) {
	c := C.gtk_source_view_new_with_buffer(buf.native())
	return wrapSourceView(glib.Take(unsafe.Pointer(c))), nil
}

// GetBuffer overriding gtk_text_view_get_buffer() to get a GtkSourceBuffer
// from a GtkSourceView
func (v *SourceView) GetBuffer() (*SourceBuffer, error) {
	c := C.gtk_text_view_get_buffer(v.toGtkTextView())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceBuffer(glib.Take(unsafe.Pointer(c))), nil
}

// SetBuffer overriding gtk_text_view_set_buffer() to set a GtkSourceBuffer
// to a GtkSourceView
func (v *SourceView) SetBuffer(buffer *SourceBuffer) {
	C.gtk_text_view_set_buffer(v.toGtkTextView(), buffer.toGtkTextBuffer())
}

// GetCompletion is a wrapper around gtk_source_view_get_completion().
func (v *SourceView) GetCompletion() (*SourceCompletion, error) {
	c := C.gtk_source_view_get_completion(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceCompletion(glib.Take(unsafe.Pointer(c))), nil
}

// GetGutter is a wrapper around gtk_source_view_get_gutter().
func (v *SourceView) GetGutter(windowType gtk.TextWindowType) (*SourceGutter, error) {
	c := C.gtk_source_view_get_gutter(v.native(), C.GtkTextWindowType(windowType))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceGutter(glib.Take(unsafe.Pointer(c))), nil
}

// SetShowLineNumbers is a wrapper around gtk_source_view_set_show_line_numbers().
func (v *SourceView) SetShowLineNumbers(show bool) {
	C.gtk_source_view_set_show_line_numbers(v.native(), gbool(show))
}

// GetShowLineNumbers is a wrapper around gtk_source_view_get_show_line_numbers().
func (v *SourceView) GetShowLineNumbers() bool {
	return gobool(C.gtk_source_view_get_show_line_numbers(v.native()))
}

// SetHighlightCurrentLine is a wrapper around gtk_source_view_set_highlight_current_line().
func (v *SourceView) SetHighlightCurrentLine(highlight bool) {
	C.gtk_source_view_set_highlight_current_line(v.native(), gbool(highlight))
}

// GetHighlightCurrentLine is a wrapper around gtk_source_view_get_highlight_current_line().
func (v *SourceView) GetHighlightCurrentLine() bool {
	return gobool(C.gtk_source_view_get_highlight_current_line(v.native()))
}

// SetShowRightMargin is a wrapper around gtk_source_view_set_show_right_margin().
func (v *SourceView) SetShowRightMargin(show bool) {
	C.gtk_source_view_set_show_right_margin(v.native(), gbool(show))
}

// GetShowRightMargin is a wrapper around gtk_source_view_get_show_right_margin().
func (v *SourceView) GetShowRightMargin() bool {
	return gobool(C.gtk_source_view_get_show_right_margin(v.native()))
}

// SetRightMarginPosition is a wrapper around gtk_source_view_set_right_margin_position().
func (v *SourceView) SetRightMarginPosition(pos uint) {
	C.gtk_source_view_set_right_margin_position(v.native(), C.guint(pos))
}

// GetRightMarginPosition is a wrapper around gtk_source_view_get_right_margin_position().
func (v *SourceView) GetRightMarginPosition() uint {
	return uint(C.gtk_source_view_get_right_margin_position(v.native()))
}

// SetAutoIndent is a wrapper around gtk_source_view_set_auto_indent().
func (v *SourceView) SetAutoIndent(enable bool) {
	C.gtk_source_view_set_auto_indent(v.native(), gbool(enable))
}

// GetAutoIndent is a wrapper around gtk_source_view_get_auto_indent().
func (v *SourceView) GetAutoIndent() bool {
	return gobool(C.gtk_source_view_get_auto_indent(v.native()))
}

// SetIndentOnTab is a wrapper around gtk_source_view_set_indent_on_tab().
func (v *SourceView) SetIndentOnTab(enable bool) {
	C.gtk_source_view_set_indent_on_tab(v.native(), gbool(enable))
}

// GetIndentOnTab is a wrapper around gtk_source_view_get_indent_on_tab().
func (v *SourceView) GetIndentOnTab() bool {
	return gobool(C.gtk_source_view_get_indent_on_tab(v.native()))
}

// SetTabWidth is a wrapper around gtk_source_view_set_tab_width().
func (v *SourceView) SetTabWidth(width uint) {
	C.gtk_source_view_set_tab_width(v.native(), C.guint(width))
}

// GetTabWidth is a wrapper around gtk_source_view_get_tab_width().
func (v *SourceView) GetTabWidth() uint {
	return uint(C.gtk_source_view_get_tab_width(v.native()))
}

// SetIndentWidth is a wrapper around gtk_source_view_set_indent_width().
func (v *SourceView) SetIndentWidth(width int) {
	C.gtk_source_view_set_indent_width(v.native(), C.gint(width))
}

// GetIndentWidth is a wrapper around gtk_source_view_get_indent_width().
func (v *SourceView) GetIndentWidth() int {
	return int(C.gtk_source_view_get_indent_width(v.native()))
}

// SetInsertSpacesInsteadOfTabs is a wrapper around gtk_source_view_set_insert_spaces_instead_of_tabs().
func (v *SourceView) SetInsertSpacesInsteadOfTabs(enable bool) {
	C.gtk_source_view_set_insert_spaces_instead_of_tabs(v.native(), gbool(enable))
}

// GetInsertSpacesInsteadOfTabs is a wrapper around gtk_source_view_get_insert_spaces_instead_of_tabs().
func (v *SourceView) GetInsertSpacesInsteadOfTabs() bool {
	return gobool(C.gtk_source_view_get_insert_spaces_instead_of_tabs(v.native()))
}

// GetVisualColumn is a wrapper around gtk_source_view_get_visual_column().
func (v *SourceView) GetVisualColumn(iter *gtk.TextIter) uint {
	return uint(C.gtk_source_view_get_visual_column(v.native(), nativeTextIter(iter)))
}

// SetSmartHomeEnd is a wrapper around gtk_source_view_set_smart_home_end().
func (v *SourceView) SetSmartHomeEnd(s SourceSmartHomeEndType) {
	C.gtk_source_view_set_smart_home_end(v.native(), C.GtkSourceSmartHomeEndType(s))
}

// GetSmartHomeEnd is a wrapper around gtk_source_view_get_smart_home_end().
func (v *SourceView) GetSmartHomeEnd() SourceSmartHomeEndType {
	return SourceSmartHomeEndType(C.gtk_source_view_get_smart_home_end(v.native()))
}

// SetMarkAttributes is a wrapper around gtk_source_view_set_mark_attributes().
func (v *SourceView) SetMarkAttributes(category string, attributes *SourceMarkAttributes, priority int) {
	cstr := C.CString(category)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_view_set_mark_attributes(
		v.native(),
		(*C.gchar)(cstr),
		attributes.native(),
		C.gint(priority))
}

// GetMarkAttributes is a wrapper around gtk_source_view_get_mark_attributes().
func (v *SourceView) GetMarkAttributes(category string) (*SourceMarkAttributes, int, error) {
	var priority int
	var cpriority *C.gint
	defer C.free(unsafe.Pointer(cpriority))

	cstr := C.CString(category)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_source_view_get_mark_attributes(
		v.native(),
		(*C.gchar)(cstr),
		cpriority)
	if c == nil {
		return nil, priority, nilPtrErr
	}

	if cpriority != nil {
		priority = int(*((*C.gint)(unsafe.Pointer(cpriority))))
	}

	return wrapSourceMarkAttributes(glib.Take(unsafe.Pointer(c))), priority, nil
}

// SetShowLineMarks is a wrapper around gtk_source_view_set_show_line_marks().
func (v *SourceView) SetShowLineMarks(show bool) {
	C.gtk_source_view_set_show_line_marks(v.native(), gbool(show))
}

// GetShowLineMarks is a wrapper around gtk_source_view_get_show_line_marks().
func (v *SourceView) GetShowLineMarks() bool {
	return gobool(C.gtk_source_view_get_show_line_marks(v.native()))
}

/*
 * GtkSourceLanguageManager (full)
 */

// SourceLanguageManager is a representation of GTK's GtkSourceLanguageManager.
// Provides access to GtkSourceLanguages
type SourceLanguageManager struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceLanguage.
func (v *SourceLanguageManager) native() *C.GtkSourceLanguageManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceLanguageManager(p)
}

func marshalSourceLanguageManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceLanguageManager(obj), nil
}

func wrapSourceLanguageManager(obj *glib.Object) *SourceLanguageManager {
	return &SourceLanguageManager{obj}
}

// SourceLanguageManagerNew is a wrapper around gtk_source_language_manager_new ().
func SourceLanguageManagerNew() (*SourceLanguageManager, error) {
	c := C.gtk_source_language_manager_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceLanguageManager(glib.Take(unsafe.Pointer(c))), nil
}

// SourceLanguageManagerGetDefault is a wrapper around gtk_source_language_manager_get_default ().
func SourceLanguageManagerGetDefault() (*SourceLanguageManager, error) {
	c := C.gtk_source_language_manager_get_default()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceLanguageManager(glib.Take(unsafe.Pointer(c))), nil
}

// SetSearchPath is a wrapper around gtk_source_language_manager_set_search_path().
// At the moment this function can be called only before the language files are
// loaded for the first time. In practice to set a custom search path for a
// GtkSourceLanguageManager, you have to call this function right after creating it.
func (v *SourceLanguageManager) SetSearchPath(dirs []string) {
	cArray := toCgcharArray(dirs)
	defer C.destroy_strings(cArray)

	C.gtk_source_language_manager_set_search_path(v.native(), cArray)
}

// GetSearchPath is a wrapper around gtk_source_language_manager_get_search_path().
func (v *SourceLanguageManager) GetSearchPath() []string {
	return toGoStringArrayNoFree(C.gtk_source_language_manager_get_search_path(v.native()))
}

// GetLanguageIds is a wrapper around gtk_source_language_manager_get_language_ids().
func (v *SourceLanguageManager) GetLanguageIds() []string {
	return toGoStringArrayNoFree(C.gtk_source_language_manager_get_language_ids(v.native()))
}

// GetLanguage is a wrapper around gtk_source_language_manager_get_language().
func (v *SourceLanguageManager) GetLanguage(id string) (*SourceLanguage, error) {
	cstr := C.CString(id)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_source_language_manager_get_language(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceLanguage(glib.Take(unsafe.Pointer(c))), nil
}

// GetGuessLanguage is a wrapper around gtk_source_language_manager_guess_language().
func (v *SourceLanguageManager) GetGuessLanguage(filename, contentType string) (*SourceLanguage, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	cstr1 := C.CString(contentType)
	defer C.free(unsafe.Pointer(cstr1))
	c := C.gtk_source_language_manager_guess_language(v.native(), (*C.gchar)(cstr), (*C.gchar)(cstr1))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceLanguage(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkSourceLanguage (full)
 */

// SourceLanguage is a representation of GTK's GtkSourceLanguage.
type SourceLanguage struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceLanguage.
func (v *SourceLanguage) native() *C.GtkSourceLanguage {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceLanguage(p)
}

func marshalSourceLanguage(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceLanguage(obj), nil
}

func wrapSourceLanguage(obj *glib.Object) *SourceLanguage {
	return &SourceLanguage{obj}
}

// GetId is a wrapper around gtk_source_language_get_id().
func (v *SourceLanguage) GetId() string {
	return goString(C.gtk_source_language_get_id(v.native()))
}

// GetName is a wrapper around gtk_source_language_get_name().
func (v *SourceLanguage) GetName() string {
	return goString(C.gtk_source_language_get_name(v.native()))
}

// GetSection is a wrapper around gtk_source_language_get_section().
func (v *SourceLanguage) GetSection() string {
	return goString(C.gtk_source_language_get_section(v.native()))
}

// GetHidden is a wrapper around gtk_source_language_get_hidden().
func (v *SourceLanguage) GetHidden() bool {
	return gobool(C.gtk_source_language_get_hidden(v.native()))
}

// GetMetadata is a wrapper around gtk_source_language_get_metadata().
func (v *SourceLanguage) GetMetadata(name string) string {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	return goString(C.gtk_source_language_get_metadata(v.native(), (*C.gchar)(cstr)))
}

// GetMimeTypes is a wrapper around gtk_source_language_get_mime_types().
func (v *SourceLanguage) GetMimeTypes() []string {
	return toGoStringArray(C.gtk_source_language_get_mime_types(v.native()))
}

// GetGlobs is a wrapper around gtk_source_language_get_globs().
func (v *SourceLanguage) GetGlobs() []string {
	return toGoStringArray(C.gtk_source_language_get_globs(v.native()))
}

// GetStyleName is a wrapper around gtk_source_language_get_style_name().
func (v *SourceLanguage) GetStyleName(styleId string) string {
	cstr := C.CString(styleId)
	defer C.free(unsafe.Pointer(cstr))

	return goString(C.gtk_source_language_get_style_name(v.native(), (*C.gchar)(cstr)))
}

// GetStyleIds is a wrapper around gtk_source_language_get_style_ids().
func (v *SourceLanguage) GetStyleIds() []string {
	return toGoStringArray(C.gtk_source_language_get_style_ids(v.native()))
}

// GetStyleFallback is a wrapper around gtk_source_language_get_style_fallback().
func (v *SourceLanguage) GetStyleFallback(styleId string) string {
	cstr := C.CString(styleId)
	defer C.free(unsafe.Pointer(cstr))

	return goString(C.gtk_source_language_get_style_fallback(v.native(), (*C.gchar)(cstr)))
}

/*
 * GtkSourceStyleSchemeManager (full)
 */

// SourceStyleSchemeManager is a representation of GTK's GtkSourceStyleSchemeManager.
type SourceStyleSchemeManager struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceStyleSchemeManager.
func (v *SourceStyleSchemeManager) native() *C.GtkSourceStyleSchemeManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyleSchemeManager(p)
}

func marshalSourceStyleSchemeManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleSchemeManager(obj), nil
}

func wrapSourceStyleSchemeManager(obj *glib.Object) *SourceStyleSchemeManager {
	return &SourceStyleSchemeManager{obj}
}

// SourceStyleSchemeManagerNew is a wrapper around gtk_source_style_scheme_manager_new().
func SourceStyleSchemeManagerNew() (*SourceStyleSchemeManager, error) {
	c := C.gtk_source_style_scheme_manager_new()
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceStyleSchemeManager(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SourceStyleSchemeManagerGetDefault is a wrapper around gtk_source_style_scheme_manager_get_default().
func SourceStyleSchemeManagerGetDefault() (*SourceStyleSchemeManager, error) {
	c := C.gtk_source_style_scheme_manager_get_default()
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceStyleSchemeManager(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SetSearchPath is a wrapper around gtk_source_style_scheme_manager_set_search_path().
func (v *SourceStyleSchemeManager) SetSearchPath(path []string) {
	cArray := toCgcharArray(path)
	defer C.destroy_strings(cArray)

	C.gtk_source_style_scheme_manager_set_search_path(v.native(), cArray)
}

// AppendSearchPath is a wrapper around gtk_source_style_scheme_manager_append_search_path().
func (v *SourceStyleSchemeManager) AppendSearchPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_style_scheme_manager_append_search_path(v.native(), (*C.gchar)(cstr))
}

// PrependSearchPath is a wrapper around gtk_source_style_scheme_manager_prepend_search_path().
func (v *SourceStyleSchemeManager) PrependSearchPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_style_scheme_manager_prepend_search_path(v.native(), (*C.gchar)(cstr))
}

// GetSearchPath is a wrapper around gtk_source_style_scheme_manager_get_search_path().
func (v *SourceStyleSchemeManager) GetSearchPath() []string {
	return toGoStringArrayNoFree(C.gtk_source_style_scheme_manager_get_search_path(v.native()))
}

// GetShemeIds is a wrapper around gtk_source_style_scheme_manager_get_scheme_ids().
func (v *SourceStyleSchemeManager) GetShemeIds() []string {
	return toGoStringArrayNoFree(C.gtk_source_style_scheme_manager_get_scheme_ids(v.native()))
}

// GetScheme is a wrapper around gtk_source_style_scheme_manager_get_scheme().
func (v *SourceStyleSchemeManager) GetScheme(id string) (*SourceStyleScheme, error) {
	cstr := C.CString(id)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_source_style_scheme_manager_get_scheme(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceStyleScheme(glib.Take(unsafe.Pointer(c))), nil
}

// ForceRescan() is a wrapper around gtk_source_style_scheme_manager_force_rescan().
func (v *SourceStyleSchemeManager) ForceRescan() {
	C.gtk_source_style_scheme_manager_force_rescan(v.native())
}

/*
 * GtkSourceStyleScheme (full)
 */

// SourceStyleScheme is a representation of GTK's GtkSourceStyleScheme.
type SourceStyleScheme struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceStyleSchemeManager.
func (v *SourceStyleScheme) native() *C.GtkSourceStyleScheme {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyleScheme(p)
}

func marshalSourceStyleScheme(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyleScheme(obj), nil
}

func wrapSourceStyleScheme(obj *glib.Object) *SourceStyleScheme {
	return &SourceStyleScheme{obj}
}

// GetId is a wrapper around gtk_source_style_scheme_get_id().
func (v *SourceStyleScheme) GetId() string {
	return goString(C.gtk_source_style_scheme_get_id(v.native()))
}

// GetName is a wrapper around gtk_source_style_scheme_get_name().
func (v *SourceStyleScheme) GetName() string {
	return goString(C.gtk_source_style_scheme_get_name(v.native()))
}

// GetDescription is a wrapper around gtk_source_style_scheme_get_description().
func (v *SourceStyleScheme) GetDescription() string {
	return goString(C.gtk_source_style_scheme_get_description(v.native()))
}

// GetAuthors is a wrapper around gtk_source_style_scheme_get_authors().
func (v *SourceStyleScheme) GetAuthors() []string {
	return toGoStringArrayNoFree(C.gtk_source_style_scheme_get_authors(v.native()))
}

// GetFilename is a wrapper around gtk_source_style_scheme_get_filename().
func (v *SourceStyleScheme) GetFilename() string {
	return goString(C.gtk_source_style_scheme_get_filename(v.native()))
}

// GetStyle is a wrapper around gtk_source_style_scheme_get_style().
func (v *SourceStyleScheme) GetStyle(styleId string) (*SourceStyle, error) {
	cstr := C.CString(styleId)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_source_style_scheme_get_style(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceStyle(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkSourceStyle (full)
 */

// SourceStyle is a representation of GTK's GtkSourceStyle.
type SourceStyle struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceStyleSchemeManager.
func (v *SourceStyle) native() *C.GtkSourceStyle {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceStyle(p)
}

func marshalSourceStyle(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceStyle(obj), nil
}

func wrapSourceStyle(obj *glib.Object) *SourceStyle {
	return &SourceStyle{obj}
}

// Copy is a wrapper around gtk_source_style_copy().
func (v *SourceStyle) Copy() (*SourceStyle, error) {
	c := C.gtk_source_style_copy(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceStyle(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkSourceMark (full)
 */

// SourceMark is a representation of GTK's GtkSourceMark.
// Mark object for GtkSourceBuffer
type SourceMark struct {
	gtk.TextMark
}

func (v *SourceMark) toTextMark() *C.GtkTextMark {
	if v == nil {
		return nil
	}
	return C.toGtkTextMark(unsafe.Pointer(v.GObject))
}

// native returns a pointer to the underlying GtkSourceStyleSchemeManager.
func (v *SourceMark) native() *C.GtkSourceMark {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceMark(p)
}

func marshalSourceMark(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceMark(obj), nil
}

func wrapSourceMark(obj *glib.Object) *SourceMark {
	return &SourceMark{gtk.TextMark{obj}}
}

// SourceMarkNew is a wrapper around gtk_source_mark_new().
func SourceMarkNew(name, category string) (*SourceMark, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	cstr1 := C.CString(category)
	defer C.free(unsafe.Pointer(cstr1))

	c := C.gtk_source_mark_new((*C.gchar)(cstr), (*C.gchar)(cstr1))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceMark(glib.Take(unsafe.Pointer(c))), nil
}

// GetBuffer overriding gtk_text_mark_get_buffer() to get a
// GtkSourceBuffer from a GtkSourceMark
func (v *SourceMark) GetBuffer() (*SourceBuffer, error) {
	c := C.gtk_text_mark_get_buffer(v.toTextMark())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceBuffer(glib.Take(unsafe.Pointer(c))), nil
}

// GetCategory is a wrapper around gtk_source_mark_get_category().
func (v *SourceMark) GetCategory() string {
	return goString(C.gtk_source_mark_get_category(v.native()))
}

// Next is a wrapper around gtk_source_mark_next().
func (v *SourceMark) Next(category string) (*SourceMark, error) {
	cstr := C.CString(category)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_source_mark_next(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceMark(glib.Take(unsafe.Pointer(c))), nil
}

// Prev is a wrapper around gtk_source_mark_prev().
func (v *SourceMark) Prev(category string) (*SourceMark, error) {
	cstr := C.CString(category)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_source_mark_prev(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceMark(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkSourceMarkAttributes (full)
 */

// SourceMarkAttributes is a representation of GTK's GtkSourceMarkAttributes.
type SourceMarkAttributes struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceStyleSchemeManager.
func (v *SourceMarkAttributes) native() *C.GtkSourceMarkAttributes {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceMarkAttributes(p)
}

func marshalSourceMarkAttributes(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceMarkAttributes(obj), nil
}

func wrapSourceMarkAttributes(obj *glib.Object) *SourceMarkAttributes {
	return &SourceMarkAttributes{obj}
}

// SourceMarkAttributesNew is a wrapper around gtk_source_mark_attributes_new().
func SourceMarkAttributesNew() (*SourceMarkAttributes, error) {
	c := C.gtk_source_mark_attributes_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceMarkAttributes(glib.Take(unsafe.Pointer(c))), nil
}

// GetBackground is a wrapper around gtk_source_mark_attributes_get_background().
func (v *SourceMarkAttributes) GetBackground() (*gdk.RGBA, bool) {
	background := gdk.NewRGBA()

	c := C.gtk_source_mark_attributes_get_background(
		v.native(),
		(*C.GdkRGBA)(unsafe.Pointer(background.Native())))

	return background, gobool(c)
}

// SetBackground is a wrapper around gtk_source_mark_attributes_set_background().
func (v *SourceMarkAttributes) SetBackground(background *gdk.RGBA) {
	C.gtk_source_mark_attributes_set_background(
		v.native(),
		(*C.GdkRGBA)(unsafe.Pointer(background.Native())))
}

// SetIconName is a wrapper around gtk_source_mark_attributes_set_icon_name().
func (v *SourceMarkAttributes) SetIconName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_source_mark_attributes_set_icon_name(v.native(), (*C.gchar)(cstr))
}

// GetIconName is a wrapper around gtk_source_mark_attributes_get_icon_name().
func (v *SourceMarkAttributes) GetIconName() string {
	return goString(C.gtk_source_mark_attributes_get_icon_name(v.native()))
}

// SetGIcon is a wrapper around gtk_source_mark_attributes_set_gicon().
func (v *SourceMarkAttributes) SetGIcon(icon *glib.Icon) {
	C.gtk_source_mark_attributes_set_gicon(
		v.native(),
		(*C.GIcon)(icon.NativePrivate()))
}

// GetGIcon is a wrapper around gtk_source_mark_attributes_get_gicon().
func (v *SourceMarkAttributes) GetGIcon() (*glib.Icon, error) {
	return toGIcon(C.gtk_source_mark_attributes_get_gicon(v.native()))
}

// SetPixBuf is a wrapper around gtk_source_mark_attributes_set_pixbuf().
func (v *SourceMarkAttributes) SetPixbuf(pixbuf *gdk.Pixbuf) {
	C.gtk_source_mark_attributes_set_pixbuf(v.native(),
		(*C.GdkPixbuf)(pixbuf.NativePrivate()))
}

// GetPixBuf is a wrapper around gtk_source_mark_attributes_get_pixbuf().
func (v *SourceMarkAttributes) GetPixbuf() (*gdk.Pixbuf, error) {
	return toPixbuf(C.gtk_source_mark_attributes_get_pixbuf(v.native()))
}

// RenderIcon is a wrapper around gtk_source_mark_attributes_render_icon().
func (v *SourceMarkAttributes) RenderIcon(widget *gtk.Widget, size int) (*gdk.Pixbuf, error) {
	return toPixbuf(C.gtk_source_mark_attributes_render_icon(v.native(), nativeWidget(widget), C.gint(size)))
}

// GetTooltipText is a wrapper around gtk_source_mark_attributes_get_tooltip_text().
func (v *SourceMarkAttributes) GetTooltipText(mark *SourceMark) string {
	return goString(C.gtk_source_mark_attributes_get_tooltip_text(v.native(), mark.native()))
}

// GetTooltipMarkup is a wrapper around gtk_source_mark_attributes_get_tooltip_markup().
func (v *SourceMarkAttributes) GetTooltipMarkup(mark *SourceMark) string {
	return goString(C.gtk_source_mark_attributes_get_tooltip_markup(v.native(), mark.native()))
}

/*
 * GtkSourceCompletion (full)
 */

// SourceCompletion is a representation of GTK's GtkSourceCompletion.
type SourceCompletion struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceCompletion.
func (v *SourceCompletion) native() *C.GtkSourceCompletion {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceCompletion(p)
}

func marshalSourceCompletion(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceCompletion(obj), nil
}

func wrapSourceCompletion(obj *glib.Object) *SourceCompletion {
	return &SourceCompletion{obj}
}

// AddProvider is a wrapper around gtk_source_completion_add_provider().
func (v *SourceCompletion) AddProvider(provider *SourceCompletionProvider) (bool, error) {
	var err *C.GError = nil
	c := C.gtk_source_completion_add_provider(v.native(), provider.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return false, errors.New(goString(err.message))
	}

	return gobool(c), nil
}

// RemoveProvider is a wrapper around gtk_source_completion_remove_provider().
func (v *SourceCompletion) RemoveProvider(provider *SourceCompletionProvider) (bool, error) {
	var err *C.GError = nil
	c := C.gtk_source_completion_remove_provider(v.native(), provider.native(), &err)
	if err != nil {
		defer C.g_error_free(err)
		return false, errors.New(goString(err.message))
	}

	return gobool(c), nil
}

// GetProviders is a wrapper around gtk_source_completion_get_providers().
func (v *SourceCompletion) GetProviders() *glib.List {

	clist := C.gtk_source_completion_get_providers(v.native())
	if clist == nil {
		return nil
	}
	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapSourceCompletionProvider(glib.Take(ptr))
	})
	runtime.SetFinalizer(glist, func(l *glib.List) {
		l.Free()
	})

	return glist
}

// Hide is a wrapper around gtk_source_completion_hide().
func (v *SourceCompletion) Hide() {
	C.gtk_source_completion_hide(v.native())
}

// GetInfoWindow is a wrapper around gtk_source_completion_get_info_window().
func (v *SourceCompletion) GetInfoWindow() (*SourceCompletionInfo, error) {
	c := C.gtk_source_completion_get_info_window(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapSourceCompletionInfo(glib.Take(unsafe.Pointer(c))), nil
}

// GetView is a wrapper around gtk_source_completion_get_view().
func (v *SourceCompletion) GetView() (*SourceView, error) {
	c := C.gtk_source_completion_get_view(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapSourceView(glib.Take(unsafe.Pointer(c))), nil
}

// CreateContext is a wrapper around gtk_source_completion_create_context().
func (v *SourceCompletion) CreateContext(position *gtk.TextIter) (*SourceCompletionContext, error) {
	c := C.gtk_source_completion_create_context(v.native(), nativeTextIter(position))
	if c == nil {
		return nil, nilPtrErr
	}

	return wrapSourceCompletionContext(glib.Take(unsafe.Pointer(c))), nil
}

// BlockInteractive is a wrapper around gtk_source_completion_block_interactive().
func (v *SourceCompletion) BlockInteractive() {
	C.gtk_source_completion_block_interactive(v.native())
}

// UnblockInteractive is a wrapper around gtk_source_completion_unblock_interactive().
func (v *SourceCompletion) UnblockInteractive() {
	C.gtk_source_completion_unblock_interactive(v.native())
}

/*
 * GtkSourceCompletionContext (full)
 */

// SourceCompletionContext is a representation of GTK's GtkSourceCompletionContext.
type SourceCompletionContext struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkSourceCompletion.
func (v *SourceCompletionContext) native() *C.GtkSourceCompletionContext {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceCompletionContext(p)
}

func marshalSourceCompletionContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceCompletionContext(obj), nil
}

func wrapSourceCompletionContext(obj *glib.Object) *SourceCompletionContext {
	return &SourceCompletionContext{glib.InitiallyUnowned{obj}}
}

// AddProposals is a wrapper around gtk_source_completion_context_add_proposals().
// Providers must ensure that they always call this function with finished set to
// TRUE once each population (even if no proposals need to be added 'nil')
// A List must be manually freed by either calling Free() or FreeFull()
func (v *SourceCompletionContext) AddProposals(provider *SourceCompletionProvider, proposals *glib.List) bool {
	var finished C.gboolean
	C.gtk_source_completion_context_add_proposals(v.native(), provider.native(), toCGList(proposals), finished)
	return gobool(finished)
}

// GetIter is a wrapper around gtk_source_completion_context_get_iter().
func (v *SourceCompletionContext) GetIter() (*gtk.TextIter, error) {
	iter := new(gtk.TextIter)
	c := C.gtk_source_completion_context_get_iter(v.native(), nativeTextIter(iter))
	if !gobool(c) {
		return nil, errors.New("unable to get active iter")
	}
	return iter, nil
}

// GetActivation is a wrapper around gtk_source_completion_context_get_activation().
func (v *SourceCompletionContext) GetActivation() SourceCompletionActivation {
	c := C.gtk_source_completion_context_get_activation(v.native())
	return SourceCompletionActivation(c)
}

/*
 * GtkSourceCompletionProvider (full)
 */

// SourceCompletionProvider is a representation of GTK's GtkSourceCompletionProvider.
type SourceCompletionProvider struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceCompletionProvider.
func (v *SourceCompletionProvider) native() *C.GtkSourceCompletionProvider {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceCompletionProvider(p)
}

func marshalSourceCompletionProvider(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceCompletionProvider(obj), nil
}

func wrapSourceCompletionProvider(obj *glib.Object) *SourceCompletionProvider {
	return &SourceCompletionProvider{obj}
}

// GetName is a wrapper around gtk_source_completion_provider_get_name().
func (v *SourceCompletionProvider) GetName() string {
	return goString(C.gtk_source_completion_provider_get_name(v.native()))
}

// GetIcon is a wrapper around gtk_source_completion_provider_get_icon().
func (v *SourceCompletionProvider) GetIcon() (*gdk.Pixbuf, error) {
	return toPixbuf(C.gtk_source_completion_provider_get_icon(v.native()))
}

// GetIconName is a wrapper around gtk_source_completion_provider_get_icon_name().
func (v *SourceCompletionProvider) GetIconName() string {
	return goString(C.gtk_source_completion_provider_get_icon_name(v.native()))
}

// GetGIcon is a wrapper around gtk_source_completion_provider_get_gicon().
func (v *SourceCompletionProvider) GetGIcon() (*glib.Icon, error) {
	return toGIcon(C.gtk_source_completion_provider_get_gicon(v.native()))
}

// Populate is a wrapper around gtk_source_completion_provider_populate().
func (v *SourceCompletionProvider) Populate(context *SourceCompletionContext) {
	C.gtk_source_completion_provider_populate(v.native(), context.native())
}

// GetActivation is a wrapper around gtk_source_completion_provider_get_activation().
func (v *SourceCompletionProvider) GetActivation(context *SourceCompletionContext) SourceCompletionActivation {
	return SourceCompletionActivation(C.gtk_source_completion_provider_get_activation(v.native()))
}

// Match is a wrapper around gtk_source_completion_provider_match().
func (v *SourceCompletionProvider) Match(context *SourceCompletionContext) bool {
	return gobool(C.gtk_source_completion_provider_match(v.native(), context.native()))
}

// GetInfoWidget is a wrapper around gtk_source_completion_provider_get_info_widget().
func (v *SourceCompletionProvider) GetInfoWidget(proposal *SourceCompletionProposal) *gtk.Widget {
	c := C.gtk_source_completion_provider_get_info_widget(v.native(), proposal.native())
	return &gtk.Widget{glib.InitiallyUnowned{glib.Take(unsafe.Pointer(c))}}
}

// UpdateInfo is a wrapper around gtk_source_completion_provider_update_info().
func (v *SourceCompletionProvider) UpdateInfo(proposal *SourceCompletionProposal, info *SourceCompletionInfo) {
	C.gtk_source_completion_provider_update_info(v.native(), proposal.native(), info.native())
}

// GetStartIter is a wrapper around gtk_source_completion_provider_get_start_iter().
func (v *SourceCompletionProvider) GetStartIter(context *SourceCompletionContext,
	proposal *SourceCompletionProposal) (*gtk.TextIter, error) {

	iter := new(gtk.TextIter)
	c := C.gtk_source_completion_provider_get_start_iter(
		v.native(),
		context.native(),
		proposal.native(),
		nativeTextIter(iter))

	if !gobool(c) {
		return nil, errors.New("unable to get start iter")
	}
	return iter, nil
}

// ActivateProposal is a wrapper around gtk_source_completion_provider_activate_proposal().
func (v *SourceCompletionProvider) ActivateProposal(proposal *SourceCompletionProposal) (*gtk.TextIter, error) {

	iter := new(gtk.TextIter)
	c := C.gtk_source_completion_provider_activate_proposal(
		v.native(),
		proposal.native(),
		nativeTextIter(iter))

	if !gobool(c) {
		return nil, errors.New("unable to get start iter")
	}
	return iter, nil
}

// GetInteractiveDelay is a wrapper around gtk_source_completion_provider_get_interactive_delay().
func (v *SourceCompletionProvider) GetInteractiveDelay(context *SourceCompletionContext) int {
	return int(C.gtk_source_completion_provider_get_interactive_delay(v.native()))
}

// GetPriority is a wrapper around gtk_source_completion_provider_get_priority().
func (v *SourceCompletionProvider) GetPriority(context *SourceCompletionContext) int {
	return int(C.gtk_source_completion_provider_get_priority(v.native()))
}

/*
 * GtkSourceCompletionInfo (full)
 */

// SourceCompletionInfo is a representation of GTK's GtkSourceCompletionInfo.
type SourceCompletionInfo struct {
	gtk.Window
}

// native returns a pointer to the underlying GtkSourceCompletionInfo.
func (v *SourceCompletionInfo) native() *C.GtkSourceCompletionInfo {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceCompletionInfo(p)
}

func marshalSourceCompletionInfo(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceCompletionInfo(obj), nil
}

func wrapSourceCompletionInfo(obj *glib.Object) *SourceCompletionInfo {

	return &SourceCompletionInfo{gtk.Window{gtk.Bin{gtk.Container{gtk.Widget{glib.InitiallyUnowned{obj}}}}}}
}

// SourceCompletionInfoNew is a wrapper around gtk_source_completion_info_new().
func SourceCompletionInfoNew() (*SourceCompletionInfo, error) {
	c := C.gtk_source_completion_info_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceCompletionInfo(glib.Take(unsafe.Pointer(c))), nil
}

// MoveToIter is a wrapper around gtk_source_completion_info_move_to_iter().
func (v *SourceCompletionInfo) MoveToIter(view *gtk.TextView, iter *gtk.TextIter) {
	C.gtk_source_completion_info_move_to_iter(v.native(),
		C.toGtkTextView(unsafe.Pointer(view.Native())),
		nativeTextIter(iter))
}

/*
 * GtkSourceCompletionProposal (full)
 */

// SourceCompletionProposal is a representation of GTK's GtkSourceCompletionProposal.
type SourceCompletionProposal struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceCompletionProposal.
func (v *SourceCompletionProposal) native() *C.GtkSourceCompletionProposal {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceCompletionProposal(p)
}

func marshalSourceCompletionProposal(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceCompletionProposal(obj), nil
}

func wrapSourceCompletionProposal(obj *glib.Object) *SourceCompletionProposal {
	return &SourceCompletionProposal{obj}
}

// GetLabel is a wrapper around gtk_source_completion_proposal_get_label().
func (v *SourceCompletionProposal) GetLabel() string {
	return toGoStringFree(C.gtk_source_completion_proposal_get_label(v.native()))
}

// GetMarkup is a wrapper around gtk_source_completion_proposal_get_markup().
func (v *SourceCompletionProposal) GetMarkup() string {
	return toGoStringFree(C.gtk_source_completion_proposal_get_markup(v.native()))
}

// GetText is a wrapper around gtk_source_completion_proposal_get_text().
func (v *SourceCompletionProposal) GetText() string {
	return toGoStringFree(C.gtk_source_completion_proposal_get_text(v.native()))
}

// GetGIcon is a wrapper around gtk_source_completion_proposal_get_gicon().
func (v *SourceCompletionProposal) GetGIcon() (*glib.Icon, error) {
	return toGIcon(C.gtk_source_completion_proposal_get_gicon(v.native()))
}

// GetInfo is a wrapper around gtk_source_completion_proposal_get_info().
func (v *SourceCompletionProposal) GetInfo() string {
	return toGoStringFree(C.gtk_source_completion_proposal_get_info(v.native()))
}

// Hash is a wrapper around gtk_source_completion_proposal_hash().
func (v *SourceCompletionProposal) Hash() uint {
	return uint(C.gtk_source_completion_proposal_hash(v.native()))
}

// Equal is a wrapper around gtk_source_completion_proposal_equal().
func (v *SourceCompletionProposal) Equal(other *SourceCompletionProposal) bool {
	return gobool(C.gtk_source_completion_proposal_equal(v.native(), other.native()))
}

/*
 * GtkSourceCompletionItem
 */

// SourceCompletionItem is a representation of GTK's GtkSourceCompletionItem.
type SourceCompletionItem struct {
	*glib.Object

	SourceCompletionProposal
}

// native returns a pointer to the underlying GtkSourceCompletionItem.
func (v *SourceCompletionItem) native() *C.GtkSourceCompletionItem {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceCompletionItem(p)
}

func marshalSourceCompletionItem(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceCompletionItem(obj), nil
}

func wrapSourceCompletionItem(obj *glib.Object) *SourceCompletionItem {
	cp := wrapSourceCompletionProposal(obj)

	return &SourceCompletionItem{obj, *cp}
}

/*
 * GtkSourceCompletionWords (full)
 */

// SourceCompletionWords is a representation of GTK's GtkSourceCompletionWords.
type SourceCompletionWords struct {
	*glib.Object

	SourceCompletionProvider
}

// native returns a pointer to the underlying GtkSourceCompletionWords.
func (v *SourceCompletionWords) native() *C.GtkSourceCompletionWords {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceCompletionWords(p)
}

func marshalSourceCompletionWords(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceCompletionWords(obj), nil
}

func wrapSourceCompletionWords(obj *glib.Object) *SourceCompletionWords {
	cp := wrapSourceCompletionProvider(obj)

	return &SourceCompletionWords{obj, *cp}
}

// SourceCompletionWordsNew is a wrapper around gtk_source_completion_words_new().
func SourceCompletionWordsNew(name string, icon *gdk.Pixbuf) (*SourceCompletionWords, error) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_source_completion_words_new((*C.gchar)(cstr), (*C.GdkPixbuf)(unsafe.Pointer(icon.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceCompletionWords(glib.Take(unsafe.Pointer(c))), nil
}

// Register is a wrapper around gtk_source_completion_words_register().
func (v *SourceCompletionWords) Register(buffer *gtk.TextBuffer) {

	C.gtk_source_completion_words_register(v.native(), (*C.GtkTextBuffer)(unsafe.Pointer(buffer.Native())))
}

// Unregister is a wrapper around gtk_source_completion_words_unregister().
func (v *SourceCompletionWords) Unregister(buffer *gtk.TextBuffer) {

	C.gtk_source_completion_words_unregister(v.native(), (*C.GtkTextBuffer)(unsafe.Pointer(buffer.Native())))
}

/*
 * GtkSourceGutter (full)
 */

// SourceGutter is a representation of GTK's GtkSourceGutter.
type SourceGutter struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceGutter.
func (v *SourceGutter) native() *C.GtkSourceGutter {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceGutter(p)
}

func marshalSourceGutter(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceGutter(obj), nil
}

func wrapSourceGutter(obj *glib.Object) *SourceGutter {
	return &SourceGutter{obj}
}

// Insert is a wrapper around gtk_source_gutter_insert().
func (v *SourceGutter) Insert(renderer *SourceGutterRenderer, position int) bool {
	return gobool(C.gtk_source_gutter_insert(
		v.native(),
		renderer.native(),
		C.gint(position)))
}

// Reorder is a wrapper around gtk_source_gutter_reorder().
func (v *SourceGutter) Reorder(renderer *SourceGutterRenderer, position int) {
	C.gtk_source_gutter_reorder(
		v.native(),
		renderer.native(),
		C.gint(position))
}

// Remove is a wrapper around gtk_source_gutter_remove().
func (v *SourceGutter) Remove(renderer *SourceGutterRenderer) {
	C.gtk_source_gutter_remove(
		v.native(),
		renderer.native())
}

// QueueDraw is a wrapper around gtk_source_gutter_queue_draw().
func (v *SourceGutter) QueueDraw() {
	C.gtk_source_gutter_queue_draw(v.native())
}

// GetRendererAtPos is a wrapper around gtk_source_gutter_get_renderer_at_pos().
func (v *SourceGutter) GetRendererAtPos(x, y int) (*SourceGutterRenderer, error) {
	c := C.gtk_source_gutter_get_renderer_at_pos(v.native(), C.gint(x), C.gint(y))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceGutterRenderer(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * GtkSourceGutterRenderer (full)
 */

// SourceGutterRenderer is a representation of GTK's GtkSourceGutterRenderer.
// Gutter cell renderer
type SourceGutterRenderer struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkSourceGutterRenderer.
func (v *SourceGutterRenderer) native() *C.GtkSourceGutterRenderer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceGutterRenderer(p)
}

func marshalSourceGutterRenderer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceGutterRenderer(obj), nil
}

func wrapSourceGutterRenderer(obj *glib.Object) *SourceGutterRenderer {
	return &SourceGutterRenderer{glib.InitiallyUnowned{obj}}
}

// Begin is a wrapper around gtk_source_gutter_renderer_begin().
func (v *SourceGutterRenderer) Begin(ct *cairo.Context, backgroundArea,
	cellArea *gdk.Rectangle, start, end *gtk.TextIter) {

	C.gtk_source_gutter_renderer_begin(
		v.native(),
		(*C.cairo_t)(unsafe.Pointer(ct.Native())),
		nativeGdkRectangle(*backgroundArea),
		nativeGdkRectangle(*cellArea),
		nativeTextIter(start), nativeTextIter(end))
}

// Draw is a wrapper around gtk_source_gutter_renderer_draw().
func (v *SourceGutterRenderer) Draw(ct *cairo.Context, backgroundArea,
	cellArea *gdk.Rectangle, start, end *gtk.TextIter, state SourceGutterRendererState) {

	C.gtk_source_gutter_renderer_draw(
		v.native(),
		(*C.cairo_t)(unsafe.Pointer(ct.Native())),
		nativeGdkRectangle(*backgroundArea),
		nativeGdkRectangle(*cellArea),
		nativeTextIter(start), nativeTextIter(end), C.GtkSourceGutterRendererState(state))
}

// GetView is a wrapper around gtk_source_gutter_renderer_get_view().
func (v *SourceGutterRenderer) GetView() (*gtk.TextView, error) {
	c := C.gtk_source_gutter_renderer_get_view(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	tv := &gtk.TextView{gtk.Container{gtk.Widget{glib.InitiallyUnowned{glib.Take(unsafe.Pointer(c))}}}}
	return tv, nil
}

// GetWindowType is a wrapper around gtk_source_gutter_renderer_get_window_type().
func (v *SourceGutterRenderer) GetWindowType() gtk.TextWindowType {
	return gtk.TextWindowType(C.gtk_source_gutter_renderer_get_window_type(v.native()))
}

// GetVisible is a wrapper around gtk_source_gutter_renderer_get_visible().
func (v *SourceGutterRenderer) GetVisible() bool {
	return gobool(C.gtk_source_gutter_renderer_get_visible(v.native()))
}

// SetVisible is a wrapper around gtk_source_gutter_renderer_set_visible().
func (v *SourceGutterRenderer) SetVisible(visible bool) {
	C.gtk_source_gutter_renderer_set_visible(v.native(), gbool(visible))
}

// GetSize is a wrapper around gtk_source_gutter_renderer_get_size().
func (v *SourceGutterRenderer) GetSize() int {
	return int(C.gtk_source_gutter_renderer_get_size(v.native()))
}

// SetSize is a wrapper around gtk_source_gutter_renderer_set_size().
func (v *SourceGutterRenderer) SetSize(size int) {
	C.gtk_source_gutter_renderer_set_size(v.native(), C.gint(size))
}

// GetPadding is a wrapper around gtk_source_gutter_renderer_get_padding().
func (v *SourceGutterRenderer) GetPadding() (int, int) {
	cx, cy := new(C.gint), new(C.gint)

	C.gtk_source_gutter_renderer_get_padding(v.native(), cx, cy)
	return int(*cx), int(*cy)
}

// SetPadding is a wrapper around gtk_source_gutter_renderer_set_padding().
func (v *SourceGutterRenderer) SetPadding(x, y int) {
	C.gtk_source_gutter_renderer_set_padding(v.native(), C.gint(x), C.gint(y))
}

// GetAlignment is a wrapper around gtk_source_gutter_renderer_get_alignment().
func (v *SourceGutterRenderer) GetAlignment() (float64, float64) {
	cx, cy := new(C.float), new(C.float)

	C.gtk_source_gutter_renderer_get_alignment(v.native(), cx, cy)
	return float64(*cx), float64(*cy)
}

// SetAlignment is a wrapper around gtk_source_gutter_renderer_set_alignment().
func (v *SourceGutterRenderer) SetAlignment(xalign, yalign float64) {
	C.gtk_source_gutter_renderer_set_alignment(
		v.native(), C.float(xalign), C.float(yalign))
}

// GetAlignmentMode is a wrapper around gtk_source_gutter_renderer_get_alignment_mode().
func (v *SourceGutterRenderer) GetAlignmentMode() SourceGutterRendererAlignmentMode {
	return SourceGutterRendererAlignmentMode(
		C.gtk_source_gutter_renderer_get_alignment_mode(v.native()))
}

// SetAlignmentMode is a wrapper around gtk_source_gutter_renderer_set_alignment_mode().
func (v *SourceGutterRenderer) SetAlignmentMode(amode SourceGutterRendererAlignmentMode) {
	C.gtk_source_gutter_renderer_set_alignment_mode(
		v.native(), C.GtkSourceGutterRendererAlignmentMode(amode))
}

// GetBackground is a wrapper around gtk_source_gutter_renderer_get_background().
func (v *SourceGutterRenderer) GetBackground() (*gdk.RGBA, bool) {
	background := gdk.NewRGBA()

	c := C.gtk_source_gutter_renderer_get_background(
		v.native(),
		(*C.GdkRGBA)(unsafe.Pointer(background.Native())))

	return background, gobool(c)
}

// SetBackground is a wrapper around gtk_source_gutter_renderer_set_background().
func (v *SourceGutterRenderer) SetBackground(background *gdk.RGBA) {
	C.gtk_source_gutter_renderer_set_background(
		v.native(),
		(*C.GdkRGBA)(unsafe.Pointer(background.Native())))
}

// Activate is a wrapper around gtk_source_gutter_renderer_activate().
func (v *SourceGutterRenderer) Activate(iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event) {
	C.gtk_source_gutter_renderer_activate(
		v.native(),
		nativeTextIter(iter),
		nativeGdkRectangle(*area),
		(*C.GdkEvent)(unsafe.Pointer(event.Native())))
}

// Activatable is a wrapper around gtk_source_gutter_renderer_query_activatable().
func (v *SourceGutterRenderer) Activatable(iter *gtk.TextIter, area *gdk.Rectangle, event *gdk.Event) bool {
	return gobool(C.gtk_source_gutter_renderer_query_activatable(
		v.native(),
		nativeTextIter(iter),
		nativeGdkRectangle(*area),
		(*C.GdkEvent)(unsafe.Pointer(event.Native()))))
}

// QueryData is a wrapper around gtk_source_gutter_renderer_query_data().
func (v *SourceGutterRenderer) QueryData(start, end *gtk.TextIter, state SourceGutterRendererState) {
	C.gtk_source_gutter_renderer_query_data(
		v.native(),
		nativeTextIter(start),
		nativeTextIter(end),
		C.GtkSourceGutterRendererState(state))
}

// QueryTooltip is a wrapper around gtk_source_gutter_renderer_query_tooltip().
func (v *SourceGutterRenderer) QueryTooltip(iter *gtk.TextIter, area *gdk.Rectangle, x, y int, tooltip *gtk.Tooltip) bool {
	return gobool(C.gtk_source_gutter_renderer_query_tooltip(
		v.native(),
		nativeTextIter(iter),
		nativeGdkRectangle(*area),
		C.gint(x),
		C.gint(y),
		(*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))))
}

// QueuDraw is a wrapper around gtk_source_gutter_renderer_queue_draw().
func (v *SourceGutterRenderer) QueuDraw() {
	C.gtk_source_gutter_renderer_queue_draw(v.native())
}

/*
 * GtkSourceGutterRendererPixbuf (full)
 */

// SourceGutterRendererPixbuf is a representation of GTK's GtkSourceGutterRendererPixbuf.
// Renders a pixbuf in the gutter
type SourceGutterRendererPixbuf struct {
	SourceGutterRenderer
}

// native returns a pointer to the underlying GtkSourceGutterRendererPixbuf.
func (v *SourceGutterRendererPixbuf) native() *C.GtkSourceGutterRendererPixbuf {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceGutterRendererPixbuf(p)
}

func marshalSourceGutterRendererPixbuf(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceGutterRendererPixbuf(obj), nil
}

func wrapSourceGutterRendererPixbuf(obj *glib.Object) *SourceGutterRendererPixbuf {
	return &SourceGutterRendererPixbuf{SourceGutterRenderer{glib.InitiallyUnowned{obj}}}
}

// SourceGutterRendererPixbufNew is a wrapper around gtk_source_gutter_renderer_pixbuf_new().
func SourceGutterRendererPixbufNew() (*SourceGutterRenderer, error) {
	c := C.gtk_source_gutter_renderer_pixbuf_new()
	if c == nil {
		return nil, nilPtrErr
	}
	// return wrapSourceGutterRendererPixbuf(glib.Take(unsafe.Pointer(c))), nil
	return wrapSourceGutterRenderer(glib.Take(unsafe.Pointer(c))), nil
}

// SetPixbuf is a wrapper around gtk_source_gutter_renderer_pixbuf_set_pixbuf().
func (v *SourceGutterRendererPixbuf) SetPixbuf(pixbuf *gdk.Pixbuf) {
	C.gtk_source_gutter_renderer_pixbuf_set_pixbuf(v.native(),
		(*C.GdkPixbuf)(pixbuf.NativePrivate()))
}

// GetPixbuf is a wrapper around gtk_source_gutter_renderer_pixbuf_get_pixbuf().
func (v *SourceGutterRendererPixbuf) GetPixbuf() (*gdk.Pixbuf, error) {
	return toPixbuf(C.gtk_source_gutter_renderer_pixbuf_get_pixbuf(v.native()))
}

// SetGIcon is a wrapper around gtk_source_gutter_renderer_pixbuf_set_gicon().
func (v *SourceGutterRendererPixbuf) SetGIcon(icon *glib.Icon) {
	C.gtk_source_gutter_renderer_pixbuf_set_gicon(
		v.native(),
		(*C.GIcon)(icon.NativePrivate()))
}

// GetGIcon is a wrapper around gtk_source_gutter_renderer_pixbuf_get_gicon().
func (v *SourceGutterRendererPixbuf) GetGIcon() (*glib.Icon, error) {
	return toGIcon(C.gtk_source_gutter_renderer_pixbuf_get_gicon(v.native()))
}

// SetIconName is a wrapper around gtk_source_gutter_renderer_pixbuf_set_icon_name().
func (v *SourceGutterRendererPixbuf) SetIconName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_gutter_renderer_pixbuf_set_icon_name(v.native(), (*C.char)(cstr))
}

// GetIconName is a wrapper around gtk_source_gutter_renderer_pixbuf_get_icon_name().
func (v *SourceGutterRendererPixbuf) GetIconName() string {
	return goString(C.gtk_source_gutter_renderer_pixbuf_get_icon_name(v.native()))
}

/*
 * GtkSourceGutterRendererText (full)
 */

// SourceGutterRendererText is a representation of GTK's GtkSourceGutterRendererText.
// Renders text in the gutter
type SourceGutterRendererText struct {
	SourceGutterRenderer
}

// native returns a pointer to the underlying GtkSourceGutterRendererText.
func (v *SourceGutterRendererText) native() *C.GtkSourceGutterRendererText {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceGutterRendererText(p)
}

func marshalSourceGutterRendererText(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceGutterRendererText(obj), nil
}

func wrapSourceGutterRendererText(obj *glib.Object) *SourceGutterRendererText {
	return &SourceGutterRendererText{SourceGutterRenderer{glib.InitiallyUnowned{obj}}}
}

// SourceGutterRendererTextNew is a wrapper around gtk_source_gutter_renderer_text_new().
func SourceGutterRendererTextNew() (*SourceGutterRenderer, error) {
	c := C.gtk_source_gutter_renderer_text_new()
	if c == nil {
		return nil, nilPtrErr
	}
	// return wrapSourceGutterRendererText(glib.Take(unsafe.Pointer(c))), nil
	return wrapSourceGutterRenderer(glib.Take(unsafe.Pointer(c))), nil
}

// SetMarkup is a wrapper around gtk_source_gutter_renderer_text_set_markup().
func (v *SourceGutterRendererText) SetMarkup(markup string) {
	cstr := C.CString(markup)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_source_gutter_renderer_text_set_markup(v.native(), (*C.char)(cstr), (C.gint)(len(markup)))
}

// SetText is a wrapper around gtk_source_gutter_renderer_text_set_text().
func (v *SourceGutterRendererText) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_source_gutter_renderer_text_set_text(v.native(), (*C.char)(cstr), (C.gint)(len(text)))
}

// Measure is a wrapper around gtk_source_gutter_renderer_text_measure().
func (v *SourceGutterRendererText) Measure(text string) (width, height int) {
	var cwidth, cheight *C.gint
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_gutter_renderer_text_measure(v.native(), (*C.char)(cstr), cwidth, cheight)

	if cwidth != nil {
		width = int(*((*C.gint)(unsafe.Pointer(cwidth))))
	}
	if cheight != nil {
		height = int(*((*C.gint)(unsafe.Pointer(cheight))))
	}
	return width, height
}

// MeasureMarkup is a wrapper around gtk_source_gutter_renderer_text_measure_markup().
func (v *SourceGutterRendererText) MeasureMarkup(markup string) (width, height int) {
	var cwidth, cheight *C.gint
	cstr := C.CString(markup)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_gutter_renderer_text_measure_markup(v.native(), (*C.char)(cstr), cwidth, cheight)

	if cwidth != nil {
		width = int(*((*C.gint)(unsafe.Pointer(cwidth))))
	}
	if cheight != nil {
		height = int(*((*C.gint)(unsafe.Pointer(cheight))))
	}
	return width, height
}

/*
 * GtkSourceUndoManager (full)
 */

// SourceUndoManager is a representation of GTK's GtkSourceUndoManager.
// Undo manager interface for GtkSourceView
type SourceUndoManager struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceUndoManager.
func (v *SourceUndoManager) native() *C.GtkSourceUndoManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceUndoManager(p)
}

func marshalSourceUndoManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceUndoManager(obj), nil
}

func wrapSourceUndoManager(obj *glib.Object) *SourceUndoManager {
	return &SourceUndoManager{obj}
}

// CanUndo is a wrapper around gtk_source_undo_manager_can_undo().
func (v *SourceUndoManager) CanUndo() bool {
	return gobool(C.gtk_source_undo_manager_can_undo(v.native()))
}

// CanRedo is a wrapper around gtk_source_undo_manager_can_redo().
func (v *SourceUndoManager) CanRedo() bool {
	return gobool(C.gtk_source_undo_manager_can_redo(v.native()))
}

// Undo is a wrapper around gtk_source_undo_manager_undo().
func (v *SourceUndoManager) Undo() {
	C.gtk_source_undo_manager_undo(v.native())
}

// Redo is a wrapper around gtk_source_undo_manager_redo().
func (v *SourceUndoManager) Redo() {
	C.gtk_source_undo_manager_redo(v.native())
}

// BeginNotUndoableAction is a wrapper around gtk_source_undo_manager_begin_not_undoable_action().
func (v *SourceUndoManager) BeginNotUndoableAction() {
	C.gtk_source_undo_manager_begin_not_undoable_action(v.native())
}

// EndNotUndoableAction is a wrapper around gtk_source_undo_manager_end_not_undoable_action().
func (v *SourceUndoManager) EndNotUndoableAction() {
	C.gtk_source_undo_manager_end_not_undoable_action(v.native())
}

// CanUndoChanged is a wrapper around gtk_source_undo_manager_can_undo_changed().
func (v *SourceUndoManager) CanUndoChanged() {
	C.gtk_source_undo_manager_can_undo_changed(v.native())
}

// CanRedoChanged is a wrapper around gtk_source_undo_manager_can_redo_changed().
func (v *SourceUndoManager) CanRedoChanged() {
	C.gtk_source_undo_manager_can_redo_changed(v.native())
}
