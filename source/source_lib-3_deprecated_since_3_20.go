// +build linux
// +build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_deprecated

/*
 * 		Library loader, please don't modify except if you know what you do.
 */

// Filename version number based on gtk

// GtkSourceView 3.24 is the latest stable GtkSourceView 3 version.

// use line below to limit compilation based on gtk_sourceview actually used.
/* +build gtk_sourceview_3_18 gtk_sourceview_3_20 gtk_sourceview_3_22 gtk_sourceview_3_24 gtk_sourceview_deprecated */

// Same copyright and license as the rest of the files in this project

/*
		- This library loader ensure that the version loaded match your Gtk version -

					this file does not and must not contain any part of code
	require:
		Gtk	> gtk_3_18: libgtksourceview-4-dev -> #cgo pkg-config: gtksourceview-4
		Gtk	< gtk_3_20: libgtksourceview-3.0-dev -> #cgo pkg-config: gtksourceview-3.0
*/

package source

// #cgo pkg-config: gtksourceview-3.0
import "C"
