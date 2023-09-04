#!/bin/bash

# Nombre del archivo Go que deseas compilar y ejecutar
ARCHIVO_CGO="maincgo.go"
ARCHIVO_GO="main.go"
ARCHIVO_C="main.c"

# Compilar el archivo Go
go build -o main_cgo.exe $ARCHIVO_CGO
go build -o main_go.exe $ARCHIVO_GO
gcc -o main_c.exe $ARCHIVO_C

# Verificar si la compilación fue exitosa
if [ $? -eq 0 ]; then
    echo "Compilación exitosa. Ejecutando el programa..."
    # Ejecutar el programa compilado
    ./main_cgo.exe
    ./main_go.exe
    ./main_c.exe
else
    echo "Error de compilación."
fi
