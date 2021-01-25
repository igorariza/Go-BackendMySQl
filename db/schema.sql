CREATE TABLE `Institution` ( 	
	`idInstitution`	INT	AUTO_INCREMENT,
	`dane_id`	VARCHAR(100),	
	`name`	VARCHAR(80),	
	`nit_id`	VARCHAR(80),		
PRIMARY KEY (`idInstitution`) ) ENGINE=InnoDB  
DEFAULT CHARSET=utf8mb4;

CREATE TABLE `Sede` ( 				
	`idSede`	INT	AUTO_INCREMENT,
  `idInstitution`	INT	, -- relationship Institution
	`code_sede`	VARCHAR(100),	
	`name`	VARCHAR(80),	
PRIMARY KEY (`idSede`),
FOREIGN KEY (`idInstitution`) 
REFERENCES Institution(`idInstitution`)) ENGINE=InnoDB  
DEFAULT CHARSET=utf8mb4;			

CREATE TABLE `Users` ( 	
	`idUser`	INT	AUTO_INCREMENT,
	`document_id`	VARCHAR(100),	
	`first_name`	VARCHAR(80),	
	`last_name`	VARCHAR(80),	
	`email`	VARCHAR(200),	
	`password`	VARCHAR(1000),	
	`phone`	VARCHAR(20),		
	`address`	VARCHAR(20),		
	`photo`	VARCHAR(1000),		
	`created_at`	VARCHAR(50),		
	`type_id`	INT,		
	`date_birth`	DATE,		
	`last_access`	DATE,		
	`rh`	VARCHAR(2),		
	`idSede`	INT, -- relationship Sede		
	`is_active`	BOOLEAN,		
PRIMARY KEY (`idUser`),
FOREIGN KEY (`idSede`) 
REFERENCES Sede(idSede)) ENGINE=InnoDB AUTO_INCREMENT=3 
DEFAULT CHARSET=utf8mb4;

-- INSERT ----

-- SELECT * FROM backendSigeGo.Institution;
INSERT INTO Institution
	(`dane_id`,	`name`)
VALUES
('176364000015',
'IE CENTRAL DE BACHILLERATO INTEGRADO');

-- SELECT * FROM backendSigeGo.Sede;
INSERT INTO backendSigeGo.Sede
	(`idInstitution`, `code_sede`, `name`)
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