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
// #include "sourcefile_since_3_14.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func init() {

	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_source_file_saver_get_type()), marshalSourceFileSaver},
		{glib.Type(C.gtk_source_file_get_type()), marshalSourceFile},
		{glib.Type(C.gtk_source_file_loader_get_type()), marshalSourceFileLoader},
		{glib.Type(C.gtk_source_file_saver_error_get_type()), marshalSourceFileSaverError},
		{glib.Type(C.gtk_source_file_saver_flags_get_type()), marshalSourceFileSaverFlags},
		{glib.Type(C.gtk_source_file_loader_error_get_type()), marshalSourceFileLoaderError},
		{glib.Type(C.gtk_source_newline_type_get_type()), marshalSourceNewlineType},
		{glib.Type(C.gtk_source_compression_type_get_type()), marshalSourceCompressionType},
	}

	glib.RegisterGValueMarshalers(tm)

	gtk.WrapMap["GtkSourceFileSaver"] = wrapSourceFileSaver
	gtk.WrapMap["GtkSourceFile"] = wrapSourceFile
	gtk.WrapMap["GtkSourceFileLoader"] = wrapSourceFileLoader
}

// SourceFileSaverError is a representation of GTK's GtkSourceFileSaverError.
type SourceFileSaverError int

const (
	SOURCE_FILE_SAVER_ERROR_INVALID_CHARS       SourceFileSaverError = C.GTK_SOURCE_FILE_SAVER_ERROR_INVALID_CHARS
	SOURCE_FILE_SAVER_ERROR_EXTERNALLY_MODIFIED SourceFileSaverError = C.GTK_SOURCE_FILE_SAVER_ERROR_EXTERNALLY_MODIFIED
)

func marshalSourceFileSaverError(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceFileSaverError(c), nil
}

// SourceFileSaverFlags is a representation of GTK's GtkSourceFileSaverFlags.
type SourceFileSaverFlags int

const (
	SOURCE_FILE_SAVER_FLAGS_NONE                     SourceFileSaverFlags = C.GTK_SOURCE_FILE_SAVER_FLAGS_NONE
	SOURCE_FILE_SAVER_FLAGS_IGNORE_INVALID_CHARS     SourceFileSaverFlags = C.GTK_SOURCE_FILE_SAVER_FLAGS_IGNORE_INVALID_CHARS
	SOURCE_FILE_SAVER_FLAGS_IGNORE_MODIFICATION_TIME SourceFileSaverFlags = C.GTK_SOURCE_FILE_SAVER_FLAGS_IGNORE_MODIFICATION_TIME
	SOURCE_FILE_SAVER_FLAGS_CREATE_BACKUP            SourceFileSaverFlags = C.GTK_SOURCE_FILE_SAVER_FLAGS_CREATE_BACKUP
)

func marshalSourceFileSaverFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceFileSaverFlags(c), nil
}

// SourceFileLoaderError is a representation of GTK's GtkSourceFileLoaderError.
type SourceFileLoaderError int

const (
	SOURCE_FILE_LOADER_ERROR_TOO_BIG                        SourceFileLoaderError = C.GTK_SOURCE_FILE_LOADER_ERROR_TOO_BIG
	SOURCE_FILE_LOADER_ERROR_ENCODING_AUTO_DETECTION_FAILED SourceFileLoaderError = C.GTK_SOURCE_FILE_LOADER_ERROR_ENCODING_AUTO_DETECTION_FAILED
	SOURCE_FILE_LOADER_ERROR_CONVERSION_FALLBACK            SourceFileLoaderError = C.GTK_SOURCE_FILE_LOADER_ERROR_CONVERSION_FALLBACK
)

func marshalSourceFileLoaderError(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceFileLoaderError(c), nil
}

// SourceNewlineType is a representation of GTK's GtkSourceNewlineType.
type SourceNewlineType int

const (
	SOURCE_NEWLINE_TYPE_LF    SourceNewlineType = C.GTK_SOURCE_NEWLINE_TYPE_LF
	SOURCE_NEWLINE_TYPE_CR    SourceNewlineType = C.GTK_SOURCE_NEWLINE_TYPE_CR
	SOURCE_NEWLINE_TYPE_CR_LF SourceNewlineType = C.GTK_SOURCE_NEWLINE_TYPE_CR_LF
)

func marshalSourceNewlineType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceNewlineType(c), nil
}

