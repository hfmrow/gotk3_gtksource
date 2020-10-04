#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GtkSourceSpaceDrawer *
toGtkSourceSpaceDrawer(void *p)
{
	return (GTK_SOURCE_SPACE_DRAWER(p));
}
