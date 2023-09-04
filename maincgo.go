package main

/*
	#include <stdio.h>

    int createDb_c(char* db_name) {
        FILE *file = fopen(db_name, "r");

        if (file == NULL) {
            file = fopen(db_name, "w");
            if (file == NULL) {
                perror("Error al crear el archivo");
                return 1;
            }
            fclose(file);
            printf("La base de datos %s fue creada\n", db_name);
        } else {
            fclose(file);
        }
        return 0;
    }
*/
import "C"
import (
	"fmt"
	"os"
	"time"
)

func main() {

	inicio := time.Now()

	for i := 0; i < 100000; i++ {
		// createDb("db_example.json")
		C.createDb_c(C.CString("db_example.json"))
	}

	// Registra el tiempo de finalización
	final := time.Now()

	// Calcula la duración en milisegundos (ms)
	duracionMs := final.Sub(inicio).Milliseconds()

	// Calcula la duración en nanosegundos (ns)
	duracionNs := final.Sub(inicio).Nanoseconds()

	// Imprime la duración en ms y ns
	fmt.Printf("La función en CGO tomó %d ms en ejecutarse\n", duracionMs)
	fmt.Printf("La función en CGO tomó %d ns en ejecutarse\n", duracionNs)

}

func createDb(db_name string) {

	if _, err := os.Stat(db_name); os.IsNotExist(err) {
		file, err := os.Create(db_name)
		if err != nil {
			fmt.Println("Error al crear la db", err)
			return
		}
		defer file.Close()
		fmt.Printf("La base de datos %s fue creada", db_name)
	} else {
		// fmt.Printf("La base de datos %s ya existe", db_name)
	}
}
