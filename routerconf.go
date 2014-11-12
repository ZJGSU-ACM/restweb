package restweb

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func LoadRouter() error {
	file, err := os.Open("config/router.conf")
	if err != nil {
		return err
	}
	br := bufio.NewReader(file)
	eof := true
	for eof {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			eof = false
		} else if err != nil {
			return err
		}
		if line == "" || line[0] == '#' { //if this line is a comment
			continue
		}
		idx := strings.Index(line, "#") //get the beginning of a comment
		if idx >= 0 {                   //if # exists
			line = line[:idx]
		}
		args := strings.Split(line, " ")
		if len(args) < 3 { //if args is to less
			return errors.New("args is to less")
		}
		TrimArgs(args)
		method, pattern := args[0], args[1]
		CA := strings.Split(args[2], ".")
		if len(CA) < 2 {
			return errors.New("args error")
		}
		contorllerName, action := CA[0], CA[1]
		AddRouter(method, pattern, contorllerName, action)
	}
	return nil
}

func TrimArgs(args []string) {
	for idx, arg := range args {
		args[idx] = strings.Trim(arg, " ") //trim space
	}
}