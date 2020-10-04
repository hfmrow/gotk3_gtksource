#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GtkSourceFileSaver *
toGtkSourceFileSaver(void *p)
{
	return (GTK_SOURCE_FILE_SAVER(p));
}

static GtkSourceFile *
toGtkSourceFile(void *p)
{
	return (GTK_SOURCE_FILE(p));
}

static GtkSourceFileLoader *
toGtkSourceFileLoader(void *p)
{
	return (GTK_SOURCE_FILE_LOADER(p));
}
