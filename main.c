#include <stdio.h>
#include <time.h>

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

int main() {
    struct timespec start_time, end_time;
    long long tiempo_ns, tiempo_ms;

    // Registra el tiempo de inicio
    clock_gettime(CLOCK_MONOTONIC, &start_time);

    // Llama a tu función

    for (size_t i = 0; i < 100000; i++)
    {
        /* code */
        createDb_c("db_example.json");
    }
    
    // Registra el tiempo de finalización
    clock_gettime(CLOCK_MONOTONIC, &end_time);

    // Calcula el tiempo transcurrido en nanosegundos
    tiempo_ns = (end_time.tv_sec - start_time.tv_sec) * 1000000000 + (end_time.tv_nsec - start_time.tv_nsec);

    // Calcula el tiempo transcurrido en milisegundos
    tiempo_ms = tiempo_ns / 1000000;

    printf("Tiempo transcurrido en C: %lld ms\n", tiempo_ms);
    printf("Tiempo transcurrido en C: %lld ns\n", tiempo_ns);

    return 0;
}