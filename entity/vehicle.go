package entity

type Vehicle struct {
	ID            string `gorm:"type: varchar(10); not null; primary key"`
	Name          string `gorm:"type: varchar(100); not null" json:"name"`
	Brand         string `gorm:"type: varchar(100); not null"`
	TypeOfVehicle string `gorm:"type: varchar(100); column:typeofvehicle; null"`
	CanDrive      bool   `gorm:"type: boolean; column:candrive; null"`
}
