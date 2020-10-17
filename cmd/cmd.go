package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mizzy/para/log"
	"golang.org/x/sync/errgroup"
)

func Run(args []string) error {
	cmd := args[0]
	args = args[1:]

	var eg errgroup.Group
	for _, arg := range args {
		arg := arg
		eg.Go(func() error {
			return run(cmd, arg)
		})
	}

	return eg.Wait()
}

func run(cmd string, arg string) error {
	cmd = strings.ReplaceAll(cmd, "#{0}", arg)
	c := exec.Command("sh", "-c", cmd)

	stdoutPipe, err := c.StdoutPipe()
	if err != nil {
		return err
	}

	stderrPipe, err := c.StderrPipe()
	if err != nil {
		return err
	}

	if err = c.Start(); err != nil {
		return fmt.Errorf("%s: %s", c, err)
	}

	logger := log.New(arg)
	go logger.Error(stderrPipe)
	go logger.Info(stdoutPipe)

	if err = c.Wait(); err != nil {
		return err
	}

	return nil
}
