package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Función recursiva para obtener todas las carpetas de un directorio
func getFolders(path string) map[string]interface{} {
	// Crear un mapa para almacenar las carpetas del directorio
	folders := make(map[string]interface{})

	// Iterar sobre los elementos del directorio
	entries, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		return folders
	}
	// fmt.Println("Verificar las entries", entries)
	for _, entry := range entries {
		// Obtener la ruta completa del elemento
		fullPath := filepath.Join(path, entry.Name())

		// Verificar si el elemento es una carpeta
		if entry.IsDir() {
			fmt.Println("Es directorio", entry.Name())
			// Agregar la carpeta al mapa y obtener las carpetas
			// contenidas en la carpeta
			folders[entry.Name()] = getFolders(fullPath)
		} else {

			b, err := ioutil.ReadFile(fullPath)
			if err != nil {
				// Si hay un error al leer el archivo, imprimimos el error y terminamos el programa
				fmt.Println(err)

			}
			fmt.Println("Encontré un archivo", entry.Name())

			// Convertimos la matriz de bytes en una cadena y la imprimimos en la consola
			str := string(b)
			// fmt.Println(str)
			folders[entry.Name()] = str

		}
	}

	// Devolver el mapa de carpetas
	return folders
}

func exportJSON(jsonData []byte) error {
	// Crear un archivo en el directorio actual
	file, err := os.Create("folders.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Escribir el JSON en el archivo
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Obtener el directorio actual
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	folders := getFolders(currentDir)

	jsonData, err := json.Marshal(folders)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Escribir el JSON en un archivo en el directorio actual
	file, err := os.Create("folders.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
}
