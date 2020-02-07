package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func matchDNA(dnas string) string {
	temp := strings.Split(dnas, " ")
	pDna := temp[0]
	vDna := temp[1]

	letters := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	pLen := len(pDna)
	vLen := len(vDna)

	var comb []string

	comb = append(comb, vDna)

	for i := 0; i < vLen; i++ {
		for j := 0; j < len(letters); j++ {
			if letters[j] != string(vDna[i]) {
				if i == 0 {
					comb = append(comb, letters[j]+string(vDna[i+1:vLen]))
				} else if i == vLen-1 {
					comb = append(comb, string(vDna[0:vLen-1])+letters[j])
				} else {
					comb = append(comb, string(vDna[0:i])+letters[j]+string(vDna[i+1:vLen]))
				}
			}
		}
	}

	foundIndices := ""

	for i := 0; i <= pLen-vLen; i++ {
		foundIdx := -1
		pDnaSub := string(pDna[i : i+vLen])

		for j := 0; j < len(comb); j++ {
			if pDnaSub == comb[j] {
				foundIdx = i
				break
			}
		}

		if foundIdx != -1 {
			if foundIndices == "" {
				foundIndices = strconv.Itoa(foundIdx)
			} else {
				foundIndices = foundIndices + " " + strconv.Itoa(foundIdx)
			}
		}
	}

	if foundIndices == "" {
		foundIndices = "No Match!"
	}

	return foundIndices
}

func main() {
	files, err := ioutil.ReadDir("./input")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileBytes, err := ioutil.ReadFile("./input/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		sliceData := strings.Split(string(fileBytes), "\n")

		times, err := strconv.Atoi(sliceData[0])

		if err != nil {
			log.Fatal(err)
		}

		// create directory if not exist
		createDir("output")

		filename := "output" + string(f.Name()[5:len(f.Name())-4]) + ".txt"

		// create file if not exists
		createFile(filename)

		// delete file content
		deleteFileContent(filename)

		for i := 1; i <= times; i++ {
			matchDNA(sliceData[i])
			foundIndeces := matchDNA(sliceData[i])

			fl, err := os.OpenFile("./output/"+filename, os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}

			defer fl.Close()

			if _, err = fl.WriteString(foundIndeces + "\n"); err != nil {
				panic(err)
			}
		}
	}
}

func createDir(dir string) {
	_, errDr := os.Stat(dir)

	if os.IsNotExist(errDr) {
		errDir := os.MkdirAll(dir, 0755)
		if errDir != nil {
			panic(errDir)
		}
	}
}

func createFile(file string) {
	_, error := os.Stat("./output/" + file)
	if error != nil {
		if os.IsNotExist(error) {
			fcrt, err := os.Create("./output/" + file)
			if err != nil {
				panic(err)
			}

			fcrt.Close()

		} else {
			panic(error)
		}
	}
}

func deleteFileContent(file string) {
	empty := []byte("")
	salah := ioutil.WriteFile("./output/"+file, empty, 0644)
	if salah != nil {
		panic(salah)
	}
}
