// +build linux
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14
// +build !gtksourceview_3_18,!gtksourceview_3_20,!gtksourceview_3_22

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
	"github.com/gotk3/gotk3/gtk"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_source_space_drawer_get_type()), marshalSourceSpaceDrawer},
		{glib.Type(C.gtk_source_space_type_flags_get_type()), marshalSourceSpaceTypeFlags},
		{glib.Type(C.gtk_source_space_location_flags_get_type()), marshalSourceSpaceLocationFlags},
	}

	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceSpaceDrawer"] = wrapSourceSpaceDrawer
}

/*
 * SourceCompletionItem
 */

// SetLabel is a wrapper around gtk_source_completion_item_set_label().
func (v *SourceCompletionItem) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_completion_item_set_label(v.native(), (*C.gchar)(cstr))
}

// SetMarkup is a wrapper around gtk_source_completion_item_set_markup().
func (v *SourceCompletionItem) SetMarkup(markup string) {
	cstr := C.CString(markup)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_completion_item_set_markup(v.native(), (*C.gchar)(cstr))
}

// SetText is a wrapper around gtk_source_completion_item_set_text().
func (v *SourceCompletionItem) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_completion_item_set_text(v.native(), (*C.gchar)(cstr))
}

// SetIcon is a wrapper around gtk_source_completion_item_set_icon().
func (v *SourceCompletionItem) SetIcon(icon *gdk.Pixbuf) {

	C.gtk_source_completion_item_set_icon(
		v.native(), (*C.GdkPixbuf)(unsafe.Pointer(icon.Native())))
}

// SetIconName is a wrapper around gtk_source_completion_item_set_icon_name().
func (v *SourceCompletionItem) SetIconName(iconName string) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_completion_item_set_icon_name(v.native(), (*C.gchar)(cstr))
}

// SetGIcon is a wrapper around gtk_source_completion_item_set_gicon().
func (v *SourceCompletionItem) SetGIcon(icon *glib.Icon) {

	C.gtk_source_completion_item_set_gicon(
		v.native(), (*C.GIcon)(unsafe.Pointer(icon.Native())))
}

// SetInfo is a wrapper around gtk_source_completion_item_set_info().
func (v *SourceCompletionItem) SetInfo(info string) {
	cstr := C.CString(info)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_completion_item_set_info(v.native(), (*C.gchar)(cstr))
}

/*
 * SourceView
 */

// GetSpaceDrawer is a wrapper around gtk_source_view_get_space_drawer().
func (v *SourceView) GetSpaceDrawer() (*SourceSpaceDrawer, error) {
	c := C.gtk_source_view_get_space_drawer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceSpaceDrawer(glib.Take(unsafe.Pointer(c))), nil
}

/*
 * SourceGutter
 */

// GetView is a wrapper around gtk_source_gutter_get_view().
func (v *SourceGutter) GetView() (*SourceView, error) {
	c := C.gtk_source_gutter_get_view(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceView(glib.Take(unsafe.Pointer(c))), nil
}

// GetWindowType is a wrapper around gtk_source_gutter_get_window_type().
func (v *SourceGutter) GetWindowType() gtk.TextWindowType {
	return gtk.TextWindowType(C.gtk_source_gutter_get_window_type(v.native()))
}

/*
 * GtkSourceSpaceDrawer
 */

// GtkSourceSpaceTypeFlags contains flags for white space types.
type SourceSpaceTypeFlags int

const (
	SOURCE_SPACE_TYPE_NONE    SourceSpaceTypeFlags = C.GTK_SOURCE_SPACE_TYPE_NONE
	SOURCE_SPACE_TYPE_SPACE   SourceSpaceTypeFlags = C.GTK_SOURCE_SPACE_TYPE_SPACE
	SOURCE_SPACE_TYPE_TAB     SourceSpaceTypeFlags = C.GTK_SOURCE_SPACE_TYPE_TAB
	SOURCE_SPACE_TYPE_NEWLINE SourceSpaceTypeFlags = C.GTK_SOURCE_SPACE_TYPE_NEWLINE
	SOURCE_SPACE_TYPE_NBSP    SourceSpaceTypeFlags = C.GTK_SOURCE_SPACE_TYPE_NBSP
	SOURCE_SPACE_TYPE_ALL     SourceSpaceTypeFlags = C.GTK_SOURCE_SPACE_TYPE_ALL
)

func marshalSourceSpaceTypeFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceSpaceTypeFlags(c), nil
}

