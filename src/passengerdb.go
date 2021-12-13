//Create User Account - Microservices
//Login is not required

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//"encoding/json"

//Passengers struct
type Passengers struct{
    Id int
	Firstname string
	Lastname string
	Emailaddress string
	Mobileno int
}

//Create New Passenger
func InsertPassenger(db *sql.DB, ID int, FN string, LN string, EM string, MNO int) {
    query := fmt.Sprintf("INSERT INTO Passengers VALUES (%d, '%s', '%s', '%s', %d)", ID, FN, LN, EM, MNO)    

    _, err := db.Query(query)  

    if err != nil {
        panic(err.Error())
    }
}

//Update Passenger Details
func EditPassenger(db *sql.DB, ID int, FN string, LN string, EM string, MNO int) {
    query := fmt.Sprintf(
        "UPDATE Passengers SET Firstname='%s', Lastname='%s', Emailaddress='%s', Mobileno=%d WHERE Id=%d", 
        FN, LN, EM, MNO, ID)
    _, err := db.Query(query)   
    if err != nil {
        panic(err.Error())
    }
}

//Retrieve all Passengers
func GetPassengers(db *sql.DB) {   
    results, err := db.Query("Select * FROM passenger_db.Passengers") 

    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
        // map this type to the record in the table
        var passenger Passengers
        err = results.Scan(&passenger.Id, &passenger.Firstname, 
                           &passenger.Lastname, &passenger.Emailaddress, &passenger.Mobileno)
        if err != nil {
            panic(err.Error()) 
        } else{
			fmt.Println(passenger.Id, passenger.Firstname, 
				passenger.Lastname, passenger.Emailaddress, passenger.Mobileno)
		}    
          

    }
}

//Get Largest ID of Passenger
func GetMaxPassengerID(db *sql.DB)(int) {   
    var emptyid int = 0
    results, err := db.Query("Select MAX(ID) FROM passenger_db.Passengers") 

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



//Create a Passenger ID for New Passenger
func createPassengerID()(int){
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")
		
	// handle error
	if err != nil {
		panic(err.Error())
	}
    var id int = 1
    var pid int = 0
    pid = GetMaxPassengerID(db)
    if pid == 0{
        id = 1
        
        return id
    } else{
        id = pid + 1
        
        return id
    }


}

