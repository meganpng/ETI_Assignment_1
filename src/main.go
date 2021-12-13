//Create User Account - Microservices
//Login is not required

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func Login(ID int, UT string){
	// Opening Current User DB
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/currentuser_db")
		
	// handle error
	if err != nil {
		panic(err.Error())
	} else{
		//Inserting User ID into Current User DB
		InsertCurrentUser(db, ID, UT)
		fmt.Println("User is Logged In")
	}
	defer db.Close() 

}



//Function for Signing Up As a Passenger
func SignUpAsPassenger(){

	fmt.Println("Please enter your first name: ")
	var fn string;
	fmt.Scanf("%s\n", &fn)

	fmt.Println("Please enter your last name: ")
	var ln string;
	fmt.Scanf("%s\n", &ln)

	fmt.Println("Please enter your email: ")
	var e string;
	fmt.Scanf("%s\n", &e)

	fmt.Println("Please enter your mobile number: ")
	var mno int;
	fmt.Scanf("%d\n", &mno)

	var id int = createPassengerID()
	var utype string = "Passenger"
	// Opening Passenger DB
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")
			
	// handle error
	if err != nil {
		panic(err.Error())
	} else{
		
		InsertPassenger(db, id, fn, ln, e, mno)
		//Logging In User
		Login(id, utype)
		fmt.Println("User is Signed Up")
		
	}
	defer db.Close() 

	
}

//Function for Signing Up As a Driver
func SignUpAsDriver(){

	fmt.Println("Please enter your first name: ")
	var fn string;
	fmt.Scanf("%s\n", &fn)

	fmt.Println("Please enter your last name: ")
	var ln string;
	fmt.Scanf("%s\n", &ln)

	fmt.Println("Please enter your email: ")
	var e string;
	fmt.Scanf("%s\n", &e)

	fmt.Println("Please enter your mobile number: ")
	var mno int;
	fmt.Scanf("%d\n", &mno)

	fmt.Println("Please enter your car license number: ")
	var cno string
	fmt.Scanf("%s\n", &cno)

	var id int = createDriverID()
	var utype string = "Driver";
	// Opening Passenger DB
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/driver_db")
			
	// handle error
	if err != nil {
		panic(err.Error())
	} else{
		
		InsertDriver(db, id, fn, ln, e, mno, cno)
		//Logging In User
		Login(id, utype)
		fmt.Println("User is Signed Up")
	}
	defer db.Close() 

	
}

//Logging Out Function
func LogOut(ID int){
		//Opening Current User DB
		db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/currentuser_db")
		
		// handle error
		if err != nil {
			panic(err.Error())
		} else{
			DeleteCurrentUser(db, ID)
			fmt.Println("User is Logged Out")
		}
		defer db.Close() 
	
}

func retrieveCurrentID()(ID int){
	ID = 0
	//Opening Current User DB
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/currentuser_db")
	
	// handle error
	if err != nil {
		panic(err.Error())
	} 

	//Getting the ID
	ID = GetCurrentUserID(db)
	return ID

}

func retrieveCurrentType()(UT string){
	UT = ""
	//Opening Current User DB
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/currentuser_db")
	
	// handle error
	if err != nil {
		panic(err.Error())
	} 

	//Getting the ID
	UT = GetCurrentUserType(db)
	return UT

}

func retrievePastTrips(ID int)(tArray []Trips){
	//Opening Trip DB
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")
	
	// handle error
	if err != nil {
		
		panic(err.Error())
	} 
	
	//Getting Slice containing Trips belonging to Current User
	tArray = GetUserTrips(db, ID)
	return tArray
}

func passengerOptions(){
	for{
		fmt.Println("Please select an option: ")
		//This allows users to book a trip
		fmt.Println("1. Book a Trip")
		//This allows drivers to view past trips
		fmt.Println("2. View Past Trips")
		//This allows the user to view their ratings given by drivers
		fmt.Println("3. View Ratings")
		fmt.Println("0. Log Out")
		var input int;
		fmt.Scanf("%d\n", &input)
		//fmt.Println(input)

		//Check for Input
		if (input == 1){
			fmt.Println("This feature is currently unavailable, please select an option again.")
			continue

		//View Past Trips Option
		}else if input == 2{
			var id int = retrieveCurrentID()
			var tripArray []Trips = retrievePastTrips(id)

			if len(tripArray) > 0{

				fmt.Println("User's Past Trips:")
				//For Loop to print trips
				for _, tripobj := range tripArray{
					coststring := fmt.Sprintf("%.2f", tripobj.TripCost)
	
					fmt.Println("Trip ID: " + fmt.Sprintf("%d",tripobj.TripID) + "\n" + "Pick Up Location: " + tripobj.Pickuplocation + 
					"\n" + "Pick Up Postal Code: " + fmt.Sprintf("%d", tripobj.Pickupcode) + "\n" + "Drop Off Location: " + 
					tripobj.Dropofflocation + "\n" + "Drop Off Postal Code: " + fmt.Sprintf("%d",tripobj.Dropoffcode) + 
					" " + "\n" + "Driver's ID: " + fmt.Sprintf("%d",tripobj.DriverID) + "\n" + "Driver's Name: " + tripobj.Drivername +
					" " + "\n" + "Trip Status: " + tripobj.TripStatus + "\n" + "Cost of Trip: $" + coststring  + "\n" +
					"Trip Started At: " + tripobj.TripStart + "\n"  + "Trip Ended At: " + tripobj.TripEnd + "\n")
				}
				
				continue

			}else{
				fmt.Println("User has not had any previous trips.")
				continue
			}

		}else if (input == 3){
			fmt.Println("This feature is currently unavailable, please select an option again.")
			continue
		}else if (input == 0){
			var id int = retrieveCurrentID()
			LogOut(id)
			break

		} else{
			fmt.Println("Invalid input, please select an option again.")
		}

	}
}

