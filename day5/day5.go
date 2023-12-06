package day5

import (
	"fmt"
	"github.com/FSpruhs/advent-of-code-2023/util"
	"strconv"
	"strings"
)

const seedToSoilKey = "seed-to-soil"
const soilToFertilizerKey = "soil-to-fertilizer"
const fertilizerToWaterKey = "fertilizer-to-water"
const waterToLightKey = "water-to-light"
const lightToTemperatureKey = "light-to-temperature"
const temperatureToHumidityKey = "temperature-to-humidity"
const humidityToLocationKey = "humidity-to-location"

type mapperValue struct {
	destination int
	source      int
	rangeLength int
}

func (m mapper) mapInput(input int) int {
	for _, value := range m.value {
		match, mappedValue := value.useMapperValue(input)
		if match {
			return mappedValue
		}
	}
	return input
}

func (m mapperValue) useMapperValue(input int) (bool, int) {
	if m.source <= input && input < m.source+m.rangeLength {
		return true, m.destination + input - m.source
	}
	return false, -1
}

type mapper struct {
	value []mapperValue
}

func Solve(filePath string) {
	input := util.ReadFile(filePath)
	seeds := strings.Split((*input)[0], " ")[1:]
	mappers := readMapperStrings(*input)
	solutionPartOne := 0
	for _, seed := range seeds {
		location := mapAll(mappers, seed)
		if solutionPartOne == 0 {
			solutionPartOne = location
		}
		if solutionPartOne > location {
			solutionPartOne = location
		}
	}
	solutionPartTwo := 0
	for i := 0; i < len(seeds); i += 2 {
		for j := util.ConvertToInt(seeds[i]); j < util.ConvertToInt(seeds[i])+util.ConvertToInt(seeds[i+1]); j++ {
			location := mapAll(mappers, strconv.Itoa(j))
			if solutionPartTwo == 0 {
				solutionPartTwo = location
			}
			if solutionPartTwo > location {
				solutionPartTwo = location
			}
		}
	}
	fmt.Printf("Solution of day 5 part 1: %d\n", solutionPartOne)
	fmt.Printf("Solution of day 5 part 2: %d\n", solutionPartTwo)
}

func mapAll(mappers map[string]mapper, seed string) int {
	soil := mappers[seedToSoilKey].mapInput(util.ConvertToInt(seed))
	fertilizer := mappers[soilToFertilizerKey].mapInput(soil)
	water := mappers[fertilizerToWaterKey].mapInput(fertilizer)
	light := mappers[waterToLightKey].mapInput(water)
	temperature := mappers[lightToTemperatureKey].mapInput(light)
	humidity := mappers[temperatureToHumidityKey].mapInput(temperature)
	return mappers[humidityToLocationKey].mapInput(humidity)
}

func readMapperStrings(input []string) map[string]mapper {
	mapperSlices := readMapperSlices(input, readMapperStringStartPosition(input))
	return initMappers(mapperSlices)
}

func initMappers(mapperSlices map[string][]string) map[string]mapper {
	mappers := map[string]mapper{}
	for key, value := range mapperSlices {
		var newMapper = mapper{value: []mapperValue{}}
		for _, slice := range value {
			newMapper.value = append(newMapper.value, toMapperValue(slice))
		}
		mappers[key] = newMapper
	}
	return mappers
}

func toMapperValue(slice string) mapperValue {
	splittedSlice := strings.Split(slice, " ")
	return mapperValue{
		destination: util.ConvertToInt(splittedSlice[0]),
		source:      util.ConvertToInt(splittedSlice[1]),
		rangeLength: util.ConvertToInt(splittedSlice[2]),
	}
}

func readMapperSlices(input []string, mapperStartIndizes map[string]int) map[string][]string {
	mapperValues := map[string][]string{}
	mapperValues[seedToSoilKey] = readSlice(input, mapperStartIndizes, seedToSoilKey)
	mapperValues[soilToFertilizerKey] = readSlice(input, mapperStartIndizes, soilToFertilizerKey)
	mapperValues[fertilizerToWaterKey] = readSlice(input, mapperStartIndizes, fertilizerToWaterKey)
	mapperValues[waterToLightKey] = readSlice(input, mapperStartIndizes, waterToLightKey)
	mapperValues[lightToTemperatureKey] = readSlice(input, mapperStartIndizes, lightToTemperatureKey)
	mapperValues[temperatureToHumidityKey] = readSlice(input, mapperStartIndizes, temperatureToHumidityKey)
	mapperValues[humidityToLocationKey] = readSlice(input, mapperStartIndizes, humidityToLocationKey)
	return mapperValues
}

func readSlice(input []string, mapperStartIndizes map[string]int, key string) []string {
	return input[mapperStartIndizes[key]+1 : mapperStartIndizes[key]+findNextNewLine(input, mapperStartIndizes[key])]
}

func findNextNewLine(input []string, start int) int {
	for index, line := range input[start:] {
		if len(line) == 0 {
			return index
		}
	}
	return len(input) - start
}

func readMapperStringStartPosition(input []string) map[string]int {
	mapperStartIndizes := map[string]int{}
	for index, line := range input {
		if strings.Contains(line, seedToSoilKey) {
			mapperStartIndizes[seedToSoilKey] = index
		}
		if strings.Contains(line, soilToFertilizerKey) {
			mapperStartIndizes[soilToFertilizerKey] = index
		}
		if strings.Contains(line, fertilizerToWaterKey) {
			mapperStartIndizes[fertilizerToWaterKey] = index
		}
		if strings.Contains(line, waterToLightKey) {
			mapperStartIndizes[waterToLightKey] = index
		}
		if strings.Contains(line, lightToTemperatureKey) {
			mapperStartIndizes[lightToTemperatureKey] = index
		}
		if strings.Contains(line, temperatureToHumidityKey) {
			mapperStartIndizes[temperatureToHumidityKey] = index
		}
		if strings.Contains(line, humidityToLocationKey) {
			mapperStartIndizes[humidityToLocationKey] = index
		}
	}
	return mapperStartIndizes
}
