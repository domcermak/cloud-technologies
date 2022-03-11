//go:build integration

package integration

import (
	"bytes"
	cl "domcermak/ctc/assignments/03/cmd/client"
	"domcermak/ctc/assignments/03/cmd/server"
	"domcermak/ctc/assignments/03/tests/helpers"
	"net"
	"strings"
	"testing"
	"time"
)

func TestClientServerIntegration(t *testing.T) {
	postgres, err := helpers.TestPostgresConnect()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(helpers.CleanupTestPostgres(postgres))

	for _, tc := range []struct {
		name                        string
		inputCommands               []string
		expectedOut, expectedErrOut string
	}{
		{
			name: "Lists data correctly from the server to the command line",
			inputCommands: []string{
				"list",
				"quit",
			},
			expectedOut: strings.Join([]string{
				"###### HELP ######",
				"",
				"# Returns all existing products",
				"> list",
				"# Returns a product by given id",
				"> get id=<number>",
				"# Deletes a product by given id",
				"> delete id=<number>",
				"# Updates a product by given id",
				"> update id=<number> [amount=<number>] [name=<text>] [price=<number>]",
				"# Quits the program",
				"> quit",
				"# Displays help text",
				"> help",
				"",
				"> [\n\t{\n\t\t\"id\": 1,\n\t\t\"name\": \"clementine\",\n\t\t\"price\": 1.38,\n\t\t\"amount\": 8\n\t},\n\t{\n\t\t\"id\": 2,\n\t\t\"name\": \"apricot\",\n\t\t\"price\": 12.3,\n\t\t\"amount\": 12\n\t},\n\t{\n\t\t\"id\": 3,\n\t\t\"name\": \"peach\",\n\t\t\"price\": 1.1,\n\t\t\"amount\": 1\n\t},\n\t{\n\t\t\"id\": 4,\n\t\t\"name\": \"star fruit\",\n\t\t\"price\": 1,\n\t\t\"amount\": 24\n\t},\n\t{\n\t\t\"id\": 5,\n\t\t\"name\": \"huckleberry\",\n\t\t\"price\": 33.9,\n\t\t\"amount\": 11\n\t},\n\t{\n\t\t\"id\": 6,\n\t\t\"name\": \"jujube\",\n\t\t\"price\": 12.89,\n\t\t\"amount\": 7\n\t}\n]",
				"> quitting...\n",
			}, "\n"),
			expectedErrOut: "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			in := bytes.NewBuffer([]byte(strings.Join(tc.inputCommands, "\n") + "\n"))
			var out, errOut bytes.Buffer

			addr, quitChan := runServer(postgres)
			defer func() {
				quitChan <- 0
			}()

			client := cl.NewClient(10*time.Second, addr)
			commandLine := cl.NewCommandLine(in, &out, &errOut, client.CommandExecutioners()...)
			commandLine.RenderAndAcceptCommands()

			helpers.Expect(tc.expectedOut, out.String(), t)
			helpers.Expect(tc.expectedErrOut, errOut.String(), t)
		})
	}
}

func runServer(pool server.Pool) (string, chan<- interface{}) {
	quitChan := make(chan interface{})
	addr := make(chan string)

	go func() {
		go func() {
			s := server.NewServer("", pool)

			listener, err := net.Listen("tcp", ":0") // find a free port
			if err != nil {
				panic(err)
			}
			addr <- listener.Addr().String()

			if err := s.Serve(listener); err != nil {
				panic(err)
			}
		}()

		// quitting this goroutine quits also the child
		// goroutine with the test server
		<-quitChan
	}()

	return <-addr, quitChan
}
