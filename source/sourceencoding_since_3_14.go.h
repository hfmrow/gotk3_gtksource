#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GtkSourceEncoding *
toGtkSourceEncoding(void *p)
{
	return (GTK_SOURCE_ENCODING_H(p));
}
