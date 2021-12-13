CREATE USER 'user'@'127.0.0.1' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'127.0.0.1';

CREATE database currentuser_db;
USE currentuser_db;

CREATE TABLE CurrentUser (Id INT NOT NULL PRIMARY KEY, Usertype VARCHAR(30)) ;

