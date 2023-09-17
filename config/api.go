package config

type (
	apiContactConfiguration struct {
		Name  string
		Email string
		URL   string
	}

	apiLicenseConfiguration struct {
		Name string
		URL  string
	}

	apiLogoConfiguration struct {
		URL     string
		Color   string
		AltText string
		Href    string
	}

	apiConfiguration struct {
		Title       string
		Description string
		Version     string
		Contact     *apiContactConfiguration
		License     *apiLicenseConfiguration
		Logo        *apiLogoConfiguration
	}
)
