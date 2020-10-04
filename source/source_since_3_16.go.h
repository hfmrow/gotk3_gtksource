#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GtkSourceStyleSchemeChooser *
toGtkSourceStyleSchemeChooser(void *p)
{
	return (GTK_SOURCE_STYLE_SCHEME_CHOOSER(p));
}

static GtkSourceStyleSchemeChooserButton *
toGtkSourceStyleSchemeChooserButton(void *p)
{
	return (GTK_SOURCE_STYLE_SCHEME_CHOOSER_BUTTON(p));
}

static GtkSourceStyleSchemeChooserWidget *
toGtkSourceStyleSchemeChooserWidget(void *p)
{
	return (GTK_SOURCE_STYLE_SCHEME_CHOOSER_WIDGET(p));
}
