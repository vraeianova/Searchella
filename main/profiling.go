package main

import (
	"fmt"
)

func main() {
	// Iniciar el perfil de CPU
	// f, _ := os.Create("cpu.prof")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	for i := 0; i < 10; i++ {
		// Código a ejecutar en cada iteración del bucle
		fmt.Println("testing", i+1)
	}
}
