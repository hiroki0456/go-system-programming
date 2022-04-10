package main

import "os"

func main() {
	file, err := os.Open("Purine_Nucleoside_Phosphorylase.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
