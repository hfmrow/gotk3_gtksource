
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static inline gchar** make_strings(int count) {
	return (gchar**)malloc(sizeof(gchar*) * count);
}

static inline void destroy_strings(gchar** strings) {
	free(strings);
}

static inline gchar* get_string(gchar** strings, int n) {
	return strings[n];
}

static inline void set_string(gchar** strings, int n, gchar* str) {
	strings[n] = str;
}

static inline gchar** next_gcharptr(gchar** s) { return (s+1); }

/*
 * Gtk Objects
 */

static GtkWidget *
toGtkWidget(void *p)
{
	return (GTK_WIDGET(p));
}

static GtkTextBuffer *
toGtkTextBuffer(void *p)
{
	return (GTK_TEXT_BUFFER(p));
}

static GtkTextView *
toGtkTextView(void *p)
{
	return (GTK_TEXT_VIEW(p));
}

static GtkTextTag *
toGtkTextTag(void *p)
{
	return (GTK_TEXT_TAG(p));
}

static GtkTextTagTable *
toGtkTextTagTable(void *p)
{
	return (GTK_TEXT_TAG_TABLE(p));
}

static GtkTextMark *
toGtkTextMark(void *p)
{
	return (GTK_TEXT_MARK(p));
}
