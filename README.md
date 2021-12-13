# ETI_Assignment_1
<h1>
Design Considerations for Microservices
</h1>

</br>

<p>One design consideration is that there is no “Login” microservice in this program. When the user (passenger) chooses the option to “Login”, they will be immediately logged in after choosing whether they are a passenger or a driver without being prompted to enter their email address and password. The Current User database is to store the ID of the current user that is logged in so that the ID may be used for the “View Past Trips” microservice to be able to retrieve the user’s past trips if the user is a passenger For there to be a current ID stored when the user is logged in, the user ID “1” which is the ID of an existing user in the Passenger as well as the Driver database will be stored in the Current User database as a placeholder.  </p>

</br>

<p>Another design consideration is that this ride-sharing platform would usually be on mobile devices, therefore not requiring a web UI. Another design consideration is that due to the lack of “Book a Trip” microservice, the Trip database has been pre-filled with trips and if the user selects the “Book a Trip” feature, they will be told the feature is unavailable. There are three prefilled trips that belong to a passenger with the User ID: 1, so it will only be triggered when the user is logged in and the “View Past Trips” option is chosen as any created users will have a different user ID. If the “View Past Trips” option is chosen while the user has signed up, there would be a message printed stating that the user has not had any past trips so far.</p>

</br>

<p>
There is a Go file for each individual database, for the Passengers database file, they have a Passengers struct with the data variables of ID, first name, last name, email address and mobile number. In the file, there are functions like a function to Create a Passenger where new passenger data are inserted into the Passenger database, a function to Edit a Passenger where passenger data, first name, last name, email address and mobile number, will be used to update the details of a passenger which the Passenger ID belongs to. There is also a function to retrieve all the Passengers in the database and print them out. Other functions for the Passenger database include Get Max Passenger ID function and createPassengerID function. The Get Max Passenger ID function retrieves the highest amount of ID as ID is an integer and returns it. The createPassengerID function gets the highest amount of ID using the Get Max Passenger function and if the highest amount is 0, it returns the new ID as 1, while if the highest amount is something else, it takes the highest amount and adds 1 to it, returning it as the new ID, like adding 1 to 9 and returning 10.</p>

</br>

<p>For the Drivers database file, they have a Driver struct with the data variables of ID, first name, last name, email address, mobile number and the user’s car license number. In the file, the functions are InsertDriver, which creates a new Driver in the database, EditDriver which updates the details of the driver which the inputted ID belongs to, as well as the GetMaxDriverID and createDriverID functions which do the same as the GetMaxPassengerID and createPassengerID functions.</p>

</br>

<p>For the Trip database file, they have a Trip struct with the data variables, TripID, which is the ID of the trip, Pickuplocation, which is the address of the pick up location, Pickupcode, which is the postal code of the pick up location, Dropofflocation, which is the address of the drop off location, Dropoffcode, which is the postal code of the drop off location, DriverID, which is the ID of the assigned driver, Drivername, the name of the driver, PassengerID, which is the ID of the passenger who took the trip, TripStatus, which is the status of the trip, which can be “Booked”, “Ongoing”, “Complete” and “Canceled”, TripCost, which is the cost of the trip, TripStart, which is when the trip started and TripEnd, which is when the trip ended. The functions include InsertTrip, which creates a new trip in the database, UpdateDriver, which updates the assigned driver details, including the ID and the name, belonging to the trip with the inputted ID and UpdateTripStatus, which updates the trip status of a specific trip. Other functions include UpdateTripCost, which updates the trip cost of a specific trip, DeleteTrip, which deletes a trip belonging to the inputted ID and GetUserTrips, which retrieves the trips belonging to the current passenger logged in.</p>

</br>

<p>For the Current User database file, they have a Current User struct with the data variables, Id and Usertype, which is the type of user logged in, either Passenger or Driver. They have functions which include, InsertCurrentUser, which puts in the ID and User Type of the current user into the database, DeleteCurrentUser, which deletes the current user data from the database, GetCurrentUserID, which retrieves the ID of the current user from the database and GetCurrentUserType which retrieves the User Type of the current user from the database.</p>

</br>

<p>For the first microservice, “Create New Account”, it will ask the user to select whether they wish to create a new Passenger or Driver account. If the user selects “Passenger”, they will be asked to input their first name, last name, email and mobile number.  If the user selects “Driver”, they will be asked to input their first name, last name, email and mobile number as well as their car license number.  For each prompt, the user’s input will be scanned and put in a variable. Once the user has completed entering their information, the data will either be put into the passenger database or the driver database according to their chosen type of account using functions InsertPassenger and InsertDriver respectively and a new ID will be created for them through the createPassengerID and createDriverID functions respectively. Afterward the Login function will be triggered, where the ID formed for them as well as the type of account will be inputted in the function call and there will be a message printed indicating that they have been successfully signed up. The user ID created for them and a string variable stating the type of account created will also be stored into the CurrentUser database until they log out.</p>

</br>

<p>For the second microservice, there is a function called retrievePastTrips which opens the Trip database and using the inputted ID, it retrieves the user’s past trips using the function GetUserTrips where the user has that ID and adds them in a slice. When the user has logged into the system as a Passenger and chosen to View Past Trips, their ID will be retrieved from the CurrentUser database and used in a function call for retrievePastTrips which returns a slice. The slice is then printed out to show the users. If there are no past trips, a message will be printed to inform the user that they have had no past trips so far. </p>

</br>

<h1>
Architecture Diagram
</h1>

<img width="800" src="architecturediagram.png" alt="My Architecture Diagram for Microservices">

<h1>
Instructions for Setting Up and Running Microservices
</h1>

<ul>
    <li>
        Make 4 SQL Connections
        <ul>
            <li>Passengers</li>
            <li>Drivers</li>
            <li>Trip</li>
            <li>CurrentUser</li>
        </ul>
    </li>
    <li>
    First run the various SQL files according to the connections
        <ul>
            <li>Example: Passenger_DB.sql for Passengers SQL connection and Driver_DB.sql for the Drivers SQL connection</li>
            <li>The databases and tables will be created</li>
            <li>In the Trip SQL file, there are INSERT statements to input 3 trips into the Trip database</li>
            <li>In the Passenger file, there is one INSERT statement to input a passenger data into the Passengers database</li>
        </ul>
    </li>
    <li>Enter the command “go run .” to run all the files</li>

</ul>
