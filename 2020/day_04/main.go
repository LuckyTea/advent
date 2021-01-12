package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr string // Birth Year
	iyr string // Issue Year
	eyr string // Expiration Year
	hgt string // Height
	hcl string // Hair Color
	ecl string // Eye Color
	pid string // Passport ID
	cid string // Country ID
}

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	var (
		countPart1 int
		countPart2 int
		p          passport
	)

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")

		for i := range fields {
			switch {
			case strings.HasPrefix(fields[i], "byr"):
				p.byr = strings.TrimPrefix(fields[i], "byr:")
			case strings.HasPrefix(fields[i], "iyr"):
				p.iyr = strings.TrimPrefix(fields[i], "iyr:")
			case strings.HasPrefix(fields[i], "eyr"):
				p.eyr = strings.TrimPrefix(fields[i], "eyr:")
			case strings.HasPrefix(fields[i], "hgt"):
				p.hgt = strings.TrimPrefix(fields[i], "hgt:")
			case strings.HasPrefix(fields[i], "hcl"):
				p.hcl = strings.TrimPrefix(fields[i], "hcl:")
			case strings.HasPrefix(fields[i], "ecl"):
				p.ecl = strings.TrimPrefix(fields[i], "ecl:")
			case strings.HasPrefix(fields[i], "pid"):
				p.pid = strings.TrimPrefix(fields[i], "pid:")
			case strings.HasPrefix(fields[i], "cid"):
				p.cid = strings.TrimPrefix(fields[i], "cid:")
			case fields[i] == "":
				if p.validate() {
					countPart1++
				}

				if p.validateV2() {
					countPart2++
				}

				p = passport{}
			}
		}
	}

	if p.validate() {
		countPart1++
	}

	if p.validateV2() {
		countPart2++
	}

	fmt.Println("answer part 1:", countPart1) // 230
	fmt.Println("answer part 2:", countPart2) // 156
}

func (p *passport) validate() bool {
	if p.byr == "" {
		return false
	}

	if p.iyr == "" {
		return false
	}

	if p.eyr == "" {
		return false
	}

	if p.hgt == "" {
		return false
	}

	if p.hcl == "" {
		return false
	}

	if p.ecl == "" {
		return false
	}

	if p.pid == "" {
		return false
	}

	return true
}

func (p *passport) validateV2() bool {
	if !validateRange(p.byr, 1920, 2002) {
		return false
	}

	if !validateRange(p.iyr, 2010, 2020) {
		return false
	}

	if !validateRange(p.eyr, 2020, 2030) {
		return false
	}

	if strings.HasSuffix(p.hgt, "cm") {
		if !validateRange(strings.TrimSuffix(p.hgt, "cm"), 150, 193) {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		if !validateRange(strings.TrimSuffix(p.hgt, "in"), 59, 76) {
			return false
		}
	} else {
		return false
	}

	if match, _ := regexp.MatchString("^#[0-9a-f]{6}$", p.hcl); !match {
		return false
	}

	if !isContains(eyeColors, p.ecl) {
		return false
	}

	if match, _ := regexp.MatchString("^[0-9]{9}$", p.pid); !match {
		return false
	}

	return true
}

func validateRange(val string, lower, upper int) bool {
	num, err := strconv.Atoi(val)
	return err == nil && num >= lower && num <= upper
}

func isContains(arr []string, val string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}

	return false
}
