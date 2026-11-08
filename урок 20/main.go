package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 1
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2
type Student struct {
	Name   string
	Grades []int
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	total := 0
	for _, grade := range s.Grades {
		total += grade
	}
	return float64(total) / float64(len(s.Grades))
}

type Person interface {
	Average() float64
}

// 3
type BankAccount struct {
	Owner   string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
	}
}

func (a *BankAccount) Withdraw(amount float64) {
	if amount > 0 && a.Balance >= amount {
		a.Balance -= amount
	}
}

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
}

// 4
type TextAnalyzer struct {
	Text string
}

func (a TextAnalyzer) CountWords() int {
	count := 0
	inWord := false
	for _, r := range a.Text {
		if r == ' ' || r == '\n' || r == '\t' || r == '\r' {
			inWord = false
		} else if !inWord {
			count++
			inWord = true
		}
	}
	return count
}

func (a TextAnalyzer) CountSentences() int {
	count := 0
	for _, r := range a.Text {
		if r == '.' || r == '!' || r == '?' {
			count++
		}
	}
	return count
}

func (a TextAnalyzer) CountVowels() int {
	const vowels = "аеёиоуыэюяАЕЁИОУЫЭЮЯaeiouAEIOU"
	count := 0
	for _, r := range a.Text {
		if unicode.IsLetter(r) && strings.ContainsRune(vowels, r) {
			count++
		}
	}
	return count
}

type Analyzer interface {
	CountWords() int
	CountSentences() int
	CountVowels() int
}

func main() {
	//1
	shapes := []Shape{
		Rectangle{Width: 4, Height: 3},
		Rectangle{Width: 5.5, Height: 2},
	}
	for i, shape := range shapes {
		fmt.Printf("Прямоугольник %d: площадь %g, периметр %g\n", i+1, shape.Area(), shape.Perimeter())
	}

	//2
	students := []Student{
		{Name: "Анна", Grades: []int{5, 4, 5}},
		{Name: "Иван", Grades: []int{3, 4, 4}},
		{Name: "Мария", Grades: []int{5, 5, 5}},
	}
	people := make([]Person, len(students))
	for i := range students {
		people[i] = students[i]
	}
	for i, person := range people {
		fmt.Printf("Студент %s: средний балл %.2f\n", students[i].Name, person.Average())
	}

	//3
	accounts := []*BankAccount{
		{Owner: "Алексей", Balance: 1000},
		{Owner: "Ольга", Balance: 500},
	}
	var account Account
	account = accounts[0]
	account.Deposit(250)
	account.Withdraw(100)
	account.Withdraw(2000)
	account = accounts[1]
	account.Deposit(300)
	account.Withdraw(150)
	for _, a := range accounts {
		fmt.Printf("Счёт %s: баланс %.2f\n", a.Owner, a.Balance)
	}

	//4
	var analyzer Analyzer = TextAnalyzer{Text: "Привет, мир! Как дела? Всё хорошо."}
	fmt.Printf("Слов: %d, предложений: %d, гласных: %d\n", analyzer.CountWords(), analyzer.CountSentences(), analyzer.CountVowels())
}
