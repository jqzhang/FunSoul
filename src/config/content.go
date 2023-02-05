package config

type Content struct {
	Env struct {
		Debug bool
	}

	DB struct {
		Host     string
		Name     string
		Password string
	}
}
