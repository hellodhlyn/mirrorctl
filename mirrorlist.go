package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/markbates/pkger"
	"gopkg.in/yaml.v2"
)

type (
	Source interface {
		EnvVars() map[string]string
	}

	Bundler struct {
		Rubygems string `yaml:"rubygems_org"`
	}

	Pipenv struct {
		PypiMirror string `yaml:"pypi_mirror"`
	}

	Pypi struct {
		IndexUrl      string `yaml:"index_url"`
		ExtraIndexUrl string `yaml:"extra_index_url"`
		TrustedHost   string `yaml:"trusted_host"`
	}

	Mirrorlist struct {
		Bundler *Bundler `yaml:"bundler"`
		Pipenv  *Pipenv  `yaml:"pipenv"`
		Pypi    *Pypi    `yaml:"pypi"`
	}
)

func GetMirrorlist(location string) (*Mirrorlist, error) {
	var locationCode string
	switch strings.ToLower(location) {
	case "kr", "korea", "south_korea":
		locationCode = "kr"
	default:
		locationCode = "default"
	}

	f, err := pkger.Open(fmt.Sprintf("/mirrorlist/%s.yaml", locationCode))
	if err != nil {
		return nil, errors.New("failed to load mirrorlist")
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.New("failed to read mirrorlist")
	}

	var mirrorlist Mirrorlist
	if yaml.Unmarshal(bytes, &mirrorlist) != nil {
		return nil, errors.New("failed to parse mirrorlist")
	}

	return &mirrorlist, nil
}

func (b *Bundler) EnvVars() map[string]string {
	return compactMap(map[string]string{
		"BUNDLE_MIRROR__RUBYGEMS__ORG": b.Rubygems,
	})
}

func (p *Pipenv) EnvVars() map[string]string {
	return compactMap(map[string]string{
		"PIPENV_PYPI_MIRROR": p.PypiMirror,
	})
}

func (p *Pypi) EnvVars() map[string]string {
	return compactMap(map[string]string{
		"PIP_INDEX_URL":       p.IndexUrl,
		"PIP_EXTRA_INDEX_URL": p.ExtraIndexUrl,
		"PIP_TRUSTED_HOST":    p.TrustedHost,
	})
}

func (ml *Mirrorlist) EnvVarsAll() map[string]string {
	sources := []Source{ml.Bundler, ml.Pipenv, ml.Pypi}
	result := map[string]string{}
	for _, source := range sources {
		for k, v := range source.EnvVars() {
			result[k] = v
		}
	}
	return result
}

func compactMap(src map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range src {
		if v != "" {
			result[k] = v
		}
	}
	return result
}
