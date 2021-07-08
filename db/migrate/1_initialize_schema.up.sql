BEGIN;

CREATE TABLE vehicles (
    ID varchar(10) NOT NULL,
    Name varchar(100) NOT NULL,
    Brand varchar(100) NOT NULL,
    TypeOfVehicle varchar(100) NULL,
    CanDrive boolean NOT NULL,
    CONSTRAINT vehicle_key PRIMARY KEY (ID)
);

COMMIT;