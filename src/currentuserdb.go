package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Current User Struct
type CurrentUser struct {
	Id int
    Usertype string
}

//Insert Current User
func InsertCurrentUser(db *sql.DB, ID int, UT string) {
    query := fmt.Sprintf("INSERT INTO CurrentUser VALUES (%d, '%s')",
        ID, UT)    

    _, err := db.Query(query)  

    if err != nil {
        panic(err.Error())
    }
}

//Delete Current User
func DeleteCurrentUser(db *sql.DB, ID int) {
    query := fmt.Sprintf(
        "DELETE FROM CurrentUser WHERE ID='%d'", ID)
    _, err := db.Query(query)   
    if err != nil {
        panic(err.Error())
    }
}

//Retrieve Current User ID
func GetCurrentUserID(db *sql.DB)(ID int) {   
    ID  = 0
    results, err := db.Query("Select Id FROM currentuser_db.CurrentUser") 

    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
        //Store ID into user variable
        var user CurrentUser
        err = results.Scan(&user.Id)
        if err != nil {
            panic(err.Error()) 
        } else{
            ID = user.Id
            //fmt.Println(ID);
			return ID;
		}    
          

    }
    return ID
}

//Retrieve Current User Type
func GetCurrentUserType(db *sql.DB)(UT string) {   
    UT  = ""
    results, err := db.Query("Select Usertype FROM currentuser_db.CurrentUser") 

    if err != nil {
        panic(err.Error())
    }

    for results.Next() {
        //Store ID into user variable
        var user CurrentUser
        err = results.Scan(&user.Usertype)
        if err != nil {
            panic(err.Error()) 
        } else{
            UT = user.Usertype
			return UT;
		}    
          

    }
    return UT
}