// SourceCompressionType is a representation of GTK's GtkSourceCompressionType.
type SourceCompressionType int

const (
	SOURCE_COMPRESSION_TYPE_NONE SourceCompressionType = C.GTK_SOURCE_COMPRESSION_TYPE_NONE
	SOURCE_COMPRESSION_TYPE_GZIP SourceCompressionType = C.GTK_SOURCE_COMPRESSION_TYPE_GZIP
)

func marshalSourceCompressionType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return SourceCompressionType(c), nil
}

/*
 * GtkSourceFile
 *
 * Gtk+ Description say:
 * This class is no longer maintained. There is work in progress
 * in the Tepl library to have a better implementation.
 */

// SourceFile is a representation of GTK's GtkSourceFile.
// On-disk representation of a GtkSourceBuffer
type SourceFile struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceFile.
func (v *SourceFile) native() *C.GtkSourceFile {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceFile(p)
}

func marshalSourceFile(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceFile(obj), nil
}

func wrapSourceFile(obj *glib.Object) *SourceFile {
	return &SourceFile{obj}
}

// TODO GtkSourceMountOperationFactory
/* I don't know how to handle this function (look like Callback function)
GMountOperation *
(*GtkSourceMountOperationFactory) (GtkSourceFile *file,
                                   gpointer userdata);
*/

