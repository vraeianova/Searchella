package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Record es la estructura que representa un registro del archivo
type Record struct {
	MessageID               string `json:"Message-ID"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	To                      string `json:"To"`
	Subject                 string `json:"Subject"`
	MimeVersion             string `json:"Mime-Version"`
	ContentType             string `json:"Content-Type"`
	ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
	XFrom                   string `json:"X-From"`
	XTo                     string `json:"X-To"`
	Xcc                     string `json:"X-cc"`
	Xbcc                    string `json:"X-bcc"`
	XFolder                 string `json:"X-Folder"`
	XOrigin                 string `json:"X-Origin"`
	XFileName               string `json:"X-FileName"`
	Message                 string `json:"Message"`
}

type Output struct {
	Index   string   `json:"index"`
	Records []Record `json:"records"`
}

var extension = ".go"

// Función recursiva para obtener todas las carpetas de un directorio

func getFolders(folderpath string) (map[string]interface{}, []Record) {
	// Crear un mapa para almacenar las carpetas del directorio y un slice de Record
	folders := make(map[string]interface{})
	var records []Record

	// Iterar sobre los elementos del directorio
	entries, err := ioutil.ReadDir(folderpath)
	if err != nil {
		fmt.Println(err)
		return folders, records
	}

	// Iterar sobre cada elemento
	for _, entry := range entries {
		// Obtener la ruta completa del elemento
		fullPath := filepath.Join(folderpath, entry.Name())
		if !strings.HasPrefix(entry.Name(), ".") {
			if entry.IsDir() {
				// Si es un directorio, llamar a la función recursivamente
				subfolders, subrecords := getFolders(fullPath)
				// Añadir el mapa de subcarpetas y el slice de subregistros al mapa y al slice actuales
				folders[entry.Name()] = subfolders
				records = append(records, subrecords...)
			} else {
				// Si es un archivo, procesarlo y añadirlo al slice de registros

				ext := path.Ext(entry.Name())
				if ext != extension {
					// fmt.Println("Entro con la extension", ext)
					fmt.Println("estoy a punto de leer la linea", entry.Name())
					file, err := os.Open(fullPath)

					// fmt.Println("archivo", file)

					if err != nil {
						fmt.Println(err)
					}

					defer file.Close()

					// var records []Record

					var currentRecord Record

					record := Record{}

					scanner := bufio.NewScanner(file)

					for scanner.Scan() {
						line := scanner.Text()
						// fmt.Println("Verificando la línea", line)

						// Divide la línea en clave y valor
						parts := strings.Split(line, ": ")
						if len(parts) != 2 {
							// fmt.Println("Verificar que entro aqui en esta linea", line)
							// La línea no tiene el formato esperado (clave: valor)
							// Añade el valor de la línea al campo Message
							currentRecord.Message += line + "\n"
							continue
						}
						key := parts[0]
						value := parts[1]

						// Crea una instancia de Record y asigna los valores de clave y valor a los campos correspondientes
						// record := Record{}
						// fmt.Println("VERIFICAR EL VALUE", value)
						switch key {
						case "Message-ID":
							record.MessageID = value
						case "Date":
							record.Date = value
						case "From":
							record.From = value
						case "To":
							record.To = value
						case "Subject":
							record.Subject = value
						case "Mime-Version":
							record.MimeVersion = value
						case "Content-Type":
							record.ContentType = value
						case "Content-Transfer-Encoding":
							record.ContentTransferEncoding = value
						case "X-From":
							record.XFrom = value
						case "X-To":
							record.XTo = value
						case "X-cc":
							record.Xcc = value
						case "X-bcc":
							record.Xbcc = value
						case "X-Folder":
							record.XFolder = value
						case "X-Origin":
							record.XOrigin = value
						case "X-FileName":
							record.XFileName = value
						}

					}
					fmt.Println(record.MessageID)

					record.Message = currentRecord.Message
					records = append(records, record)

				}
			}
		}
		output := Output{
			Index:   "messages",
			Records: records,
		}

		outputJSON, err := json.Marshal(output)
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println("Verificar output", string(outputJSON))
		err = ioutil.WriteFile("output.json", outputJSON, 0644)
		if err != nil {
			fmt.Println(err)

		}
	}

	// Devolver el mapa de carpetas y el slice de registros
	return folders, records
}

func main() {
	// Obtener el directorio actual
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	getFolders(currentDir)

}
