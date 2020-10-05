# Golang GtkSourceView binding for use with gotk3

This repository must be used with [gotk3: Go bindings for GTK3](https://github.com/gotk3/gotk3) only.

---

Example of GtkSourceView usage in Go: [SourceViewSimpleExample](https://github.com/hfmrow/gotk3_gtksource/tree/main/SourceViewSimpleExample)

---

##### installation:

- Ubuntu linux only* see at bottom why.
- Designed to be used with gtk3 >= 3.16 and gtksourceview >=3.18-4
- Be sure you have [gotk3](https://github.com/gotk3/gotk3/wiki#installation) installed correctly and working right.
- Use classic command: `go get github.com/hfmrow/gotk3_gtksource/source`
- Golang GtkSourceView follow the same rules that gotk3 for compiling differents version of gtksourceview.

> ##### Information about [libgtksourceview](https://packages.ubuntu.com/search?lang=en&keywords=libgtksourceview) versions:
> 
> Since libgtksourceview-3.0-dev start at GTK >= 3.16, minimal GtkSourceView version is limited to: libgtksourceview >= 3.18

> ##### Library Installation (Ubuntu):
> 
> ##### [**xenial (16.04LTS)**](https://packages.ubuntu.com/xenial/libgtksourceview-3.0-dev), [**bionic (18.04LTS)**](https://packages.ubuntu.com/bionic/libgtksourceview-3.0-dev)
> 
> ```bash
> $ sudo apt install libgtksourceview-3.0-dev
> ```
> 
> ##### [**focal (20.04LTS)**](https://packages.ubuntu.com/focal/libgtksourceview-4-dev) this version of Ubuntu may use libgtksourceview-3.0-dev too
> 
> ```bash
> $ sudo apt install libgtksourceview-4-dev
> ```

> *To install targeting your version of GtkSourceView:*
> 
> ```shell
>     $ go get -tags gtksourceview_4 github.com/hfmrow/gotk3_gtksource/source
> or:
>     $ go get -tags gtksourceview_3_18 github.com/hfmrow/gotk3_gtksource/source
> ```
> 
> *To rebuild the package for another GtkSourceView version:*
> 
> ```shell
> $ go install -tags gtksourceview_X_XX github.com/hfmrow/gotk3_gtksource/source
> ```

### Gotk3 GtkSourceView wrapping progression

##### what for ?, check this out: [GtkSourceView](https://wiki.gnome.org/Projects/GtkSourceView)

---

- [x] [GtkSourceView Initialization and Finalization](https://developer.gnome.org/gtksourceview/stable/gtksourceview-4.0-GtkSourceView-Initialization-and-Finalization.html)

---

- [x] [GtkSourceBuffer](https://developer.gnome.org/gtksourceview/stable/GtkSourceBuffer.html) — Subclass of GtkTextBuffer
  
- [x] [GtkSourceView](https://developer.gnome.org/gtksourceview/stable/GtkSourceView.html) — Subclass of GtkTextView
  
- [x] [GtkSourceLanguage](https://developer.gnome.org/gtksourceview/stable/GtkSourceLanguage.html) — Represents a syntax highlighted language
  
- [x] [GtkSourceLanguageManager](https://developer.gnome.org/gtksourceview/stable/GtkSourceLanguageManager.html) — Provides access to GtkSourceLanguages
  

---

- [x] [GtkSourceStyle](https://developer.gnome.org/gtksourceview/stable/GtkSourceStyle.html) — Represents a style
  
- [x] [GtkSourceStyleScheme](https://developer.gnome.org/gtksourceview/stable/GtkSourceStyleScheme.html) — Controls the appearance of GtkSourceView
  
- [x] [GtkSourceStyleSchemeManager](https://developer.gnome.org/gtksourceview/stable/GtkSourceStyleSchemeManager.html) — Provides access to GtkSourceStyleSchemes
  
- [x] [GtkSourceStyleSchemeChooser](https://developer.gnome.org/gtksourceview/stable/GtkSourceStyleSchemeChooser.html) — Interface implemented by widgets for choosing style schemes
  
- [x] [GtkSourceStyleSchemeChooserButton](https://developer.gnome.org/gtksourceview/stable/GtkSourceStyleSchemeChooserButton.html) — A button to launch a style scheme selection dialog
  
- [x] [GtkSourceStyleSchemeChooserWidget](https://developer.gnome.org/gtksourceview/stable/GtkSourceStyleSchemeChooserWidget.html) — A widget for choosing style schemes
  

---

- [x] [GtkSourceCompletion](https://developer.gnome.org/gtksourceview/stable/GtkSourceCompletion.html) — Main Completion Object
  
- [x] [GtkSourceCompletionContext](https://developer.gnome.org/gtksourceview/stable/GtkSourceCompletionContext.html) — The context of a completion
  
- [x] [GtkSourceCompletionInfo](https://developer.gnome.org/gtksourceview/stable/GtkSourceCompletionInfo.html) — Calltips object
  
- [x] [GtkSourceCompletionItem](https://developer.gnome.org/gtksourceview/stable/GtkSourceCompletionItem.html) — Simple implementation of GtkSourceCompletionProposal
  
- [x] [GtkSourceCompletionProposal](https://developer.gnome.org/gtksourceview/stable/GtkSourceCompletionProposal.html) — Completion proposal interface
  
- [x] [GtkSourceCompletionProvider](https://developer.gnome.org/gtksourceview/stable/GtkSourceCompletionProvider.html) — Completion provider interface
  
- [x] [GtkSourceCompletionWords](https://developer.gnome.org/gtksourceview/stable/GtkSourceCompletionWords.html) — A GtkSourceCompletionProvider for the completion of words
  

---

- [x] [GtkSourceGutter](https://developer.gnome.org/gtksourceview/stable/GtkSourceGutter.html) — Gutter object for GtkSourceView
  
- [x] [GtkSourceGutterRenderer](https://developer.gnome.org/gtksourceview/stable/GtkSourceGutterRenderer.html) — Gutter cell renderer
  
- [x] [GtkSourceGutterRendererPixbuf](https://developer.gnome.org/gtksourceview/stable/GtkSourceGutterRendererPixbuf.html) — Renders a pixbuf in the gutter
  
- [x] [GtkSourceGutterRendererText](https://developer.gnome.org/gtksourceview/stable/GtkSourceGutterRendererText.html) — Renders text in the gutter
  
- [x] [GtkSourceMark](https://developer.gnome.org/gtksourceview/stable/GtkSourceMark.html) — Mark object for GtkSourceBuffer
  
- [x] [GtkSourceMarkAttributes](https://developer.gnome.org/gtksourceview/stable/GtkSourceMarkAttributes.html) — The source mark attributes object
  

---

- [ ] [GtkSourcePrintCompositor](https://developer.gnome.org/gtksourceview/stable/GtkSourcePrintCompositor.html) — Compose a GtkSourceBuffer for printing

---

- [x] [GtkSourceSearchContext](https://developer.gnome.org/gtksourceview/stable/GtkSourceSearchContext.html) — Search context
  
- [x] [GtkSourceSearchSettings](https://developer.gnome.org/gtksourceview/stable/GtkSourceSearchSettings.html) — Search settings
  

---

- [x] [GtkSourceEncoding](https://developer.gnome.org/gtksourceview/stable/GtkSourceEncoding.html) — Character encoding

---

- [x] [GtkSourceMap](https://developer.gnome.org/gtksourceview/stable/GtkSourceMap.html) — Widget that displays a map for a specific GtkSourceView
  
- [ ] [GtkSourceRegion](https://developer.gnome.org/gtksourceview/stable/GtkSourceRegion.html) — Region utility
  
- [x] [GtkSourceSpaceDrawer](https://developer.gnome.org/gtksourceview/stable/GtkSourceSpaceDrawer.html) — Represent white space characters with symbols
  
- [x] [GtkSourceTag](https://developer.gnome.org/gtksourceview/stable/GtkSourceTag.html) — A tag that can be applied to text in a GtkSourceBuffer
  
- [x] [GtkSourceUndoManager](https://developer.gnome.org/gtksourceview/stable/GtkSourceUndoManager.html) — Undo manager interface for GtkSourceView
  
- [x] [GtkSourceUtils](https://developer.gnome.org/gtksourceview/stable/gtksourceview-4.0-GtkSourceUtils.html) — Utility functions
  
- [x] [Version Information](https://developer.gnome.org/gtksourceview/stable/gtksourceview-4.0-Version-Information.html) — Macros and functions to check the GtkSourceView version
  

---

#### Not fully wrapped but usable (implementation stopped)

> - async operations are not wrapped
> 
> This [class](https://developer.gnome.org/gtksourceview/stable/GtkSourceEncoding.html#GtkSourceEncoding.description) is no longer maintained, patches are not accepted. There is a better implementation in the [Tepl](https://wiki.gnome.org/Projects/Tepl) library.

- [x] [GtkSourceFile](https://developer.gnome.org/gtksourceview/stable/GtkSourceFile.html) — On-disk representation of a GtkSourceBuffer
- [x] [GtkSourceFileLoader](https://developer.gnome.org/gtksourceview/stable/GtkSourceFileLoader.html) — Load a file into a GtkSourceBuffer
- [x] [GtkSourceFileSaver](https://developer.gnome.org/gtksourceview/stable/GtkSourceFileSaver.html) — Save a GtkSourceBuffer into a file

---

#### Informations, Documentation: [GtkSourceView](https://developer.gnome.org/gtksourceview/4.2/)

---

**(*)** Why Ubuntu Linux only ... because i really don't know how to explain and deploy it on window and darwin, i'm sorry. If you are able to doing that, you're welcome. All information are available at [Ubuntu - libgtksourceview](https://packages.ubuntu.com/search?lang=en&keywords=libgtksourceview). Theoretically, **sourceview 4** may work on amd64 arm64 armhf i386 ppc64el s390x, **sourceview 3.0** amd64 arm64 armhf i386 powerpc ppc64el s390x both depending on OS version.
