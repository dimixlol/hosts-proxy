package api

type (
	newPersistentSiteRequest struct {
		Host string `json:"host" description:"fqdn" binding:"required"`
		IP   string `json:"ip" description:"ipv4 address" binding:"required"`
	}

	newPersistentSiteResponse struct {
		IP   string `json:"ip"`
		Host string `json:"host"`
		Slug string `json:"slug"`
	}
	getURLBySlugRequest struct {
		Slug string `json:"slug" binding:"required" example:"abasicslug"`
	}
)
