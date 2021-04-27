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

/*****************************************\
* Some pre-defined tags that could be used *
*******************************************/

// TextTag: Is a structure that own some methods to handle gtk.TextTag
// and make more user-friendly the usage of it in a TextView. Some
// explanations and definitions are availables trought the official
// manual: https://developer.gnome.org/gtk3/stable/GtkTextTag.html
// type TextTagList struct {
// 	TagList  []tagStore
// 	Buff     *gtk.TextBuffer
// 	tagTable *gtk.TextTagTable
// }

// // TextTagProp: tag' properties record.
// type TextTagProp struct {
// 	Name  string
// 	Value interface{}
// }

// // TextTagList: Complete tag record.
// type TextTagRecord struct {
// 	TagName string
// 	Props   []TextTagProp
// }

// // tagStore: Used to track already added tags.
// type tagStore struct {
// 	Name string
// 	Tag  *gtk.TextTag
// }

// // TextTagPropNew: Convenient function to create and fill a new TextTagProp struct.
// // i.e: Modify("blue_foreground",
// //         TextTagPropNew("foreground", "#0022FF"),
// //         TextTagPropNew("strikethrough", true))
// func TextTagPropNew(name string, value interface{}) (outProp map[string]interface{}) {
// 	outProp = make(map[string]interface{})
// 	outProp[name] = value
// 	return
// }

// // TextTagNew: Create and initialize a new TextTagNew structure
// func TextTagListNew(buff *gtk.TextBuffer) (tt *TextTagList, err error) {
// 	tt = new(TextTagList)
// 	err = tt.Init(buff)
// 	return
// }

// // Add: Adding tag with property (could be multiple).
// // ie: Add("theName", []TextTagProp{
// //         {Name: "weight", Value: pango.WEIGHT_BOLD},
// //         {Name: "size", Value: 15 * pango.SCALE}}...)
// func (tt *TextTagList) Add(name string, tagProps ...TextTagProp) (tag *gtk.TextTag, err error) {

// 	props := make([]TextTagProp, len(tagProps))
// 	for idx, prop := range tagProps {
// 		props[idx] = prop
// 	}
// 	return tt.createTag(name, props...)
// }

// // Remove: stored tag from TagTable
// func (tt *TextTagList) Remove(name string) {
// 	for idx, tag := range tt.TagList {
// 		if tag.Name == name {
// 			tt.tagTable.Remove(tag.Tag)                                  // Remove from buffer
// 			tt.TagList = append(tt.TagList[:idx], tt.TagList[idx+1:]...) // Remove from store list
// 			break
// 		}
// 	}
// }

// // RemoveAllTagsFromTagsList: remove all defaults and user defined tags,
// // (contained into "TagsList")
// func (tt *TextTagList) RemoveAllTagsFromTagsList() {
// 	for _, tag := range tt.TagList {
// 		tt.Buff.RemoveTagByName(tag.Name, tt.Buff.GetStartIter(), tt.Buff.GetEndIter())
// 	}
// }

// // Modify: stored tag from TagTable
// func (tt *TextTagList) Modify(name string, tagProps ...TextTagProp) (err error) {
// 	for _, tag := range tt.TagList {
// 		if tag.Name == name {
// 			for _, prop := range tagProps {
// 				if err = tag.Tag.SetProperty(prop.Name, prop.Value); err != nil {
// 					return
// 				}
// 			}
// 			break
// 		}
// 	}
// 	return
// }

// // ApplyTagsByName: Apply one or more Tags to the previously initialized TextBuffer
// func (tt *TextTagList) ApplyTagsByName(startIter, endIter *gtk.TextIter, tags ...string) {
// 	for _, name := range tags {
// 		tt.Buff.ApplyTagByName(name, startIter, endIter)
// 	}
// }

// // ApplyTagsByNameAtOffets: Apply one or more Tags to the previously initialized TextBuffer
// // At one or multiples offsets (character position)
// func (tt *TextTagList) ApplyTagsByNameAtOffets(offsets [][]int, tags ...string) {
// 	var startIter, endIter *gtk.TextIter
// 	for _, offset := range offsets {
// 		startIter = tt.Buff.GetIterAtOffset(offset[0])
// 		endIter = tt.Buff.GetIterAtOffset(offset[1])
// 		for _, name := range tags {
// 			tt.Buff.ApplyTagByName(name, startIter, endIter)
// 		}
// 	}
// }

// // RemoveTagsByName: from TextBuffer (clear tags)
// func (tt *TextTagList) RemoveTagsByName(startIter, endIter *gtk.TextIter, tags ...string) {
// 	for _, name := range tags {
// 		tt.Buff.RemoveTagByName(name, startIter, endIter)
// 	}
// }

// // RemoveAllTags: from TextBuffer (clear tags)
// func (tt *TextTagList) RemoveAllTags(startIter, endIter *gtk.TextIter) {
// 	tt.Buff.RemoveAllTags(startIter, endIter)
// }

// // Init: initialize the default TextBuffer Tags set.
// func (tt *TextTagList) Init(buff *gtk.TextBuffer) (err error) {
// 	if tt.tagTable == nil { // Do it only if not previously done (only one time)
// 		if tt.tagTable, err = buff.GetTagTable(); err == nil {
// 			tt.Buff = buff

