package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var a int = 4
var b int = 6

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/add", addCall)
	http.HandleFunc("/sub", subCall)
	http.HandleFunc("/div", divCall)
	http.HandleFunc("/mul", mulCall)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Do you want to provide values for a and b? (y/n): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "y" || choice == "Y" {
		fmt.Print("Enter values for a and b (space-separated): ")
		input, _ := reader.ReadString('\n')
		values := strings.Fields(input)
		if len(values) == 2 {
			a, _ = strconv.Atoi(values[0])
			b, _ = strconv.Atoi(values[1])
		} else {
			fmt.Println("Invalid input. Using default values for a and b.")
		}
	}

	server()
}

func server() {
	fmt.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Use /add /mul /sub /div to perform operations.")
}

func addCall(w http.ResponseWriter, r *http.Request) {
	result := a + b
	fmt.Fprintf(w, "%d + %d = %d", a, b, result)
}

func subCall(w http.ResponseWriter, r *http.Request) {
	result := a - b
	fmt.Fprintf(w, "%d - %d = %d", a, b, result)
}

func mulCall(w http.ResponseWriter, r *http.Request) {
	result := a * b
	fmt.Fprintf(w, "%d * %d = %d", a, b, result)
}

func divCall(w http.ResponseWriter, r *http.Request) {
	if b == 0 {
		fmt.Fprintf(w, "Division by zero is undefined")
		return
	}
	result := a / b
	fmt.Fprintf(w, "%d / %d = %d", a, b, result)
}
