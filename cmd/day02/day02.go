package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Kreonn/AoC2020/internal/tools"
)

// Policy represents the policy of a password
type Policy struct {
	min  int
	max  int
	char string
}

// Password represents a password with its policy
type Password struct {
	policy Policy
	text   string
}

// IsValid1 returns the validity of the password for part 1
func (p *Password) IsValid1() bool {
	var total int

	for _, char := range p.text {
		if string(char) == p.policy.char {
			total++
		}
	}

	return total <= p.policy.max && total >= p.policy.min
}

// IsValid2 returns the validity of the password for part 2
func (p *Password) IsValid2() bool {
	chars := strings.Split(p.text, "")

	first := chars[p.policy.min-1]
	second := chars[p.policy.max-1]
	return first != second && (first == p.policy.char || second == p.policy.char)
}

// ParsePassword converts a string to Password type
func ParsePassword(passtr string) *Password {
	var password Password
	var err error

	res := strings.Split(passtr, ": ")
	password.text = res[1]

	res = strings.Split(res[0], " ")
	password.policy.char = res[1]

	res = strings.Split(res[0], "-")
	password.policy.min, err = strconv.Atoi(res[0])
	if err != nil {
		panic(err)
	}
	password.policy.max, err = strconv.Atoi(res[1])
	if err != nil {
		panic(err)
	}

	return &password
}

func main() {
	data, err := tools.GetData("day02.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var passwords []*Password
	for _, pass := range data.Lines {
		passwords = append(passwords, ParsePassword(pass))
	}

	var total1 int
	for _, password := range passwords {
		if password.IsValid1() {
			total1++
		}
	}
	fmt.Printf("Result for part 1 : %d\n", total1)

	var total2 int
	for _, password := range passwords {
		if password.IsValid2() {
			total2++
		}
	}
	fmt.Printf("Result for part 2 : %d\n", total2)
}
