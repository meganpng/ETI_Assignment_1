CREATE USER 'user'@'127.0.0.1' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'127.0.0.1';

CREATE database trip_db;
USE trip_db;

CREATE TABLE Trip (Id INT NOT NULL PRIMARY KEY, Pickuplocation VARCHAR(30), Pickupcode INT, Dropofflocation VARCHAR(30), 
Dropoffcode INT, DriverID VARCHAR(8), Drivername VARCHAR(30), PassengerID VARCHAR(8), TripStatus VARCHAR(8),
TripCost decimal, TripStart VARCHAR(30), TripEnd VARCHAR(30));

INSERT INTO Trip values(1, "391A Orchard Road", 238873, 
	"83 Punggol Central", 828761, 1, "John Tan", 1, "Complete", 19.00,
	 "12 December 2021, 09:45", "12 December 2021, 10:11");

INSERT INTO Trip values(2, "83 Punggol Central", 828761, 
	"9 Bishan Pl", 579837, 2, "Tai Yang Lee", 1, "Complete", 17.00,
	 "12 December 2021, 012:14", "12 December 2021, 12:31");

INSERT INTO Trip values(3, "9 Bishan Pl", 579837, 
	"204 Hougang Street 21", 530204, 3, "Jaiden Chen", 1, "Complete", 15.00,
	 "12 December 2021, 14:53", "12 December 2021, 15:05");
     
