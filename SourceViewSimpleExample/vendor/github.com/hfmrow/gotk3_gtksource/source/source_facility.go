package source

// #include <glib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// #include <gio/gio.h>
// #include <gtksourceview/gtksource.h>
// #include "source_facility.go.h"
// #include "source_facility.gtk.go.h"
import "C"
import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

/*
 * Helpers Type conversions
 */

// toCgcharArray: give a string[] and get a ** gchar. Return nil for empty
// input slice. C.destroy_strings(cArray) may be used after using if needed.
func toCgcharArray(in []string) (cArray **C.gchar) {
	if len(in) > 0 {
		cArray = C.make_strings(C.int(len(in) + 1))
		for i, str := range in {
			cstr := C.CString(str)
			C.set_string(cArray, C.int(i), (*C.gchar)(cstr))
		}
		C.set_string(cArray, C.int(len(in)), nil)
	}
	return
}

// toGoStringArrayNoFree: get []string from ** C.gchar and don't
// free it, as the returned string is owned by called function
// and should not be freed or modified.
func toGoStringArrayNoFree(c **C.gchar) []string {
	var strs []string
	if c != nil {
		for *c != nil {
			strs = append(strs, C.GoString((*C.char)(*c)))
			c = C.next_gcharptr(c)
		}
	}
	return strs
}

// toGoStringFree: C.g_free(*C.gchar) and return goString
// replace nil with the return of an empty string
func toGoStringFree(c *C.gchar) string {
	var s string
	if c != nil {
		s = C.GoString((*C.char)(c))
		defer C.g_free((C.gpointer)(c))
	}

	return s
}

// toCGList:  A List must be manually freed by either calling Free() or FreeFull()
func toCGList(glist *glib.List) *C.GList {
	if glist == nil {
		return nil
	}
	return (*C.GList)(unsafe.Pointer(glist.Native()))
}

/*
 * Gtk Native
 */

// nativeWidget:
func nativeWidget(v *gtk.Widget) *C.GtkWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWidget(p)
}

// nativeTextTag:
func nativeTextTag(v *gtk.TextTag) *C.GtkTextTag {
	if v == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	// return (*C.GtkTextTag)(unsafe.Pointer(v))
	return C.toGtkTextTag(p)
}

// nativeTextTagTable
func nativeTextTagTable(v *gtk.TextTagTable) *C.GtkTextTagTable {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkTextTagTable(p)
}

// nativeTextIter
func nativeTextIter(v *gtk.TextIter) *C.GtkTextIter {
	if v == nil {
		return nil
	}
	return (*C.GtkTextIter)(unsafe.Pointer(v))
}

/*
 *	Native to gotk3
 */

// toPixbuff:
func toPixbuf(c *C.GdkPixbuf) (*gdk.Pixbuf, error) {
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}

	p := &gdk.Pixbuf{obj}
	runtime.SetFinalizer(p, func(_ interface{}) { obj.Unref() })
	return p, nil
}

// toGIcon:
func toGIcon(c *C.GIcon) (*glib.Icon, error) {
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	i := &glib.Icon{obj}
	runtime.SetFinalizer(i, func(_ interface{}) { obj.Unref() })
	return i, nil
}

// toGFile:
func toGFile(c *C.GFile) (*glib.File, error) {
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	i := &glib.File{obj}
	runtime.SetFinalizer(i, func(_ interface{}) { obj.Unref() })
	return i, nil
}

// toTextTag:
func toTextTag(c *C.GtkTextTag) *gtk.TextTag {
	if c == nil {
		return nil
	}
	return &gtk.TextTag{glib.Take(unsafe.Pointer(c))}
}

/*
 *	Imported functions
 */

// Here there are a few methods (the minimum possible) of glib, gio ...
// which are needed to implement some methods in GtkSourceView for example.

/*
 * GAsyncReadyCallback for use inside gtk package
 *
 * Definition copied from the glib directory
 */

// AsyncReadyCallback is a representation of GAsyncReadyCallback
// type AsyncReadyCallback func(object *glib.Object, res *glib.AsyncResult, data uintptr)
type asyncReadyCallbackData struct {
	fn       glib.AsyncReadyCallback
	userData uintptr
}

var (
	asyncReadyCallbackRegistry = struct {
		sync.RWMutex
		next int
		m    map[int]asyncReadyCallbackData
	}{
		next: 1,
		m:    make(map[int]asyncReadyCallbackData),
	}
)

func registerAsyncReadyCallback(fn glib.AsyncReadyCallback, userData uintptr) int {
	asyncReadyCallbackRegistry.Lock()
	id := asyncReadyCallbackRegistry.next
	asyncReadyCallbackRegistry.next++
	asyncReadyCallbackRegistry.m[id] = asyncReadyCallbackData{fn: fn, userData: userData}
	asyncReadyCallbackRegistry.Unlock()

	return id
}

/*
 * Export
 * for exported functions, an underscore is added at start to avoid a cgo multiple declaration error.
 */

//export _goAsyncReadyCallbacks
func _goAsyncReadyCallbacks(sourceObject *C.GObject, res *C.GAsyncResult, userData C.gpointer) {
	id := int(uintptr(userData))

	asyncReadyCallbackRegistry.Lock()
	r := asyncReadyCallbackRegistry.m[id]
	//delete(asyncReadyCallbackRegistry.m, id)
	asyncReadyCallbackRegistry.Unlock()

	var source *glib.Object
	if sourceObject != nil {
		source = glib.Take(unsafe.Pointer(sourceObject))
	}

	r.fn(source, &glib.AsyncResult{glib.Take(unsafe.Pointer(res))})
}
