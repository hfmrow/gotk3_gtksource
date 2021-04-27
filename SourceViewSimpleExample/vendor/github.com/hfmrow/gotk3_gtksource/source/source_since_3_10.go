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
// #include "source_facility.go.h"
// #include "source_since_3_10.go.h"
import "C"
import (
	"errors"
	"log"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_source_search_context_get_type()), marshalSourceSearchContext},
		{glib.Type(C.gtk_source_search_settings_get_type()), marshalSourceSearchSettings},
	}

	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceSearchContext"] = wrapSourceSearchContext
	gtk.WrapMap["GtkSourceSearchSettings"] = wrapSourceSearchSettings
}

/*
 * GtkSourceUtils
 * Since: 3.10 but this project start at gtk_3_16
 */

// SourceUtilsUnescapeSearchText is a wrapper around gtk_source_utils_unescape_search_text().
// Warning: the escape and unescape functions are not reciprocal! For example,
// escape (unescape (\)) = \\. So avoid cycles such as:
// search entry -> unescape -> search settings -> escape -> search entry.
// The original search entry text may be modified.
func SourceUtilsUnescapeSearchText(text string) string {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	return goString(
		C.gtk_source_utils_unescape_search_text((*C.gchar)(cstr)))
}

// SourceUtilsEscapeSearchText is a wrapper around gtk_source_utils_escape_search_text().
// Warning: the escape and unescape functions are not reciprocal! For example,
// escape (unescape (\)) = \\. So avoid cycles such as:
// search entry -> unescape -> search settings -> escape -> search entry.
// The original search entry text may be modified.
func SourceUtilsEscapeSearchText(text string) string {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	return goString(
		C.gtk_source_utils_escape_search_text((*C.gchar)(cstr)))
}

/*
 * GtkSourceSearchContext
 * Since: 3.10 but this project start at gtk_3_16
 */

// SourceSearchContext is a representation of GTK's GtkSourceSearchContext.
type SourceSearchContext struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceSearchContext.
func (v *SourceSearchContext) native() *C.GtkSourceSearchContext {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceSearchContext(p)
}

func marshalSourceSearchContext(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceSearchContext(obj), nil
}

func wrapSourceSearchContext(obj *glib.Object) *SourceSearchContext {
	return &SourceSearchContext{obj}
}

