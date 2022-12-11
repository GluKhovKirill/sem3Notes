package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type People struct {
	ID   int
	Name string
}
type Note struct {
	name    string
	surname string
	text    string
}

func clearShell() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type strA struct {
	name    string
	surname string
}

type strB struct {
	name string
	age  int
}

func main() {

	var name, surname, text, isContinue string

	scanner := bufio.NewScanner(os.Stdin)
	var sl []Note
	flag := false
	for true {
		if flag {
			break
		}
		clearShell()
		fmt.Println("Введите имя:")
		scanner.Scan()
		name = scanner.Text()
		fmt.Println("Введите фамилию:")
		scanner.Scan()
		surname = scanner.Text()
		fmt.Println("Введите текст:")
		scanner.Scan()
		text = scanner.Text()
		clearShell()
		fmt.Println("Surname: "+surname+"\nName: "+name+"\ntext: ", text)
		row := Note{name, surname, text}
		sl = append(sl, row)

		for true {
			fmt.Print("Продолжить?\ny-да\tn-nope\ta-показать все\n>>")
			scanner.Scan()
			isContinue = scanner.Text()

			if isContinue == "n" {
				flag = true
				break
			} else if isContinue == "a" {
				clearShell()
				for i := 0; i < len(sl); i++ {
					note := sl[i]
					fmt.Println("Surname: "+note.surname+"\nName: "+note.name+"\ntext: ", note.text)
					fmt.Println("--------------------------")

				}
			} else if isContinue == "y" {
				break
			}

		}

	}
	clearShell()
	for i := 0; i < len(sl); i++ {
		note := sl[i]
		fmt.Println("Surname: "+note.surname+"\nName: "+note.name+"\ntext: ", note.text)
		fmt.Println("--------------------------")

	}

	//fmt.Println(surname + ", " + name + ": " + text)

}
