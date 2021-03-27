// textViewNumbered.go

/*
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2019 H.F.M - TextView Numbered  library "https//github/hfmrow"
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php

	This source code is a part of a personal library using gotk3.

	Allow using a scrolled TextView with line numbers (left sided) with some formatting text
	controls

	There is a limitation regarding some types of text containing non-printable characters
	(some fonts print Utf-8 codes instead of character), does not work properly, the height
	of the line numbers is not correctly synchronized with the text. These characters print
	an oversized entry in the textView column, so the numbers are shifted relative to the
	lines of text.
*/

package sourceView

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	gimc "github.com/hfmrow/gtk3Import/misc"
	gitvtt "github.com/hfmrow/gtk3Import/textView/textTag"

	"github.com/hfmrow/gotk3_gtksource/source"
)

type SourceViewStruct struct {
	Window    gtk.IWindow
	View      *source.SourceView
	Buffer    *source.SourceBuffer
	Map       *source.SourceMap
	SearchSet *source.SourceSearchSettings
	SearchCtx *source.SourceSearchContext

	Language           *source.SourceLanguage
	LanguageManager    *source.SourceLanguageManager
	StyleScheme        *source.SourceStyleScheme
	StyleSchemeManager *source.SourceStyleSchemeManager

	// User style and language paths
	UserStylePath,
	UserLanguagePath string

	StyleShemeIds,
	LanguageIds []string

	DefaultStyleShemeId,
	DefaultLanguageId string

	FontSize int
	FontFamily,
	colorBgRangeName,
	ColorBgRangeSet,

	TxtFgCol,
	TxtBgCol,
	SelFgCol,
	SelBgCol string

	langAndStyleInitialized bool

	/*
	 * Search settings
	 */
	TextSearch,
	previousTextSearch,
	TextReplace string

	// Search options
	UseRegexp,
	WrapAround,
	CaseSensitive,
	WordBoundaries,
	HighlightFound bool

	// Search results
	IterStart,
	IterEnd *gtk.TextIter
	HasWrappedAround,
	HasFound,
	SearchOk bool

	OccurencesCount int

	SearchError error

	styleSchemeChooserWidget *source.SourceStyleSchemeChooserWidget
	dialog                   *gtk.Dialog
}

// SourceViewStructNew:
func SourceViewStructNew(sourceView *source.SourceView, sourceMap *source.SourceMap, parentWindow ...gtk.IWindow) (svs *SourceViewStruct, err error) {

	svs = new(SourceViewStruct)

	if len(parentWindow) > 0 {
		svs.Window = parentWindow[0]
	}

	svs.View = sourceView
	svs.Map = sourceMap

	if svs.Buffer, err = svs.View.GetBuffer(); err == nil {

		svs.View.SetLeftMargin(4)
		svs.View.SetShowLineNumbers(true)

		// css
		svs.TxtFgCol = "#331111"
		svs.TxtBgCol = "#F8F8F8"
		svs.SelFgCol = "#152727"
		svs.SelBgCol = "#CBEBEB"

		svs.FontSize = 12
		svs.FontFamily = `"Liberation Mono", sans-serif`
		svs.colorBgRangeName = "highlightBCColorSoftLightGreen"
		svs.ColorBgRangeSet = "#E6FFE6" // lightgreen
		svs.UpdateCss()

		// search
		svs.WrapAround = true
		// svs.CaseSensitive=
		svs.HighlightFound = true

		// Language & style, Set default
		svs.DefaultStyleShemeId = "classic"
		svs.DefaultLanguageId = "go"

		// Make a tag to indicate found element (when HighlightFound not checked)
		// tag := make(map[string]interface{})
		// tag["background"] = "#ABF6FF"
		// markFound = svs.Buffer.CreateTag("markFound", tag)

	}
	return
}

