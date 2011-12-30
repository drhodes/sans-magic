package main

import (
	. "sansmagic"
    "flag"
    "http"
    "log"
	"io"
	"os"
)

var (
    ErrLineTooLong          = &http.ProtocolError{"header line too long"}
    ErrHeaderTooLong        = &http.ProtocolError{"header too long"}
    ErrShortBody            = &http.ProtocolError{"entity body too short"}
    ErrNotSupported         = &http.ProtocolError{"feature not supported"}
    ErrUnexpectedTrailer    = &http.ProtocolError{"trailer header without chunked transfer encoding"}
    ErrMissingContentLength = &http.ProtocolError{"missing ContentLength in HEAD response"}
    ErrNotMultipart         = &http.ProtocolError{"request Content-Type isn't multipart/form-data"}
    ErrMissingBoundary      = &http.ProtocolError{"no multipart boundary param Content-Type"}
	
	//Errors introduced by the HTTP server.
    ErrWriteAfterFlush = os.NewError("Conn.Write called after Flush")
    ErrBodyNotAllowed  = os.NewError("http: response status code does not allow body")
    ErrHijacked        = os.NewError("Conn has been hijacked")
    ErrContentLength   = os.NewError("Conn.Write wrote more than the declared Content-Length")
    ErrPersistEOF = &http.ProtocolError{"persistent connection closed"}
    ErrPipeline   = &http.ProtocolError{"pipeline error"}

	//DefaultClient is the default Client and is used by Get, Head, and Post.
	DefaultClient = &http.Client{}

	//DefaultServeMux is the default ServeMux used by Serve.
	DefaultServeMux = http.NewServeMux()

	// ErrBodyReadAfterClose is returned when reading a Request Body 
	// after the body has been closed. This typically happens when the 
	// body is read after an HTTP Handler calls WriteHeader or Write on 
	// its ResponseWriter.
	ErrBodyReadAfterClose = os.NewError("http: invalid Read on closed request Body")

	// ErrHandlerTimeout is returned on ResponseWriter Write calls in 
	// handlers which have timed out.
	ErrHandlerTimeout = os.NewError("http: Handler timeout")

	// ErrMissingFile is returned by FormFile when the provided file 
	// field name is either not present in the request or not a file 
	// field.
	ErrMissingFile = os.NewError("http: no such file")

	ErrNoCookie = os.NewError("http: named cookied not present")
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
	
func main() {
	p := Person.UpdateField("Name", String8{"Derek"})
	log.Print(Person)
	log.Print(p)

	post := Post.UpdateTable("Author", p)
	a,b,c := post.Insert()

	log.Print(a)
	log.Print(b)
	log.Print(c)

	addr := flag.String("addr", ":8080", "http service address")
    flag.Parse()

	rtr := Router(
		new(Homepage),
		new(Homepage),
		new(Homepage),
		)

    //http.Handle("/", http.HandlerFunc(QR))
    //http.Handle("/Homepage", new(Homepage).Route() )
	
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.String())
	}
}



