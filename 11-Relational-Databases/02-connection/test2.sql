create schema test2;
use test2;
CREATE TABLE `amigos` (
  `aID` int NOT NULL AUTO_INCREMENT,
  `aName` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`aID`)
);

INSERT INTO `amigos` (`aID`, `aName`) VALUES (1,'Jorge'),(2,'Felipe'),(3,'Padre'),(4,'Alberto');

select *
from amigos;