package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ArchivoInstrucciones []byte
var InstruccionesGo string

func main() {
	var NombreArchivo string
	fmt.Println("Ingrese el nombre del archivo con su extensión")
	fmt.Scanln(&NombreArchivo)
	CargarInstrucciones()
	NuevoArchivo(NombreArchivo)
	GenerarInstrucciones(NombreArchivo)
	fmt.Println(InstruccionesGo)
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
	Archivo, err := ioutil.ReadFile("INCISO-2_3/data_sources/Nuevo_" + NombreArchivo)
	if err != nil {
		fmt.Println("Error al abrir archivo")
	} else {
		LineaData := Tokenizador(Archivo, "\n")
		for _, Linea := range LineaData {
			Token := Tokenizador([]byte(Linea), " ")
			for _, Palabra := range Token {
				AnalizarPalabra(strings.TrimSpace(Palabra), Contador)
			}
			Contador++
		}
	}
}

/*
Función que sirve para analizar si en una palabra se encuentra algún símbolo
*/
func AnalizarPalabra(Palabra string, Contador int) {
	var Instruccion string
	var PalabraTMP string
	var SimboloTMP string
	var LetraRuna rune

	if strings.ContainsAny(Palabra, "{}()/*\"") {
		LetrasPalabra := Tokenizador([]byte(Palabra), "")
		for _, Letra := range LetrasPalabra {
			LetraRuna = rune(Letra[0])
			switch unicode.IsLetter(LetraRuna) {
			case true:

				if SimboloTMP != "" {
					SimboloTMP = SimboloTMP + "," + strconv.Itoa(Contador)
					InstruccionesGO(SimboloTMP)
					SimboloTMP = ""
				}

				if PalabraTMP != "" {
					PalabraTMP = PalabraTMP + string(LetraRuna)
				} else {
					PalabraTMP = string(LetraRuna)
				}

				break
			case false:
				if PalabraTMP != "" {
					PalabraTMP = PalabraTMP + "," + strconv.Itoa(Contador)
					InstruccionesGO(PalabraTMP)
					PalabraTMP = ""
				}

				switch LetraRuna {
				case '/':
					SimboloTMP = SimboloTMP + string(LetraRuna)
					break
				case '*':
					SimboloTMP = SimboloTMP + string(LetraRuna)
					break
				default:
					SimboloTMP = string(LetraRuna) + "," + strconv.Itoa(Contador)
					InstruccionesGO(SimboloTMP)
					SimboloTMP = ""
				}

			}
		}
		IndexLetra := strings.LastIndex(Palabra, string(LetraRuna))

		if IndexLetra == len(Palabra) {
			if PalabraTMP != "" {
				PalabraTMP = PalabraTMP + "," + strconv.Itoa(Contador)
				InstruccionesGO(PalabraTMP)
				PalabraTMP = ""
			}

			if SimboloTMP != "" {
				SimboloTMP = SimboloTMP + "," + strconv.Itoa(Contador)
				InstruccionesGO(SimboloTMP)
				SimboloTMP = ""
			}
		}

	} else {
		Instruccion = Palabra + "," + strconv.Itoa(Contador)
		InstruccionesGO(Instruccion)
	}
}

func InstruccionesGO(Instruccion string) {
	if InstruccionesGo != "" {
		InstruccionesGo = InstruccionesGo + "\n" + Instruccion
	} else {
		InstruccionesGo = Instruccion
	}
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
