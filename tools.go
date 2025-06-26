//go:build tools
// +build tools

package tools

// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

import (
	_ "github.com/aarondl/sqlboiler/v4"
	_ "github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-psql"
)
