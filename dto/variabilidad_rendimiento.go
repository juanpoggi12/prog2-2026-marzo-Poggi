package dto

import "time"

type VariabilidadRendimientoRequest struct {
	Atleta       string    `json:"atleta" binding:"required"`
	Rendimientos []float64 `json:"rendimientos" binding:"required"`
}

type VariabilidadRendimientoResponse struct {
	Atleta       string    `json:"atleta"`
	Variabilidad float64   `json:"variabilidad"`
	Metodo       string    `json:"metodo"`
	Maximo       int       `json:"maximo"`
	Minimo       int       `json:"minimo"`
	Fecha        time.Time `json:"fecha"`
}
