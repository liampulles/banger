package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/liampulles/banger/pkg/file"

	"github.com/liampulles/banger/pkg/config"
)

func InitCommand(args []string) error {
	libraryRoot, configPath, err := setupInit(args)
	if err != nil {
		return err
	}

	if !file.DoesExist(libraryRoot) {
		return fmt.Errorf("library root must already exist: %s", libraryRoot)
	}
	if !file.IsDir(libraryRoot) {
		return fmt.Errorf("library root must be a directory: %s", libraryRoot)
	}

	cfg := config.NewConfigImpl(libraryRoot)
	cfgService := config.NewConfigServiceImpl(configPath)
	if err = cfgService.Save(cfg); err != nil {
		return fmt.Errorf("could not save config: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Config created at %s.\n", configPath)
	return nil
}

func setupInit(args []string) (string, string, error) {
	defaultCfgPath, err := config.DefaultConfigPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not get default config path - continuing: %v", err)
	}

	flagSet := flag.NewFlagSet("playlist", flag.ContinueOnError)
	configPath := flagSet.String("configPath", defaultCfgPath, "location of config file. Defaults to standard user config location.")

	if err := flagSet.Parse(args); err != nil {
		return "", "", fmt.Errorf("could not parse flags: %w", err)
	}
	if configPath == nil || *configPath == "" {
		return "", "", fmt.Errorf("no config directory specified")
	}

	if len(flagSet.Args()) != 1 {
		return "", "", fmt.Errorf("one positional argument required for library root directory")
	}
	libraryRoot := flagSet.Arg(0)
	if libraryRoot == "" {
		return "", "", fmt.Errorf("one positional argument required for library root directory")
	}

	return libraryRoot, *configPath, nil
}
