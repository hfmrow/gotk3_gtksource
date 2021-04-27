// +build linux
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14
// +build !gtksourceview_3_18,!gtksourceview_3_20

// Filename version number based on gtk_sourceview

// Limit > gtk_3_14 because libgtksourceview-3.0 start at gtk >= 3.16
// Latest gtksourceview_3_18 start with gtk_3_16
// Latest gtksourceview_4_0 start with gtk_3_22

// Same copyright and license as the rest of the files in this project

package source

// #include <gtk/gtk.h>
// #include <gtksourceview/gtksource.h>
// #include "source.go.h"
// #include "source_facility.go.h"
import "C"
import (
	"github.com/gotk3/gotk3/gtk"
)

/*
 * SourceStyle
 */

// Apply is a wrapper around gtk_source_style_apply().
func (v *SourceStyle) Apply(tag *gtk.TextTag) {
	C.gtk_source_style_apply(v.native(), nativeTextTag(tag))
}
