package ui

import (
	"io"
	"io/fs"
	"net/http"
	"strings"
)

type spaHandler struct {
	fileServer http.Handler
	fs         http.FileSystem
	index      []byte
}

// NewHandler returns an http.Handler that serves a Nuxt SPA from the embedded FS.
// It serves exact file matches from the FS and falls back to index.html for all
// other paths so that client-side routing works correctly.
func NewHandler(fsys fs.FS) (http.Handler, error) {
	sub, err := fs.Sub(fsys, "dist")
	if err != nil {
		return nil, err
	}

	hfs := http.FS(sub)

	f, err := hfs.Open("/index.html")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	index, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return &spaHandler{
		fileServer: http.FileServer(hfs),
		fs:         hfs,
		index:      index,
	}, nil
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
	}

	f, err := h.fs.Open(upath)
	if err == nil {
		fi, statErr := f.Stat()
		f.Close()
		if statErr == nil && !fi.IsDir() {
			h.fileServer.ServeHTTP(w, r)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(h.index)
}
