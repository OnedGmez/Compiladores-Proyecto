package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var ArchivoInstrucciones []byte

func main() {
	var NombreArchivo string
	fmt.Println("Ingrese el nombre del archivo con su extensión")
	fmt.Scanln(&NombreArchivo)
	CargarInstrucciones()
	NuevoArchivo(NombreArchivo)
	GenerarInstrucciones(NombreArchivo)
}

/*
Función que servirá para llenar la variable que contiene las intrucciones GO y su equivalente en lenguaje Natural
*/
func CargarInstrucciones() {
	Data, err := ioutil.ReadFile("INCISO-2_3/data_sources/Go_instrucciones.csv")
	if err != nil {
		fmt.Println("Error al abrir archivo")
	} else {
		ArchivoInstrucciones = Data
	}
}

func GenerarInstrucciones(NombreArchivo string) {
	Contador := 0
	Archivo, err := ioutil.ReadFile("Nuevo_" + NombreArchivo)
	if err != nil {
		fmt.Println("Error al abrir archivo")
	} else {
		LineaData := Tokenizador(Archivo, "\n")
		for _, Linea := range LineaData {
			Token := Tokenizador([]byte(Linea), " ")
			for _, Palabra := range Token {
				BuscarSim(Palabra)
			}
			Contador++
		}
	}
}

/*
Función que sirve para analizar si en una palabra se encuentra algún símbolo
*/
func BuscarSim(Palabra string) {

}

/*
Función que nos servirá para separar la data según se necesite
*/
func Tokenizador(Cadena []byte, Delimitador string) []string {
	Tokenizado := strings.Split(string(Cadena), Delimitador)
	return Tokenizado
}

/*
Función para quitar los espacios innecesarios en el archivo .go
*/
func LimpiezaData(NombreArchivo string) string {
	var DataNueva string
	Data, err := ioutil.ReadFile("INCISO-2_3/data_sources/" + NombreArchivo)
	if err != nil {
		return "No se puede abrir el archivo"
	} else {
		DataProc := Tokenizador(Data, "\n")
		for _, Linea := range DataProc {
			if Linea != "" {
				if DataNueva != "" {
					DataNueva = DataNueva + "\n" + Linea
				} else {
					DataNueva = Linea
				}
			}
		}
	}
	return DataNueva
}

/*
Función que nos generará el archivo nuevo sin espacios
*/
func NuevoArchivo(NombreArchivo string) {
	DataNueva := LimpiezaData(NombreArchivo)
	NuevoNombre := "Nuevo_" + NombreArchivo
	err := ioutil.WriteFile("INCISO-2_3/data_sources/"+NuevoNombre, []byte(DataNueva), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
