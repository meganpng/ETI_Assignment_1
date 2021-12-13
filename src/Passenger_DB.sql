CREATE USER 'user'@'127.0.0.1' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'127.0.0.1';
CREATE database passenger_db;
USE passenger_db;

CREATE TABLE Passengers (Id INT NOT NULL PRIMARY KEY, Firstname VARCHAR(30), Lastname VARCHAR(30), Emailaddress VARCHAR(30), Mobileno INT);
INSERT INTO Passengers values(1, "Jane", "Tan", "jtan@gmail.com", 81234567);




