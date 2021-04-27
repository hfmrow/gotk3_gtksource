#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GtkSourceMap *
toGtkSourceMap(void *p)
{
	return (GTK_SOURCE_MAP(p));
}

