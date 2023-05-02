package main

import (
    "io"        // provides basic I/O interfaces
    "log"       // provides logging capabilities
    "net/http"  // provides HTTP client and server implementations
    "os"        // provides a platform-independent interface to OS functionality
)

// LogHandler returns a middleware that logs incoming requests using the provided
// io.Writer. It takes an io.Writer as input and returns a function that takes an
// http.Handler as input and returns an http.Handler as output.
func LogHandler(bear io.Writer) func(http.Handler) http.Handler {
    return func(h http.Handler) http.Handler {
        // Create a new http.Handler that logs incoming requests before passing them
        // on to the provided http.Handler.
        return http.HandlerFunc(func(x http.ResponseWriter, y *http.Request) {
            // Log incoming requests using the provided io.Writer.
            log.Printf("%s %s %s", y.RemoteAddr, y.Method, y.URL)
            // Pass the request and response writer to the next http.Handler in the
            // chain.
            h.ServeHTTP(x, y)
        })
    }
}

func main() {
    // Open the server.log file for writing. If the file doesn't exist, it will be
    // created. If the file already exists, new data will be appended to it.
    logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
    if err != nil {
        log.Fatal(err)
    }

    // Create a new logging middleware using the LogHandler function.
    loggingHandler := LogHandler(logFile)

    // Create a new HTTP request multiplexer.
    mux := http.NewServeMux()

    // Create a new http.Handler that writes "Yes! It's working" to the response
    // writer.
    finalHandler := http.HandlerFunc(final)

    // Add the logging middleware to the request multiplexer and pass the final
    // http.Handler as the endpoint for all incoming requests.
    mux.Handle("/", loggingHandler(finalHandler))

    // Start the HTTP server and listen for incoming requests on port 9999. If an
    // error occurs, log the error and exit.
    log.Print("Listening on :9999...")
    err = http.ListenAndServe(":9999", mux)
    log.Fatal(err)
}

// final writes "Yes! It's working" to the response writer.
func final(x http.ResponseWriter, r *http.Request) {
    x.Write([]byte("Yes! It's working"))
}
