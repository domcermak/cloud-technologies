package client

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"domcermak/ctc/assignments/03/cmd/common"
)

type CommandExecutionerFn func(options map[string]interface{}) (string, error)

type CommandExecutioner struct {
	Execute                  CommandExecutionerFn
	Description, CommandName string
	RequiredArgs             map[string]interface{}
	OptionalArgs             map[string]interface{}
	hideInHelp               bool
	kill                     bool
}

type CommandLine struct {
	InReader     *bufio.Reader
	Out          io.Writer
	ErrOut       io.Writer
	execCommands []CommandExecutioner
}

func NewCommandLine(in io.Reader, out io.Writer, errOut io.Writer, commands ...CommandExecutioner) *CommandLine {
	baseCommands := []CommandExecutioner{
		noopCommand(),
		quitCommand(),
	}
	newCommands := make([]CommandExecutioner, len(commands)+len(baseCommands)+1)
	for i, command := range commands {
		newCommands[i] = command
	}
	for i, command := range baseCommands {
		newCommands[i+len(commands)] = command
	}
	newCommands[len(commands)+len(baseCommands)] = helpCommand(newCommands)

	return &CommandLine{
		InReader:     bufio.NewReader(in),
		Out:          out,
		ErrOut:       errOut,
		execCommands: newCommands,
	}
}

func (command CommandExecutioner) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("# %s\n", command.Description))
	builder.WriteString(fmt.Sprintf("> %s", command.CommandName))

	for _, key := range common.SortedKeys(command.RequiredArgs) {
		builder.WriteString(fmt.Sprintf(" %v=%v", key, command.RequiredArgs[key]))
	}

	for _, key := range common.SortedKeys(command.OptionalArgs) {
		builder.WriteString(fmt.Sprintf(" [%v=%v]", key, command.OptionalArgs[key]))
	}

	return builder.String()
}

func (cmd *CommandLine) RenderAndAcceptCommands() {
	exec, _ := cmd.findFor("help")
	msg, _ := exec.Execute(nil)
	cmd.displayMessage(msg)

	for {
		commandName, args, err := cmd.parseCommand()
		if err != nil {
			cmd.displayError(err)
			continue
		}

		command, err := cmd.findFor(commandName)
		if err != nil {
			cmd.displayError(err)
			continue
		}
		message, err := command.Execute(args)
		if err != nil {
			cmd.displayError(err)
			continue
		}
		cmd.displayMessage(message)

		if command.kill {
			return
		}
	}
}

func (cmd *CommandLine) parseCommand() (string, map[string]interface{}, error) {
	cmd.displayCommandLine()

	line, err := cmd.InReader.ReadString('\n')
	if err != nil {
		return "", nil, err
	}

	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return "", nil, nil
	}

	items := strings.Fields(trimmed)
	commandName, args := items[0], strings.Join(items[1:], " ")

	re := regexp.MustCompile("(\\w+)=(\"[^\"]+\"|[\\w0-9]+)")
	matches := re.FindAllStringSubmatch(args, -1)

	mappedArgs := make(map[string]interface{})
	for _, match := range matches {
		mappedArgs[match[1]] = strings.TrimSuffix(strings.TrimPrefix(match[2], "\""), "\"")
	}

	return commandName, mappedArgs, nil
}

func (cmd *CommandLine) findFor(name string) (CommandExecutioner, error) {
	for _, exec := range cmd.execCommands {
		if exec.CommandName == name {
			return exec, nil
		}
	}

	return CommandExecutioner{}, errors.Errorf("command not found: `%v`", name)
}

func (cmd *CommandLine) displayCommandLine() {
	if _, err := fmt.Fprint(cmd.Out, "> "); err != nil {
		panic(err)
	}
}

func (cmd *CommandLine) displayMessage(msg string) {
	if msg == "" {
		return
	}
	if msg == "\n" {
		msg = ""
	}

	if _, err := fmt.Fprintln(cmd.Out, msg); err != nil {
		panic(err)
	}
}

func (cmd *CommandLine) displayError(err error) {
	if _, err := fmt.Fprintf(cmd.ErrOut, "Error: %v", err); err != nil {
		panic(err)
	}
	cmd.displayMessage("\n")
}

func noopCommand() CommandExecutioner {
	return CommandExecutioner{
		hideInHelp: true,
		Execute: func(options map[string]interface{}) (string, error) {
			return "", nil
		},
	}
}

func quitCommand() CommandExecutioner {
	return CommandExecutioner{
		Execute: func(_ map[string]interface{}) (string, error) {
			return "quitting...", nil
		},
		Description: "Quits the program",
		CommandName: "quit",
		hideInHelp:  false,
		kill:        true,
	}
}

func helpCommand(commands []CommandExecutioner) CommandExecutioner {
	return CommandExecutioner{
		Description: "Displays help text",
		CommandName: "help",
		Execute: func(_ map[string]interface{}) (string, error) {
			builder := strings.Builder{}
			builder.WriteString("###### HELP ######\n\n")

			for _, command := range commands {
				if command.hideInHelp {
					continue
				}
				builder.WriteString(command.String())
				builder.WriteString("\n")
			}

			return builder.String(), nil
		},
	}
}
