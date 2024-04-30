package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Status struct {
	enduranceCounter int  //countdown
	restCounter      int  //countdown
	flying           bool //countdown
	resting          bool //countdown
	distance         int  //var
	points           int  //var
}

type Reindeer struct {
	speed     int //const
	endurance int //const
	rest      int //const
	status    Status
}

type Reindeers map[string]Reindeer

func readStrings(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, err
}

func (r Reindeers) parseReindeer(s string) {
	var name string
	reindeer := Reindeer{}

	fmt.Sscanf(s, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &reindeer.speed, &reindeer.endurance, &reindeer.rest)

	reindeer.status.enduranceCounter = reindeer.endurance
	reindeer.status.flying = true
	r[name] = reindeer

}

func (r *Reindeers) secondLoop() {
	for name, reindeer := range *r {
		if reindeer.status.flying {
			reindeer.status.distance += reindeer.speed
			reindeer.status.enduranceCounter--
			if reindeer.status.enduranceCounter == 0 {
				reindeer.status.flying = false
				reindeer.status.resting = true
				reindeer.status.restCounter = reindeer.rest
			}
		} else if reindeer.status.resting {
			reindeer.status.restCounter--
			if reindeer.status.restCounter == 0 {
				reindeer.status.flying = true
				reindeer.status.resting = false
				reindeer.status.enduranceCounter = reindeer.endurance
			}
		}
		(*r)[name] = reindeer
	}
}

func (r *Reindeers) calculatePoints() {
	distance := map[string]int{}
	var name string
	var reindeer Reindeer

	for name, reindeer = range *r {
		distance[name] = reindeer.status.distance
	}

	distanceSlice := make([]int, 0)
	for _, reindeer = range *r {
		distanceSlice = append(distanceSlice, reindeer.status.distance)
	}

	for name, reindeer = range *r {
		if reindeer.status.distance == slices.Max(distanceSlice) {
			reindeer.status.points++
			(*r)[name] = reindeer
		}
	}
}

func (r Reindeers) printStatus() {
	for name, reindeer := range r {
		if reindeer.status.resting {
			fmt.Printf("%s is resting for %d more seconds and has gone %d km\n", name, reindeer.status.restCounter, reindeer.status.distance)
		} else if reindeer.status.flying {
			fmt.Printf("%s is flying for %d more seconds and has gone %d km\n", name, reindeer.status.enduranceCounter, reindeer.status.distance)
		}
	}
	fmt.Println("")
}

func main() {
	input, err := readStrings("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	reindeers := Reindeers{}

	for _, s := range input {
		reindeers.parseReindeer(s)
	}
	for i := 1; i <= 2503; i++ {
		reindeers.secondLoop()
		reindeers.calculatePoints()
	}
	reindeers.printStatus()

	maxDistance := 0
	maxPoints := 0
	for _, reindeer := range reindeers {
		if reindeer.status.distance > maxDistance {
			maxDistance = reindeer.status.distance
		}
		if reindeer.status.points > maxPoints {
			maxPoints = reindeer.status.points
		}
	}

	fmt.Printf("The furthest distance travelled is %d\n", maxDistance)
	fmt.Printf("The most points earned is %d\n", maxPoints)
}
