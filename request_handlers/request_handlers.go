package requesthandlers

import (
	"encoding/json"
	"fmt"
	//"fmt"
	"net/http"
	Config "pkg/config"
	Helper "pkg/globals/utility"
	Auth "pkg/globals/authentications"
	//"regexp"
	"strings"
	"unicode"
)

func Auth_Token(w http.ResponseWriter, r *http.Request) {
	authInfo := r.Header.Get("Authorization")
	info := strings.Split(authInfo, " ")

	if authInfo == "" || len(info) != 2 {
		http.Error(w, "Unauthorized /auth/token", http.StatusUnauthorized)
		return
	}

	base64String := info[1]
	auth, err := Helper.DecodeBase64(base64String)
	if err != nil {
		http.Error(w, "Bad Reuqest /auth/token", http.StatusBadRequest)
		return 
	}

	accInfo := strings.Split(auth, ":")
	if len(accInfo) != 2 {
		http.Error(w, "Bad Reuqest /auth/token", http.StatusBadRequest)
		return 
	}
	//username := accInfo[0]
	password := accInfo[1]

	if !isValidAPIAuth(password) || password != Config.PW_1 {
		http.Error(w, "Unauthorized /auth/token", http.StatusUnauthorized)
		return
	}

	fmt.Println("swswsw")
	res := make(map[string]interface{})
	if password == Config.PW_1 {
		res["user_id"] = Config.REQUEST
		res["key"] = Config.KEY
	} else {
		res["user_id"] = Config.REQUEST2
		res["key"] = Config.KEY2
	}

	json.NewEncoder(w).Encode(res)
}

func GenerateJWT(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("secret_key")
	id := r.FormValue("user_id")

	if key == "" || id == "" {
		http.Error(w, "Bad Reuqest /auth/token", http.StatusBadRequest)
		return
	}

	jwt, err := Auth.GenerateToken(id)
	if err != nil {
		http.Error(w, "Internal Server Error /auth/token", http.StatusInternalServerError)
		return
	}
	res := map[string]interface{} {
		"token": jwt,
	}

	json.NewEncoder(w).Encode(res)
}

func GetCustomerInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("JWT pass"))
}

func UploadImges(w http.ResponseWriter, r *http.Request) {

	filePath := r.FormValue("file_dir")

	if filePath == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	uploadPath := Config.FILES_DIR + "/" + filePath
	if err := Helper.CreateDirIfNotExist(uploadPath); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	Helper.UploadFiles(w, r, uploadPath, 30)
}

func isValidAPIAuth(input string) bool {
    // Check for minimum length of 12 characters
    if len(input) < 12 {
        return false
    }

    hasUppercase := false
    hasLowercase := false
    hasDigit := false
    hasSpecial := false

    for _, char := range input {
        switch {
        case unicode.IsUpper(char):
            hasUppercase = true
        case unicode.IsLower(char):
            hasLowercase = true
        case unicode.IsDigit(char):
            hasDigit = true
        case strings.ContainsRune("@$!%*?&", char):
            hasSpecial = true
        }
    }

    return hasUppercase && hasLowercase && hasDigit && hasSpecial
}
