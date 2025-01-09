package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/andrew-lawlor/cc-simple-auth-server/token"
	"github.com/andrew-lawlor/cc-simple-auth-server/user"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Replace all println statements with calls to logging module (once it's ready).

func Login(w http.ResponseWriter, r *http.Request) {
	if !bearerAuth(w, r) {
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 148576)
	u, err := parseInput(r)
	if err != nil {
		msg := err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	user, err := user.GetUser(u.UserName)
	if err != nil {
		http.Error(w, "That user doesn't exist.", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		http.Error(w, "Incorrect password.", http.StatusUnauthorized)
		return
	}
	fmt.Println("User logged in: " + user.UserName)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if !bearerAuth(w, r) {
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 148576)
	u, err := parseInput(r)
	if err != nil {
		msg := err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	fmt.Println(u)
	hashedPW := HashPW(u.Password)
	w.Header().Set("Content-Type", "application/json")
	// Send success resp.
	if user.NewUser(u.UserName, u.DisplayName, hashedPW, u.Email) {
		w.WriteHeader(http.StatusOK)
		usr, _ := user.GetUser(u.UserName)
		// Clear password to send over the wire.
		usr.Password = ""
		json.NewEncoder(w).Encode(usr)
		return
	} else { // Send failed response.
		http.Error(w, "User Registration Failed.", http.StatusInternalServerError)
		return
	}
}

func HashPW(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashed)
}

func parseInput(r *http.Request) (user.User, error) {
	var u user.User
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&u)
	return u, err
}

func bearerAuth(w http.ResponseWriter, r *http.Request) bool {
	authorization := r.Header.Get("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorization, "Bearer", "", 1))
	if token.IsTokenValid(idToken) {
		return true
	} else {
		http.Error(w, "Invalid token.", http.StatusUnauthorized)
		return false
	}
}
