package main

import (
	"log"
	"net/http"
)

func main() {
	// Start a web server with the two endpoints.get and set cookie handler.
	mux := http.NewServeMux()
	mux.HandleFunc("/set", setCookieHandler)
	mux.HandleFunc("/get", getCookieHandler)

	log.Print("Listening...")
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}
//creating setCookieHandler to set the cookies
func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize two new cookies with different values and attributes.
    //cookie1 set to message value
	cookie1 := http.Cookie{
		Name:     "cookieOne",
		Value:    "Hello, this is cookie number one!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
    //cookie2 set to message value
	cookie2 := http.Cookie{
		Name:     "cookieTwo",
		Value:    "Hi, this is cookie number two!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// Use the http.SetCookie() function to send the cookies to the client.
	// Behind the scenes this adds a `Set-Cookie` header to the response
	// containing the necessary cookie data.
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)

	// Write a HTTP response as normal.
	w.Write([]byte("cookies set!"))
}
//creating getCookieHandler to get the cookies
func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve all cookies from the request using r.Cookies().
	cookies := r.Cookies()

	// Check if any cookies were found.
	if len(cookies) == 0 {
		log.Println("No cookies found")
		http.Error(w, "No cookies found", http.StatusBadRequest)
		return
	}

	// Loop through the cookies and echo out their values in the response body.
	for _, cookie := range cookies {
		w.Write([]byte(cookie.Name + ": " + cookie.Value + "\n"))
	}
}