// SourceFileNew is a wrapper around gtk_source_file_new().
func SourceFileNew() (*SourceFile, error) {
	c := C.gtk_source_file_new()
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceFile(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// GetLocation is a wrapper around gtk_source_file_get_location().
func (v *SourceFile) GetLocation() (*glib.File, error) {
	c := C.gtk_source_file_get_location(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := &glib.File{glib.Take(unsafe.Pointer(c))}
	return e, nil
}

// SetLocation is a wrapper around gtk_source_file_set_location().
func (v *SourceFile) SetLocation(location *glib.File) {
	C.gtk_source_file_set_location(
		v.native(),
		(*C.GFile)(location.NativePrivate()))
}

// GetEncoding is a wrapper around gtk_source_file_get_encoding().
func (v *SourceFile) GetEncoding() (*SourceEncoding, error) {
	c := C.gtk_source_file_get_encoding(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := (*SourceEncoding)(unsafe.Pointer(c))
	return e, nil
}

// GetNewlineType is a wrapper around gtk_source_file_get_newline_type().
func (v *SourceFile) GetNewlineType() SourceNewlineType {
	return SourceNewlineType(C.gtk_source_file_get_newline_type(v.native()))
}

// GetCompressionType is a wrapper around gtk_source_file_get_compression_type().
func (v *SourceFile) GetCompressionType() SourceCompressionType {
	return SourceCompressionType(C.gtk_source_file_get_compression_type(v.native()))
}

// TODO require GtkSourceMountOperationFactory
/*
void
gtk_source_file_set_mount_operation_factory
                               (GtkSourceFile *file,
                                GtkSourceMountOperationFactory callback,
                                gpointer user_data,
                                GDestroyNotify notify);
*/

/*
 * GtkSourceFileLoader
 *
 * Gtk+ Description say:
 * This class is no longer maintained. There is work in progress
 * in the Tepl library to have a better implementation.
 */

// SourceFileLoader is a representation of GTK's GtkSourceFileLoader.
// Load a file into a GtkSourceBuffer
type SourceFileLoader struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceFileLoader.
func (v *SourceFileLoader) native() *C.GtkSourceFileLoader {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceFileLoader(p)
}

func marshalSourceFileLoader(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceFileLoader(obj), nil
}

func wrapSourceFileLoader(obj *glib.Object) *SourceFileLoader {
	return &SourceFileLoader{obj}
}

// SourceFileLoaderNew is a wrapper around gtk_source_file_loader_new().
func SourceFileLoaderNew(buffer *SourceBuffer, file *SourceFile) (*SourceFileLoader, error) {
	c := C.gtk_source_file_loader_new(buffer.native(), file.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceFileLoader(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SourceFileLoaderNewFromStream is a wrapper around gtk_source_file_loader_new_from_stream().
func SourceFileLoaderNewFromStream(buffer *SourceBuffer, file *SourceFile, stream *glib.InputStream) (*SourceFileLoader, error) {
	c := C.gtk_source_file_loader_new_from_stream(
		buffer.native(), file.native(),
		(*C.GInputStream)(unsafe.Pointer(stream.Native())))
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceFileLoader(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SetCandidateEncodings is a wrapper around gtk_source_file_loader_set_candidate_encodings().
// The list must be Free with (*glib.SList).Free().
func (v *SourceFileLoader) SetCandidateEncodings(candidate *glib.SList) {
	C.gtk_source_file_loader_set_candidate_encodings(
		v.native(),
		(*C.GSList)(cGSList(candidate)))
}

// GetBuffer is a wrapper around gtk_source_file_loader_get_buffer().
func (v *SourceFileLoader) GetBuffer() (*SourceBuffer, error) {
	c := C.gtk_source_file_loader_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceBuffer(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// GetFile is a wrapper around gtk_source_file_loader_get_file().
func (v *SourceFileLoader) GetFile() (*SourceFile, error) {
	c := C.gtk_source_file_loader_get_file(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceFile(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// GetFileLocation is a wrapper around gtk_source_file_loader_get_location().
func (v *SourceFileLoader) GetFileLocation() (*glib.File, error) {
	return toGFile(C.gtk_source_file_loader_get_location(v.native()))
}

// GetInputStream is a wrapper around gtk_source_file_loader_get_input_stream().
func (v *SourceFileLoader) GetInputStream() (*glib.InputStream, error) {
	c := C.gtk_source_file_loader_get_input_stream(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := &glib.InputStream{glib.Take(unsafe.Pointer(c))}
	return e, nil
}

// TODO Require GFileProgressCallback, GAsyncReadyCallback, GDestroyNotify
/*
void
gtk_source_file_loader_load_async (GtkSourceFileLoader *loader,
                                   gint io_priority,
                                   GCancellable *cancellable,
                                   GFileProgressCallback progress_callback,
                                   gpointer progress_callback_data,
                                   GDestroyNotify progress_callback_notify,
                                   GAsyncReadyCallback callback,
                                   gpointer user_data);
*/

// LoadFinish is a wrapper around gtk_source_file_loader_load_finish().
func (v *SourceFileLoader) LoadFinish(result *glib.AsyncResult) (bool, error) {
	var gerr *C.GError

	c := C.gtk_source_file_loader_load_finish(v.native(), (*C.GAsyncResult)(unsafe.Pointer(result.Native())), &gerr)
	if gerr != nil {
		defer C.g_error_free(gerr)
		return false, errors.New(goString(gerr.message))
	}
	return gobool(c), nil
}

// GetEncoding is a wrapper around gtk_source_file_loader_get_encoding().
func (v *SourceFileLoader) GetEncoding() (*SourceEncoding, error) {
	c := C.gtk_source_file_loader_get_encoding(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := (*SourceEncoding)(unsafe.Pointer(c))
	return e, nil
}

// GetNewLineType is a wrapper around gtk_source_file_loader_get_newline_type().
func (v *SourceFileLoader) GetNewLineType() SourceNewlineType {
	return SourceNewlineType(C.gtk_source_file_loader_get_newline_type(v.native()))
}

// GetCompressionType is a wrapper around gtk_source_file_loader_get_compression_type().
func (v *SourceFileLoader) GetCompressionType() SourceCompressionType {
	return SourceCompressionType(C.gtk_source_file_loader_get_compression_type(v.native()))
}

/*
 * GtkSourceFileSaver
 *
 * Gtk+ Description say:
 * This class is no longer maintained. There is work in progress
 * in the Tepl library to have a better implementation.
 */

// SourceFileSaver is a representation of GTK's GtkSourceFileSaver.
// Save a GtkSourceBuffer into a file
type SourceFileSaver struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkSourceFileSaver.
func (v *SourceFileSaver) native() *C.GtkSourceFileSaver {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkSourceFileSaver(p)
}

func marshalSourceFileSaver(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapSourceFileSaver(obj), nil
}

func wrapSourceFileSaver(obj *glib.Object) *SourceFileSaver {
	return &SourceFileSaver{obj}
}

// SourceFileSaverNew is a wrapper around gtk_source_file_saver_new().
func SourceFileSaverNew(sourceBuffer *SourceBuffer, file *SourceFile) (*SourceFileSaver, error) {
	c := C.gtk_source_file_saver_new(sourceBuffer.native(), file.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceFileSaver(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// SourceFileSaverNewWithTarget is a wrapper around gtk_source_file_saver_new_with_target().
func SourceFileSaverNewWithTarget(buffer *SourceBuffer, file *SourceFile, target *glib.File) (*SourceFileSaver, error) {
	c := C.gtk_source_file_saver_new_with_target(
		buffer.native(),
		file.native(),
		(*C.GFile)(target.NativePrivate()))
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceFileSaver(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// GetBuffer is a wrapper around gtk_source_file_saver_get_buffer().
func (v *SourceFileSaver) GetBuffer() (*SourceBuffer, error) {
	c := C.gtk_source_file_saver_get_buffer(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceBuffer(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// GetFile is a wrapper around gtk_source_file_saver_get_file().
func (v *SourceFileSaver) GetFile() (*SourceFile, error) {
	c := C.gtk_source_file_saver_get_file(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := wrapSourceFile(glib.Take(unsafe.Pointer(c)))
	return e, nil
}

// GetLocation is a wrapper around gtk_source_file_saver_get_location().
func (v *SourceFileSaver) GetLocation() (*glib.File, error) {
	c := C.gtk_source_file_saver_get_location(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := &glib.File{glib.Take(unsafe.Pointer(c))}
	return e, nil
}

// SetEncoding is a wrapper around gtk_source_file_saver_set_encoding().
func (v *SourceFileSaver) SetEncoding(encoding *SourceEncoding) {
	C.gtk_source_file_saver_set_encoding(v.native(), encoding.native())
}

// GetEncoding is a wrapper around gtk_source_file_saver_get_encoding().
func (v *SourceFileSaver) GetEncoding() (*SourceEncoding, error) {
	c := C.gtk_source_file_saver_get_encoding(v.native())
	if c == nil {
		return nil, nilPtrErr
	}

	e := (*SourceEncoding)(unsafe.Pointer(c))
	return e, nil
}

// SetNewlineType is a wrapper around gtk_source_file_saver_set_newline_type().
func (v *SourceFileSaver) SetNewlineType(newlineType SourceNewlineType) {
	C.gtk_source_file_saver_set_newline_type(v.native(),
		C.GtkSourceNewlineType(newlineType))
}

// GetNewlineType is a wrapper around gtk_source_file_saver_get_newline_type().
func (v *SourceFileSaver) GetNewlineType() SourceNewlineType {
	return SourceNewlineType(C.gtk_source_file_saver_get_newline_type(v.native()))
}

// SetCompressionType is a wrapper around gtk_source_file_saver_set_compression_type().
func (v *SourceFileSaver) SetCompressionType(compressionType SourceCompressionType) {
	C.gtk_source_file_saver_set_compression_type(v.native(),
		C.GtkSourceCompressionType(compressionType))
}

// GetCompressionType is a wrapper around gtk_source_file_saver_get_compression_type().
func (v *SourceFileSaver) GetCompressionType() SourceCompressionType {
	return SourceCompressionType(C.gtk_source_file_saver_get_compression_type(v.native()))
}

// SetFlags is a wrapper around gtk_source_file_saver_set_flags().
func (v *SourceFileSaver) SetFlags(flags SourceFileSaverFlags) {
	C.gtk_source_file_saver_set_flags(v.native(),
		C.GtkSourceFileSaverFlags(flags))
}

// GetFlags is a wrapper around gtk_source_file_saver_get_flags().
func (v *SourceFileSaver) GetFlags() SourceFileSaverFlags {
	return SourceFileSaverFlags(C.gtk_source_file_saver_get_flags(v.native()))
}

// TODO require GFileProgressCallback, GAsyncResult, GDestroyNotify, GAsyncReadyCallback
/*
void
gtk_source_file_saver_save_async (GtkSourceFileSaver *saver,
                                  gint io_priority,
                                  GCancellable *cancellable,
                                  GFileProgressCallback progress_callback,
                                  gpointer progress_callback_data,
                                  GDestroyNotify progress_callback_notify,
                                  GAsyncReadyCallback callback,
                                  gpointer user_data);
*/

// ReplaceAll is a wrapper around gtk_source_file_saver_save_finish().
func (v *SourceFileSaver) SaveFinish(resulst *glib.AsyncResult) error {
	var err *C.GError = nil

	C.gtk_source_file_saver_save_finish(
		v.native(), (*C.GAsyncResult)(unsafe.Pointer(resulst.Native())), &err)

	if err != nil {
		defer C.g_error_free(err)
		return errors.New(goString(err.message))
	}
	return nil
}
