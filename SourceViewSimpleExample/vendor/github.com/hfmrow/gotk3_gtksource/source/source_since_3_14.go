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
import "C"

/*
 * SourceBuffer
 */

// SetImplicitTrailingNewLine is a wrapper around gtk_source_buffer_set_implicit_trailing_newline().
func (v *SourceBuffer) SetImplicitTrailingNewLine(impTrailNewLine bool) {
	C.gtk_source_buffer_set_implicit_trailing_newline(
		v.native(),
		gbool(impTrailNewLine))
}

// GetImplicitTrailingNewLine is a wrapper around gtk_source_buffer_get_implicit_trailing_newline().
func (v *SourceBuffer) GetImplicitTrailingNewLine() bool {
	return gobool(
		C.gtk_source_buffer_get_implicit_trailing_newline(
			v.native()))
}
