package main

import (
	"fmt"
	"io/ioutil"
)

var ArchivoInstrucciones []byte

func main() {
	TraerInstrucciones()

}

/*
Función que servirá para llenar la variable que contiene las intrucciones GO y su equivalente en lenguaje Natural
*/
func TraerInstrucciones() {
	Data, err := ioutil.ReadFile("INCISO-2_3/data_sources/Go_instrucciones.csv")
	if err != nil {
		fmt.Println("Error al abrir archivo")
	} else {
		ArchivoInstrucciones = Data
	}
}
