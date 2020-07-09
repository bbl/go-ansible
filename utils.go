package ansible

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func Exec(command string) error {
	logrus.Debug("Executing cmd: ", command)

	cmdArr := strings.Fields(command)
	cmdName := cmdArr[0]
	cmdArr = cmdArr[1:]

	cmd := exec.Command(cmdName, cmdArr...)
	cmdReader, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	cmdErrReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	errScanner := bufio.NewScanner(cmdErrReader)
	go func() {
		for errScanner.Scan() {
			logrus.Warn(errScanner.Text())
		}
	}()

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			logrus.Debug(scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
