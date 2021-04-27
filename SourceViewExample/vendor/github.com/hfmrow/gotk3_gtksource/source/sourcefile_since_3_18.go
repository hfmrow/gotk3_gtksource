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

/*
 * SourceFile
 */

// CheckFileOnDisk is a wrapper around gtk_source_file_check_file_on_disk().
func (v *SourceFile) CheckFileOnDisk() {
	C.gtk_source_file_check_file_on_disk(v.native())
}

// IsLocal is a wrapper around gtk_source_file_is_local().
func (v *SourceFile) IsLocal() bool {
	return gobool(C.gtk_source_file_is_local(v.native()))
}

// IsExternallyModified is a wrapper around gtk_source_file_is_externally_modified().
func (v *SourceFile) IsExternallyModified() bool {
	return gobool(C.gtk_source_file_is_externally_modified(v.native()))
}

// IsDeleted is a wrapper around gtk_source_file_is_deleted().
func (v *SourceFile) IsDeleted() bool {
	return gobool(C.gtk_source_file_is_deleted(v.native()))
}

// IsReadOnly is a wrapper around gtk_source_file_is_readonly().
func (v *SourceFile) IsReadOnly() bool {
	return gobool(C.gtk_source_file_is_readonly(v.native()))
}
