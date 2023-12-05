// Package day05 is our package for the 5th AoC day
package day05

import (
	"fmt"
	"strconv"
	"strings"
)

// RunPart1 is for the first star of the day
func RunPart1(input []string) int {
	minSeed := -1
	currentMap := ""
	var seeds []int
	var seedToSoil [][]int
	var soilToFertilizer [][]int
	var fertilizerToWater [][]int
	var waterToLight [][]int
	var lightToTemperature [][]int
	var temperatureToHumidity [][]int
	var humidityToLocation [][]int
	for _, line := range input {
		if line == "" {
			continue
		}
		_, isNumber := strconv.Atoi(line[0:1])
		if isNumber != nil {
			currentMap = strings.Split(line, " ")[0]
			if currentMap != "seeds:" {
				continue
			}
		}
		switch currentMap {
		case "seeds:":
			seedLine := strings.Split(line, ":")
			seedLine = strings.Split(strings.TrimSpace(seedLine[1]), " ")
			for _, v := range seedLine {
				number, _ := strconv.Atoi(v)
				seeds = append(seeds, number)
			}
		case "seed-to-soil":
			soilLine := strings.Split(line, " ")
			var newMap []int
			for _, v := range soilLine {
				number, _ := strconv.Atoi(v)
				newMap = append(newMap, number)
			}
			seedToSoil = append(seedToSoil, newMap)
		case "soil-to-fertilizer":
			fertilizer := strings.Split(line, " ")
			var newMap []int
			for _, v := range fertilizer {
				number, _ := strconv.Atoi(v)
				newMap = append(newMap, number)
			}
			soilToFertilizer = append(soilToFertilizer, newMap)
		case "fertilizer-to-water":
			waterLine := strings.Split(line, " ")
			var newMap []int
			for _, v := range waterLine {
				number, _ := strconv.Atoi(v)
				newMap = append(newMap, number)
			}
			fertilizerToWater = append(fertilizerToWater, newMap)
		case "water-to-light":
			lightLine := strings.Split(line, " ")
			var newMap []int
			for _, v := range lightLine {
				number, _ := strconv.Atoi(v)
				newMap = append(newMap, number)
			}
			waterToLight = append(waterToLight, newMap)
		case "light-to-temperature":
			temperatureLine := strings.Split(line, " ")
			var newMap []int
			for _, v := range temperatureLine {
				number, _ := strconv.Atoi(v)
				newMap = append(newMap, number)
			}
			lightToTemperature = append(lightToTemperature, newMap)
		case "temperature-to-humidity":
			humidityLine := strings.Split(line, " ")
			var newMap []int
			for _, v := range humidityLine {
				number, _ := strconv.Atoi(v)
				newMap = append(newMap, number)
			}
			temperatureToHumidity = append(temperatureToHumidity, newMap)
		case "humidity-to-location":
			locationLine := strings.Split(line, " ")
			var newMap []int
			for _, v := range locationLine {
				number, _ := strconv.Atoi(v)
				newMap = append(newMap, number)
			}
			humidityToLocation = append(humidityToLocation, newMap)
		}
	}
	for _, seed := range seeds {
		for _, soil := range seedToSoil {
			dest := soil[0]
			source := soil[1]
			length := soil[2]
			if seed >= source && seed <= source+length {
				seed += dest - source
				break
			}
		}
		for _, fertilizer := range soilToFertilizer {
			dest := fertilizer[0]
			source := fertilizer[1]
			length := fertilizer[2]
			if seed >= source && seed <= source+length {
				seed += dest - source
				break
			}
		}
		for _, water := range fertilizerToWater {
			dest := water[0]
			source := water[1]
			length := water[2]
			if seed >= source && seed <= source+length {
				seed += dest - source
				break
			}
		}
		for _, light := range waterToLight {
			dest := light[0]
			source := light[1]
			length := light[2]
			if seed >= source && seed <= source+length {
				seed += dest - source
				break
			}
		}
		for _, temperature := range lightToTemperature {
			dest := temperature[0]
			source := temperature[1]
			length := temperature[2]
			if seed >= source && seed <= source+length {
				seed += dest - source
				break
			}
		}
		for _, humidity := range temperatureToHumidity {
			dest := humidity[0]
			source := humidity[1]
			length := humidity[2]
			if seed >= source && seed <= source+length {
				seed += dest - source
				break
			}
		}
		for _, location := range humidityToLocation {
			dest := location[0]
			source := location[1]
			length := location[2]
			if seed >= source && seed <= source+length {
				seed += dest - source
				break
			}
		}
		if minSeed == -1 || minSeed > seed {
			minSeed = seed
		}
	}
	return minSeed
}

// RunPart2 is for the second star of the day
func RunPart2(input []string) int {
	for _, line := range input {
		fmt.Println(line)
	}
	return 0
}
