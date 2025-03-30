package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/dustin/go-humanize"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(`
██████╗ ██╗   ██╗██████╗  ██████╗ ███████╗████████╗     ██████╗ █████╗ ██╗      ██████╗██╗   ██╗██╗      █████╗ ████████╗ ██████╗ ██████╗ 
██╔══██╗██║   ██║██╔══██╗██╔════╝ ██╔════╝╚══██╔══╝    ██╔════╝██╔══██╗██║     ██╔════╝██║   ██║██║     ██╔══██╗╚══██╔══╝██╔═══██╗██╔══██╗
██████╔╝██║   ██║██║  ██║██║  ███╗█████╗     ██║       ██║     ███████║██║     ██║     ██║   ██║██║     ███████║   ██║   ██║   ██║██████╔╝
██╔══██╗██║   ██║██║  ██║██║   ██║██╔══╝     ██║       ██║     ██╔══██║██║     ██║     ██║   ██║██║     ██╔══██║   ██║   ██║   ██║██╔══██╗
██████╔╝╚██████╔╝██████╔╝╚██████╔╝███████╗   ██║       ╚██████╗██║  ██║███████╗╚██████╗╚██████╔╝███████╗██║  ██║   ██║   ╚██████╔╝██║  ██║
╚═════╝  ╚═════╝ ╚═════╝  ╚═════╝ ╚══════╝   ╚═╝        ╚═════╝╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ ╚══════╝╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝`)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select the percentages that you want to use")

		necAmt := readAndConvert("Percentage for Necessities: ")
		wantsAmt := readAndConvert("Percentage for Wants: ")
		savingsAmt := readAndConvert("Percentage for Savings: ")
		incomeAmt := readAndConvert("Incomes: ")

		necPer, wantPer, savPer, err := makeCalculation(necAmt, wantsAmt, savingsAmt, incomeAmt)
		if err != nil {
			println(err.Error())
			continue
		}

		fmt.Println("Your percentages are: ")

		fmt.Printf("%-20s%-20v%-20s\n", "------------------", "-----------", "------")
		fmt.Printf("%-20s%-20s%-20s\n", "Category", "Percentage", "Amount")
		fmt.Printf("%-20s%-20v%-20s\n", "------------------", "-----------", "------")
		fmt.Printf("%-20s%%%-20v%-20s\n", "Necessities", necAmt, humanize.Commaf(necPer))
		fmt.Printf("%-20s%%%-20v%-20s\n", "Wants", wantsAmt, humanize.Commaf(wantPer))
		fmt.Printf("%-20s%%%-20v%-20s\n", "Savings", savingsAmt, humanize.Commaf(savPer))
		fmt.Printf("%-20s%-20v%-20s\n", "------------------", "-----------", "------")

		fmt.Println("\nDo you want to calculate a different income? (y/n)")
		restartProgram, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if restartProgram = strings.TrimSuffix(restartProgram, "\n"); restartProgram != "y" {
			println("Have a nice day 💙")
			return
		}

	}

}

func makeCalculation(nec, want, sav, income int) (necPer, wantPer, savPer float64, err error) {
	if nec+want+sav != 100 {
		return 0.0, 0.0, 0.0, errors.New("Total of percentages cannot be more than 100%, please try again 🫡")
	}
	necPer = float64(income) * (float64(nec) / 100.0)
	wantPer = float64(income) * (float64(want) / 100.0)
	savPer = float64(income) * (float64(sav) / 100.0)

	return necPer, wantPer, savPer, nil
}

func readAndConvert(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSuffix(input, "\n")
	value, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Invalid input")
	}
	return value
}
