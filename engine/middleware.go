package engine

import (
	"encoding/json"
	"net/http"
	"strings"
)

// AuthenticationHandler middleware to check for the jwt header for every request
func AuthenticationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMap := make(map[string]interface{})

		if IsDebug() {
			next.ServeHTTP(w, r)
			return
		}

		if pathStartsWith(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		userID, ok := ctx.Value(ContextAuth).(int)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			errorMap["code"] = http.StatusBadRequest
			errorMap["message"] = "no user id header provided"
			json.NewEncoder(w).Encode(errorMap)
			return
		}

		// The key and name env vars must be set
		if !hasKey() || !hasName() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			errorMap["code"] = http.StatusServiceUnavailable
			errorMap["message"] = "application has not been setup properly"
			json.NewEncoder(w).Encode(errorMap)
			return
		}

		// Grabs the token from the request header and verifies that it has the jwt token
		token := r.Header.Get("Authorization")
		if token == "" || len(strings.Split(token, " ")) < 2 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNetworkAuthenticationRequired)
			errorMap["code"] = http.StatusNetworkAuthenticationRequired
			errorMap["message"] = "authorization header is required"
			json.NewEncoder(w).Encode(errorMap)
			return
		}

		// Verifies that the token is valid and returns us the app_name
		verify := VerifyJwt(userID, strings.Split(token, " ")[1])
		if verify["is_valid"] == false {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			errorMap["code"] = http.StatusUnauthorized
			errorMap["message"] = "unauthorized request"
			json.NewEncoder(w).Encode(errorMap)
			return
		}

		// Runs the handler function and sets a application-token for the response header
		next.ServeHTTP(w, r)
	})
}

func pathStartsWith(url string) bool {
	// Opens the file and gets all its contents
	paths := []string{"/login"}

	// Iterates through the file checking if the url starts with any of the paths
	// in the file
	for _, path := range paths {
		if path == "" {
			continue
		}
		if strings.HasPrefix(url, path) {
			return true
		}
	}
	return false
}