// ComboboxHandling: 'defaultValues' if present must be: Language followed by Style
// they're the main options (stored values) that will be updated each time one of them are changed.
func (svs *SourceViewStruct) ComboboxHandling(cbxtStyle, cbxtLang *gtk.ComboBoxText, defaultLang, defaultStyle *string) {

	// Setting Language and style scheme
	if currentLanguage := svs.SetLanguage(svs.DefaultLanguageId); currentLanguage != nil {
		if currentStyleScheme := svs.SetStyleScheme(svs.DefaultStyleShemeId); currentStyleScheme != nil {

			// Fill comboboxes with languages and styles
			for _, id := range svs.LanguageIds {
				cbxtLang.AppendText(id)
			}
			for _, id := range svs.StyleShemeIds {
				cbxtStyle.AppendText(id)
			}

			// Just indicate id must be set as first model column.
			cbxtLang.SetIDColumn(0)
			cbxtStyle.SetIDColumn(0)

			// Set ComboBox current values display.
			cbxtLang.SetActiveID(svs.DefaultLanguageId)
			cbxtStyle.SetActiveID(svs.DefaultStyleShemeId)

			// Signals changed
			cbxtLang.Connect("changed", func(cbxt *gtk.ComboBoxText) {
				*defaultLang = cbxt.GetActiveID()
				svs.SetLanguage(*defaultLang)
			})
			cbxtStyle.Connect("changed", func(cbxt *gtk.ComboBoxText) {
				*defaultStyle = cbxt.GetActiveID()
				svs.SetStyleScheme(*defaultStyle)
			})

		}
	}
}

// SetLanguage:
func (svs *SourceViewStruct) SetLanguage(id string) *source.SourceLanguage {

	var err error

	if err = svs.initLanguageAndStyleScheme(); err == nil {

		// Set Language scheme to sourceBuffer
		if !isExistSlice(svs.LanguageIds, id) {
			id = svs.DefaultLanguageId
		}

		if svs.Language, err = svs.LanguageManager.GetLanguage(id); err == nil {
			svs.Buffer.SetLanguage(svs.Language)
		}

	}

	if err != nil {
		return nil
	}

	return svs.Language
}

// SetStyleScheme:
func (svs *SourceViewStruct) SetStyleScheme(id string) *source.SourceStyleScheme {

	var err error

	if err = svs.initLanguageAndStyleScheme(); err == nil {

		// Set style scheme to sourceBuffer
		if !isExistSlice(svs.StyleShemeIds, id) {
			id = svs.DefaultStyleShemeId
		}

		if svs.StyleScheme, err = svs.StyleSchemeManager.GetScheme(id); err == nil {
			svs.Buffer.SetStyleSheme(svs.StyleScheme)
		}
	}

	if err != nil {
		return nil
	}

	return svs.StyleScheme
}

// initLanguageAndStyleScheme: Initialize Language and StylScheme
func (svs *SourceViewStruct) initLanguageAndStyleScheme() (err error) {

	if !svs.langAndStyleInitialized {

		var searchPaths []string

		if svs.StyleSchemeManager, err = source.SourceStyleSchemeManagerGetDefault(); err == nil {

			// Setting User defined StyleScheme (if exist)
			// There is another (simplest) method to doing that, (AppendSearchPath)
			// but 'LanguageManager' doesn't have it. For clarity, i use same way
			// for both in this part.
			if _, err = os.Stat(svs.UserStylePath); !os.IsNotExist(err) {
				searchPaths = svs.StyleSchemeManager.GetSearchPath()
				searchPaths = append(searchPaths, svs.UserStylePath)
				svs.StyleSchemeManager.SetSearchPath(searchPaths)
			}
			err = nil

			// getting available list of ids
			svs.StyleShemeIds = svs.StyleSchemeManager.GetShemeIds()

			if svs.LanguageManager, err = source.SourceLanguageManagerGetDefault(); err == nil {

				// Setting User defined Language  (if exist)
				if _, err = os.Stat(svs.UserLanguagePath); !os.IsNotExist(err) {
					searchPaths = svs.LanguageManager.GetSearchPath()
					searchPaths = append(searchPaths, svs.UserLanguagePath)
					svs.LanguageManager.SetSearchPath(searchPaths)
				}
				err = nil

				// getting available list of ids
				svs.LanguageIds = svs.LanguageManager.GetLanguageIds()

				// Simply indicate the initialization was already done.
				svs.langAndStyleInitialized = true
			}
		}
	}

	return
}

