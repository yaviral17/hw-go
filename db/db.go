package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/yaviral17/hw-go/auth"
	"github.com/yaviral17/hw-go/models"
	"github.com/yaviral17/hw-go/myLogs"
)

var DB *sql.DB

func InitDB(dataSourceName string) error {
	myLogs.MyInfoLog("Connecting to database...")
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		myLogs.MyErrorLog(fmt.Sprintf("Error opening database: %q", err))
		return err
	}

	// Verify the connection
	err = DB.Ping()
	if err != nil {
		myLogs.MyErrorLog(fmt.Sprintf("Error verifying connection with database: %q", err))
		return err
	}

	myLogs.MySuccessLog("Connected to database successfully 🚀")
	return nil
}

func GetDB() *sql.DB {
	return DB
}

func CreateUser(ctx context.Context, user models.UserRegister) (uuid.UUID, error) {
	log.Println("Creating user...with password: ", user.PasswordHash)
	saltedPassword := auth.SaltPassword(user.PasswordHash)
	hashedPassword, er := auth.Encrypt(saltedPassword)

	if er != nil {
		myLogs.MyErrorLog("Error hashing password")
		return uuid.Nil, er
	}

	query := `
	INSERT INTO "user" (
		first_name, last_name, dob, mobile, email, username, password_hash, gender, links, bio
	) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
	RETURNING id`

	// Generate a UUID for the user ID
	var userID uuid.UUID

	err := DB.QueryRow(query,
		user.FirstName,
		user.LastName,
		user.DOB,
		user.Mobile,
		user.Email,
		user.Username,
		hashedPassword,
		user.Gender,
		pq.Array(user.Links),
		user.Bio,
	).Scan(&userID)

	if err != nil {
		log.Println("Error creating user: ", err)
		return uuid.Nil, err
	}

	return userID, nil
}

func GetUserByLoginCredentials(ctx context.Context, user models.UserLogin) (models.User, error) {

	isEmail := strings.ContainsRune(user.Username, '@')

	saltedPassword := auth.SaltPassword(user.Password)
	hashedPassword, er := auth.Encrypt(saltedPassword)

	if er != nil {
		myLogs.MyErrorLog("Error hashing password")
		return models.User{}, er
	}
	var query string
	if isEmail {
		query = `SELECT * FROM "user" WHERE email = $1 AND password_hash = $2`
	} else {
		query = `SELECT * FROM "user" WHERE username = $1 AND password_hash = $2`
	}

	var u models.User
	err := DB.QueryRowContext(ctx, query, user.Username, hashedPassword).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.DOB,
		&u.Mobile,
		&u.Email,
		&u.Username,
		&u.PasswordHashed,
		&u.WorkUploaded,
		&u.WorkDone,
		&u.WorkScore,
		&u.TotalWorkScore,
		&u.Bio,
		&u.ProfilePicture,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Gender,
		pq.Array(&u.Links),
	)

	if err != nil {
		log.Println("Error getting user by login credentials: ", err)
		return models.User{}, err
	}

	return u, nil
}
