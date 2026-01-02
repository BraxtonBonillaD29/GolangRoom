package main

import "fmt"

func main() {
	fmt.Println("Paso 2: Creando canales de trabajo")

	jobs := make(chan int, 10) //buffer de 10

	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Println("Enviando trabajo:", i)
	}

	close(jobs)

	fmt.Println(jobs)
	fmt.Println("Canal de trabajos creado y cerrado")
}
