package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Trips Struct
type Trips struct{
	TripID int
	Pickuplocation string
	Pickupcode int
	Dropofflocation string
	Dropoffcode int
	DriverID int
	Drivername string
	PassengerID int
	TripStatus string
	TripCost float64
	TripStart string
	TripEnd string
}

//Input New Trip Data into Trip DB
func InsertTrip(db *sql.DB, ID int, PL string, PU int, DL string, 
	DO int, DID int, DN string, PID int, TS string, TC float64, TST string, TE string) {
    query := fmt.Sprintf("INSERT INTO Trip VALUES (%d, '%s', %d, '%s', %d, %d, '%s', %d,'%s', %f, '%s', '%s')",
        ID, PL, PU, DL, DO, DID, DN, PID, TS, TC, TST, TE)    

    _, err := db.Query(query)  

    if err != nil {
        panic(err.Error())
    }
}

//Update Assigned Driver to Trip
func UpdateDriver(db *sql.DB, ID int, DID int, DN string) {
    query := fmt.Sprintf(
        "UPDATE Trip SET DriverID=%d, Drivername='%s' WHERE Id=%d", 
        DID, ID)
    _, err := db.Query(query)   
    if err != nil {
        panic(err.Error())
    }
}

//Update Trip Status
func UpdateTripStatus(db *sql.DB, ID int, TS string) {
    query := fmt.Sprintf(
        "UPDATE Trip SET TripStatus='%s' WHERE Id=%d", 
        TS, ID)
    _, err := db.Query(query)   
    if err != nil {
        panic(err.Error())
    }
}

//Update Cost of Trip
func UpdateTripCost(db *sql.DB, ID int, TC float64) {
    query := fmt.Sprintf(
        "UPDATE Trip SET TripCost=%f WHERE Id=%d", 
        TC, ID)
    _, err := db.Query(query)   
    if err != nil {
        panic(err.Error())
    }
}

//Delete Trip Information
func DeleteTrip(db *sql.DB, ID int) {
    query := fmt.Sprintf(
        "DELETE FROM Trip WHERE ID=%d", ID)
    _, err := db.Query(query)   
    if err != nil {
        panic(err.Error())
    }
}

//Retrieve Trips of Current User
func GetUserTrips(db *sql.DB, PID int)(tripArray []Trips) {  
    results, err := db.Query("SELECT * FROM trip_db.Trip WHERE PassengerID = ?", PID) 
	
    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
        // map this type to the record in the table
        var trip Trips = Trips{0, "", 0, "", 0, 0, "", 0, "", 0.0, "", ""}
        err = results.Scan(&trip.TripID, &trip.Pickuplocation, &trip.Pickupcode, &trip.Dropofflocation,
                           &trip.Dropoffcode, &trip.DriverID, &trip.Drivername, &trip.PassengerID, 
						   &trip.TripStatus, &trip.TripCost, &trip.TripStart, &trip.TripEnd)
						   
        if err != nil {
            panic(err.Error()) 
        } else{

			tripArray = append(tripArray, trip)
		}    
          
    }

	return tripArray
}

