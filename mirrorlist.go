package main

type (
	Source interface {
		EnvVars() map[string]string
	}

	Bundler struct {
		Rubygems string
	}

	Pipenv struct {
		PypiMirror string
	}

	Pypi struct {
		IndexUrl      string
		ExtraIndexUrl string
		TrustedHost   string
	}

	Mirrorlist struct {
		Bundler *Bundler
		Pipenv  *Pipenv
		Pypi    *Pypi
	}
)

var (
	// Default
	defaultBundler = Bundler{
		Rubygems: "https://rubygems.org",
	}
	defaultPipenv = Pipenv{
		PypiMirror: "https://pypi.org/simple",
	}
	defaultPypi = Pypi{
		IndexUrl: "https://pypi.org/simple",
	}
	defaultMirror = Mirrorlist{
		Bundler: &defaultBundler,
		Pipenv:  &defaultPipenv,
		Pypi:    &defaultPypi,
	}

	// Korea (korea)
	krBundler = Bundler{
		Rubygems: "http://mirror.kakao.com/rubygem",
	}
	krPipenv = Pipenv{
		PypiMirror: "http://mirror.kakao.com/pypi/simple",
	}
	krPypi = Pypi{
		IndexUrl:      "http://mirror.kakao.com/pypi/simple",
		ExtraIndexUrl: "https://pypi.org/simple",
		TrustedHost:   "mirror.kakao.com",
	}
	krMirror = Mirrorlist{
		Bundler: &krBundler,
		Pipenv:  &krPipenv,
		Pypi:    &krPypi,
	}
)

func GetMirrorlist(location string) *Mirrorlist {
	switch location {
	case "kr", "korea":
		return &krMirror
	default:
		return &defaultMirror
	}
}

func (b *Bundler) EnvVars() map[string]string {
	return compactMap(map[string]string{
		"BUNDLER_RUBYGEMS__ORG": b.Rubygems,
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
