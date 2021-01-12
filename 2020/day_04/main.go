package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		count int
		p     passport
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
					count++
				}

				p = passport{}
			}
		}
	}

	if p.validate() {
		count++
	}

	fmt.Println("count:", count) // 230
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
