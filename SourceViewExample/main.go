// main.go

/*
	Source file auto-generated on Mon, 12 Oct 2020 04:10:23 using Gotk3ObjHandler v1.6.5 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE
*/

/*******************************************************************************************************/
/* NOTICE: Unless otherwise stated in the various sections, most of this file will not be modified    */
/* during an update. This Gotk3ObjHandler's project format is intended to be an abbreviated version  */
/* for demonstration purposes only and not for a full project.                                      */
/***************************************************************************************************/

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hfmrow/gotk3_gtksource/source"

	gitvsv "github.com/hfmrow/gtk3Import/textView/sourceView"
)

var (
	// Structure that hold some GtkSourceView methods. (see below for explanations)
	svs                 *gitvsv.SourceViewStruct
	SourceViewStructNew = gitvsv.SourceViewStructNew
	// GtkSourceView objects
	currentLanguage          *source.SourceLanguage
	currentStyleScheme       *source.SourceStyleScheme
	styleSchemeChooserWidget *source.SourceStyleSchemeChooserWidget
	// Gtk3 Objects
	markFound *gtk.TextTag
	Dialog    *gtk.Dialog
)

/* Start code */
func mainApplication() {

	var err error

	// Set iconApp
	mainObjects.Window.SetIconName("computer-symbolic")

	// Set Link label
	mainObjects.LinkButton.SetUri("https://" + Repository)
	mainObjects.LinkButton.SetLabel(Repository)

	// SourceViewStruct is a structure that integrates most of the common features
	// of GtkSource but is not exhaustive, I created this library to simplify the
	// interactions and facilitate the implementation of this component and thus
	// clarify the steps necessary for its uses.
	// There is no obligation to use a structure, but in this case, for this example,
	// i prefere using it for my convenience. All Operations are described in library
	// to help you understand how to do.
	if svs, err = SourceViewStructNew(mainObjects.View, mainObjects.Map, mainObjects.Window); err == nil {

		// Loading default file to be displayed
		mainObjects.FileChooserButton.SetFilename("main.go")
		FileChooserButtonFileSet(mainObjects.FileChooserButton)

		// Make a tag to indicate found element (when HighlightFound not checked)
		tag := make(map[string]interface{})
		tag["background"] = "#ABF6FF"
		markFound = svs.Buffer.CreateTag("markFound", tag)

		// Language & style, add a personal version for Golang (directory content)
		svs.UserStylePath = "assets/langAndstyle"
		svs.UserLanguagePath = "assets/langAndstyle"

		// Setting Language and style scheme
		if currentLanguage = svs.SetLanguage(mainOptions.DefaulLanguage); currentLanguage != nil {
			if currentStyleScheme = svs.SetStyleScheme(mainOptions.DefaultStyle); currentStyleScheme != nil {

				// In case where there is no user defined style or language available, we set it
				// to main options because it's used below to set comboboxes current values.
				mainOptions.DefaulLanguage = currentLanguage.GetId()
				mainOptions.DefaultStyle = currentStyleScheme.GetId()

				// Set options to GtkObjects
				mainOptions.UpdateObjects()

				// Fill comboboxes with languages and styles
				for _, id := range svs.LanguageIds {
					mainObjects.ComboBoxTextLanguage.AppendText(id)
				}
				for _, id := range svs.StyleShemeIds {
					mainObjects.ComboBoxTextStyle.AppendText(id)
				}

				// Just indicate id must be set as first model column.
				mainObjects.ComboBoxTextLanguage.SetIDColumn(0)
				mainObjects.ComboBoxTextStyle.SetIDColumn(0)

				// Set ComboBox current values display.
				mainObjects.ComboBoxTextLanguage.SetActiveID(mainOptions.DefaulLanguage)
				mainObjects.ComboBoxTextStyle.SetActiveID(mainOptions.DefaultStyle)

				// This ensure the options displayed by (checkboxes, comboboxes)
				// are defined on the GtkSourceView structure.
				setOptions()
			}
		}
	}
	if err != nil {

		dispStatusbar(err.Error())
	}
}

// Sections below is not relevant part of GtkSourceView,
// but handle interface building part.
func main() {

	/* Create temp directory */
	doTempDir = false

	/* Init & read options file */
	mainOptions = new(MainOpt)
	mainOptions.Init()
	mainOptions.Read()

	/* Init gtk display */
	mainStartGtk(fmt.Sprintf("%s %s  %s %s %s",
		Name,
		Vers,
		"©"+YearCreat,
		Creat,
		LicenseAbrv),
		mainOptions.MainWinWidth,
		mainOptions.MainWinHeight, true)
}

