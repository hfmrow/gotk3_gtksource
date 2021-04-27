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
import "github.com/gotk3/gotk3/gtk"

// SourceChangeCaseType is a representation of GTK's GtkSourceChangeCaseType. Since: 3.12
type SourceChangeCaseType int

const (
	SOURCE_CHANGE_CASE_LOWER  SourceChangeCaseType = C.GTK_SOURCE_CHANGE_CASE_LOWER
	SOURCE_CHANGE_CASE_UPPER  SourceChangeCaseType = C.GTK_SOURCE_CHANGE_CASE_UPPER
	SOURCE_CHANGE_CASE_TOGGLE SourceChangeCaseType = C.GTK_SOURCE_CHANGE_CASE_TOGGLE
	SOURCE_CHANGE_CASE_TITLE  SourceChangeCaseType = C.GTK_SOURCE_CHANGE_CASE_TITLE
)

/*
 * SourceBuffer
 */

// ChangeCase is a wrapper around gtk_source_buffer_change_case().
func (v *SourceBuffer) ChangeCase(caseType SourceChangeCaseType, start, end *gtk.TextIter) {
	C.gtk_source_buffer_change_case(v.native(), C.GtkSourceChangeCaseType(caseType),
		nativeTextIter(start), nativeTextIter(end))
}
