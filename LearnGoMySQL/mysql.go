package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// References URL: https://github.com/go-sql-driver/mysql

// To install MySQL driver:
// $ go get -u github.com/go-sql-driver/mysql
// To solve inconsistent vendor dependencies/constraints:
// $ go mod vendor

// For this example we use MySQL via Docker
// $ docker pull mysql
// $ docker run --name mysql -p 3306:3306 -e MYSQL_DATABASE=mysql-db -e MYSQL_USER=mysql-user -e MYSQL_PASSWORD=p4ssw0rd -e MYSQL_ROOT_PASSWORD=r00t-p4ssw0rd -d mysql

// Create a table and data for example purposes
// $ docker exec -it mysql /bin/bash
// $ mysql mysql-db -uroot -pr00t-p4ssw0rd
// mysql> create table users (firstname varchar(30), lastname varchar(30), username varchar(30));
// mysql> insert into users (firstname, lastname, username) values ('Rob', 'Pike', 'rob');
// mysql> insert into users (firstname, lastname, username) values ('Ken', 'Thompson', 'ken');
// mysql> insert into users (firstname, lastname, username) values ('Robert', 'Griesemet', 'gri');
// mysql> select * from users;
// +-----------+-----------+----------+
// | firstname | lastname  | username |
// +-----------+-----------+----------+
// | Rob       | Pike      | rob      |
// | Ken       | Thompson  | ken      |
// | Robert    | Griesemet | gri      |
// +-----------+-----------+----------+

// If you use DBeaver to test connection:
// - Edit Driver Settings
// - Driver properties
//      allowPublicKeyRetrieval = TRUE

type user struct {
	FirstName string `alias:"firstname"`
	LastName  string `alias:"lastname"`
	Username  string `alias:"username"`
}

type Users []user

func main() {
	//db, err := sql.Open("mysql", "mysql-user:p4ssw0rd@tcp(localhost:3306)/mysql-db?parseTime=true")
	//db, err := sql.Open("mysql", "mysql-user:p4ssw0rd@/mysql-db")
	db, err := sql.Open("mysql", "root:r00t-p4ssw0rd@/mysql-db")
	if err != nil {
		panic(err)
	}

	users := getUsers(db)

	if users == nil {
		fmt.Println("No users found")
	} else {
		for _, myUser := range users {
			fmt.Println("> ", myUser.FirstName, myUser.LastName, "'", myUser.Username, "'")
		}
	}

	if !addUser(db, user{"Rafael", "Hernamperez", "rafinguer"}) {
		fmt.Println("User could not be added")
	}

	defer db.Close()
}

// This function retrieves all the users using SELECT * FROM users
func getUsers(db *sql.DB) Users {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("Could not get data from users table: " + string(err.Error()))
		return nil
	}

	cols, _ := rows.Columns()
	fmt.Println("Columns detected: ", cols)

	users := Users{}

	for rows.Next() {
		u := user{}
		err := rows.Scan(&u.FirstName, &u.LastName, &u.Username)
		if err != nil {
			log.Fatal("Error scanning row: ", err)
		}
		fmt.Println(u)
		users = append(users, u)
	}

	defer rows.Close()
	return users
}

// This function adds a new user to the table using INSERT
func addUser(db *sql.DB, u user) bool {
	res, err := db.Exec("INSERT INTO users (firstname, lastname, username) VALUES (?, ?, ?)", u.FirstName, u.LastName, u.Username)
	if err != nil {
		log.Fatal("Could not add user: ", err)
		return false
	}

	ra, _ := res.RowsAffected()
	fmt.Println("Rows affected: ", ra)

	return true
}
