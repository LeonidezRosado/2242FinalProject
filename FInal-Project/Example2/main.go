// Read/Write
package main

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
)

var (
	ErrValueTooLong = errors.New("cookie value is too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

func main() {
	// Start a web server with the two endpoints.get and set cookie handler.
	mux := http.NewServeMux()
	mux.HandleFunc("/set", setCookieHandler)
	mux.HandleFunc("/get", getCookieHandler)

	log.Print("Listening...")
	err := http.ListenAndServe(":4040", mux)
	if err != nil {
		log.Fatal(err)
	}
}

// creating setCookieHandler to set the cookies
func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize two new cookies with different values and attributes.
	//cookie1 set to message value
	cookie1 := http.Cookie{
		Name:     "cookieOne",
		Value:    "Hello, this is cookie number one! 果",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	//cookie2 set to message value
	cookie2 := http.Cookie{
		Name:     "cookieTwo",
		Value:    "Hi, this is cookie number two! 果",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	// Write the cookies to the response.
	if err := Write(w, cookie1); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := Write(w, cookie2); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Write a HTTP response as normal.
	w.Write([]byte("cookies set!"))
}

// creating getCookieHandler to get the cookies
func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve all cookies from the request using r.Cookies().
	cookies := r.Cookies()
	// Check if any cookies were found.
	if len(cookies) == 0 {
		http.Error(w, "No cookies found", http.StatusBadRequest)
		return
	}
	// Loop through the cookies and echo out their values in the response body.
	for _, cookie := range cookies {
		value, err := Read(cookie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte(cookie.Name + ": " + value + "\n"))
	}
}
func Write(w http.ResponseWriter, cookie http.Cookie) error {
	// Ensure the cookie value is not too long.
	if len(cookie.Value) > 4096 {
		return ErrValueTooLong
	}
	// Encode the cookie value in base64 format.
	encodedValue := base64.StdEncoding.EncodeToString([]byte(cookie.Value))
	// Set the encoded value to the cookie.
	cookie.Value = encodedValue
	// Write the cookie to the response.
	http.SetCookie(w, &cookie)
	return nil
}
func Read(cookie *http.Cookie) (string, error) {
	// Decode the cookie value from base64 format.
	decodedValue, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return "", ErrInvalidValue
	}
	return string(decodedValue), nil
}
