CREATE database driver_db;
USE driver_db;

CREATE TABLE Drivers (Id INT NOT NULL PRIMARY KEY, Firstname VARCHAR(30), Lastname VARCHAR(30), Emailaddress VARCHAR(30), Mobileno INT, Carlicenseno VARCHAR(8));
INSERT INTO driver_db.drivers values(1, "John", "Tan", "jtan@gmail.com", 81234567, "SKW3784S");
