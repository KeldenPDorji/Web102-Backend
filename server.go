package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // Define routes and their corresponding handlers
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/login", loginHandler) // Handles the /login route

    // Serve static files from the "static" directory
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Start the server
    fmt.Println("Server listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the about page")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        // Parse the form data
        r.ParseForm()
        username := r.Form.Get("username")
        password := r.Form.Get("password")

        // Check if the username and password match the expected values
        if username == "admin" && password == "root" {
            fmt.Fprintf(w, "successful")
        } else {
            fmt.Fprintf(w, "Login failed")
        }
    } else {
        // Serve the login.html file from the static directory
        http.ServeFile(w, r, "static/login.html")
    }
}
// rock86