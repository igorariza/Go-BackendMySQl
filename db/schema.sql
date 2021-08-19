CREATE TABLE `institution` ( 	
	`idInstitution`	INT	AUTO_INCREMENT,
	`dane_id`	VARCHAR(100),	
	`name`	VARCHAR(80),	
	`nit_id`	VARCHAR(80),		
PRIMARY KEY (`idInstitution`) ) ENGINE=InnoDB;

CREATE TABLE `sede` ( 				
	`idSede`	INT	AUTO_INCREMENT,
  `idInstitution`	INT	, -- relationship Institution
	`code_sede`	VARCHAR(100),	
	`name_sede`	VARCHAR(80),	
PRIMARY KEY (`idSede`),
FOREIGN KEY (`idInstitution`) 
REFERENCES institution(`idInstitution`)) ENGINE=InnoDB;		

CREATE TABLE `users` ( 	
	`idUser`	INT	AUTO_INCREMENT,
	`document_id`	VARCHAR(100),	
	`first_name`	VARCHAR(80),	
	`last_name`	VARCHAR(80),	
	`email`	VARCHAR(200),	
	`password`	VARCHAR(1000),	
	`phone`	VARCHAR(20),		
	`address`	VARCHAR(20),		
	`photo`	VARCHAR(1000),		
	`created_at`	VARCHAR(100),		
	`type_id`	VARCHAR(2),		
	`date_birth`	VARCHAR(100),		
	`last_access`	VARCHAR(100),		
	`rh`	VARCHAR(2),		
	`idSede`	VARCHAR(2), -- relationship Sede		
	`is_active`	VARCHAR(2),		
PRIMARY KEY (`idUser`),
FOREIGN KEY (`idSede`) 
REFERENCES sede(idSede)) AUTO_INCREMENT=3;

-- INSERT ----

-- SELECT * FROM backendSigeGo.Institution;
INSERT INTO institution
	(`dane_id`,	`name`)
VALUES
('176364000015',
'IE CENTRAL DE BACHILLERATO INTEGRADO');

-- SELECT * FROM backendSigeGo.Sede;
INSERT INTO sede
	(`idInstitution`, `code_sede`, `name_sede`)
VALUES
(1,'17636400001502','MANUELA BELTRAN');

-- SELECT * FROM backendSigeGo.Users;
-- Json
{
    "document_id":"11223344",
    "first_name": "Admin",
    "last_name": "SIGE",
    "email": "adminsige@gmail.com",
    "password": "11223344",
    "phone": "",
    "address": "",
    "photo": "",
    "created_at": "",
    "type_id": "",
    "date_birth": "",
    "last_access": "",
    "rh": "",    
    "idSede": "1",
    "is_active": ""
}	
-- SELECT * from Users WHERE email ='adminsige@gmail.com'

-- QUERY

SELECT Users.idUser, Users.document_id, Users.first_name, Users.last_name, Users.email, Users.password, Users.phone, Users.address, Users.photo, Users.created_at, Users.type_id, Users.date_birth, Users.rh, Users.idSede, Users.is_active, Sede.name_sede 
FROM Users
INNER JOIN Sede ON Sede.idSede = Users.idSede
WHERE Users.email = "adminsige@gmail.com"