package middleware

import (
	"database/sql"
	"fmt"
	"go-lambda-postgres/models"
	"log"

	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {

	db, err := sql.Open("postgres", "Place your DB connection string here")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

func InsertUser(user models.User) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO users (firstname, lastname, createdtime, modifiedtime) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, user.FirstName, user.LastName, user.CreatedTime, user.ModifiedTime).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

func GetUser(id int64) (models.User, error) {

	db := createConnection()

	defer db.Close()

	var user models.User

	sqlStatement := `SELECT * FROM users WHERE id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CreatedTime, &user.ModifiedTime)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, err
}

func GetAllUsers() ([]models.User, error) {

	db := createConnection()

	defer db.Close()

	var users []models.User

	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.CreatedTime, &user.ModifiedTime)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		users = append(users, user)

	}

	return users, err
}

func UpdateUser(id int64, user models.User) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE users SET firstname=$2, lastname=$3, modifiedtime=$4 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, user.FirstName, user.LastName, user.ModifiedTime)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteUser(id int64) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM users WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
