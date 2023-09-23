package modules

import (
	"breakfast/config"
	"breakfast/types"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

type TemplateData struct {
	Version string

	Vars map[string]string

	PreviousVersions []types.Version

	Http struct {
		Host string
		Url  string
	}

	ContainerName string
}

func MustGetPrepareTemplates(targetVersion string) (TemplateData, *template.Template, *template.Template) {
	version := types.Version{}
	lowerVersions := []types.Version{}

	for _, v := range config.Main.Versions {
		if v.Version == targetVersion {
			version = v
		}

		if v.Version < targetVersion {
			lowerVersions = append(lowerVersions, v)
		}
	}

	if version.Version == "" {
		log.Fatalf("the requested version (%v) is not found among the given versions.", targetVersion)
	}

	sort.Slice(lowerVersions, func(i, j int) bool {
		return lowerVersions[i].Version > lowerVersions[j].Version
	})

	for _, vm := range config.Main.VersionModifiers {
		// replace the wildcard character '*' with the proper format in the regex
		regexStr := vm.Format
		regexStr = fmt.Sprintf("^%s$", regexStr)
		regexStr = strings.ReplaceAll(regexStr, ".", "\\.")
		regexStr = strings.ReplaceAll(regexStr, "*", ".+")

		// compile the regex pattern
		regexPattern := regexp.MustCompile(regexStr)

		// test if the input string matches the regex pattern
		if regexPattern.MatchString(version.Version) {
			for key, value := range vm.Vars {
				version.Vars[key] = value
			}
		}
	}

	templateData := TemplateData{
		Version:          version.Version,
		Vars:             version.Vars,
		PreviousVersions: lowerVersions,
		ContainerName:    os.Getenv("HOSTNAME"),
	}

	var longTemplate, shortTemplate *template.Template

	longTemplate = template.Must(template.ParseFiles(config.Main.Paths.LongTemplate))

	if config.Main.Paths.ShortTemplate != "" {
		shortTemplate = template.Must(template.ParseFiles(config.Main.Paths.ShortTemplate))
	}

	return templateData, longTemplate, shortTemplate
}
