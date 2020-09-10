package commands

var Registered = map[string]Command{
	"playlist": PlaylistCommand,
	"init":     InitCommand,
}

type Command func(args []string) error
