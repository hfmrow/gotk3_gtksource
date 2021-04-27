// objHandler.go

/*
	Source file auto-generated on Mon, 12 Oct 2020 23:25:41 using Gotk3ObjHandler v1.6.5 ©2018-20 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE
*/

package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

/*
 * Some functions used (not necessarily related to sourceView) in this example.
 */

func dispStatusbar(inStr string) {
	ctxId := mainObjects.Statusbar.GetContextId("infos")
	mainObjects.Statusbar.Push(ctxId, inStr)
}

// scrollAndTag: This will scroll to found position and apply a highlight tag
// in case where 'CheckButtonHighlightFound' is not toggled on.
// And display found items count and current viewed.
func scrollAndTag() {

	svs.ScrollToIter(svs.IterStart)

	if !svs.HighlightFound {

		// Clear previous tag.
		svs.Buffer.RemoveTagByName("markFound", svs.Buffer.GetStartIter(), svs.Buffer.GetEndIter())

		// Apply tag to new position.
		svs.Buffer.ApplyTagByName("markFound", svs.IterStart, svs.IterEnd)
	}

	// Display count and current found item (if available)
	// in fact, occurrences count return -1 when buffer is not entirely read (see GtkSourceView docupentation)
	mainObjects.LabelCurr.SetLabel(fmt.Sprintf("%v", svs.SearchCtx.GetOccurencePosition(svs.IterStart, svs.IterEnd)))
	if count := svs.SearchCtx.GetOccurencesCount(); count > -1 {
		mainObjects.LabelCount.SetLabel(fmt.Sprintf("%v", count))
	}
}

/*
 * Buttons
 */

func ButtonReplaceClicked() {

	if svs.SearchOk {

		if ok, err := svs.SearchCtx.Replace(svs.IterStart, svs.IterEnd, svs.TextReplace); err == nil {
			if ok {

				dispStatusbar("Replaced")
			}
		} else {

			dispStatusbar(err.Error())
		}
	} else {

		dispStatusbar("Nothing has been searched at this time, please perform a search operation before trying to replacing something")
	}
}

func ButtonReplaceAllClicked() {

	if svs.SearchOk {

		if count, err := svs.SearchCtx.ReplaceAll(svs.TextReplace); err == nil {
			if count > 0 {

				dispStatusbar(fmt.Sprintf("%v occurence(s) replaced", count))
			}
		} else {

			dispStatusbar(err.Error())
		}
	} else {

		dispStatusbar("Nothing has been searched at this time, please perform a search operation before trying to replacing something")
	}
}

/* Search functions handling */

/* Async version of search function */
func ButtonForwardAsyncClicked() {

	if len(svs.TextSearch) != 0 {

		// Set iter, If it's first time search, set iter to the first of buffer
		iter := svs.IterEnd
		if iter == nil {
			iter = svs.Buffer.GetStartIter()
		}

		if err := svs.SearchAsync(

			// This is a callback function that executed when async search is done.
			// Inside we display the count of found items and their positions,
			// we perform a scroll to this position too.
			func() {

				if svs.HasFound {

					scrollAndTag()

				} else {

					// To avoid invalid iter when nothing is found at buffer end
					// and search button is clicked
					svs.IterEnd = iter
					svs.IterStart = nil
				}
			},
			iter,
			false); err != nil {

			dispStatusbar(err.Error())
			return
		}
	}
}

func ButtonBackwardAsyncClicked() {

	if len(svs.TextSearch) != 0 {

		// Set iter, If it's first time search, set iter to the first of buffer
		iter := svs.IterStart
		if iter == nil {
			iter = svs.Buffer.GetEndIter()
		}

		if err := svs.SearchAsync(

			// This is a callback function that executed when async search is done.
			// Inside we display the count of found items and their positions,
			// we perform a scroll to this position too.
			func() {

				if svs.HasFound {

					scrollAndTag()

				} else {

					// To avoid invalid iter when nothing is found at buffer start
					// and search button is clicked
					svs.IterEnd = nil
					svs.IterStart = iter
				}
			},
			iter,
			true); err != nil {

			dispStatusbar(err.Error())
			return
		}
	}
}

/* Sync version of search function */
func ButtonForwardClicked() {

	if len(svs.TextSearch) != 0 {

		// Set iter, If it's first time search, set iter to the first of buffer
		iter := svs.IterEnd
		if svs.IterEnd == nil {
			iter = svs.Buffer.GetStartIter()
		}

		if err := svs.Search(iter, false); err != nil {

			dispStatusbar(err.Error())
			return
		}

		if svs.HasFound {

			scrollAndTag()

		} else {

			// To avoid invalid iter when nothing is found at buffer end
			// and search button is clicked
			svs.IterEnd = iter
			svs.IterStart = nil
		}
	}
}

