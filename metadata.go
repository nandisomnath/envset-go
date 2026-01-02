package main

import (
	"fmt"
	"strings"
)

// this file contains all the metadata for this package
type Metadata struct {
	Name        string
	Version     string
	Description string
	Author      string
	Maintainers []string
	License     string
}

var App = Metadata{
	Name:        "envset",
	Version:     "0.1.0",
	Description: "A cli tool to manage the shell env",
	Author:      "Somnath Nandi",
	Maintainers: []string{"Somnath Nandi"},
	License:     "MIT",
}

func (m Metadata) InfoString() string {
	var lines []string

	// Name and version on first line
	lines = append(lines, fmt.Sprintf("%s %s", m.Name, m.Version))

	// Description
	if m.Description != "" {
		lines = append(lines, m.Description)
	}

	// Author
	if m.Author != "" {
		lines = append(lines, "Author: "+m.Author)
	}

	// Maintainers (compact)
	if len(m.Maintainers) > 0 {
		if len(m.Maintainers) == 1 {
			lines = append(lines, "Maintainer: "+m.Maintainers[0])
		} else {
			lines = append(lines, fmt.Sprintf("Maintainers: %s (and %d more, use --maintainers to see)", m.Maintainers[0], len(m.Maintainers)-1))
		}
	}

	// License
	if m.License != "" {
		lines = append(lines, "License: "+m.License)
	}

	return strings.Join(lines, "\n")
}
