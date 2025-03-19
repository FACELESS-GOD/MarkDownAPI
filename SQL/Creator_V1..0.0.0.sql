CREATE database MarkDownAPIV2;

use MarkDownAPIV2 ; 

Create table User_Data(
	Userid int NOT NULL ,
	FirstName varchar(255),
	LastName varchar(255),
    Email varchar(255),
    Title varchar(255)    
);

Create table User_Cred(
	Userid int NOT NULL AUTO_INCREMENT,
	UserName varchar(255),
	Password varchar(255),
    constraint unique(Userid,UserName),
    PRIMARY KEY (Userid)
);

Create table Token_Storage(
	TokenID int NOT NULL AUTO_INCREMENT,
	UserID int,
    ISLoggedIN int,
	Token varchar(255),
    DateTime datetime,
    constraint unique(TokenID),
    PRIMARY KEY (Userid)
);