func ButtonBackwardClicked() {

	if len(svs.TextSearch) != 0 {

		// Set iter, If it's first time search, set iter to the last of buffer
		iter := svs.IterStart
		if svs.IterStart == nil {
			iter = svs.Buffer.GetEndIter()
		}

		if err := svs.Search(iter, true); err != nil {

			dispStatusbar(err.Error())
			return
		}

		if svs.HasFound {

			scrollAndTag()

		} else {

			// To avoid invalid iter when nothing is found at buffer start
			// and search button is clicked
			svs.IterStart = iter
			svs.IterEnd = nil
		}
	}
}

func ButtonStyleClicked() {

	if s := svs.StyleChooserDialog(); s != "" {

		mainObjects.ComboBoxTextStyle.SetActiveID(s)
	} else {

		dispStatusbar("Could not set another StyleScheme.")
	}
}

func FileChooserButtonFileSet(chooser *gtk.FileChooserButton) {

	if err := svs.LoadSource(chooser.GetFilename(), false); err != nil {

		dispStatusbar(err.Error())
	}
}

/*
 * Checkboxes
 */

func CheckButtonBracketsToggled(chk *gtk.CheckButton) {
	svs.Buffer.SetHighlightMatchingBrackets(chk.GetActive())
}

func CheckButtonHighlightToggled(chk *gtk.CheckButton) {
	svs.Buffer.SetHighlightSyntax(chk.GetActive())
}

func CheckButtonHighlightLineToggled(chk *gtk.CheckButton) {
	svs.View.SetHighlightCurrentLine(chk.GetActive())
}

func CheckButtonNumbersToggled(chk *gtk.CheckButton) {
	svs.View.SetShowLineNumbers(chk.GetActive())
}

func CheckButtonShowMarksToggled(chk *gtk.CheckButton) {
	svs.View.SetShowLineMarks(chk.GetActive())
}

func CheckButtonShowMapToggled(chk *gtk.CheckButton) {

	mainOptions.MainWinWidth, mainOptions.MainWinHeight = mainObjects.Window.GetSize()

	if chk.GetActive() {

		mainObjects.Paned.SetPosition(mainOptions.MainWinWidth - mainOptions.PanedWidth)
	} else {

		mainObjects.Paned.SetPosition(mainOptions.MainWinWidth)
	}

	svs.Map.SetVisible(chk.GetActive())
}

func CheckButtonWrapToggled(chk *gtk.CheckButton) {

	if chk.GetActive() {

		svs.View.SetWrapMode(gtk.WRAP_WORD_CHAR)
	} else {

		svs.View.SetWrapMode(gtk.WRAP_NONE)
	}
}

// Search options
func CheckButtonWrapAroundToggled(chk *gtk.CheckButton) {
	svs.WrapAround = chk.GetActive()
}

func CheckButtonWordBoundariesToggled(chk *gtk.CheckButton) {
	svs.WordBoundaries = chk.GetActive()
}

func CheckButtonRegexpToggled(chk *gtk.CheckButton) {
	svs.UseRegexp = chk.GetActive()
}

func CheckButtonCaseSensitiveToggled(chk *gtk.CheckButton) {
	svs.CaseSensitive = chk.GetActive()
}

func CheckButtonHighlightFoundToggled(chk *gtk.CheckButton) {
	svs.HighlightFound = chk.GetActive()
}

/*
 * Comboboxes
 */

func ComboBoxTextLanguageChanged(cbx *gtk.ComboBoxText) {
	mainOptions.DefaulLanguage = cbx.GetActiveID()
	svs.SetLanguage(mainOptions.DefaulLanguage)
}

func ComboBoxTextStyleChanged(cbx *gtk.ComboBoxText) {
	mainOptions.DefaultStyle = cbx.GetActiveID()
	svs.SetStyleScheme(mainOptions.DefaultStyle)
}

/*
 * Entries
 */

func EntryReplaceChanged(e *gtk.Entry) {
	svs.TextReplace, _ = e.GetText()

	// Clear labels for count and current search
	mainObjects.LabelCurr.SetLabel("")
	mainObjects.LabelCount.SetLabel("")
}

func EntrySearchChanged(e *gtk.Entry) {
	svs.TextSearch, _ = e.GetText()

	// Clear labels for count and current search
	mainObjects.LabelCurr.SetLabel("")
	mainObjects.LabelCount.SetLabel("")
}
