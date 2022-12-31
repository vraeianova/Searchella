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
func getFolders(folderpath string) map[string]interface{} {
	// Crear un mapa para almacenar las carpetas del directorio
	folders := make(map[string]interface{})

	// Iterar sobre los elementos del directorio
	entries, err := ioutil.ReadDir(folderpath)

	// fmt.Println("Verificar todas las entries", entries)

	if err != nil {
		fmt.Println(err)
		return folders
	}

	// execFile, err := os.Executable()
	// if err != nil {
	// 	// Handle error
	// }
	// fmt.Println("entries", entries)
	var records []Record
	// outerLoop:
	for _, entry := range entries {
		// Obtener la ruta completa del elemento
		fullPath := filepath.Join(folderpath, entry.Name())
		if !strings.HasPrefix(entry.Name(), ".") {
			if entry.IsDir() {
				// fmt.Println("entrando al path", fullPath)
				getFolders(fullPath)
			} else {
				// fmt.Println("Verificar", fullPath)

				// if entry.Name() != ""
				// file, err := os.Open(fullPath)

				ext := path.Ext(entry.Name())
				// fmt.Println("ext", ext, "extension", extension)
				if ext != extension {
					// fmt.Println("Entro con la extension", ext)
					// fmt.Println("estoy a punto de leer la linea", fullPath)

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
					record.Message = currentRecord.Message
					records = append(records, record)

					// fmt.Println("===========fullypathy===========", fullPath)
					// fmt.Println("~~~~~~~Verificando el record~~~~~~~", record)

					// outputJSON, err := json.Marshal(records)
					// if err != nil {
					// 	fmt.Println(err)

					// }
					// fmt.Println("VERIFICAR OUTPUT EN JSON", string(outputJSON))
					// fmt.Println("VERIFICANDO EL ARREGLO DE RECORDS", records)
					// output := Output{
					// 	Index:   "messages",
					// 	Records: records,
					// }

					// outputJSON, err := json.Marshal(records)
					// if err != nil {
					// 	fmt.Println(err)

					// }
					output := Output{
						Index:   "messages",
						Records: records,
					}

					outputJSON, err := json.Marshal(output)
					if err != nil {
						fmt.Println(err)

					}
					fmt.Println("Verificar output", string(outputJSON))
					// break outerLoop
					// // Exporta la cadena de caracteres de formato JSON a un archivo
					err = ioutil.WriteFile("output.json", outputJSON, 0644)
					if err != nil {
						fmt.Println(err)

					}
				}

			}
		}

	}
	// fmt.Println("Verificar records", records)

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

	getFolders(currentDir)

	// jsonData, err := json.Marshal(folders)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Escribir el JSON en un archivo en el directorio actual
	// file, err := os.Create("folders.json")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()

	// _, err = file.Write(jsonData)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
