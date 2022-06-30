package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Person struct {
	Name   string
	Status string
}

// ByStatus реализует sort.Interface
// основанный на поле Status.
type ByStatus []Person

func (a ByStatus) Len() int { return len(a) }
func (a ByStatus) Less(i, j int) bool { // Less сообщает должен ли элемента с индексом i
	// быть отсортированным перед элементом с индексом j.
	if a[i].Status == "rat" && a[j].Status != "rat" { // rat woman/child man captain
		return true
	}
	if (a[i].Status == "woman" || a[i].Status == "child") && (a[j].Status == "man" || a[j].Status == "captain") {
		return true
	}
	if a[i].Status == "man" && a[j].Status == "captain" {
		return true
	}
	return false
}
func (a ByStatus) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main() {
	//
	//family := []Person{
	//	{"Alice", "rat"},
	//	{"Eve", "child"},
	//	{"Bob", "rat"},
	//	{"Karl", "captain"},
	//	{"Polik", "woman"},
	//	{"Mas", "child"},
	//	{"Masik", "man"},
	//}
	//sort.Sort(ByStatus(family))
	//fmt.Println(family)

	snumber, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	snumber = strings.Trim(snumber, "\n")

	number, err := strconv.Atoi(snumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(number) // remove

	people := make([]Person, number)

	for i := 0; i < number; i++ {
		fmt.Scan(&people[i].Name, &people[i].Status)
	}
	sort.Stable(ByStatus(people))
	fmt.Println(people)
}