/********************************************************/
/* This section preserve user modifications on update. */
/* Signals & Property implementations:                */
/* initialise signals used by gtk objects ...        */
/****************************************************/
// signalsPropHandler: initialise signals used by gtk objects ...
func signalsPropHandler() {
	mainObjects.ButtonBackward.Connect("clicked", ButtonBackwardClicked)
	mainObjects.ButtonBackwardAsync.Connect("clicked", ButtonBackwardAsyncClicked)
	mainObjects.ButtonForward.Connect("clicked", ButtonForwardClicked)
	mainObjects.ButtonForwardAsync.Connect("clicked", ButtonForwardAsyncClicked)
	mainObjects.ButtonReplace.Connect("clicked", ButtonReplaceClicked)
	mainObjects.ButtonReplaceAll.Connect("clicked", ButtonReplaceAllClicked)
	mainObjects.ButtonStyle.Connect("clicked", ButtonStyleClicked)
	mainObjects.CheckButtonBrackets.Connect("toggled", CheckButtonBracketsToggled)
	mainObjects.CheckButtonCaseSensitive.Connect("toggled", CheckButtonCaseSensitiveToggled)
	mainObjects.CheckButtonHighlight.Connect("toggled", CheckButtonHighlightToggled)
	mainObjects.CheckButtonHighlightFound.Connect("toggled", CheckButtonHighlightFoundToggled)
	mainObjects.CheckButtonHighlightLine.Connect("toggled", CheckButtonHighlightLineToggled)
	mainObjects.CheckButtonNumbers.Connect("toggled", CheckButtonNumbersToggled)
	mainObjects.CheckButtonRegexp.Connect("toggled", CheckButtonRegexpToggled)
	mainObjects.CheckButtonShowMap.Connect("toggled", CheckButtonShowMapToggled)
	mainObjects.CheckButtonShowMarks.Connect("toggled", CheckButtonShowMarksToggled)
	mainObjects.CheckButtonWordBoundaries.Connect("toggled", CheckButtonWordBoundariesToggled)
	mainObjects.CheckButtonWrap.Connect("toggled", CheckButtonWrapToggled)
	mainObjects.CheckButtonWrapAround.Connect("toggled", CheckButtonWrapAroundToggled)
	mainObjects.ComboBoxTextLanguage.Connect("changed", ComboBoxTextLanguageChanged)
	mainObjects.ComboBoxTextStyle.Connect("changed", ComboBoxTextStyleChanged)
	mainObjects.EntryReplace.Connect("changed", EntryReplaceChanged)
	mainObjects.EntrySearch.Connect("changed", EntrySearchChanged)
	mainObjects.FileChooserButton.Connect("file-set", FileChooserButtonFileSet)
	mainObjects.LabelCount.Connect("notify", blankNotify)
	mainObjects.LabelCurr.Connect("notify", blankNotify)
	mainObjects.LinkButton.Connect("notify", blankNotify)
	mainObjects.Map.Connect("notify", blankNotify)
	mainObjects.Paned.Connect("notify", blankNotify)
	mainObjects.Statusbar.Connect("notify", blankNotify)
	mainObjects.View.Connect("notify", blankNotify)
	mainObjects.Window.Connect("notify", blankNotify)
}

/******************************/
/* Main options declarations */
/****************************/
type MainOpt struct {
	MainWinWidth,
	MainWinHeight,
	MainWinPosX,
	MainWinPosY int

	Brackets,
	Highlight,
	HighlightLine,
	Numbers,
	ShowMap,
	ShowMarks,
	Wrap,

	WrapAround,
	WordBoundaries,
	Regexp,
	CaseSensitive,
	HighlightFound bool

	PanedWidth int

	DefaulLanguage,
	DefaultStyle string
}

// Main options initialisation
func (opt *MainOpt) Init() {

	// Set default options
	opt.MainWinWidth = 800
	opt.MainWinHeight = 600

	opt.DefaulLanguage = "go-hfmrow"
	opt.DefaultStyle = "hfmrow"

	opt.PanedWidth = 120
}

// Variables -> Objects.
func (opt *MainOpt) UpdateObjects() {
	mainObjects.Window.Resize(opt.MainWinWidth, opt.MainWinHeight)
	mainObjects.Window.Move(opt.MainWinPosX, opt.MainWinPosY)

	mainObjects.CheckButtonBrackets.SetActive(opt.Brackets)
	mainObjects.CheckButtonHighlight.SetActive(opt.Highlight)
	mainObjects.CheckButtonHighlightLine.SetActive(opt.HighlightLine)
	mainObjects.CheckButtonNumbers.SetActive(opt.Numbers)
	mainObjects.CheckButtonShowMap.SetActive(opt.ShowMap)
	mainObjects.CheckButtonShowMarks.SetActive(opt.ShowMarks)
	mainObjects.CheckButtonWrap.SetActive(opt.Wrap)

	mainObjects.CheckButtonWrapAround.SetActive(opt.WrapAround)
	mainObjects.CheckButtonWordBoundaries.SetActive(opt.WordBoundaries)
	mainObjects.CheckButtonRegexp.SetActive(opt.Regexp)
	mainObjects.CheckButtonCaseSensitive.SetActive(opt.CaseSensitive)
	mainObjects.CheckButtonHighlightFound.SetActive(opt.HighlightFound)

	mainObjects.Paned.SetPosition(opt.MainWinWidth - opt.PanedWidth)
}

