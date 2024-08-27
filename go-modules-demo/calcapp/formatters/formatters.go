package formatters

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func Red(message string) {
	fmt.Println(
		chalk.Red,
		message,
		chalk.ResetColor,
	)

}

func Green(message string) {
	fmt.Println(
		chalk.Green,
		message,
		chalk.ResetColor,
	)
}
