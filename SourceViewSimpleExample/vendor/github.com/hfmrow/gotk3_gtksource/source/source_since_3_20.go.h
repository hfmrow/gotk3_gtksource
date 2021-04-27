#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GtkSourceTag *
toGtkSourceTag(void *p)
{
	return (GTK_SOURCE_TAG(p));
}