// Objects -> Variables.
func (opt *MainOpt) UpdateOptions() {
	opt.MainWinWidth, opt.MainWinHeight = mainObjects.Window.GetSize()
	opt.MainWinPosX, opt.MainWinPosY = mainObjects.Window.GetPosition()

	opt.Brackets = mainObjects.CheckButtonBrackets.GetActive()
	opt.Highlight = mainObjects.CheckButtonHighlight.GetActive()
	opt.HighlightLine = mainObjects.CheckButtonHighlightLine.GetActive()
	opt.Numbers = mainObjects.CheckButtonNumbers.GetActive()
	opt.ShowMap = mainObjects.CheckButtonShowMap.GetActive()
	opt.ShowMarks = mainObjects.CheckButtonShowMarks.GetActive()
	opt.Wrap = mainObjects.CheckButtonWrap.GetActive()

	opt.WrapAround = mainObjects.CheckButtonWrapAround.GetActive()
	opt.WordBoundaries = mainObjects.CheckButtonWordBoundaries.GetActive()
	opt.Regexp = mainObjects.CheckButtonRegexp.GetActive()
	opt.CaseSensitive = mainObjects.CheckButtonCaseSensitive.GetActive()
	opt.HighlightFound = mainObjects.CheckButtonHighlightFound.GetActive()

	opt.PanedWidth = opt.MainWinWidth - mainObjects.Paned.GetPosition()
	if opt.PanedWidth < 10 {
		opt.PanedWidth = 120
	}
}

var (
	// Project declarations
	Name         = "GtkSourceView example"
	Vers         = "v1.0"
	Descr        = "GtkSourceView example"
	Creat        = "H.F.M"
	YearCreat    = "2020"
	LicenseShort = "This program comes with absolutely no warranty.\nSee the The MIT License (MIT) for details:\nhttps://opensource.org/licenses/mit-license.php"
	LicenseAbrv  = "License (MIT)"
	Repository   = "github.com/hfmrow/gotk3_gtksource"

	// Common declarations
	absoluteRealPath, optFilename = getAbsRealPath()
	mainOptions                   *MainOpt
	tempDir                       string
	doTempDir                     bool
	mainObjects                   *MainControlsObj
)

/********************************************************/
/* This section preserve user modifications on update. */
/* Main structure Declaration: You may add your own   */
/* declarations (gotk3 objects only) here.           */
/****************************************************/
type MainControlsObj struct {
	ButtonBackward            *gtk.Button
	ButtonBackwardAsync       *gtk.Button
	ButtonForward             *gtk.Button
	ButtonForwardAsync        *gtk.Button
	ButtonReplace             *gtk.Button
	ButtonReplaceAll          *gtk.Button
	ButtonStyle               *gtk.Button
	CheckButtonBrackets       *gtk.CheckButton
	CheckButtonCaseSensitive  *gtk.CheckButton
	CheckButtonHighlight      *gtk.CheckButton
	CheckButtonHighlightFound *gtk.CheckButton
	CheckButtonHighlightLine  *gtk.CheckButton
	CheckButtonNumbers        *gtk.CheckButton
	CheckButtonRegexp         *gtk.CheckButton
	CheckButtonShowMap        *gtk.CheckButton
	CheckButtonShowMarks      *gtk.CheckButton
	CheckButtonWordBoundaries *gtk.CheckButton
	CheckButtonWrap           *gtk.CheckButton
	CheckButtonWrapAround     *gtk.CheckButton
	ComboBoxTextLanguage      *gtk.ComboBoxText
	ComboBoxTextStyle         *gtk.ComboBoxText
	EntryReplace              *gtk.Entry
	EntrySearch               *gtk.Entry
	FileChooserButton         *gtk.FileChooserButton
	LabelCount                *gtk.Label
	LabelCurr                 *gtk.Label
	LinkButton                *gtk.LinkButton
	mainUiBuilder             *gtk.Builder
	Map                       *source.SourceMap
	Paned                     *gtk.Paned
	Statusbar                 *gtk.Statusbar
	View                      *source.SourceView
	Window                    *gtk.Window
}

