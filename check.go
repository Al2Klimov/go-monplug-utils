package go_monplug_utils

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func ExecuteCheck(
	onTerminal func() (output string),
	check func() (output string, perfdata PerfdataCollection, errors map[string]error),
) (exit int) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Print(onTerminal())
		return 3
	}

	output, perfdata, errsChk := check()
	if errsChk != nil {
		for context, err := range errsChk {
			fmt.Printf("%s: %s\n", context, err.Error())
		}

		return 3
	}

	if _, errFP := fmt.Print(output + perfdata.String()); errFP != nil {
		return 3
	}

	return int(perfdata.calcStatus())
}
