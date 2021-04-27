
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

/*
 * GAsyncReadyCallback for use inside gtk package
 */

extern void _goAsyncReadyCallbacks (GObject *source_object,
									GAsyncResult *res,
									gpointer user_data);

static inline void _gtk_source_search_context_forward_async(GtkSourceSearchContext *search,
                                                            const GtkTextIter *iter,
                                                            GCancellable *cancellable,
                                                            gpointer user_data) {
    gtk_source_search_context_forward_async(search, iter, cancellable, (GAsyncReadyCallback)(_goAsyncReadyCallbacks), user_data);
}

static inline void _gtk_source_search_context_backward_async(GtkSourceSearchContext *search,
                                                            const GtkTextIter *iter,
                                                            GCancellable *cancellable,
                                                            gpointer user_data) {
    gtk_source_search_context_backward_async(search, iter, cancellable, (GAsyncReadyCallback)(_goAsyncReadyCallbacks), user_data);
}


