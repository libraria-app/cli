/*
Copyright © 2023 libraria-app
*/
package commands

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/libraria-app/cli/internal/librariacli"
	"github.com/libraria-app/cli/internal/utils/print"
	"github.com/spf13/cobra"
)

const fileLangPlaceholder = "{lang}"
const (
	jsonFormat = "json"
	xmlFormat  = "xml"
	iosFormat  = "strings"
)
const (
	projectIdFlag = "projectId"
	languageFlag  = "language"
	pathFlag      = "path"
)

func newExportCommand(lcli *librariacli.Lcli) *cobra.Command {
	exportCmd := &cobra.Command{
		Use:   "export",
		Short: "Export translation to the file",
		Long: `Export command will creates a new or updates exising file by the provided path with the exported project terms translations 
by the user API Key.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runExport(lcli, cmd); err != nil {
				print.Error(fmt.Sprintf("error: %v", err))
			}
		},
	}

	exportCmd.Flags().String(projectIdFlag, "", "The Project ID")
	exportCmd.Flags().String(languageFlag, "", "The language")
	exportCmd.Flags().String(pathFlag, "", "The filename to export with {lang} placeholder")

	exportCmd.MarkFlagRequired(projectIdFlag)
	exportCmd.MarkFlagRequired(languageFlag)
	exportCmd.MarkFlagRequired(pathFlag)

	return exportCmd
}

func runExport(lcli *librariacli.Lcli, cmd *cobra.Command) error {
	filename := cmd.Flag(pathFlag).Value.String()
	if !validatePath(filename) {
		return fmt.Errorf("path should contain {lang} placeholder and format one of: .json, .xml, .strings")
	}
	format := fetchFileFormat(filename)
	if format == "" {
		return fmt.Errorf("file format should be one of: .json, .xml, .strings")
	}

	var qs = []*survey.Question{
		{
			Name:     "apiKey",
			Prompt:   &survey.Input{Message: "Please, input your Libraria API key:"},
			Validate: survey.Required,
		},
	}
	answer := struct {
		ApiKey string
	}{}

	if err := survey.Ask(qs, &answer); err != nil {
		return fmt.Errorf("asking error: %w", err)
	}

	projectId := cmd.Flag(projectIdFlag).Value.String()
	language := cmd.Flag(languageFlag).Value.String()

	s := lcli.GetService()
	response, err := s.ExportTerms(answer.ApiKey, projectId, language, format)
	if err != nil {
		return fmt.Errorf("export terms: %w", err)
	}

	filename = strings.ReplaceAll(filename, fileLangPlaceholder, language)
	if err := createFilepath(filename); err != nil {
		return fmt.Errorf("create filepath: %w", err)
	}

	perm, err := fetchFilePerm(filename)
	if err != nil {
		return fmt.Errorf("fetch file permission: %w", err)
	}

	if err := os.WriteFile(filename, response, perm); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	print.Info(fmt.Sprintf("The file %s is successfully generated", filename))

	return nil
}

func createFilepath(filename string) error {
	dir, _ := path.Split(filename)

	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return fmt.Errorf("create dir: %w", err)
			}
		} else {
			return fmt.Errorf("find path: %w", err)
		}
	}

	return nil
}

func fetchFilePerm(filename string) (os.FileMode, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return os.ModePerm, nil
		}
		return os.ModeIrregular, fmt.Errorf("find file: %w", err)
	}
	return fi.Mode().Perm(), nil
}

func validatePath(path string) bool {
	key := strings.TrimSpace(path)
	key = strings.ToLower(key)
	if !strings.Contains(key, fileLangPlaceholder) {
		return false
	}
	return strings.HasSuffix(key, fmt.Sprintf(".%s", jsonFormat)) ||
		strings.HasSuffix(key, fmt.Sprintf(".%s", xmlFormat)) ||
		strings.HasSuffix(key, fmt.Sprintf(".%s", iosFormat))
}

func fetchFileFormat(path string) string {
	if path == "" {
		return ""
	}
	parts := strings.Split(path, ".")
	if len(parts) == 1 {
		return ""
	}
	ext := strings.ToLower(parts[len(parts)-1])

	formats := map[string]string{
		jsonFormat: jsonFormat,
		xmlFormat:  xmlFormat,
		iosFormat:  "ios",
	}

	return formats[ext]
}