/******************************************************************/
/* This section preserve user modification on update.            */
/* GtkObjects initialisation: You may add your own declarations */
/* as you  wish, best way is to add them by grouping  same     */
/* objects names (below first declaration).                   */
/*************************************************************/
func gladeObjParser() {
	mainObjects.ButtonBackward = loadObject("ButtonBackward").(*gtk.Button)
	mainObjects.ButtonBackwardAsync = loadObject("ButtonBackwardAsync").(*gtk.Button)
	mainObjects.ButtonForward = loadObject("ButtonForward").(*gtk.Button)
	mainObjects.ButtonForwardAsync = loadObject("ButtonForwardAsync").(*gtk.Button)
	mainObjects.ButtonReplace = loadObject("ButtonReplace").(*gtk.Button)
	mainObjects.ButtonReplaceAll = loadObject("ButtonReplaceAll").(*gtk.Button)
	mainObjects.ButtonStyle = loadObject("ButtonStyle").(*gtk.Button)
	mainObjects.CheckButtonBrackets = loadObject("CheckButtonBrackets").(*gtk.CheckButton)
	mainObjects.CheckButtonCaseSensitive = loadObject("CheckButtonCaseSensitive").(*gtk.CheckButton)
	mainObjects.CheckButtonHighlight = loadObject("CheckButtonHighlight").(*gtk.CheckButton)
	mainObjects.CheckButtonHighlightFound = loadObject("CheckButtonHighlightFound").(*gtk.CheckButton)
	mainObjects.CheckButtonHighlightLine = loadObject("CheckButtonHighlightLine").(*gtk.CheckButton)
	mainObjects.CheckButtonNumbers = loadObject("CheckButtonNumbers").(*gtk.CheckButton)
	mainObjects.CheckButtonRegexp = loadObject("CheckButtonRegexp").(*gtk.CheckButton)
	mainObjects.CheckButtonShowMap = loadObject("CheckButtonShowMap").(*gtk.CheckButton)
	mainObjects.CheckButtonShowMarks = loadObject("CheckButtonShowMarks").(*gtk.CheckButton)
	mainObjects.CheckButtonWordBoundaries = loadObject("CheckButtonWordBoundaries").(*gtk.CheckButton)
	mainObjects.CheckButtonWrap = loadObject("CheckButtonWrap").(*gtk.CheckButton)
	mainObjects.CheckButtonWrapAround = loadObject("CheckButtonWrapAround").(*gtk.CheckButton)
	mainObjects.ComboBoxTextLanguage = loadObject("ComboBoxTextLanguage").(*gtk.ComboBoxText)
	mainObjects.ComboBoxTextStyle = loadObject("ComboBoxTextStyle").(*gtk.ComboBoxText)
	mainObjects.EntryReplace = loadObject("EntryReplace").(*gtk.Entry)
	mainObjects.EntrySearch = loadObject("EntrySearch").(*gtk.Entry)
	mainObjects.FileChooserButton = loadObject("FileChooserButton").(*gtk.FileChooserButton)
	mainObjects.LabelCount = loadObject("LabelCount").(*gtk.Label)
	mainObjects.LabelCurr = loadObject("LabelCurr").(*gtk.Label)
	mainObjects.LinkButton = loadObject("LinkButton").(*gtk.LinkButton)
	mainObjects.Map = loadObject("Map").(*source.SourceMap)
	mainObjects.Paned = loadObject("Paned").(*gtk.Paned)
	mainObjects.Statusbar = loadObject("Statusbar").(*gtk.Statusbar)
	mainObjects.View = loadObject("View").(*source.SourceView)
	mainObjects.Window = loadObject("Window").(*gtk.Window)
}

/*******************************/
/* Gtk3 Window Initialisation */
/*****************************/
func mainStartGtk(winTitle string, width, height int, center bool) {
	mainObjects = new(MainControlsObj)
	gtk.Init(nil)
	if err := newBuilder(mainGlade); err == nil {

		/* Init tempDir and Remove it on quit if requested. */
		if doTempDir {
			tempDir = tempMake(Name)
			defer os.RemoveAll(tempDir)
		}
		/* Parse Gtk objects */
		gladeObjParser()

		/* Update gtk conctrols with stored values into mainOptions */
		mainOptions.UpdateObjects()

		/* Start main ... */
		mainApplication()

		/* Objects Signals initialisations */
		signalsPropHandler()

		/* Set Window Properties */
		if center {
			mainObjects.Window.SetPosition(gtk.WIN_POS_CENTER)
		}
		mainObjects.Window.SetTitle(winTitle)
		mainObjects.Window.SetDefaultSize(width, height)
		mainObjects.Window.Connect("delete-event", windowDestroy)
		mainObjects.Window.ShowAll()

		/*	Start Gui loop */
		gtk.Main()
	} else {
		log.Fatal("Builder initialisation error.", err.Error())
	}
}

/*******************************************************/
/* Functions declarations, used to initialize objects */
/*****************************************************/
// newBuilder: initialise builder with glade xml string
func newBuilder(varPath interface{}) (err error) {
	var Gtk3Interface []byte
	if Gtk3Interface, err = getBytesFromVarAsset(varPath); err == nil {
		if mainObjects.mainUiBuilder, err = gtk.BuilderNew(); err == nil {
			err = mainObjects.mainUiBuilder.AddFromString(string(Gtk3Interface))
		}
	}
	return err
}