// SourceSearchContextNew is a wrapper around gtk_source_search_context_new().
func SourceSearchContextNew(buffer *SourceBuffer, settings *SourceSearchSettings) (*SourceSearchContext, error) {
	c := C.gtk_source_search_context_new(buffer.native(), settings.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceSearchContext(glib.Take(unsafe.Pointer(c))), nil
}

// GetBuffer is a wrapper around gtk_source_search_context_get_buffer().
func (v *SourceSearchContext) GetBuffer() (*SourceBuffer, error) {
	c := C.gtk_source_search_context_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceBuffer(glib.Take(unsafe.Pointer(c))), nil
}

// GetSettings is a wrapper around gtk_source_search_context_get_settings().
func (v *SourceSearchContext) GetSettings() (*SourceSearchSettings, error) {
	c := C.gtk_source_search_context_get_settings(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceSearchSettings(glib.Take(unsafe.Pointer(c))), nil
}

// GetHighLight is a wrapper around gtk_source_search_context_get_highlight().
func (v *SourceSearchContext) GetHighLight() bool {
	return gobool(C.gtk_source_search_context_get_highlight(v.native()))
}

// SetHighLight is a wrapper around gtk_source_search_context_set_highlight().
func (v *SourceSearchContext) SetHighLight(highlight bool) {
	C.gtk_source_search_context_set_highlight(
		v.native(),
		gbool(highlight))
}

// GetOccurencesCount is a wrapper around gtk_source_search_context_get_occurrences_count().
func (v *SourceSearchContext) GetOccurencesCount() int {
	c := C.gtk_source_search_context_get_occurrences_count(v.native())
	return int(c)
}

// GetOccurencePosition is a wrapper around gtk_source_search_context_get_occurrence_position().
func (v *SourceSearchContext) GetOccurencePosition(start, end *gtk.TextIter) int {

	c := C.gtk_source_search_context_get_occurrence_position(
		v.native(),
		nativeTextIter(start),
		nativeTextIter(end))

	return int(c)
}

// ForwardAsync is a wrapper around gtk_source_search_context_forward_async().
func (v *SourceSearchContext) ForwardAsync(
	iter *gtk.TextIter, cancellable *glib.Cancellable, callback glib.AsyncReadyCallback, userData ...uintptr) {
	var userDataPtr uintptr
	var c *C.GCancellable = nil

	if cancellable != nil {
		c = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	switch {
	case len(userData) == 1:
		userDataPtr = userData[0]
	case len(userData) > 1:
		log.Fatal("ForwardAsync: Only one argument is allowed for userData.")
	}

	id := registerAsyncReadyCallback(callback, userDataPtr)

	C._gtk_source_search_context_forward_async(
		v.native(),
		nativeTextIter(iter),
		c,
		C.gpointer(uintptr(id)))
}

// BackwardAsync is a wrapper around gtk_source_search_context_backward_async().
func (v *SourceSearchContext) BackwardAsync(
	iter *gtk.TextIter, cancellable *glib.Cancellable, callback glib.AsyncReadyCallback, userData ...uintptr) {
	var userDataPtr uintptr
	var c *C.GCancellable = nil

	if cancellable != nil {
		c = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	switch {
	case len(userData) == 1:
		userDataPtr = userData[0]
	case len(userData) > 1:
		log.Fatal("BackwardAsync: Only one argument is allowed for userData.")
	}

	id := registerAsyncReadyCallback(callback, userDataPtr)

	C._gtk_source_search_context_backward_async(
		v.native(),
		nativeTextIter(iter),
		c,
		C.gpointer(uintptr(id)))
}

// ReplaceAll is a wrapper around gtk_source_search_context_replace_all().
func (v *SourceSearchContext) ReplaceAll(replace string) (int, error) {
	var err *C.GError = nil
	cstr := C.CString(replace)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gtk_source_search_context_replace_all(v.native(), (*C.gchar)(cstr), C.gint(len(replace)), &err)

	if err != nil {
		defer C.g_error_free(err)
		return int(c), errors.New(goString(err.message))
	}
	return int(c), nil
}

// ReplaceAll is a wrapper around gtk_source_search_context_get_regex_error().
func (v *SourceSearchContext) GetRegexError() error {

	if err := C.gtk_source_search_context_get_regex_error(v.native()); err != nil {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}

/*
 * GtkSourceSearchSettings (full)
 * Since: 3.10 but this project start at gtk_3_16 with gtk_sourceview_3_18
 */

// SourceSearchSettings is a representation of GTK's GtkSourceSearchSettings.
type SourceSearchSettings struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceSearchSettings.
func (v *SourceSearchSettings) native() *C.GtkSourceSearchSettings {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceSearchSettings(p)
}

func marshalSourceSearchSettings(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceSearchSettings(obj), nil
}

func wrapSourceSearchSettings(obj *glib.Object) *SourceSearchSettings {
	return &SourceSearchSettings{obj}
}

// SourceSearchSettingsNew is a wrapper around gtk_source_search_settings_new().
func SourceSearchSettingsNew() (*SourceSearchSettings, error) {
	c := C.gtk_source_search_settings_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceSearchSettings(glib.Take(unsafe.Pointer(c))), nil
}

// GetSearchText is a wrapper around gtk_source_search_settings_get_search_text().
func (v *SourceSearchSettings) GetSearchText() string {
	return goString(C.gtk_source_search_settings_get_search_text(v.native()))
}

// SetSearchText is a wrapper around gtk_source_search_settings_set_search_text().
func (v *SourceSearchSettings) SetSearchText(search string) {
	cstr := C.CString(search)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_search_settings_set_search_text(v.native(), (*C.char)(cstr))
}

// GetCaseSensitive is a wrapper around gtk_source_search_settings_get_case_sensitive().
func (v *SourceSearchSettings) GetCaseSensitive() bool {
	return gobool(C.gtk_source_search_settings_get_case_sensitive(v.native()))
}

// SetCaseSensitive is a wrapper around gtk_source_search_settings_set_case_sensitive().
func (v *SourceSearchSettings) SetCaseSensitive(caseSensitive bool) {
	C.gtk_source_search_settings_set_case_sensitive(v.native(),
		gbool(caseSensitive))
}

// GetWordBoundaries is a wrapper around gtk_source_search_settings_get_at_word_boundaries().
func (v *SourceSearchSettings) GetWordBoundaries() bool {
	return gobool(C.gtk_source_search_settings_get_at_word_boundaries(v.native()))
}

// SetWordBoundaries is a wrapper around gtk_source_search_settings_set_at_word_boundaries().
func (v *SourceSearchSettings) SetWordBoundaries(set bool) {
	C.gtk_source_search_settings_set_at_word_boundaries(v.native(),
		gbool(set))
}

// GetWrapAround is a wrapper around gtk_source_search_settings_get_wrap_around().
func (v *SourceSearchSettings) GetWrapAround() bool {
	return gobool(C.gtk_source_search_settings_get_wrap_around(v.native()))
}

// SetWrapAround is a wrapper around gtk_source_search_settings_set_wrap_around().
func (v *SourceSearchSettings) SetWrapAround(wrapAround bool) {
	C.gtk_source_search_settings_set_wrap_around(v.native(),
		gbool(wrapAround))
}

// GetRegexEnabled is a wrapper around gtk_source_search_settings_get_regex_enabled().
func (v *SourceSearchSettings) GetRegexEnabled() bool {
	return gobool(C.gtk_source_search_settings_get_regex_enabled(v.native()))
}

// SetRegexEnabled is a wrapper around gtk_source_search_settings_set_regex_enabled().
func (v *SourceSearchSettings) SetRegexEnabled(enabled bool) {
	C.gtk_source_search_settings_set_regex_enabled(v.native(),
		gbool(enabled))
}
