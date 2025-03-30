package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dustin/go-humanize"
)

func main() {
	fmt.Println(`
██████  ██    ██ ██████   ██████  ███████ ████████     ███████  ██████ ██   ██ ███████ ██████  ██    ██ ██      ███████ ██████  
██   ██ ██    ██ ██   ██ ██       ██         ██        ██      ██      ██   ██ ██      ██   ██ ██    ██ ██      ██      ██   ██ 
██████  ██    ██ ██   ██ ██   ███ █████      ██        ███████ ██      ███████ █████   ██   ██ ██    ██ ██      █████   ██████  
██   ██ ██    ██ ██   ██ ██    ██ ██         ██             ██ ██      ██   ██ ██      ██   ██ ██    ██ ██      ██      ██   ██ 
██████   ██████  ██████   ██████  ███████    ██        ███████  ██████ ██   ██ ███████ ██████   ██████  ███████ ███████ ██   ██`)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Select the percentages that you want to use")
		fmt.Println("Percentage for Necessities: ")

		necAmt, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Percentage for Wants: ")
		wantsAmt, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Percentage for Savings / Debt: ")
		savingsAmt, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Income: ")
		incomeAmt, err := reader.ReadString('\n')

		necAmtInt, err := strconv.Atoi(necAmt[:len(necAmt)-1])
		wantsAmtInt, err := strconv.Atoi(wantsAmt[:len(wantsAmt)-1])
		savingsAmtInt, err := strconv.Atoi(savingsAmt[:len(savingsAmt)-1])
		incomeAmtInt, err := strconv.Atoi(incomeAmt[:len(incomeAmt)-1])

		necPer, wantPer, savPer := makeCalculation(necAmtInt, wantsAmtInt, savingsAmtInt, incomeAmtInt)

		fmt.Println("Your percentages are: ")

		fmt.Printf(`
	Necessities: %s,
	Wants: %s,
	Savings / Debt: %s,
    Income: %s,
`, humanize.Commaf(necPer), humanize.Commaf(wantPer), humanize.Commaf(savPer), humanize.Comma(int64(incomeAmtInt)))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Do you want to calculate a different income? (y/n)")
		restartProgram, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if restartProgram[:1] != "y" {
			println("Have a nice day :)")
			return
		}

	}

}

func makeCalculation(nec, want, sav, income int) (necPer, wantPer, savPer float64) {
	necPer = float64(income) * (float64(nec) / 100.0)
	wantPer = float64(income) * (float64(want) / 100.0)
	savPer = float64(income) * (float64(sav) / 100.0)

	return
}
