// example.go

/*
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2020 H.F.M - SourceView simple example
	This program and the mapped GtkSourceView library comes with absolutely no warranty.
	See the The MIT License (MIT) for details: https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hfmrow/gotk3_gtksource/source"

	"github.com/gotk3/gotk3/gtk"
)

var (
	sv     *source.SourceView
	buff   *source.SourceBuffer
	slm    *source.SourceLanguageManager
	lng    *source.SourceLanguage
	sssm   *source.SourceStyleSchemeManager
	scheme *source.SourceStyleScheme
	style  *source.SourceStyle

	win  *gtk.Window
	scrl *gtk.ScrolledWindow
	grid *gtk.Grid
	box  *gtk.Box

	lbl,
	lblLang,
	lblStyle,
	lblVersion *gtk.Label

	cbtxStyle,
	cbtxlang *gtk.ComboBoxText

	chkbtnHl,
	chkbtnLn,
	chkbtnMb,
	chkbtnWp,
	chkbtnMu,
	chkbtnClH *gtk.CheckButton

	tag *gtk.TextTag

	currStyle,
	currLang,
	filename string
)

func main() {

	gtk.Init(nil)

	// Initialize SourceView on enter. WORK only with gtkSourceView >=v4.0
	source.SourceInit() // comment this line on error

	win := setupWindow("Gotk3 SourceView simple usage Example")
	win.ShowAll()

	// Apply css for desired fonts to sourceView only
	CssWidget(`
	* {font: 12px "Liberation Mono", sans-serif;
	}
	`, &sv.Widget)

	// Init red tag
	props := make(map[string]interface{})
	props["background"] = "red"
	tag = buff.CreateTag("red", props)

	// init lang and style
	currStyle = "hfmrow"
	currLang = "go-hfmrow"
	getSourceComp()

	filename = "example.go"
	textDisplay(filename)

	gtk.Main()
}

// Create and initialize the window and Gtk objects
func setupWindow(title string) *gtk.Window {
	var err error

	win, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	// Cleaning SourceView on exit. WORK only with gtkSourceView >=v4.0
	win.Connect("delete-event", source.SourceFinalize) // comment this line on error

	// make window
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetPosition(gtk.WIN_POS_CENTER)
	width, height := 800, 600
	win.SetDefaultSize(width, height)

	// Box that contain others objects
	box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Labels
	if lbl, err = gtk.LabelNew(""); err == nil {
		lbl.SetMarkup(`<a href="https://github.com/hfmrow?tab=repositories">hfmrow' repository</a>`)
		if lblLang, err = gtk.LabelNew(" Language"); err == nil {
			if lblStyle, err = gtk.LabelNew("Style"); err == nil {
				lblVersion, err = gtk.LabelNew("Current GtkSource")
			}
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	// Button to load example text
	btnDisp, err := gtk.ButtonNew()
	btnDisp.Connect("clicked", func() { textDisplay() })
	btnDisp.SetLabel("Choose file")
	if err != nil {
		log.Fatal(err)
	}
	// Button to refresh
	btnRfs, err := gtk.ButtonNew()
	btnRfs.Connect("clicked", func() { refresh() })
	btnRfs.SetLabel("Reload style sheme and language")
	if err != nil {
		log.Fatal(err)
	}

	// GtkSourceView implementation
	if sv, err = source.SourceViewNew(); err == nil {
		ma, mi, mc, _ := sv.GetVersion()

		// Display GtkSourceView version informations
		lblVersion.SetLabel(fmt.Sprintf("%v: v%v.%v.%v", lblVersion.GetLabel(), ma, mi, mc))

		// Store sourceBuffer for further usage
		buff, err = sv.GetBuffer()

		sv.SetProperty("visible", true)
		sv.SetProperty("can_focus", true)
		sv.SetProperty("left_margin", 2)
		sv.SetProperty("right_margin", 2)
		sv.SetProperty("show_line_numbers", true)
		sv.SetProperty("highlight_current_line", true)

	}
	if err != nil {
		log.Fatal(err)
	}

	// ScrolledWindow View
	if scrl, err = gtk.ScrolledWindowNew(nil, nil); err == nil {

		scrl.SetProperty("visible", true)
		scrl.SetProperty("can-focus", true)
		scrl.SetProperty("hexpand", true)
		scrl.SetProperty("vexpand", true)
		scrl.SetProperty("shadow_type", gtk.SHADOW_IN)
	}
	if err != nil {
		log.Fatal(err)
	}

	scrl.Add(sv)

	// ComboboxText for style and language choice
	if cbtxStyle, err = gtk.ComboBoxTextNew(); err == nil {
		if cbtxlang, err = gtk.ComboBoxTextNew(); err == nil {

			cbtxStyle.Connect("changed", func(c *gtk.ComboBoxText) {

				if scheme, err = sssm.GetScheme(c.GetActiveText()); err == nil {
					buff.SetStyleSheme(scheme)
				}
			})

			cbtxlang.Connect("changed", func(c *gtk.ComboBoxText) {

				if lng, err = slm.GetLanguage(c.GetActiveText()); err == nil {
					buff.SetLanguage(lng)
				}
			})
		}

		cbtxStyle.SetProperty("id-column", 0)
		cbtxlang.SetProperty("id-column", 0)
	}
	if err != nil {
		log.Fatal(err)
	}

	initCheckButtons()

	// Grid for multi-component
	if grid, err = gtk.GridNew(); err == nil {
		grid.SetProperty("visible", true)
		grid.SetProperty("hexpand", true)
		grid.SetProperty("vexpand", false)
		grid.SetProperty("column-spacing", 2)
		grid.SetColumnHomogeneous(true)
		grid.SetRowHomogeneous(true)

		// Packing
		grid.Attach(lblLang, 0, 0, 1, 1)
		grid.AttachNextTo(cbtxlang, lblLang, gtk.POS_RIGHT, 1, 1)
		grid.AttachNextTo(lblStyle, lblLang, gtk.POS_BOTTOM, 1, 1)
		grid.AttachNextTo(cbtxStyle, lblStyle, gtk.POS_RIGHT, 1, 1)

		grid.AttachNextTo(chkbtnWp, lblStyle, gtk.POS_BOTTOM, 1, 1)
		grid.AttachNextTo(chkbtnHl, chkbtnWp, gtk.POS_RIGHT, 1, 1)
		grid.AttachNextTo(chkbtnLn, chkbtnWp, gtk.POS_BOTTOM, 1, 1)
		grid.AttachNextTo(chkbtnMb, chkbtnLn, gtk.POS_RIGHT, 1, 1)
		grid.AttachNextTo(chkbtnMu, chkbtnHl, gtk.POS_RIGHT, 1, 1)
		grid.AttachNextTo(chkbtnClH, chkbtnMb, gtk.POS_RIGHT, 1, 1)

		grid.AttachNextTo(btnDisp, chkbtnLn, gtk.POS_BOTTOM, 1, 1)
		grid.AttachNextTo(btnRfs, btnDisp, gtk.POS_RIGHT, 1, 1)
		grid.AttachNextTo(lblVersion, btnRfs, gtk.POS_RIGHT, 1, 1)

		box.Add(lbl)
		box.Add(scrl)
		box.Add(grid)
		win.Add(box)
	}
	if err != nil {
		log.Fatal(err)
	}
	return win
}

// refresh:
func refresh() {

	currStyle = cbtxStyle.GetActiveText()
	currLang = cbtxlang.GetActiveText()

	cbtxlang.RemoveAll()
	cbtxStyle.RemoveAll()

	getSourceComp()
}

// getSourceComp: Get GtkSourceView components (Style scheme and Language type)
func getSourceComp() {
	var err error

	// Get languages
	if slm, err = source.SourceLanguageManagerNew(); err == nil {

		// Adding my own language mapper for Go
		sp := slm.GetSearchPath()
		sp = append(sp, "langAndstyle")
		slm.SetSearchPath(sp)

		// Get style scheme
		if sssm, err = source.SourceStyleSchemeManagerNew(); err == nil {

			// Adding my own Style scheme designed for Go
			sssm.AppendSearchPath("langAndstyle")

			// Get languages and Fill cbtx lang
			languageIds := slm.GetLanguageIds()
			for _, lngId := range languageIds {

				if lngId != "def" { // Skip definition file that cause Warning !
					cbtxlang.AppendText(lngId)
				}

			}
			//  Get styles and Fill cbtx style
			shemeIds := sssm.GetShemeIds()
			for _, sytleId := range shemeIds {

				cbtxStyle.AppendText(sytleId)
			}

			cbtxlang.SetActiveID(currLang)
			cbtxStyle.SetActiveID(currStyle)
		}

	}

	if err != nil {
		log.Fatal(err)
	}
}

// textDisplay: populate GtkSourceView
func textDisplay(file ...string) {
	var err error
	var fName string
	var ok bool

	if len(file) > 0 {
		fName = file[0]
	}

	if _, err = os.Stat(fName); os.IsNotExist(err) {

		filename, ok, err = FileChooser(fName)
	} else {
		ok = true
	}

	if ok {
		var data []byte
		if data, err = ioutil.ReadFile(filename); err == nil {

			if chkbtnMu.GetActive() {
				s, e := buff.GetStartIter(), buff.GetEndIter()
				buff.Delete(s, e)
				buff.InsertMarkup(s, string(data))
			} else {
				buff.SetText(string(data))
				sv.SetLeftMargin(2)
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

// initCheckButtons: Set Checkbuttons
func initCheckButtons() {
	var err error

	// Checkbox wrap
	if chkbtnWp, err = gtk.CheckButtonNew(); err != nil {
		log.Fatal(err)
	}
	chkbtnWp.SetLabel("Wrap mode")
	chkbtnWp.Connect("toggled", func(chk *gtk.CheckButton) {
		switch chk.GetActive() {
		case true:
			sv.SetWrapMode(gtk.WRAP_WORD_CHAR)
		case false:
			sv.SetWrapMode(gtk.WRAP_NONE)
		}
	})

	// Checkbox highlight
	if chkbtnHl, err = gtk.CheckButtonNew(); err != nil {
		log.Fatal(err)
	}
	chkbtnHl.SetLabel("Highlight")
	chkbtnHl.Connect("toggled", func(chk *gtk.CheckButton) {
		buff.SetHighlightSyntax(chk.GetActive())
	})

	// Checkbox line numbers
	if chkbtnLn, err = gtk.CheckButtonNew(); err != nil {
		log.Fatal(err)
	}
	chkbtnLn.SetLabel("Show numbers")
	chkbtnLn.Connect("toggled", func(chk *gtk.CheckButton) {
		sv.SetShowLineNumbers(chk.GetActive())
	})

	// Checkbox currline highlight
	if chkbtnClH, err = gtk.CheckButtonNew(); err != nil {
		log.Fatal(err)
	}
	chkbtnClH.SetLabel("Highlight Current Line")
	chkbtnClH.Connect("toggled", func(chk *gtk.CheckButton) {
		sv.SetHighlightCurrentLine(chk.GetActive())
	})

	// Checkbox matching brackets
	if chkbtnMb, err = gtk.CheckButtonNew(); err != nil {
		log.Fatal(err)
	}
	chkbtnMb.SetLabel("Matching brackets")
	chkbtnMb.Connect("toggled", func(chk *gtk.CheckButton) {
		buff.SetHighlightMatchingBrackets(chk.GetActive())
	})

	// Checkbox load as markup
	if chkbtnMu, err = gtk.CheckButtonNew(); err != nil {
		log.Fatal(err)
	}
	chkbtnMu.SetLabel("Load Markup file")
	chkbtnMu.SetTooltipText("Next loaded file will contain markup elements")

	// Define the state of the checkboxes according to their assignment
	chkbtnHl.SetActive(buff.GetHighlightSyntax())
	chkbtnMb.SetActive(buff.GetHighlightMatchingBrackets())
	chkbtnLn.SetActive(sv.GetShowLineNumbers())
	chkbtnClH.SetActive(sv.GetHighlightCurrentLine())

	switch chkbtnWp.GetActive() {
	case true:
		sv.SetWrapMode(gtk.WRAP_WORD_CHAR)
	case false:
		sv.SetWrapMode(gtk.WRAP_NONE)
	}
}

// CssWidget: read from data and apply css to widget.
func CssWidget(css string, wdgt *gtk.Widget) {

	var err error
	var cssProv *gtk.CssProvider

	if cssProv, err = gtk.CssProviderNew(); err == nil {

		if err = cssProv.LoadFromData(css); err == nil {
			var styleContext *gtk.StyleContext

			if styleContext, err = wdgt.GetStyleContext(); err == nil {
				styleContext.AddProvider(cssProv, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
			}
		}
	}
	if err != nil {
		log.Printf("CssWdgScnLoad: %s\n", err.Error())
	}
}

// FileChooser: Display a file chooser dialog.
func FileChooser(filename string) (outFilename string, result bool, err error) {
	var fileChooser *gtk.FileChooserDialog

	if fileChooser, err = gtk.FileChooserDialogNewWith2Buttons("Choose file", win, gtk.FILE_CHOOSER_ACTION_SAVE,
		"Cancel", gtk.RESPONSE_CANCEL, "Ok", gtk.RESPONSE_ACCEPT); err != nil {
		return
	}

	fileChooser.SetCurrentName(filepath.Base(filename))
	fileChooser.SetDoOverwriteConfirmation(false)
	fileChooser.SetModal(true)
	fileChooser.SetSkipPagerHint(true)
	fileChooser.SetSkipTaskbarHint(true)
	fileChooser.SetKeepAbove(false)

	switch int(fileChooser.Run()) {
	case -3:
		result = true
		outFilename = fileChooser.GetFilename()
	}

	fileChooser.Destroy()
	return
}