// StyleChooserDialog: I don't know how to get result from a SourceStyleSchemeChooserButton object,
// so i made my own dialog who can give me a result that permit to set new styleScheme.
func (svs *SourceViewStruct) StyleChooserDialog() string {
	var err error

	if svs.dialog == nil {

		// Init SourceStyleSchemeChooserWidget (Only first time).
		if svs.styleSchemeChooserWidget, err = source.SourceStyleSchemeChooserWidgetNew(); err == nil {

			svs.styleSchemeChooserWidget.SetStyleScheme(svs.StyleScheme)

			// Create Dialog.
			if svs.dialog, err = gtk.DialogNewWithButtons("Choose scheme style",
				svs.Window, gtk.DIALOG_DESTROY_WITH_PARENT,
				[]interface{}{"Cancel", gtk.RESPONSE_CANCEL},
				[]interface{}{"Accept", gtk.RESPONSE_ACCEPT}); err == nil {

				// Add widget to Dialog.
				if box, err := svs.dialog.GetContentArea(); err == nil {
					box.Add(svs.styleSchemeChooserWidget)
				}
			}
		}
	}
	svs.styleSchemeChooserWidget.ShowAll()

	// Launch Dialog.
	if resp := svs.dialog.Run(); resp == gtk.RESPONSE_ACCEPT {
		if svs.StyleScheme, err = svs.styleSchemeChooserWidget.GetStyleScheme(); err == nil {
			svs.Buffer.SetStyleSheme(svs.StyleScheme)
		}
	}
	svs.dialog.Hide()

	// On error just stop.
	if err != nil {
		log.Fatal(err)
	}

	// Return ComboBox Style name (id).
	return svs.StyleScheme.GetId()
}

// initSearch: Init SearchSettings & SearchContext, set options
func (svs *SourceViewStruct) initSearch() (err error) {

	// First time only init SearchSettings and SearchContext
	if svs.SearchCtx == nil {
		if svs.SearchSet, err = source.SourceSearchSettingsNew(); err == nil {
			svs.SearchCtx, err = source.SourceSearchContextNew(svs.Buffer, svs.SearchSet)
		}
	}

	// Set options
	if err == nil {
		svs.SearchSet.SetSearchText(svs.TextSearch)
		svs.SearchSet.SetRegexEnabled(svs.UseRegexp)
		svs.SearchSet.SetWordBoundaries(svs.WordBoundaries)
		svs.SearchSet.SetCaseSensitive(svs.CaseSensitive)
		svs.SearchSet.SetWrapAround(svs.WrapAround)
		svs.SearchCtx.SetHighLight(svs.HighlightFound)

		// At least one search operation has been done
		svs.SearchOk = true
	}
	return
}

// Search: Simple search method.
func (svs *SourceViewStruct) Search(iter *gtk.TextIter, backward bool) (err error) {

	// Init search engine
	if err = svs.initSearch(); err == nil {

		switch {

		case !backward:

			// Do a Forward search
			svs.IterStart, svs.IterEnd, svs.HasWrappedAround, svs.HasFound = svs.SearchCtx.Forward(iter)

		case backward:

			// Do a Backward search
			svs.IterStart, svs.IterEnd, svs.HasWrappedAround, svs.HasFound = svs.SearchCtx.Backward(iter)

		}
	}

	return
}

