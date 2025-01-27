package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

// Struct para organizar as informações necessárias de cada cidade
type WeatherInfo struct {
	numEntries int
	totalTemperature float64
	min float64
	max float64
}
// Método que retorna a temperatura média de cada cidade
func (w WeatherInfo) getMediumTemperature () float64{
	return w.totalTemperature / float64(w.numEntries)
}

func main() {
	// Grava o tempo do ínicio de execução
	inicio := time.Now()	

	// Abre o arquivo com os dados de cada estação de tempo
	file, err := os.Open("./measurements_1000000.csv")
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	// Criar um leitor CSV com delimitador personalizado
	reader := csv.NewReader(file)
	reader.Comma = ';'
	// Lê todas as linhas
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo CSV: %v\n", err)
		return
	}
	
	// hash map para manter os valores de cada cidade
	citiesMap := make(map[string]WeatherInfo)
	// Slice de strings que receberá o nome das
	// cidades para imprimi-las em ordem alfabética
	cityNames := []string{}

	// Itera sobre todas as linhas do arquivo
	for _, record := range records[1:] {
		// Armazena valores da linha
		cityName := record[0]
		temperature, _ := strconv.ParseFloat(record[1], 64)
		_, ok := citiesMap[cityName]
		if !ok {
			// Atualiza o slice
			cityNames = append(cityNames, cityName)
			// Cria uma estrutura com as informações da cidade
			w := WeatherInfo{
				numEntries: 1, 
				totalTemperature: temperature, 
				min: temperature, 
				max: temperature, 
			}
			// Armazena no hash map
			citiesMap[cityName] = w
			continue 
		}

		// Atualiza valores no hashMap
		city := citiesMap[cityName]
		city.numEntries++
		city.totalTemperature += temperature
		if (city.min > temperature) {
			city.min = temperature
		}
		if (city.max < temperature) {
			city.max = temperature
		}
		citiesMap[cityName] = city
	}
	// Print All cities info
	sort.Strings(cityNames)
	for _, key := range cityNames {
		medium := citiesMap[key].getMediumTemperature()
		fmt.Printf("%s=%.1f/%.1f/%.1f, ", 
			key, 
			citiesMap[key].min, 
			medium,
			citiesMap[key].max,
		)
	}
	fmt.Print("}")
	fmt.Println("Tempo de Execução: ", time.Since(inicio))
}