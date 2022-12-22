package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {
	// Abre el archivo
	file, err := os.Open("2.")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Crea un slice de Record para almacenar los registros
	var records []Record

	var currentRecord Record

	record := Record{}

	// Crea un escáner para leer el archivo línea a línea
	scanner := bufio.NewScanner(file)

	// Lee cada línea del archivo
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Verificando la línea", line)
		// Divide la línea en clave y valor
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			fmt.Println("Verificar que entro aqui en esta linea", line)
			// La línea no tiene el formato esperado (clave: valor)
			// Añade el valor de la línea al campo Message
			currentRecord.Message += line + "\n"
			continue
		}
		key := parts[0]
		value := parts[1]

		// Crea una instancia de Record y asigna los valores de clave y valor a los campos correspondientes
		// record := Record{}
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

		// Agrega la instancia de Record al slice

	}
	record.Message = currentRecord.Message
	records = append(records, record)

	// fmt.Println("Verificando la salida de records", records)
	// fmt.Println("Verificando la salida de currentRecord", currentRecord.Message)
	fmt.Println("Verificando la salida de records con el mensaje", records)

	// Verifica si hubo algún error al leer el archivo
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	output := Output{
		Index:   "messages",
		Records: records,
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Exporta la cadena de caracteres de formato JSON a un archivo
	err = ioutil.WriteFile("output.json", outputJSON, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
