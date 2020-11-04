// textTag.go

/*
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2019 H.F.M - TextTag structure
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package textTag

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"

	"github.com/hfmrow/gotk3_gtksource/source"
)

// TextTagPropNew: Convenient function to create and fill a new TextTagProp struct.
// i.e: Modify("blue_foreground",
//         TextTagPropNew("foreground", "#0022FF"),
//         TextTagPropNew("strikethrough", true))
func TextTagPropNew(name string, value interface{}) (outProp map[string]interface{}) {
	outProp = make(map[string]interface{})
	outProp[name] = value
	return
}

// CreateTag: create tag with properties and add it to buffer.
// Check whether tag already exist in this case, only add or
// update property value.
func TagCreateIfNotExists(buff interface{}, tagName string, props map[string]interface{}) (tag *gtk.TextTag) {
	var ok bool

	if tag, ok = TagLookupIfExists(buff, tagName); ok {
		for name, value := range props {
			tag.SetProperty(name, value)
		}
	} else {
		switch b := buff.(type) {
		case *gtk.TextBuffer:
			tag = b.CreateTag(tagName, props) // add tag & properties
		case *source.SourceBuffer:
			tag = b.CreateTag(tagName, props) // add tag & properties
		}
	}
	return
}

// RemoveTag: from entire buffer if exists.
func TagRemoveIfExists(buff interface{}, tagName string) {
	if tag, ok := TagLookupIfExists(buff, tagName); ok {

		switch b := buff.(type) {
		case *gtk.TextBuffer:
			b.RemoveTag(tag,
				b.GetStartIter(),
				b.GetEndIter())
		case *source.SourceBuffer:
			b.RemoveTag(tag,
				b.GetStartIter(),
				b.GetEndIter())
		}
	}
	return
}

// LookupExistingTag:
func TagLookupIfExists(buff interface{}, tagName string) (tag *gtk.TextTag, ok bool) {
	var err error
	var tagTable *gtk.TextTagTable

	switch b := buff.(type) {
	case *gtk.TextBuffer:
		tagTable, err = b.GetTagTable()
	case *source.SourceBuffer:
		tagTable, err = b.GetTagTable()
	}
	if err == nil {
		if tag, _ := tagTable.Lookup(tagName); tag != nil { // Check whether it exists
			return tag, true
		}
	} else {
		fmt.Printf("Unable to get TagTable for %s: %s\n", tagName, err.Error())
	}
	return
}
