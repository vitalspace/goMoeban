package main

import "C"
import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createDb(db_name string) {

	if !fileExists("db_example.json") {
		file, err := os.Create(db_name)
		if err != nil {
			fmt.Println("Error al crear la db", err)
			return
		}
		defer file.Close()
		fmt.Printf("La base de datos %s fue creada", db_name)
	} else {
		fmt.Println("ya existe")
	}
}

func Write(dbName, collection, newObject string) bool {
	// Leer el contenido del archivo JSON
	data, err := os.ReadFile(dbName)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return false
	}

	// Crear un mapa para representar el objeto JSON
	var root map[string]interface{}

	// Decodificar el contenido JSON en el mapa
	if err := json.Unmarshal(data, &root); err != nil {
		fmt.Println("Error al analizar el archivo JSON:", err)
		return false
	}

	// Verificar si la colección existe
	collectionObj, exists := root[collection]

	if !exists {
		// Si la colección no existe, crearla como un slice vacío
		root[collection] = []interface{}{}
		collectionObj = root[collection]
	}

	// Crear un objeto JSON a partir de la cadena proporcionada
	var newJSON interface{}
	if err := json.Unmarshal([]byte(newObject), &newJSON); err != nil {
		fmt.Println("Error al analizar el nuevo objeto JSON:", err)
		return false
	}

	// Agregar el nuevo objeto a la colección correspondiente
	collectionSlice, ok := collectionObj.([]interface{})
	if !ok {
		fmt.Println("Error: La colección no es un slice de JSON")
		return false
	}
	collectionSlice = append(collectionSlice, newJSON)
	root[collection] = collectionSlice

	// Codificar el mapa de nuevo como JSON
	updatedData, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		fmt.Println("Error al codificar el JSON actualizado:", err)
		return false
	}

	// Guardar los cambios en el archivo JSON
	if err := os.WriteFile(dbName, updatedData, os.ModePerm); err != nil {
		fmt.Println("Error al escribir el archivo JSON:", err)
		return false
	}

	return true
}

func main() {

	inicio := time.Now()

	// createDb("db_example.json")

	dbName := "db_example.json"
	collection := "example_collection"
	newObject := `{"name": "John", "age": 30}`

	if Write(dbName, collection, newObject) {
		fmt.Println("Objeto agregado exitosamente.")
	} else {
		fmt.Println("No se pudo agregar el objeto.")
	}

	// Registra el tiempo de finalización
	final := time.Now()

	// Calcula la duración en milisegundos (ms)
	duracionMs := final.Sub(inicio).Milliseconds()

	// Calcula la duración en nanosegundos (ns)
	duracionNs := final.Sub(inicio).Nanoseconds()

	// Imprime la duración en ms y ns
	fmt.Printf("La función GO tomó %d ms en ejecutarse\n", duracionMs)
	fmt.Printf("La función GO tomó %d ns en ejecutarse\n", duracionNs)

}
