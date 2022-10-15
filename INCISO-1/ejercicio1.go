package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var path = "HolaMundo.py"
var x int
var can string

func main() {
	x = 0
	fmt.Println("valor", x)
	crearArchivo()

	readedprogram, err := ioutil.ReadFile("codigo_fuente.txt")

	if err != nil {
		fmt.Println("Error leyendo archivo")
	} else {
		tokens := tokenizador(readedprogram, "\n")

		for i := 0; i < len(tokens); i++ {

			cadena := []byte(tokens[i])
			lexema := tokenizador(cadena, " ")
			for _, token := range lexema {
				fmt.Println("valor cambio", token)
				fmt.Println("valor actual de cambio", x)

				if x == 0 {
					if token == "std::cout" {
						x = 1
						fmt.Println("valor actual", x)
						fmt.Println(token)
						can += "print "
						//escribeArchivoln("print")
						fmt.Println()
					}

				}
				if x == 1 {
					if token == "<<" {
						x = 2
						fmt.Println("valor actual", x)
						fmt.Println(token)
						can += "( "
						//escribeArchivoln("(")
						fmt.Println()
					}

				}

				if x == 2 {
					if leer_tsimbolos(token) == false {
						x = 3
						fmt.Println("valor actual", x)
						fmt.Println(token)
						can += token + " "
						//escribeArchivoln(token)
						fmt.Println()
					}
				}

				if x == 3 {
					x = 0
					fmt.Println("valor actual", x)
					fmt.Println(token)
					can += ")"
					//escribeArchivoln(")")

				}

			}

		}

		escribeArchivo(can)
	}

}

//func generar_token

func leer_tsimbolos(token string) bool {
	readedprogram, err := ioutil.ReadFile("tabla_simbolos.txt")

	if err != nil {
		fmt.Println("Error leyendo archivo")
	} else {
		if strings.Contains(string(readedprogram), token) {
			return true
		}
	}
	return false
}

func tokenizador(cadena []byte, delimitador string) []string {
	/*leer := bufio.NewReader(os.Stdin)*/
	texto := string(cadena)
	tokens := strings.Split(texto, delimitador)
	return tokens
}

func crearArchivo() {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func escribeArchivoln(lexema string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	fmt.Println(lexema)
	_, err = file.WriteString(lexema)
	_, err = file.WriteString(" ")
	if existeError(err) {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
}

func escribeArchivo(lexema string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(lexema)
	if existeError(err) {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
}

func escribeArchivoca(lexema string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(lexema)
	if existeError(err) {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
}
