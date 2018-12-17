package cmd

type (
	// Command represent a command details  the client sent to the server
	Command struct {
		Name string
		Args []string
	}
)

func New(name string, args []string) *Command {
	return &Command{Name: name, Args: args}
}
