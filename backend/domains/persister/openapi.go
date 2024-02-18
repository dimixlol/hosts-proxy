package persister

import (
	"github.com/dimixlol/hosts-proxy/config"
	"github.com/wI2L/fizz/openapi"
)

func constructOpenApiInfo() *openapi.Info {
	return &openapi.Info{
		Title:       config.Configuration.API.Title,
		Description: config.Configuration.API.Description,
		Version:     config.Configuration.Version,
		Contact: &openapi.Contact{
			Name:  config.Configuration.API.Contact.Name,
			URL:   config.Configuration.API.Contact.URL,
			Email: config.Configuration.API.Contact.Email,
		},
		License: &openapi.License{
			Name: config.Configuration.API.License.Name,
			URL:  config.Configuration.API.License.URL,
		},
		TermsOfService: "string",
		XLogo: &openapi.XLogo{
			URL:             config.Configuration.API.Logo.URL,
			BackgroundColor: config.Configuration.API.Logo.Color,
			AltText:         config.Configuration.API.Logo.AltText,
			Href:            config.Configuration.API.Logo.Href,
		},
	}
}
