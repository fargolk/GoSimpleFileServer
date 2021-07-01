package main

import (
    "fmt"
    "log"
    "net/http"
	"io/ioutil"
)

func uploadMusic(writer http.ResponseWriter, req *http.Request) {
    
	fmt.Println("Upload Music Requested ...")

    req.ParseMultipartForm(10 << 20)
   
    file, handler, err := req.FormFile("musicFile")

    if err != nil {
        fmt.Println("Error Getting Music File")
        fmt.Println(err)
        return
    }
    defer file.Close()
	
    fmt.Printf("Music File: %+v\n", handler.Filename)
    fmt.Printf("Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
	
	
	werr := ioutil.WriteFile("./media/music/" + handler.Filename, fileBytes, 0644)
	if werr != nil {
		fmt.Println(werr)
		return
	}

    fmt.Fprintf(writer, "Successfully Uploaded File\n")
}


func uploadMovie(writer http.ResponseWriter, req *http.Request) {
    
	fmt.Println("Upload Movie Requested ...")

    req.ParseMultipartForm(10 << 20)
    
    file, handler, err := req.FormFile("movieFile")
    if err != nil {
        fmt.Println("Error Getting Movie File")
        fmt.Println(err)
        return
    }
	
    defer file.Close()
    fmt.Printf("Movie File: %+v\n", handler.Filename)
    fmt.Printf("Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
	
	werr := ioutil.WriteFile("./media/movie/" + handler.Filename, fileBytes, 0644)
	if werr != nil {
		fmt.Println(werr)
		return
	}

    fmt.Fprintf(writer, "Successfully Uploaded File\n")
}



func startServer() {
	
	fileServer := http.FileServer(http.Dir("./media"))
	staticFileServer := http.FileServer(http.Dir("./static"))
	
	http.Handle("/media/", http.StripPrefix("/media", fileServer))
	http.Handle("/", staticFileServer)
	http.HandleFunc("/uploadMusic", uploadMusic)
	http.HandleFunc("/uploadMovie", uploadMovie)

	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func main() {
    fmt.Printf("Starting server at port 8080\n")
	startServer()
}