func driverOptions(){
	for{
		fmt.Println("You do not have any bookings.")
		fmt.Println("Please select an option: ")
		//View Demand Map allows Drivers to see the demand for drivers in Singapore
		//It uses coloured areas like a heat map to show various levels of demand in different parts of Singapore
		fmt.Println("1. View Demand Map")
		//This allows drivers to view their ratings by passengers
		fmt.Println("2. View Ratings")
		fmt.Println("0. Log Out")
		var input int;
		fmt.Scanf("%d\n", &input)
		//fmt.Println(input)

		//Check for Input
		if (input == 1){
			fmt.Println("This feature is currently unavailable, please select an option again.")
			continue

		//View Past Trips Option
		}else if input == 2{
			
			fmt.Println("This feature is currently unavailable, please select an option again.")
				continue

		}else if (input == 0){
			var id int = retrieveCurrentID()
			LogOut(id)
			break

		} else{
			fmt.Println("Invalid input, please select an option again.")
		}

	}
}

func logInOption()(choice int){
	choice = 0
	for{
		fmt.Println("Log in either as a passenger or a driver.")
		fmt.Println("1. Passenger")
		fmt.Println("2. Driver")
		fmt.Println("0. Go Back")
		var input int;
		fmt.Scanf("%d\n", &input)
		var id int = 1;
		//Log In as Passenger
		if (input == 1){
			var utype string = "Passenger"
			Login(id, utype)
			fmt.Println("Logged in as a Passenger")
			choice = 1
			return choice

		//Log In as Driver
		}else if input == 2{
			var utype string = "Driver"
			Login(id, utype)
			fmt.Println("Logged in as a Driver")
			choice = 2
			return choice

		}else if (input == 0){
			fmt.Println("Going Back...")
			break

		} else{
			fmt.Println("Invalid input, please select an option again.")
		}

	}
	return choice
}

func signUpOption()(choice int){
	choice = 0
	for{
		fmt.Println("Sign Up either as a passenger or a driver.")
		fmt.Println("1. Passenger")
		fmt.Println("2. Driver")
		fmt.Println("0. Go Back")
		var input int;
		fmt.Scanf("%d\n", &input)
		//fmt.Println(input)

		//Check for Input
		if (input == 1){
			choice = 1 
			SignUpAsPassenger()
			return choice

		//View Past Trips Option
		}else if input == 2{
			choice = 2
			SignUpAsDriver()
			return choice

		}else if (input == 0){
			fmt.Println("Going back...")
			break

		} else{
			fmt.Println("Invalid input, please select an option again.")
		}

	}
	return choice
}

//Adding Trips
func addTrips(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/trip_db")

	// handle error
	if err != nil {
		panic(err.Error())
	} 
	
	InsertTrip(db, 1, "391A Orchard Road", 238873, 
	"83 Punggol Central", 828761, 1, "John Tan", 1, "Complete", 19.00,
	 "12 December 2021, 09:45", "12 December 2021, 10:11")
	InsertTrip(db, 2, "83 Punggol Central", 828761, 
	"9 Bishan Pl", 579837, 1, "Tai Yang Lee", 1, "Complete", 17.00,
	 "12 December 2021, 012:14", "12 December 2021, 12:31")
	 InsertTrip(db, 3, "9 Bishan Pl", 579837, 
	"204 Hougang Street 21", 530204, 3, "Jaiden Chen", 3, "Complete", 15.00,
	 "12 December 2021, 14:53", "12 December 2021, 15:05")
	 defer db.Close()
}

//Adding Passenger
func addPassenger(){
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/passenger_db")

	// handle error
	if err != nil {
		panic(err.Error())
	} 
	
	InsertPassenger(db, 1, "Jane", "Tan", "jtan@gmail.com", 81234567)
	defer db.Close()
}


//Main Function
func main() {

	for{
		fmt.Println("Please login or sign up if you do not have an account: ")
		fmt.Println("1. Login")
		fmt.Println("2. Sign Up")
		var input int;
		fmt.Scanf("%d\n", &input)
		if(input == 1){
			var logtype int =logInOption()
			if logtype == 1{
				passengerOptions()
				continue
			} else if logtype == 2{
				driverOptions()
				continue
			}
			
		} else if(input == 2){
			var signtype int = signUpOption()
			if signtype == 1{
				passengerOptions()
				continue
			} else if signtype == 2{
				driverOptions()
				continue
			}
			
		} else{
			fmt.Println("Invalid input, please select an option again.")
		}
	}

    
}