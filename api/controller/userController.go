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
)

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
		generateToken(w, user.UserID, user.Username, user.UserType)
		sendSuccessResponse(w, "Login Success", nil)
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

	//Read From Request Body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed")
		return
	}
	//Ambil password lama dan baru dari form
	oldpassword := r.Form.Get("old_password")
	newpassword := r.Form.Get("new_password")

	//User id ambil pakai cookie
	userid := getUserIdFromCookie(r)

	//Password lama user dari database untuk dibandingkan
	var password = getUserPassword(userid)

	//Hash password lama yang diinput untuk verifikasi
	hash := sha256.Sum256([]byte(oldpassword))
	passwordHash := hex.EncodeToString(hash[:])

	if password == passwordHash {
		//hash password baru
		hash2 := sha256.Sum256([]byte(newpassword))
		passwordHash2 := hex.EncodeToString(hash2[:])
		query := "UPDATE user SET password = '" + passwordHash2 + "' WHERE userId = " + strconv.Itoa(userid)
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
