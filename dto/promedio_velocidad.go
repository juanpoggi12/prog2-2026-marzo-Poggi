package dto

import "time"

type PromedioVelocidadRequest struct {
	Atleta      string    `json:"atleta" binding:"required"`
	Velocidades []float64 `json:"velocidades" binding:"required"`
}

type PromedioVelocidadResponse struct {
	ID                string    `json:"id"`
	Atleta            string    `json:"atleta"`
	PromedioVelocidad float64   `json:"promedio_velocidad"`
	CantidadRegistros int       `json:"cantidad_registros"`
	Fecha             time.Time `json:"fecha"`
}
