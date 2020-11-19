package smoke

import (
	"os/exec"
	"strings"

	"github.com/argoproj/pkg/errors"
)

// RunCmd is a function to run generic shell commands
func RunCmd(name string, args ...string) ([]byte, error) {

	output, error := exec.Command(name, args...).Output()
	errors.CheckError(error)

	return output, error

}

// OutputContains is a function that checks if a given output contains a specific expected string
func OutputContains(output string, expected string) bool {
	return strings.Contains(output, expected)
}