// GtkSourceSpaceLocationFlags contains flags for white space locations.
type SourceSpaceLocationFlags int

const (
	SOURCE_SPACE_LOCATION_NONE        SourceSpaceLocationFlags = C.GTK_SOURCE_SPACE_LOCATION_NONE
	SOURCE_SPACE_LOCATION_LEADING     SourceSpaceLocationFlags = C.GTK_SOURCE_SPACE_LOCATION_LEADING
	SOURCE_SPACE_LOCATION_INSIDE_TEXT SourceSpaceLocationFlags = C.GTK_SOURCE_SPACE_LOCATION_INSIDE_TEXT
	SOURCE_SPACE_LOCATION_TRAILING    SourceSpaceLocationFlags = C.GTK_SOURCE_SPACE_LOCATION_TRAILING
	SOURCE_SPACE_LOCATION_ALL         SourceSpaceLocationFlags = C.GTK_SOURCE_SPACE_LOCATION_ALL
)

func marshalSourceSpaceLocationFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceSpaceLocationFlags(c), nil
}

// SourceSpaceDrawer is a representation of GTK's GtkSourceSpaceDrawer.
// Represent white space characters with symbols
type SourceSpaceDrawer struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceSpaceDrawer.
func (v *SourceSpaceDrawer) native() *C.GtkSourceSpaceDrawer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceSpaceDrawer(p)
}

func marshalSourceSpaceDrawer(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceSpaceDrawer(obj), nil
}

func wrapSourceSpaceDrawer(obj *glib.Object) *SourceSpaceDrawer {
	return &SourceSpaceDrawer{obj}
}

// SourceSpaceDrawerNew is a wrapper around gtk_source_space_drawer_new().
func SourceSpaceDrawerNew() (*SourceSpaceDrawer, error) {
	c := C.gtk_source_space_drawer_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapSourceSpaceDrawer(glib.Take(unsafe.Pointer(c))), nil
}

// SetTypesForLocations is a wrapper around gtk_source_space_drawer_set_types_for_locations().
func (v *SourceSpaceDrawer) SetTypesForLocations(locations SourceSpaceLocationFlags, types SourceSpaceTypeFlags) {
	C.gtk_source_space_drawer_set_types_for_locations(
		v.native(), C.GtkSourceSpaceLocationFlags(locations), C.GtkSourceSpaceTypeFlags(types))
}

// GetTypesForLocations is a wrapper around gtk_source_space_drawer_get_types_for_locations().
func (v *SourceSpaceDrawer) GetTypesForLocations(locations SourceSpaceLocationFlags) SourceSpaceTypeFlags {
	c := C.gtk_source_space_drawer_get_types_for_locations(
		v.native(), C.GtkSourceSpaceLocationFlags(locations))
	return SourceSpaceTypeFlags(c)
}

// SetMatrix is a wrapper around gtk_source_space_drawer_set_matrix().
func (v *SourceSpaceDrawer) SetMatrix(matrix *glib.Variant) {
	C.gtk_source_space_drawer_set_matrix(
		v.native(), (*C.GVariant)(unsafe.Pointer(matrix.Native())))
}

// GetMatrix is a wrapper around gtk_source_space_drawer_get_matrix().
func (v *SourceSpaceDrawer) GetMatrix() (*glib.Variant, error) {
	c := C.gtk_source_space_drawer_get_matrix(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return glib.TakeVariant(unsafe.Pointer(c)), nil
}

// SetEnableMatrix is a wrapper around gtk_source_space_drawer_set_enable_matrix().
func (v *SourceSpaceDrawer) SetEnableMatrix(enable bool) {
	C.gtk_source_space_drawer_set_enable_matrix(v.native(), gbool(enable))
}

// GetEnableMatrix is a wrapper around gtk_source_space_drawer_get_enable_matrix().
func (v *SourceSpaceDrawer) GetEnableMatrix() bool {
	return gobool(C.gtk_source_space_drawer_get_enable_matrix(v.native()))
}

// BindMatrixSetting is a wrapper around gtk_source_space_drawer_bind_matrix_setting().
func (v *SourceSpaceDrawer) BindMatrixSetting(
	settings *glib.Settings, key string, flag glib.SettingsBindFlags) {

	cstr := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_source_space_drawer_bind_matrix_setting(v.native(),
		(*C.GSettings)(unsafe.Pointer(settings.Native())),
		(*C.gchar)(cstr), C.GSettingsBindFlags(flag))
}
