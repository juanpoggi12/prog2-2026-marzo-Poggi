package dto

type Proyeccion struct {
	Semana      int     `json:"semana"`
	Rendimiento float64 `json:"rendimiento"`
}
type ProyeccionMejoraResponse struct {
	RendimientoInicial float64      `json:"rendimiento_inicial"`
	TasaMejoraSemanal  float64      `json:"tasa_mejora_semanal"`
	Proyeccion         []Proyeccion `json:"proyeccion"`
}
