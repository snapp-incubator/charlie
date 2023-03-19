package runtime

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/amirhnajafiz/DJaaS/internal/component/filesystem"
)

type Runtime struct {
	Cfg    Config
	System filesystem.System
}

func (r *Runtime) Process(operator, path, input string) error {
	cmd := exec.Command(operator, path)
	cmd.Stdin = strings.NewReader(input)

	var out strings.Builder

	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	result := out.String()
	result = strings.Trim(result, "\n")

	fmt.Println(result == input)

	return nil
}