// Search: perform asynchronous search, callback can be nil.
// Cancel callback not used here, but is implemented and you can use it
// with your own function. Same with UserData, as uintPtr.
func (svs *SourceViewStruct) SearchAsync(callback func(), iter *gtk.TextIter, backward bool) (err error) {

	// Init search engine
	if err = svs.initSearch(); err == nil {

		switch {

		case !backward:

			// Do a ForwardAsync search
			svs.SearchCtx.ForwardAsync(iter, nil,
				func(object *glib.Object, res *glib.AsyncResult) {

					if svs.IterStart,
						svs.IterEnd,
						svs.HasWrappedAround,
						svs.HasFound,
						svs.SearchError = svs.SearchCtx.ForwardFinish(res); svs.SearchError != nil {

						if errIn := svs.SearchCtx.GetRegexError(); errIn != nil {
							svs.SearchError = fmt.Errorf("[%v][%v]", svs.SearchError, errIn)
						}

					} else {
						svs.OccurencesCount = svs.SearchCtx.GetOccurencesCount()
					}
					if callback != nil {

						// callback function can be used to signify the end of
						// the search and check for errors via 'svs.SearchError'.
						// All results are stored in the structure and can be easily accessed.
						callback()
					}
				})

		case backward:

			// Do a BackwardAsync search
			svs.SearchCtx.BackwardAsync(iter, nil,
				func(object *glib.Object, res *glib.AsyncResult) {

					if svs.IterStart,
						svs.IterEnd,
						svs.HasWrappedAround,
						svs.HasFound,
						svs.SearchError = svs.SearchCtx.BackwardFinish(res); svs.SearchError != nil {

						if errIn := svs.SearchCtx.GetRegexError(); errIn != nil {
							svs.SearchError = fmt.Errorf("[%v][%v]", svs.SearchError, errIn)
						}

					} else {
						svs.OccurencesCount = svs.SearchCtx.GetOccurencesCount()
					}
					if callback != nil {

						// callback function can be used to signify the end of
						// the search and check for errors via 'svs.SearchError'.
						// All results are stored in the structure and can be easily accessed.
						callback()
					}
				})
		}
	}

	return
}

// LoadSource: Load file and search if requested.
func (svs *SourceViewStruct) LoadSource(filename string, doSearch ...bool) (err error) {
	var data []byte
	var search bool

	if len(doSearch) > 0 {
		search = doSearch[0]
	}

	if data, err = ioutil.ReadFile(filename); err == nil {
		svs.SetText(string(data))

		if search {
			err = svs.SearchAsync(nil, svs.Buffer.GetStartIter(), false)
		}
	}
	return
}

/*
 * These methods below are not really part of GtkSource and don't
 * require special attention unless you want them to.
 */

// doColored: Add desired font and some colors to the numbers column and selected text.
func (svs *SourceViewStruct) UpdateCss() {

	gimc.CssWdgScnLoad(`
* {
	font: `+fmt.Sprintf("%d", svs.FontSize)+`px `+svs.FontFamily+`;
}

* text {
	color: `+svs.TxtFgCol+`;
	background-color: `+svs.TxtBgCol+`;
}
* text selection {
	background-color: `+svs.SelBgCol+`;
	color: `+svs.SelFgCol+`;
}`, &svs.View.Widget)
}

// ScrollToLine: Scroll to line and return the corresponding iter
func (svs *SourceViewStruct) ScrollToLine(line int) (iter *gtk.TextIter) {

	if line > 0 && line < svs.Buffer.GetLineCount() {
		iter = svs.Buffer.GetIterAtLine(line)

		svs.View.ScrollToIter(iter, 0.0, true, 0.5, 0.5)
	}

	return
}

// ScrollToIter: Scroll to iter and return the corresponding line
func (svs *SourceViewStruct) ScrollToIter(iter *gtk.TextIter) int {

	if !iter.InRange(svs.Buffer.GetStartIter(), svs.Buffer.GetEndIter()) {
		fmt.Println("Iter is out of range")
		return -1
	}

	svs.View.ScrollToIter(iter, 0.0, true, 0.5, 0.5)

	return iter.GetLine()
}

