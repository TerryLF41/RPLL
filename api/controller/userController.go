package controller

import (
	"RPLL/api/model"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM user"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var user model.User
	var userList []model.User
	for rows.Next() {
		if err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.ProfileDesc, &user.JoinDate,
			&user.UserType, &user.BanStatus, &user.ProfilePicture); err != nil {
			log.Println(err)
			return
		} else {
			userList = append(userList, user)
		}
	}

	// Kirim response ke client
	// Response dibuat dengan factory di responseHandler
	if len(userList) >= 1 {
		sendSuccessResponse(w, "Successfully retrieved userlist", userList)
	} else {
		sendErrorResponse(w, "Failed to retrieve userlist")
	}
}

// Login untuk user
func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Read dari request body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if email == "" || password == "" {
		sendErrorResponse(w, "There are some empty input")
		return
	}

	// Hash password
	hash := sha256.Sum256([]byte(password))
	passwordHash := hex.EncodeToString(hash[:])

	// Digunakan untuk mengecek apakah user di ban atau tidak
	// jika di ban, tidak bisa login
	var isbanned bool

	// Query berdasarkan pasangan email dan password yang sama dengan input user
	query := "SELECT userId,userName,email,profilePicture,profileDesc,userType,banstatus FROM user WHERE Email ='" + email + "' && Password='" + passwordHash + "'"
	var user model.User
	err1 := db.QueryRow(query).Scan(&user.UserID, &user.Username, &user.Email, &user.ProfilePicture, &user.ProfileDesc, &user.UserType, &isbanned)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			sendErrorResponse(w, "Wrong login credentials")
			return
		}
		log.Println(err1)
		sendErrorResponse(w, "Error")
		return
	} else {
		if isbanned {
			sendErrorResponse(w, "User is banned")
			return
		}
		userData := map[string]interface{}{
			"userId":         user.UserID,
			"username":       user.Username,
			"email":          user.Email,
			"profilePicture": user.ProfilePicture,
			"profileDesc":    user.ProfileDesc,
			"userType":       user.UserType,
		}

		token := generateToken(w, user.UserID, user.Username, user.UserType)
		sendSuccessLoginResponse(w, "Login Success", userData, token)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//User id ambil pakai cookie
	userid := getUserIdFromCookie(r)
	if userid == -1 {
		sendErrorResponse(w, "No login activity before")
	} else {
		resetUserToken(w)
		sendSuccessResponse(w, "Logout success", nil)
	}

}

// Fungsi untuk register user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	//Read From Request Body
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Failed")
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")

	if username == "" || email == "" || password == "" {
		sendErrorResponse(w, "There are some empty input")
		return
	}

	// Hash password yang diinput user dengan SHA256
	hash := sha256.Sum256([]byte(password))
	passwordHash := hex.EncodeToString(hash[:])

	userFactory := model.NewUserModelFactory()

	// Create a new user instance
	newUser := userFactory.CreateUser(
		0,
		username,
		passwordHash,
		email,
		"",
		time.Now(),
		0,
		false,
		"",
		nil,
	)

	// Query untuk cek apakah ada user yang terdaftar dengan email yang diinput
	// jika ada, maka tidak bisa register
	query := "SELECT userId FROM user WHERE email ='" + email + "'"
	err1 := db.QueryRow(query).Scan(&newUser.UserID)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			_, errQuery := db.Exec("INSERT INTO user(username, email, password) values(?,?,?)",
				newUser.Username,
				newUser.Email,
				newUser.Password,
			)
			if errQuery != nil {
				log.Println(errQuery)
				sendErrorResponse(w, "Register failed")
			} else {
				sendSuccessResponse(w, "Register success", nil)
			}
		}
	} else {
		sendErrorResponse(w, "The email is already registered")
	}
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	userId, err1 := strconv.Atoi(vars["userId"])
	if err1 != nil {
		http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
		return
	}
	//Read From Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	//Ambil password lama dan baru dari form
	oldpassword := r.Form.Get("old_password")
	newpassword := r.Form.Get("new_password")

	//Password lama user dari database untuk dibandingkan
	var password = getUserPassword(userId)

	//Hash password lama yang diinput untuk verifikasi
	hash := sha256.Sum256([]byte(oldpassword))
	passwordHash := hex.EncodeToString(hash[:])

	if password == passwordHash {
		//hash password baru
		hash2 := sha256.Sum256([]byte(newpassword))
		passwordHash2 := hex.EncodeToString(hash2[:])
		query := "UPDATE user SET password = '" + passwordHash2 + "' WHERE userId = " + strconv.Itoa(userId)
		_, errQuery := db.Exec(query)

		if errQuery != nil {
			sendErrorResponse(w, "Failed to change password!")
		} else {
			sendSuccessResponse(w, "Password successfully changed!", nil)
		}
	} else {
		sendErrorResponse(w, "Password does not match!")
	}
}

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	userId := vars["userId"]

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}

	userName := r.Form.Get("username")
	email := r.Form.Get("email")
	description := r.Form.Get("description")
	profilePicture := r.Form.Get("profilePicture")

	sqlStatement := `
		UPDATE user 
		SET username = ?, email = ?, profileDesc = ?, profilePicture =?

		WHERE userId = ?`

	_, errQuery := db.Exec(sqlStatement,
		userName,
		email,
		description,
		profilePicture,
		userId,
	)
	if errQuery == nil {
		sendSuccessResponse(w, "Successfully updated user information", nil)
	} else {
		sendErrorResponse(w, "Failed to update user information")
	}
}

func BanUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	vars := mux.Vars(r)
	userId := vars["userId"]
	banStatus, err := strconv.Atoi(r.Form.Get("banStatus"))
	if err != nil {
		http.Error(w, "Unable to convert string to int", http.StatusInternalServerError)
		return
	}
	var sqlStatement = `
		UPDATE user 
		SET banStatus = ?
		WHERE userId = ?`

	_, errQuery := db.Exec(sqlStatement,
		banStatus,
		userId,
	)

	if errQuery == nil {
		sendSuccessResponse(w, "Successfully Banned/Unbanned the user", nil)
	} else {
		sendErrorResponse(w, "Failed to Ban/Unban the user")
	}
}

func getUserPassword(id int) string {
	db := connect()
	defer db.Close()

	var password string
	query := "SELECT password FROM user WHERE userId = " + strconv.Itoa(id)
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return ""
	}

	for rows.Next() {
		if err := rows.Scan(&password); err != nil {
			log.Println(err)
			return ""
		}
	}
	return password
}
