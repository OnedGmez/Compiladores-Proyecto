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
	CargarInstrucciones()
	var DataNueva string
	var NombreArchivo string
	fmt.Println("Ingrese el nombre del archivo con su extensión")
	fmt.Scanln(&NombreArchivo)
	DataNueva = LimpiezaData(NombreArchivo)
	NuevoArchivo("Nuevo", NombreArchivo, []byte(DataNueva))
	GenerarInstrucciones(NombreArchivo)

	DataNueva = BuscarIgual()
	NombreArchivo = strings.Trim(NombreArchivo, ".go")
	NuevoArchivo("Instrucciones", NombreArchivo, []byte(DataNueva))
}

/*
def: Función que servirá para llenar la variable que contiene las instrucciones GO y su equivalente en lenguaje Natural
*/
func CargarInstrucciones() {
	Data, err := ioutil.ReadFile("../data_sources/Go_instrucciones.csv")
	if err != nil {
		fmt.Println("Error al abrir archivo")
	} else {
		ArchivoInstrucciones = Data
	}
}

/*
def: Función que nos ayudará a leer el archivo de GO para analizarlo posteriormente, separando la información del archivo en líneas y luego en palabras
* NombreArchivo (in): Parámetro que almacenará el nombre del archivo al que le aplicará el proceso
*/
func GenerarInstrucciones(NombreArchivo string) {
	Contador := 0
	Archivo, err := ioutil.ReadFile("../data_sources/Nuevo_" + NombreArchivo)
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
def: Función que nos ayudará a comparar el código del .go con las instrucciones de la tabla de símbolos (.csv) cargada en el arranque de la ejecución
return: Devuelve un string formateado con la respectiva instrucción en lenguaje natural y su número de línea del archivo .go
* GoInstruc: Variable que nos ayudará a almacenar el código en lenguaje go y el respectivo número de línea
* InstrucNat: Variable que nos ayudará a almacenar el equivalente en lenguaje natural de las palabras reservadas de go
: Texto:: Lógica aplicada al contenido del archivo Go que no es ninguna palabra reservada o que no se encontró en la tabla de símbolos (Variable con la información del .csv)
*/
func BuscarIgual() string {
	var DataRetorno string
	var GoInstruc []string
	var InstrucNat []string
	LineasInstruccionesGO := Tokenizador([]byte(InstruccionesGo), "\n")
	LineasInstrucciones := Tokenizador([]byte(ArchivoInstrucciones), "\n")
	for _, InstrucGo := range LineasInstruccionesGO {
		for _, Instruc := range LineasInstrucciones {
			GoInstruc = Tokenizador([]byte(InstrucGo), ",")
			InstrucNat = Tokenizador([]byte(Instruc), ";")

			if GoInstruc[0] == InstrucNat[0] {
				Instruccion := string(GoInstruc[1]) + "\t" + InstrucNat[1]
				if DataRetorno != "" {
					DataRetorno = DataRetorno + "\n" + Instruccion
				} else {
					DataRetorno = Instruccion
				}
				GoInstruc[0] = ""
				break
			}
		}

		if GoInstruc[0] != "" {
			Instruccion := string(GoInstruc[1]) + "\t" + ":Texto:" + GoInstruc[0]
			if DataRetorno != "" {
				DataRetorno = DataRetorno + "\n" + Instruccion
			} else {
				DataRetorno = Instruccion
			}
		}
	}
	return DataRetorno
}

/*
def: Función que sirve para analizar si en una palabra se encuentra algún símbolo, separándola carácter por carácter
* Palabra (in): Parámetro que recibe la palabra que vamos a analizar para saber si existe un símbolo, si no, agregarlo directamente temporal que creará la variable definitiva con todas las instrucciones
* Contador (in): Parámetro que almacena el número de línea a la que pertenece la palabra que está siendo analizada
* PalabraTMP: Variable para almacenar temporalmente las letras de la palabra que está siendo analizada
* SimboloTMP: Variable para almacenar temporalmente los símbolos de la palabra que está siendo analizada
* IndexLetra: Variable que funciona como un iterador sobre las letras de la palabra
* LetraRuna: Variable que utilizamos para convertir el carácter (string) en rune type para poderla analizarlo si era o no una letra
*/
func AnalizarPalabra(Palabra string, Contador int) {
	var Instruccion string
	var PalabraTMP string
	var SimboloTMP string
	var LetraRuna rune
	var IndexLetra int = 0

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

				if len(Palabra)-1 == IndexLetra {
					if PalabraTMP != "" {
						PalabraTMP = PalabraTMP + "," + strconv.Itoa(Contador)
						InstruccionesGO(PalabraTMP)
						PalabraTMP = ""
					}
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
			IndexLetra++
		}

	} else {
		Instruccion = Palabra + "," + strconv.Itoa(Contador)
		InstruccionesGO(Instruccion)
	}
}

/*
def: Función que nos ayuda únicamente a reducir la redundancia de código al momento de generar la variable que utilizaremos para guardar el código de go palabra por palabra
*/
func InstruccionesGO(Instruccion string) {
	if InstruccionesGo != "" {
		InstruccionesGo = InstruccionesGo + "\n" + Instruccion
	} else {
		InstruccionesGo = Instruccion
	}
}

/*
def: Función que nos servirá para separar la data según se necesite
return: Devuelve un arreglo de strings, es decir, la cadena recibida ya separada.
* Cadena (in): Parámetro que nos sirve para recibir la información que queremos separar
* Delimitar (in): Parámetro que nos sirve para recibir el o los caracteres que servirán para separar
*/
func Tokenizador(Cadena []byte, Delimitador string) []string {
	Tokenizado := strings.Split(string(Cadena), Delimitador)
	return Tokenizado
}

/*
def: Función que nos servirá para quitar los espacios innecesarios en el archivo .go
return: Devuelve una variable de tipo string que contiene todo el código go sin espacios y respetando la sintaxis
* NombreArchivo (in): Parámetro que almacenará el nombre del archivo al que le aplicará el proceso
*/
func LimpiezaData(NombreArchivo string) string {
	var DataNueva string
	Data, err := ioutil.ReadFile("../data_sources/" + NombreArchivo)
	if err != nil {
		return "No se puede abrir el archivo"
	} else {
		DataProc := Tokenizador(Data, "\n")
		for _, Linea := range DataProc {
			if (strings.TrimSpace(Linea)) != "" {
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
def: Función que nos generará el archivo que deseamos
* Prefijo (in): Parámetro que nos ayudará a recibir un string que queremos concatenar al nombre del archivo
* NombreArchivo (in): Parámetro que almacenará el nombre que tendrá el archivo, teniendo en cuenta el prefijo
* Data (in): Parámetro que almacenará la información que contendrá el nuevo archivo
*/
func NuevoArchivo(Prefijo string, NombreArchivo string, Data []byte) {
	NuevoNombre := Prefijo + "_" + NombreArchivo
	err := ioutil.WriteFile("../data_sources/"+NuevoNombre, Data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
