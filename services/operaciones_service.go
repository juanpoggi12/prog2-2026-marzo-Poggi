package services

import (
	"context"
	"errors"
	"prog2-2026-marzo-Poggi/database"
	"prog2-2026-marzo-Poggi/dto"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type OperacionesServicesInterface interface {
	CalcularPromedioVelocidad(request dto.PromedioVelocidadRequest) (*mongo.InsertOneResult, dto.PromedioVelocidadResponse, error)
	CalcularVariabilidadRendimiento(request dto.VariabilidadRendimientoRequest) (*mongo.InsertOneResult, dto.VariabilidadRendimientoResponse, error)
	CalcularProyeccionMejora(rendimientoActual float64, tasaMejora float64, semanas int) (dto.ProyeccionMejoraResponse, error)
}

type OperacionesService struct {
	db database.DB
}

func NewOperacionesService(db database.DB) *OperacionesService {
	return &OperacionesService{
		db: db,
	}
}

func (service *OperacionesService) CalcularPromedioVelocidad(request dto.PromedioVelocidadRequest) (*mongo.InsertOneResult, dto.PromedioVelocidadResponse, error) {
	collection := service.db.GetClient().Database("prog2-2026-marzo-Poggi").Collection("operaciones")

	var promedio float64
	var velocidades = request.Velocidades

	for _, velocidad := range velocidades {
		if velocidad < 0 {
			return nil, dto.PromedioVelocidadResponse{}, errors.New("la velocidad no puede ser negativa")
		}
		promedio += velocidad
	}
	var cantidadRegistros = float64(len(velocidades))
	promedio = promedio / cantidadRegistros

	var response dto.PromedioVelocidadResponse
	response.Atleta = request.Atleta
	response.PromedioVelocidad = promedio
	response.CantidadRegistros = int(cantidadRegistros)
	response.Fecha = time.Now()

	result, err := collection.InsertOne(context.TODO(), response)
	return result, response, err

}

func (service *OperacionesService) CalcularVariabilidadRendimiento(request dto.VariabilidadRendimientoRequest) (*mongo.InsertOneResult, dto.VariabilidadRendimientoResponse, error) {
	collection := service.db.GetClient().Database("prog2-2026-marzo-Poggi").Collection("operaciones")

	rendimientos := request.Rendimientos
	var rendimientoMayor float64
	rendimientoMayor = 0

	var rendimientoMenor float64
	rendimientoMenor = 0

	for _, rendimiento := range rendimientos {
		if rendimiento < 0 {
			return nil, dto.VariabilidadRendimientoResponse{}, errors.New("el vaor del rendimiento no puede ser negativo")
		}
		if rendimiento > rendimientoMayor {
			rendimientoMayor = rendimiento
		}
		if rendimiento < rendimientoMenor {
			rendimientoMenor = rendimiento
		}
	}

	variabilidad := rendimientoMayor - rendimientoMenor

	var response dto.VariabilidadRendimientoResponse
	response.Atleta = request.Atleta
	response.Variabilidad = variabilidad
	response.Metodo = "diferencia_max-min"
	response.Maximo = int(rendimientoMayor)
	response.Minimo = int(rendimientoMenor)
	response.Fecha = time.Now()

	result, err := collection.InsertOne(context.TODO(), response)
	return result, response, err
}

func (service *OperacionesService) CalcularProyeccionMejora(rendimientoActual float64, tasaMejora float64, semanas int) (dto.ProyeccionMejoraResponse, error) {
	var proyecciones []dto.Proyeccion
	var response dto.ProyeccionMejoraResponse
	rendimiento := rendimientoActual

	for i := 1; i <= semanas; i++ {
		proyeccion := dto.Proyeccion{}
		proyeccion.Semana = i
		proyeccion.Rendimiento = rendimiento * (1 + tasaMejora/100)
		rendimiento = proyeccion.Rendimiento
		proyecciones = append(proyecciones, proyeccion)
	}

	response.RendimientoInicial = rendimientoActual
	response.TasaMejoraSemanal = tasaMejora
	response.Proyeccion = proyecciones
	return response, nil
}
