package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var ArchivoInstrucciones string
var ArchivoInstruccionesLenguaje string
var InstruccionesGo string
var Reconstruccion string

func main() {
	CargarInstrucciones()
	var NombreArchivo string
	fmt.Println("Ingrese el nombre del archivo con su extensión")
	fmt.Scanln(&NombreArchivo)
	LeerInstrucciones(NombreArchivo)
	DataGo_Instrucciones()
	NuevoArchivo("Reconstruccion", NombreArchivo, []byte(Reconstruccion))
}

func CargarInstrucciones() {
	Data, err := ioutil.ReadFile("./data_sources/Go_instrucciones.csv")
	if err != nil {
		fmt.Println("Error al abrir archivo")
	} else {
		ArchivoInstrucciones = strings.ReplaceAll(string(Data), " ", "")
	}
}

func LeerInstrucciones(NombreArchivo string) {
	DataInstrucciones, err := ioutil.ReadFile("INCISO-2_3/data_sources/" + NombreArchivo)
	if err != nil {
		fmt.Println("Error al abrir archivo")
	} else {
		ArchivoInstruccionesLenguaje = strings.ReplaceAll(string(DataInstrucciones), " ", "")
	}
}

func Tokenizador(Cadena []byte, Delimitador string) []string {
	Tokenizado := strings.Split(string(Cadena), Delimitador)
	return Tokenizado
}

func DataGo_Instrucciones() {
	contador := 0
	var NumLinea int = 0
	byteArchivoInstrucciones := []byte(ArchivoInstrucciones)
	byteArchivoInstruccionesLenguaje := []byte(ArchivoInstruccionesLenguaje)
	ArrayInstrucciones := Tokenizador(byteArchivoInstrucciones, "\n")
	ArrayInstruccionesLenguaje := Tokenizador(byteArchivoInstruccionesLenguaje, "\n")
	for _, LineaNat := range ArrayInstruccionesLenguaje {
		TokenLineaNat := Tokenizador([]byte(LineaNat), "\t") //instrucciones
		contador = 0
		for _, Linea := range ArrayInstrucciones {
			TokenLinea := Tokenizador([]byte(Linea), ";") //csv
			if string(TokenLineaNat[1]) == string(TokenLinea[1]) {
				LineaTMP, err := strconv.Atoi(TokenLineaNat[0])
				if err == nil {
					if LineaTMP == NumLinea {
						LlenarRecons(string(TokenLinea[0]))
					} else {
						NumLinea = LineaTMP
						LlenarRecons(string("\n" + TokenLinea[0]))
					}
				}
			}
			contador++
		}
		if strings.Contains(string(TokenLineaNat[1]), ":Texto:") {
			Texto := strings.ReplaceAll(string(TokenLineaNat[1]), ":Texto:", "")
			LlenarRecons(Texto)
		}
	}
}

func LlenarRecons(Data string) {
	if Reconstruccion != "" {
		Reconstruccion = Reconstruccion + " " + Data
	} else {
		Reconstruccion = Data
	}
}

func NuevoArchivo(Prefijo string, NombreArchivo string, Data []byte) {
	NuevoNombre := Prefijo + "_" + NombreArchivo
	err := ioutil.WriteFile("INCISO-2_3/data_sources/"+NuevoNombre, Data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
