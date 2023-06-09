package models

type AssignmentResponse struct {
	ShipmentAddress string `json:"shipent_address"`
	Driver          Driver `json:"driver"`
}

type Driver struct {
	Name string  `json:"name"`
	SS   float32 `json:"ss"`
}
