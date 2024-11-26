package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
	"os/exec"
)

type Player struct {
	p1, p2 rune
	num    int
}

var player Player

func InitArray() [7][7]rune {
	var GameArray [7][7]rune
	for i := 0; i < len(GameArray); i++ {
		for j := 0; j < len(GameArray[i]); j++ {
			GameArray[i][j] = ' ' // fill the array with spaces
		}
	}
	return GameArray
}

func DoGrid(GameArray [7][7]rune) {
	abs := 0
	for arrI := 0; arrI < 13; arrI++ {
		if arrI%2 == 0 {
			for arrJ := 0; arrJ < len(GameArray); arrJ++ {
				fmt.Print("|")
				fmt.Print(" ")
				if GameArray[abs][arrJ] == player.p1 {
					fmt.Print("\033[33m" + string(GameArray[abs][arrJ]) + "\033[0m") // prints the X in yellow into the grid
				} else if GameArray[abs][arrJ] == player.p2 {
					fmt.Print("\033[35m" + string(GameArray[abs][arrJ]) + "\033[0m") // prints the O in magenta into the grid
				} else if GameArray[abs][arrJ] == ' ' {
					fmt.Print(" ")
				}
				fmt.Print(" ")
			}
			fmt.Print("|")
			fmt.Print(" ")
			fmt.Print("\n")
			abs++
		} else {
			for dash := 0; dash < 29; dash++ {
				fmt.Print("-")
			}
			fmt.Print("\n")
		}
	}
}

func Gravity(GameArray [7][7]rune, posX int) int {
	for i := len(GameArray) - 1; i >= 0; i-- {
		if GameArray[i][posX-1] == ' ' {
			return i // return the y position of the free location on the column
		}
	}
	return -1 // if the column does not have any free spaces left
}

func FillingTab(GameArray [7][7]rune, player Player) [7][7]rune {
	var posX int
	var element rune
	if player.num%2 == 0 {
		element = player.p1
	} else if player.num%2 == 1 {
		element = player.p2
	}
	fmt.Printf("Player %c's turn. Please enter a number between 1 and 7 : \n", element)
	_, err := fmt.Scan(&posX)
	if err != nil {
		return GameArray
	}
	if posX == -1 { // Writing -1 when the player is asked to write restarts the game
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("\033[32m" + "The game has been reinitialised" + "\033[0m")
		Game()
	}
	if posX == -2 { // Writing -2 when the player is asked to write go back to the main menu
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		Menu()
	}
	z01.PrintRune('\n')
	if int(posX) <= 7 && int(posX) > 0 {
		posY := Gravity(GameArray, posX)
		if posY == -1 {
			DoGrid(GameArray)
			fmt.Println("\033[31m" + "Oh snap! There's not any space left on this column" + "\033[0m") // ANSI code for the red colour, then prints the content of the error, ANSI code for the white colour
			return FillingTab(GameArray, player)
		}
		GameArray[posY][posX-1] = element
		return GameArray
	} else {
		DoGrid(GameArray)
		fmt.Println("\033[31m" + "Oh snap! The number you've written is not between 1 and 7" + "\033[0m") // ANSI code for the red colour, then prints the content of the error, ANSI code for the white colour
		return FillingTab(GameArray, player)
	}
}

