CREATE TABLE `Employee` (
    `ID` int NOT NULL AUTO_INCREMENT,
    `Name` text NOT NULL,
    `Gender` varchar(255) NOT NULL,
    `Phone` varchar(255) DEFAULT NULL,
    `Email` varchar(255) DEFAULT NULL,
    `Address` json DEFAULT NULL,
    `VacationDays` decimal(10,0) DEFAULT NULL,
    PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;;