// loadObject: Load GtkObject to be transtyped ...
func loadObject(name string) (newObj glib.IObject) {
	var err error
	if newObj, err = mainObjects.mainUiBuilder.GetObject(name); err != nil {
		fmt.Printf("Unable to load %s object, maybe it was deleted from the Glade file ... : %s\n%s\n",
			name, err.Error(),
			fmt.Sprint("An update with GOH may avoid this issue."))
		os.Exit(1)
	}
	return newObj
}

// WindowDestroy: is the triggered handler when closing/destroying the gui window.
func windowDestroy() {
	/* Doing something before quit. Put something here ... */
	if err := mainOptions.Write(); err != nil { /* Update mainOptions with values of gtk conctrols and write to file */
		fmt.Printf("%s\n%v\n", "Writing options error.", err)
	}
	gtk.MainQuit()
}

// getBytesFromVarAsset: Get []byte representation from file or asset, depending on type
func getBytesFromVarAsset(varPath interface{}) (outBytes []byte, err error) {
	switch reflect.TypeOf(varPath).String() {
	case "string":
		return ioutil.ReadFile(varPath.(string))
	case "[]uint8":
		return varPath.([]byte), err
	}
	return
}

// HexToBytes: Convert Gzip Hex to []byte used for embedded binary in source code
func HexToBytes(varPath string, gzipData []byte) (outByte []byte) {
	r, err := gzip.NewReader(bytes.NewBuffer(gzipData))
	if err == nil {
		var bBuffer bytes.Buffer
		if _, err = io.Copy(&bBuffer, r); err == nil {
			if err = r.Close(); err == nil {
				return bBuffer.Bytes()
			}
		}
	} else {
		fmt.Printf("An error occurred while reading: %s\n%s\n", varPath, err.Error())
	}
	return
}

/*******************************/
/* Simplified files Functions */
/*****************************/
// Read Options from file (Options structure part)
func (opt *MainOpt) Read() (err error) {
	var textFileBytes []byte
	if textFileBytes, err = ioutil.ReadFile(optFilename); err == nil {
		err = json.Unmarshal(textFileBytes, &opt)
	}
	if err != nil {
		fmt.Printf("Error while reading options file: %s\n", err.Error())
	}
	return
}

// Write Options to file (Options structure part)
func (opt *MainOpt) Write() (err error) {
	var jsonData []byte
	var out bytes.Buffer
	opt.UpdateOptions()
	if jsonData, err = json.Marshal(&opt); err == nil {
		if err = json.Indent(&out, jsonData, "", "\t"); err == nil {
			err = ioutil.WriteFile(optFilename, out.Bytes(), os.ModePerm)
		}
	}
	return err
}

// Make temporary directory
func tempMake(prefix string) (dir string) {
	var err error
	if dir, err = ioutil.TempDir("", prefix+"-"); err != nil {
		log.Fatal(err)
	}
	return dir + string(os.PathSeparator)
}

// Retrieve current realpath and options filename
func getAbsRealPath() (absoluteRealPath, optFilename string) {
	var base string

	var setExt = func(filename, ext string) (out string) {
		return filename[:len(filename)-len(path.Ext(filename))] + ext
	}

	if absoluteBaseName, err := os.Executable(); err == nil {
		absoluteRealPath, base = filepath.Split(absoluteBaseName)
		optFilename = setExt(filepath.Join(absoluteRealPath, base), ".opt")
	} else {
		log.Fatal(err)
	}
	return
}

// Used as a fake function in signals section
func blankNotify() {}

