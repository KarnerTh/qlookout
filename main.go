package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/KarnerTh/query-lookout/core/orchestration"
)

//go:embed all:web/build
var webFiles embed.FS

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	setupCore()
	setupWeb()
	setupQuitWatcher()
}

// Keep program running until SIGINT or SIGTERM
func setupQuitWatcher() {
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	slog.Info("Service stopped - see you soon 👋")
}

func setupCore() {
	slog.Info("Starting core")
	go orchestration.Setup()
}

func setupWeb() {
	slog.Info("Starting web")

	webSubFiles, err := fs.Sub(webFiles, "web/build")
	if err != nil {
		slog.Error("Could not get web files", slog.Any("error", err))
		return
	}

	httpFS := http.FS(webSubFiles)
	frontendFS := http.FileServer(httpFS)
	serveIndex := serveFileContents("index.html", httpFS)
	http.Handle("/", intercept404(frontendFS, serveIndex))

	slog.Info("Listening on 63000")
	err = http.ListenAndServe(":63000", nil)
	if err != nil {
		slog.Error("Could not start server", slog.Any("error", err))
	}
}

// Source: https://hackandsla.sh/posts/2021-11-06-serve-spa-from-go/
func intercept404(handler, on404 http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hookedWriter := &hookedResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(hookedWriter, r)

		if hookedWriter.got404 {
			on404.ServeHTTP(w, r)
		}
	})
}

type hookedResponseWriter struct {
	http.ResponseWriter
	got404 bool
}

func (hrw *hookedResponseWriter) WriteHeader(status int) {
	if status == http.StatusNotFound {
		// Don't actually write the 404 header, just set a flag.
		hrw.got404 = true
	} else {
		hrw.ResponseWriter.WriteHeader(status)
	}
}

func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
	if hrw.got404 {
		// No-op, but pretend that we wrote len(p) bytes to the writer.
		return len(p), nil
	}

	return hrw.ResponseWriter.Write(p)
}

func serveFileContents(file string, files http.FileSystem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Restrict only to instances where the browser is looking for an HTML file
		if !strings.Contains(r.Header.Get("Accept"), "text/html") {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 not found")

			return
		}

		// Open the file and return its contents using http.ServeContent
		index, err := files.Open(file)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s not found", file)

			return
		}

		fi, err := index.Stat()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s not found", file)

			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeContent(w, r, fi.Name(), fi.ModTime(), index)
	}
}
