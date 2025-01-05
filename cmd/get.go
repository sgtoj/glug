package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	// "path/pathfile"

	"github.com/spf13/cobra"

	// "github.com/sgtoj/glug/internal/downloader"
	"github.com/sgtoj/glug/internal/downloader"
	"github.com/sgtoj/glug/internal/registry"
	"github.com/sgtoj/glug/internal/unpacker"
	// "github.com/sgtoj/glug/internal/unpacker"
)

var getCmd = &cobra.Command{
	Use:   "get [tool_name]",
	Short: "get a tool by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		toolName := args[0]

		configFile := "./data/registry.lua"
		registryData, err := registry.BuildToolRegistry(configFile)
		if err != nil {
			return fmt.Errorf("failed to build registry: %w", err)
		}

		tool, ok := registryData[toolName]
		if !ok {
			return fmt.Errorf("tool not found in registry: %s", toolName)
		}

		toolVer, err := tool.GetVersion()
		if err != nil {
			return fmt.Errorf("unable to get version: %w", err)
		}
		log.Println("tool version", toolVer)

		toolUrl, err := tool.GetUrl()
		if err != nil {
			return fmt.Errorf("unable to get url: %w", err)
		}

		tmpDir, err := downloader.CreateTmpDir()
		if err != nil {
			return fmt.Errorf("unable to create tmp directory: %w", err)
		}
		// defer downloader.CleanUpTmpDir(tmpDir)

		tmpFile, err := downloader.Download(toolUrl, tmpDir)
		if err != nil {
			return fmt.Errorf("unable to download file: %w", err)
		}
		log.Println("file downloaded to tmp directory", tmpFile)

		tmpUnpackBaseDir := filepath.Join(tmpDir, "unpacked")
		tmpUnpackDir, err := unpacker.Unpack(tmpFile, tmpUnpackBaseDir)
		if err != nil {
			return fmt.Errorf("unpacked failed: %w", err)
		}

		binDir := "./tmp/bin"
		binPath, err := downloader.Move(tmpUnpackDir, binDir, tool)
		if err != nil {
			return fmt.Errorf("unable to move binary to bin directory: %w", err)
		}
		log.Println("tool installed", binPath)

		return nil
	},
}