// Embedded Glade file
var mainGlade = HexToBytes("mainGlade", []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x5d\x5d\x73\xa3\xb8\x12\x7d\x9f\x5f\xa1\xcb\xeb\xad\xf8\x33\x53\x75\xeb\x96\xcd\xd4\x64\x6a\x33\xfb\x90\xd9\xda\x9a\xcc\x6e\x1e\x5d\xb2\x68\x83\xd6\xb2\xc4\x0a\x61\x3b\xfb\xeb\xb7\x00\x3b\xb1\xcb\x7c\x09\x70\x06\x88\xde\x88\x43\xb7\x41\x7d\x4e\xf7\x69\x21\xac\xd9\xa7\xfd\x86\xa1\x2d\xc8\x80\x0a\x3e\xb7\xc6\x83\x91\x85\x80\x13\xe1\x50\xee\xce\xad\x3f\x7e\xdc\xdf\xfc\xcf\xfa\x64\x7f\x98\xfd\xe7\xe6\x06\x7d\x05\x0e\x12\x2b\x70\xd0\x8e\x2a\x0f\xb9\x0c\x3b\x80\xa6\x83\xc9\x64\x30\x41\x37\x37\xf6\x87\x19\xe5\x0a\xe4\x0a\x13\xb0\x3f\x20\x34\x93\xf0\x77\x48\x25\x04\x88\xd1\xe5\xdc\x72\xd5\xfa\xbf\xd6\xeb\x17\x4d\x07\x93\x91\x35\x4c\x3f\x2f\x10\xa1\x24\xb0\xa5\xb0\x3b\x31\xb8\x1d\x1c\xcf\x17\xcb\xbf\x80\x28\x44\x18\x0e\x82\xb9\xf5\x55\xad\x9f\x28\x77\xc4\xce\x42\xd4\x99\x5b\x87\xe3\xe8\x44\x84\x66\xbe\x14\x3e\x48\xf5\x8c\x38\xde\xc0\xdc\x22\x98\x2f\x56\x82\x84\x81\x65\xdf\x63\x16\xc0\x6c\x78\x3c\xe1\x70\x3e\xf1\x28\x73\x90\x7a\xf6\x61\x6e\x29\xaa\x18\x2c\xb1\x3c\xf8\x8a\xbc\x31\x4c\xc0\x13\xcc\x01\x39\x3c\x18\x0c\x63\x8b\x53\xeb\x97\xb3\x2f\x2e\xf3\x4e\xec\x5f\x7c\x5d\x5e\xdb\x96\x06\x74\xc9\xc0\xb2\x7f\xc8\xf0\xe2\xc2\xaa\xdc\x4c\x9a\xcd\x06\x4b\x97\xf2\x05\x83\x95\xb2\xec\x89\x86\x85\xa4\xae\xa7\x69\xa2\x84\xaf\x67\xb0\x14\x4a\x89\x4d\x49\x1b\x0f\xf6\x3e\xe6\x4e\xf9\xe1\xda\xea\x1a\x08\x49\x81\x2b\xac\xa8\xe0\x96\xbd\x05\xa9\x28\xc1\x2c\xd5\xf0\x2c\xee\xe9\xb1\x7f\xa0\x7c\x7d\x17\x2a\x25\x78\x02\xd3\x93\xbf\x4f\x0d\x2f\xaf\x82\xe1\x25\x30\x0b\x29\x89\x79\xc0\xb0\xc2\x4b\x06\x73\xeb\x19\x02\xcb\x5e\xc6\xe6\x69\x17\x54\x09\x5e\x05\x10\xd3\x31\x93\x40\x80\x6e\x21\x58\x38\xb0\xc2\x21\x53\x7a\xd6\x5b\xcc\xa8\xcb\x2d\x3b\x50\x58\xaa\xf2\x5f\xc9\x28\xac\x2c\x9b\x0b\x5e\xfa\x8b\x42\x49\x2d\xdb\x53\xca\x0f\xfe\x3f\x1c\xba\x54\x79\xe1\x72\x40\xc4\x66\xe8\xad\x36\x52\xec\xd2\xbd\xcc\x86\x49\x68\xcf\x3e\xf3\x31\x59\x53\xee\xe6\x7f\xdb\x11\x7e\x99\x5c\x4d\x33\x5a\x51\xc6\xf4\x46\xcf\x17\x01\x4d\x20\x3b\xca\xba\x83\x8b\xcb\x3d\xcb\x62\x65\x11\xfd\x3b\xe6\xe0\x24\x60\x4e\x0e\xdb\x83\xbf\xc2\xdc\x50\x29\x3f\xe4\x8f\xf6\xf4\xe3\x48\xd7\x68\x11\x80\x26\x35\x76\xd4\x81\x85\x87\xb9\x53\x38\x90\x17\x01\x4c\x0f\xe2\x23\x91\x82\x31\x70\xce\xaa\x66\xed\x30\xd6\x08\x65\xe5\x70\x56\x0e\x69\x9a\x61\xe0\x61\x47\xec\x16\x91\x0e\xb0\x6c\xca\x73\x4d\x53\x47\x3a\x63\xb4\x63\x5d\xf3\x67\xac\x6b\x22\xde\xc4\x47\x69\xa6\x35\xc6\xbd\xe6\xd8\xa7\x96\x20\x58\xa9\x45\x52\xa3\x33\x8a\x73\x6e\x66\x8e\x94\x43\x59\xf3\xb4\xf4\x9a\x91\xa3\xf2\x0d\xd2\x73\x72\x7a\xe5\x08\xe8\x3f\x50\x90\x97\xd3\x41\x22\x29\x5f\x17\x0d\x69\x4a\xba\xcd\xbc\x1d\xc3\xda\xf6\xb2\xf6\x1b\xf6\x13\xd2\x46\x07\x6f\xcb\xd9\x42\x64\xa6\xca\x0d\x87\xc6\x62\xb1\xa2\xf9\x1b\x72\x3e\xa5\x1f\x10\x5c\x04\x3e\x26\x15\x87\x2e\x6e\x1c\xed\x28\xb9\xb6\x36\xd9\xe8\x43\xbb\xd9\x5c\x53\x5f\xc6\xea\xc8\x96\x3a\x2a\x76\x7c\x5d\x15\xfb\x08\x3e\x96\x58\x09\x79\x65\xf9\xaa\xa5\xfa\x8f\x1d\x10\x70\xa7\xac\x49\x61\xcf\x5d\xb5\xef\xee\x54\xcb\x33\xb9\x2e\x58\xbe\x4a\xea\x74\x0d\x27\x25\x25\x45\xca\xad\x5d\x4b\x48\x68\x0b\xad\x12\xb7\x99\xab\x23\xb4\xbf\x90\x08\x16\x6e\xf8\xc2\x13\x1b\xe1\x02\x07\x51\x4a\xfe\x68\x28\x8a\x2f\x62\xb3\x14\x77\x62\xff\x03\xf6\x2a\x11\x15\xa7\x9f\x3c\xaa\xe7\x68\x84\xdb\x26\x31\xb2\xab\x65\x5e\x01\xcc\xd4\x16\x58\x29\x4c\xbc\xcc\xfc\x9e\x6d\xad\x84\xff\x62\x3c\x2a\xba\xe4\xec\xc2\x3c\xcc\x88\x97\x4e\x1c\x4f\x27\xf4\x92\x63\x9d\xd8\x65\xcf\xec\xc5\x5e\xf4\x55\xcf\xcf\xeb\x10\xf5\x26\xfc\xea\xcc\xe1\x5d\x15\x8b\xa3\x0e\x63\xf1\x21\x46\x53\x07\x1a\x93\x6c\xd8\x3f\x60\xee\x86\xd8\xfd\x39\xb1\x9f\x74\x38\xf6\xf9\xf5\xe4\x38\xac\xef\xa6\xa4\x4c\x3b\x1c\xca\x7b\xca\xe0\x8b\x27\x44\x00\xf2\xb4\xba\x5c\x7e\xdc\x01\xa6\x13\x09\x58\xc1\x62\x15\x3f\x30\xad\xea\x24\x7e\x0a\x9b\x9a\x2e\x92\xf1\x40\x2b\xca\x7e\x0e\xcc\x6e\x3b\x0c\xb3\xac\xb6\xb7\x95\x38\x2a\xa9\xff\x2b\xb7\xc5\xf5\xda\xe3\x36\x6b\x12\x6d\x71\xbd\xa3\x8e\xf2\x2c\xfb\x63\xd3\xe8\xac\x3f\x73\x56\x7e\x50\x2a\xd2\xee\x1a\x13\xf5\x3d\xec\xaa\xeb\x4f\xeb\x17\x5f\xa9\x46\x22\xcb\x18\xe1\xd6\xe5\xb0\xd2\xc3\x56\xeb\x89\x48\x6a\x57\x25\x76\x7a\x33\x19\xf9\x21\xc8\x50\xa0\x1e\x90\xb3\x76\xf8\xe4\x83\x5f\xa9\xeb\xb1\x64\xf9\x50\x86\x3f\x9d\x26\xe1\xc5\x5b\xd1\x2d\xd4\x86\x40\x03\x7d\x72\xb9\x5e\xb9\x14\xa0\x52\x41\x55\x66\xa5\x4c\x9e\x03\x47\xe2\xdd\x82\x72\x87\x92\x44\x0b\x94\xbb\xa7\xbc\x4a\x57\x5c\xed\xea\x57\xbc\xa2\x24\x3f\x29\x73\x0b\xf9\x15\x79\x98\x83\xff\x26\xb9\xf1\x5b\xb8\x59\xc6\xda\xb8\x01\x66\x3c\x7a\x62\x87\x78\xe2\xd0\x90\xa3\xa7\xe4\x18\xd7\x25\xc7\xb8\x3b\xe4\xb8\x93\x98\xac\x41\x35\xc3\x8e\x6f\x58\x11\x8f\x72\x17\x2d\x0f\x5e\x0d\x45\x4c\xfd\xa8\xd4\x9b\xb7\x89\x22\x2f\x6a\xe8\x81\x72\x68\x56\x5f\x21\x46\x39\x18\x92\x18\x92\x74\xbe\x8e\x44\xc2\x28\x7b\xd5\x56\x05\x91\xf5\x0d\xfb\x86\x18\x46\x60\x75\xa2\xfb\x28\x9a\xe9\xbd\x36\x50\x6b\xc3\x8b\x00\x57\x20\xab\x78\x38\x7b\xd9\xe9\xb6\x86\x87\xc3\xcb\x4f\x95\x5c\x68\x4d\xfc\xd4\x9e\xfc\xa9\xf5\x1e\xd3\x9b\xf2\x6c\xf2\x06\x2a\x2d\x2d\x1a\xa5\x43\xd9\xae\xea\x25\xd7\x41\x93\xf5\x4b\xae\x4d\xff\x63\x2a\x58\x06\xb3\xa6\xdd\x21\xc7\x93\x6c\x48\xd7\x45\x8e\x0c\x23\x4c\xb3\xd3\x46\x46\x5c\xbe\x76\xae\xe1\xdd\x2c\x95\x33\xcf\x0c\x1b\x7f\xdc\x97\xb7\xd8\xb0\xb5\x0d\x45\x8e\x28\x02\x2c\x89\x67\x26\x2e\x1b\x2e\xdc\xbf\x70\x25\x9f\x93\x92\x1d\x1f\x26\xc3\x6c\xb5\xb9\xb8\xb6\x5d\x9a\xb5\x2c\xc2\x97\x0b\xe0\xef\x85\xdc\x61\xe9\x34\xa2\xc9\x0e\xbe\xfa\x23\xcb\xaa\x7b\x2a\xbb\x30\xde\xcc\x1c\x37\x87\xe1\xcf\xc1\x33\x27\x4d\x02\x19\xc5\x1e\x0d\x9c\x5b\x05\xe7\x71\xef\xe1\x7c\x87\xc9\xba\xb1\x9c\x7c\x74\x66\x50\xdc\xaf\xa4\x3c\xe9\x0a\x8a\x9b\xcb\xca\x47\x8f\x26\x2d\xf7\x31\x2d\x4f\xda\xdd\x0b\x7d\x87\x78\x56\xe7\x3d\x37\x43\xb7\xbd\x6f\x86\xf4\x83\x9c\x9d\xad\x0e\xbe\x4c\x9a\x6a\x55\x9a\xba\xed\xbd\x7a\x3c\xe0\xee\x33\x63\x4d\xc2\x18\x61\xc6\x0c\x94\xfb\x05\xe5\x0e\x2d\xbb\x7f\x12\xd2\xb9\x13\x21\x77\xb0\xa4\xd0\xcc\x83\xf5\xc8\x25\x7a\xf5\x69\x9e\x25\xf6\xf4\x59\xe2\xf4\x1d\xad\x2e\x7e\x92\xd8\xff\x2c\x23\x4c\x37\xf6\x8c\x1d\x25\xfe\x0c\x3d\x0c\x3d\x3a\xbf\xae\xf8\x65\xa9\xfc\x7d\x63\x14\x79\x5d\x7d\x7f\x6f\x68\x62\x68\xd2\x0b\xb1\xf5\x05\x07\xf0\x08\x3c\xa0\x8a\x6e\x9b\xe9\x85\x23\x8f\x28\x38\xba\x34\x24\xe9\x29\x49\x26\xef\x88\x24\xdf\xc1\x85\xbd\xdf\x50\x8b\xed\xc2\xde\x90\xc2\x90\xa2\x13\x02\x2b\x67\x31\x5e\x6b\x57\x92\x55\xf9\x7d\xd0\x72\xe3\x53\x6d\xb9\x5d\x23\x03\xd5\xd4\x60\x69\x96\x72\x11\x72\x55\xce\x6d\x11\xb1\xca\x91\xab\x99\x47\xa9\xa8\x89\x97\x53\x0a\x89\x56\x48\x36\x03\xa8\x4b\x40\x85\x52\x0e\x5a\x05\xa8\xf1\xbb\x01\x54\xb2\xb3\x55\x74\x18\xd3\xba\xcb\x00\xd3\x7f\xfb\x10\xe5\xfc\xc6\x56\x3d\x2f\xda\xaf\x10\x76\x2b\x57\x8e\x3b\x06\xed\x50\x4a\x83\xec\xbe\x23\x7b\xdc\x29\x64\xf7\xe0\x5d\xdd\xea\x3d\xc7\xbb\xf8\x61\xfa\xb6\xfc\x80\xe6\xb8\x1d\xbb\xcf\xbc\xd9\x8e\x22\xd3\x2b\x6f\x3f\xa3\xb0\x0a\x83\x25\x96\x49\x81\x79\xfd\xb3\x45\xbb\x8c\x14\xef\x19\x8b\xaa\xed\x1b\x9b\x63\x16\xcf\xd9\x58\xf6\x78\xa4\x69\x07\x11\x38\xb4\xad\xae\xb1\x6b\x4e\xcd\xdf\x61\xb8\xd8\x00\xca\xc7\x84\x72\xb7\x0f\x9b\xf4\xdc\x56\xa5\xd4\xf9\x3d\x9e\xfc\xf3\xf5\x1f\xb3\xe1\xc9\x16\xd4\xff\x06\x00\x00\xff\xff\x55\x5c\xe9\x01\xdb\x7a\x00\x00"))