func VictoryCheck(GameArray [7][7]rune, player Player, num int) int {
	occurrences := 0
	if IsDraw(GameArray) == true {
		fmt.Println("Draw")
		return 1
	}
	for i := len(GameArray) - 1; i >= 0; i-- { // Row check
		for j := 0; j < len(GameArray[i]); j++ {
			if num%2 == 0 {
				if GameArray[i][j] == player.p1 {
					occurrences++
				} else if GameArray[i][j] == player.p2 || GameArray[i][j] == ' ' {
					occurrences = 0
				}
				if occurrences == 4 {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 1 has won horizontally" + "\033[0m")
					return 0
				}
			} else {
				if GameArray[i][j] == player.p2 {
					occurrences++
				} else if GameArray[i][j] == player.p1 || GameArray[i][j] == ' ' {
					occurrences = 0
				}
				if occurrences == 4 {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 2 has won horizontally" + "\033[0m")
					return 0
				}
			}
		}
		occurrences = 0
	}
	for i := len(GameArray) - 1; i >= 0; i-- { // Column check
		for j := 0; j < len(GameArray[i]); j++ {
			if num%2 == 0 {
				if GameArray[j][i] == player.p1 {
					occurrences += 1
				} else if GameArray[j][i] == player.p2 || GameArray[i][j] == ' ' {
					occurrences = 0
				}
				if occurrences == 4 {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 1 has won vertically" + "\033[0m")
					return 0
				}
			} else {
				if GameArray[j][i] == player.p2 {
					occurrences += 1
				} else if GameArray[j][i] == player.p1 || GameArray[i][j] == ' ' {
					occurrences = 0
				}
				if occurrences == 4 {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 2 has won vertically" + "\033[0m")
					return 0
				}
			}
		}
		occurrences = 0
	}
	for i := len(GameArray) - 1; i >= 0; i-- { // Diagonal check (bottom left to top right)
		for j := 0; j < len(GameArray[i]); j++ {
			if i >= 3 && j < len(GameArray[i])-4 {
				if player.p1 == GameArray[i][j] && player.p1 == GameArray[i-1][j+1] && player.p1 == GameArray[i-2][j+2] && player.p1 == GameArray[i-3][j+3] {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 1 has won diagonally" + "\033[0m")
					return 0
				}
				if player.p2 == GameArray[i][j] && player.p2 == GameArray[i-1][j+1] && player.p2 == GameArray[i-2][j+2] && player.p2 == GameArray[i-3][j+3] {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 2 has won diagonally" + "\033[0m")
					return 0
				}
			}
		}
	}
	for i := len(GameArray) - 1; i >= 0; i-- { // Diagonal check (bottom right to top left)
		for j := len(GameArray[i]) - 1; j >= 0; j-- {
			if i >= 3 && j > 3 {
				if player.p1 == GameArray[i][j] && player.p1 == GameArray[i-1][j-1] && player.p1 == GameArray[i-2][j-2] && player.p1 == GameArray[i-3][j-3] {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 1 has won diagonally" + "\033[0m")
					return 0
				}
				if player.p2 == GameArray[i][j] && player.p2 == GameArray[i-1][j-1] && player.p2 == GameArray[i-2][j-2] && player.p2 == GameArray[i-3][j-3] {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 2 has won diagonally" + "\033[0m")
					return 0
				}
			}
		}
	}
	for i := 0; i < len(GameArray); i++ { // Diagonal check (top right to bottom left)
		for j := len(GameArray) - 1; j >= 0; j-- {
			if i < len(GameArray[i])-4 && j > 3 {
				if player.p1 == GameArray[i][j] && player.p1 == GameArray[i+1][j-1] && player.p1 == GameArray[i+2][j-2] && player.p1 == GameArray[i+3][j-3] {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 1 has won diagonally" + "\033[0m")
					return 0
				}
				if player.p2 == GameArray[i][j] && player.p2 == GameArray[i+1][j-1] && player.p2 == GameArray[i+2][j-2] && player.p2 == GameArray[i+3][j-3] {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 2 has won diagonally" + "\033[0m")
					return 0
				}
			}
		}
	}
	for i := 0; i < len(GameArray); i++ { // Diagonal check (top left to bottom right)
		for j := 0; j < len(GameArray[i]); j++ {
			if i < len(GameArray)-4 && j < len(GameArray[i])-4 {
				if player.p1 == GameArray[i][j] && player.p1 == GameArray[i+1][j+1] && player.p1 == GameArray[i+2][j+2] && player.p1 == GameArray[i+3][j+3] {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 1 has won diagonally" + "\033[0m")
					return 0
				}
				if player.p2 == GameArray[i][j] && player.p2 == GameArray[i+1][j+1] && player.p2 == GameArray[i+2][j+2] && player.p2 == GameArray[i+3][j+3] {
					fmt.Println("\033[32m" + "\n\n\n\nPlayer 2 has won diagonally" + "\033[0m")
					return 0
				}
			}
		}
	}
	return -1
}

func IsDraw(GameArray [7][7]rune) bool {
	for i := 0; i < len(GameArray); i++ {
		for j := 0; j < len(GameArray[i]); j++ {
			if GameArray[i][j] == ' ' { // checks if there is any spaces left on the grid
				return false
			}
		}
	}
	return true
}

func Menu() {
	var choice string
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run() // Clears the terminal (lines 238-240)
	mainMenuAscii := "  __  __       _                                    \n |  \\/  |     (_)                                   \n | \\  / | __ _ _ _ __    _ __ ___   ___ _ __  _   _ \n | |\\/| |/ _` | | '_ \\  | '_ ` _ \\ / _ \\ '_ \\| | | |\n | |  | | (_| | | | | | | | | | | |  __/ | | | |_| |\n |_|  |_|\\__,_|_|_| |_| |_| |_| |_|\\___|_| |_|\\__,_|\n                                                    \n                                                    "
	fmt.Print(mainMenuAscii)
	fmt.Println("\n" + "\033[32m" + "start" + "\033[0m" + " : launch a game")
	fmt.Println("\033[31m" + "quit / exit" + "\033[0m" + " : exit the game")
	fmt.Printf("\nWrite below which action you want to go :\n")
	_, err := fmt.Scan(&choice)
	if err != nil {
		main()
	}
	if choice == "start" {
		Game() // launches the game
	}
	if choice == "quit" || choice == "exit" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()  // Clears the terminal (lines 254-256)
		os.Exit(3) // Exit the program
	}

}

func Game() {
	var endGame string
	slice := InitArray()
	for {
		DoGrid(slice)
		if VictoryCheck(slice, player, player.num-1) != -1 {
			DoGrid(slice)
			fmt.Println("Enter anything to go back to the main menu :\n")
			_, err := fmt.Scan(&endGame)
			if err != nil {
				Menu()
			}
			Menu()
		}
		slice = FillingTab(slice, player)
		player.num += 1
	}
}

func main() {
	player.num = 0
	player.p1 = 'X'
	player.p2 = 'O'
	Menu()
}