// GetLineNumber: Get current line number at the cursor position
func (svs *SourceViewStruct) GetCurrentLineNb() int {

	iter := svs.Buffer.GetIterAtMark(svs.Buffer.GetInsert())

	if iter != svs.Buffer.GetEndIter() {
		return iter.GetLine()
	}
	return -1
}

// SetCursorAtLine: Set current line number & place cursor on it.
func (svs *SourceViewStruct) SetCursorAtLine(line int) (iter *gtk.TextIter) {

	iter = svs.Buffer.GetIterAtLine(line)
	svs.Buffer.PlaceCursor(iter)
	svs.View.GrabFocus()
	return
}

// SelectRange: select Lines
func (svs *SourceViewStruct) SelectRange(startLine, endLine int) {

	if startLine <= endLine && startLine > 0 {
		startIter := svs.Buffer.GetIterAtLine(startLine)
		endIter := svs.Buffer.GetIterAtOffset(svs.Buffer.GetIterAtLine(endLine).GetOffset() - 1)
		svs.Buffer.SelectRange(startIter, endIter)
	}
}

// ColorBgRange: apply colored background to lines range.
func (svs *SourceViewStruct) ColorBgRange(startLine, endLine int) {

	if startLine <= endLine && startLine > 0 {
		tag := gitvtt.TagCreateIfNotExists(
			svs.Buffer, svs.colorBgRangeName,
			map[string]interface{}{"background": svs.ColorBgRangeSet})

		startIter := svs.Buffer.GetIterAtLine(startLine)
		endIter := svs.Buffer.GetIterAtOffset(svs.Buffer.GetIterAtLine(endLine).GetOffset() - 1)
		svs.Buffer.ApplyTag(tag, startIter, endIter)
	}
}

// GetText: retrieve text from buffer.
func (svs *SourceViewStruct) GetText() (text string) {

	var err error

	if text, err = svs.Buffer.GetText(svs.Buffer.GetStartIter(), svs.Buffer.GetEndIter(), true); err != nil {
		fmt.Printf("SourceViewStruct:GetText: %s\n", err.Error)
	}
	return
}

// GetText: update css and set text to buffer.
func (svs *SourceViewStruct) SetText(text string) {

	svs.UpdateCss()
	svs.Buffer.SetText(text)
}

// GetOccurences: Using TimeoutAdd function, callback will be executed
// when occurences count is available.
func (svs *SourceViewStruct) GetOccurences(callback func(occurences int)) {

	glib.TimeoutAdd(uint(64), func() bool {

		callback(svs.SearchCtx.GetOccurencesCount())

		return svs.SearchCtx.GetOccurencesCount() < 0
	})

	/* example:
	svs.GetOccurences(func(occ int) {
		fmt.Println(occ)
	})
	*/
}

// RunAfterEvents: Execute event pending before f()
// usually used when Gtk does not refresh in time and command
// miss to be executed (ScrollTo ...)
func (svs *SourceViewStruct) RunAfterEvents(f func()) {

	for gtk.EventsPending() {
		gtk.MainIteration() // Wait for pending events (until widget redrawn)
	}

	if f != nil {
		f()
	}
	// glib.IdleAdd(func() {
	// 	var count int
	// 	glib.TimeoutAdd(uint(64), func() bool {
	// 		count++
	// 		svs.View.ScrollToIter(iter, 0.0, true, 0.5, 0.5)
	// 		return count <= 5
	// 	})
	// })
}

// BringToFront: Set window position to be over all others windows
// without staying on top whether another window come to be selected.
func (svs *SourceViewStruct) BringToFront() {
	if svs.Window != nil {
		svs.Window.ToWindow().Deiconify()
		svs.Window.ToWindow().ShowAll()
		svs.Window.ToWindow().GrabFocus()
	}
}

// IsExistSlice: if exist then  ...
func isExistSlice(slice []string, item string) bool {
	for _, mainRow := range slice {
		if mainRow == item {
			return true
		}
	}
	return false
}
