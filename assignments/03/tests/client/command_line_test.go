package client

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"domcermak/ctc/assignments/03/cmd/client"
	"domcermak/ctc/assignments/03/tests/helpers"
)

func TestCommandLine_RenderAndAcceptCommands(t *testing.T) {
	for _, tc := range []struct {
		name                        string
		commandExecutioners         []client.CommandExecutioner
		inputCommands               []string
		expectedOut, expectedErrOut string
	}{
		{
			name:                "With only built in commands",
			commandExecutioners: []client.CommandExecutioner{},
			inputCommands: []string{
				"help",
				"",                     // noop
				"invalid_command id=1", // causes error
				"quit",
			},
			expectedOut: strings.Join([]string{
				"###### HELP ######",
				"",
				"# Quits the program",
				"> quit",
				"# Displays help text",
				"> help",
				"",
				"> ###### HELP ######",
				"",
				"# Quits the program",
				"> quit",
				"# Displays help text",
				"> help",
				"",
				"> > ",
				"> quitting...\n",
			}, "\n"),
			expectedErrOut: strings.Join([]string{
				"Error: command not found: `invalid_command`",
			}, "\n"),
		},
		{
			name: "With additional commands",
			commandExecutioners: []client.CommandExecutioner{
				{
					CommandName:  "list",
					Description:  "List command description",
					OptionalArgs: map[string]interface{}{},
					RequiredArgs: map[string]interface{}{},
					Execute: func(options map[string]interface{}) (string, error) {
						return "list message", nil
					},
				},
				{
					CommandName: "get",
					Description: "Get command description",
					OptionalArgs: map[string]interface{}{
						"name": "<text>",
					},
					RequiredArgs: map[string]interface{}{
						"id": "<number>",
					},
					Execute: func(options map[string]interface{}) (string, error) {
						id, ok := options["id"]
						if !ok {
							return "", errors.New("get: missing required arguments")
						}

						helpers.Expect(id, "1", t)
						if name, ok := options["name"]; ok {
							helpers.Expect(name, "hello", t)
						}

						return "get success message", nil
					},
				},
			},
			inputCommands: []string{
				"list",
				"get",
				`get id=1 name=hello  `,
				"quit",
			},
			expectedOut: strings.Join([]string{
				"###### HELP ######",
				"",
				"# List command description",
				"> list",
				"# Get command description",
				"> get id=<number> [name=<text>]",
				"# Quits the program",
				"> quit",
				"# Displays help text",
				"> help",
				"",
				"> list message",
				"> ", // errors out
				"> get success message",
				"> quitting...\n",
			}, "\n"),
			expectedErrOut: strings.Join([]string{
				"Error: get: missing required arguments",
			}, "\n"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			in := bytes.NewBuffer([]byte(strings.Join(tc.inputCommands, "\n") + "\n"))
			var out, errOut bytes.Buffer

			commandLine := client.NewCommandLine(in, &out, &errOut, tc.commandExecutioners...)
			commandLine.RenderAndAcceptCommands()

			helpers.Expect(tc.expectedOut, out.String(), t)
			helpers.Expect(tc.expectedErrOut, errOut.String(), t)
		})
	}
}
