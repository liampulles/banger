package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/liampulles/banger/pkg/library"

	"github.com/liampulles/banger/pkg/wire"

	"github.com/liampulles/banger/playlists"

	"github.com/liampulles/banger/pkg/config"
)

func PlaylistCommand(args []string) error {
	playlist, context, err := setupPlaylist(args)
	if err != nil {
		return err
	}

	playlistCol, ok := playlists.Registered[playlist]
	if !ok {
		return fmt.Errorf("no playlist for arg %s - valid playlists: %s", playlist, playlistsStr(playlists.Registered))
	}

	tracks, err := context.LibraryService().PipeAllTracks()
	if err != nil {
		return fmt.Errorf("could not load tracks: %w", err)
	}

	collected := playlistCol(tracks)

	fmt.Print(formatOutput(collected))
	return nil
}

func setupPlaylist(args []string) (string, wire.Context, error) {
	defaultCfgPath, err := config.DefaultConfigPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not get default config path - continuing: %v\n", err)
	}

	flagSet := flag.NewFlagSet("playlist", flag.ContinueOnError)
	configPath := flagSet.String("configPath", defaultCfgPath, "location of config file. Defaults to standard user config location.")

	if err := flagSet.Parse(args); err != nil {
		return "", nil, fmt.Errorf("could not parse flags: %w", err)
	}
	if configPath == nil || *configPath == "" {
		return "", nil, fmt.Errorf("no config directory specified")
	}
	if len(flagSet.Args()) != 1 {
		return "", nil, fmt.Errorf("one positional argument required for playlist - valid playlists: %s", playlistsStr(playlists.Registered))
	}

	playlist := flagSet.Arg(0)
	context, err := wire.WireContext(*configPath)
	if err != nil {
		return "", nil, fmt.Errorf("could not wire context: %w", err)
	}

	return playlist, context, nil
}

func playlistsStr(all map[string]playlists.PlaylistCollector) string {
	var allNames []string
	for name := range all {
		allNames = append(allNames, name)
	}
	return strings.Join(allNames, ", ")
}

func formatOutput(collected []library.Track) string {
	output := ""
	for _, track := range collected {
		output += fmt.Sprintf("%s\n", track.Location())
	}
	return output
}