// 			// got this issue when using it:
// 			// panic: runtime error: cgo argument has Go pointer to Go pointer
// 			// github.com/gotk3/gotk3/glib.(*Value).SetPointer.func1(0xc000198468, 0xc000198450)
// 			// github.com/gotk3/gotk3/glib/glib.go:1365 +0x5a
// 			// github.com/gotk3/gotk3/glib.(*Value).SetPointer(0xc000198468, 0xc000198450)
// 			// github.com/gotk3/gotk3/glib/glib.go:1365 +0x35
// 			// github.com/gotk3/gotk3/glib.GValue(0x8d4240, 0xc000198450, 0x7fe7d81145c0, 0x1, 0x0)
// 			// github.com/gotk3/gotk3/glib/glib.go:1105 +0x4eb
// 			// github.com/gotk3/gotk3/glib.(*Object).SetProperty(0xc000198458, 0x967332, 0xa, 0x8d4240, 0xc000198450, 0x0, 0x0)
// 			// github.com/gotk3/gotk3/glib/glib.go:682 +0xba

// 			// bgcRGBA := gdk.NewRGBA()
// 			// if !bgcRGBA.Parse("#FFFF99") {
// 			// 	log.Fatal("error parsing RGBA")
// 			// }
// 			// rgba := bgcRGBA.Floats()
// 			// bgcRGBA.SetColors(rgba[0], rgba[1], rgba[2], 0.5)

// 			list := []TextTagList{
// 				// {TagName: "yellow_background", Props: []TextTagProp{{Name: "background", Value: bgcRGBA}}},

// 				{TagName: "normal", Props: []TextTagProp{{Name: "style", Value: pango.STYLE_NORMAL}}},
// 				{TagName: "oblique", Props: []TextTagProp{{Name: "style", Value: pango.STYLE_OBLIQUE}}},
// 				{TagName: "italic", Props: []TextTagProp{{Name: "style", Value: pango.STYLE_ITALIC}}},
// 				{TagName: "bold", Props: []TextTagProp{{Name: "weight", Value: pango.WEIGHT_BOLD}}},
// 				{TagName: "underline", Props: []TextTagProp{{Name: "underline", Value: pango.UNDERLINE_SINGLE}}},
// 				{TagName: "double_underline", Props: []TextTagProp{{Name: "underline", Value: pango.UNDERLINE_DOUBLE}}},
// 				{TagName: "strikethrough", Props: []TextTagProp{{Name: "strikethrough", Value: true}}},
// 				{TagName: "heading", Props: []TextTagProp{{Name: "weight", Value: pango.WEIGHT_BOLD}, {Name: "size", Value: 15 * pango.SCALE}}},
// 				{TagName: "big", Props: []TextTagProp{{Name: "size", Value: 20 * pango.SCALE}}},
// 				// Cannot use pango variables, ... unable to convert type ...
// 				{TagName: "xx-small", Props: []TextTagProp{{Name: "scale", Value: 0.5787037037037}}}, /* pango.SCALE_XX_SMALL */
// 				{TagName: "x-small", Props: []TextTagProp{{Name: "scale", Value: 0.6444444444444}}},  /* pango.SCALE_X_SMALL */
// 				{TagName: "small", Props: []TextTagProp{{Name: "scale", Value: 0.8333333333333}}},    /* pango.SCALE_SMALL */
// 				{TagName: "medium", Props: []TextTagProp{{Name: "scale", Value: 1.0}}},               /* pango.SCALE_MEDIUM */
// 				{TagName: "large", Props: []TextTagProp{{Name: "scale", Value: 1.2}}},                /* pango.SCALE_LARGE */
// 				{TagName: "x-large", Props: []TextTagProp{{Name: "scale", Value: 1.4399999999999}}},  /* pango.SCALE_X_LARGE */
// 				{TagName: "xx-large", Props: []TextTagProp{{Name: "scale", Value: 1.728}}},           /* pango.SCALE_XX_LARGE */
// 				{TagName: "superscript", Props: []TextTagProp{{Name: "rise", Value: 10 * pango.SCALE}, {Name: "size", Value: 8 * pango.SCALE}}},
// 				{TagName: "subscript", Props: []TextTagProp{{Name: "rise", Value: -10 * pango.SCALE}, {Name: "size", Value: 8 * pango.SCALE}}},
// 				{TagName: "monospace", Props: []TextTagProp{{Name: "family", Value: "monospace"}}},
// 				{TagName: "blue_foreground", Props: []TextTagProp{{Name: "foreground", Value: "#3050FF"}}},   // soft blue
// 				{TagName: "yellow_background", Props: []TextTagProp{{Name: "background", Value: "#FFFFAF"}}}, // light yellow
// 				{TagName: "foreground_color", Props: []TextTagProp{{Name: "foreground", Value: "#0022DF"}}},  // deep blue
// 				{TagName: "background_color", Props: []TextTagProp{{Name: "background", Value: "#AAFFAA"}}},  // light green
// 			}
// 			for _, tag := range list {
// 				if _, err = tt.createTag(tag.TagName, tag.Props...); err != nil {
// 					return
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// // createTag: create tag with properties and add it to buffer.
// // I do not use the native gtk3' "TextBuffer.CreateTag" function
// // because I need to use a structure with some methods instead
// // of a map.
// // TODO rewrite with version below
// func (tt *TextTagList) createTag(name string, props ...TextTagProp) (tag *gtk.TextTag, err error) {
// 	if tag, err = gtk.TextTagNew(name); err == nil {
// 		tt.tagTable.Add(tag)
// 		for _, prop := range props {
// 			if err = tag.SetProperty(prop.Name, prop.Value); err != nil {
// 				err = errors.New(fmt.Sprintf("%s, %s: %v\n", err.Error(), prop.Name, prop.Value))
// 				break
// 			}
// 		}
// 		tt.TagList = append(tt.TagList, tagStore{Name: name, Tag: tag})
// 	}
// 	return
// }

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
