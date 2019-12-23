package main

import (
    "html/template"
    "net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type EmployeeDetails struct {
    EmpId   string
    Name string
}

func main() {

    http.HandleFunc("/", homepage)
	http.HandleFunc("/onsubmit", onsubmit)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func homepage(w http.ResponseWriter, r *http.Request){
    details := EmployeeDetails{
            EmpId: "1",
            Name: "dummy"}

    t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  fmt.Println("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, details) //execute the template and pass it the EmployeeDetails struct to fill in the gaps
    if err != nil { // if there is an error
  	  fmt.Println("template executing error: ", err) //log it
  	}
}

func onsubmit(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()

        empdetails := EmployeeDetails{
            EmpId: r.FormValue("empid"),
            Name: r.FormValue("name")}
		
		
		t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
		if err != nil { // if there is an error
		  fmt.Println("template parsing error: ", err) // log it
		}

		err = t.Execute(w, empdetails) //execute the template and pass it the details struct to fill in the gaps
		if err != nil { // if there is an error
		  fmt.Println("template executing error: ", err) //log it
		}
			
		// Configure the database connection 
		db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/testdb")

		if err != nil {
        fmt.Println(err)
		}
		if err := db.Ping(); err != nil {
        fmt.Println(err)
		}
		
				
		// Switch to created db
		fmt.Println("******Connect to db")
		if _, err := db.Exec(`USE testdb`); err != nil {
			fmt.Println(err)
		}
				
		// Create a new table
		fmt.Println("******Create table")
		query := `
			CREATE TABLE IF NOT EXISTS employees (
				id INT AUTO_INCREMENT,
				empid TEXT NOT NULL,
				name TEXT NOT NULL,
				PRIMARY KEY (id)
			);`

		if _, err := db.Exec(query); err != nil {
			fmt.Println(err)
		}
				
		// Insert a new employee
		fmt.Println("******Insert into table")
		insertStmt, _ := db.Prepare("INSERT employees SET empid=?, name=?")
		
		_, insertErr := insertStmt.Exec(empdetails.EmpId,empdetails.Name)
		if insertErr != nil {
			fmt.Println(insertErr)
		}
		        
    }