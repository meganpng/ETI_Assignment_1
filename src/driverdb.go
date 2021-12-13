//Create User Account - Microservices

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Drivers Struct
type Drivers struct {
	Id int
	Firstname string
	Lastname string
	Emailaddress string
	Mobileno int
	Carlicenseno string
}

//Creating New Driver into Database
func InsertDriver(db *sql.DB, ID int, FN string, LN string, EM string, MNO int, CARNO string) {
    query := fmt.Sprintf("INSERT INTO Drivers VALUES (%d, '%s', '%s', '%s', %d, '%s')",
        ID, FN, LN, EM, MNO, CARNO)    

    _, err := db.Query(query)  

    if err != nil {
        panic(err.Error())
    }
}

//Update Driver 
func EditDriver(db *sql.DB, ID int, FN string, LN string, EM string, MNO int, CARNO string) {
    query := fmt.Sprintf(
        "UPDATE Drivers SET Firstname='%s', Lastname='%s', Emailaddress='%s', Mobileno=%d, Carlicenseno='%s' WHERE Id=%d", 
        FN, LN, EM, MNO, CARNO, ID)
    _, err := db.Query(query)   
    if err != nil {
        panic(err.Error())
    }
}

//Get Largest ID of Passenger
func GetMaxDriverID(db *sql.DB)(int) {   
    var emptyid int = 0
    results, err := db.Query("Select MAX(ID) FROM driver_db.Drivers") 

    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
        // map this type to the record in the table
        var id int
        err = results.Scan(&id)
        if err != nil {
            panic(err.Error()) 
        } else{
			return id
		}    
          

    }
    return emptyid
}

//Create a Driver ID for New Driver
func createDriverID()(int){
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
		
	// handle error
	if err != nil {
		panic(err.Error())
	}
    var id int = 1
    var pid int = 0
    pid = GetMaxDriverID(db)
    if pid == 0{
        id = 1
        
        return id
    } else{
        id = pid + 1
        
        return id
    }


}