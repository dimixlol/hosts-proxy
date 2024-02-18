package api

import (
	"github.com/dimixlol/hosts-proxy/utils"
	"github.com/wI2L/fizz"
	"net/http"
	"strconv"
)

var (
	newPersistentSiteSpec = []fizz.OperationOption{
		fizz.ID("CreatePersistentSite"),
		fizz.Summary("Create Site"),
		fizz.Description("Create persistent site mapping with host to ip"),
		fizz.StatusDescription("Successful Response"),
		fizz.Response(strconv.Itoa(http.StatusBadRequest), "Invalid IP/Host", nil, nil, &utils.UnsuccessfulResponse{Status: http.StatusBadRequest, Err: "invalid data passed"}),
	}
	getURLBySlugSpec = []fizz.OperationOption{
		fizz.ID("GetURLBySlug"),
		fizz.Summary("Get URL By Slug"),
		fizz.Description("Get URL By Slug"),
		fizz.StatusDescription("Successful Response"),
		fizz.Response(strconv.Itoa(http.StatusBadRequest), "Invalid slug", nil, nil, &utils.UnsuccessfulResponse{Status: http.StatusBadRequest, Err: "invalid slug"}),
		fizz.Response(strconv.Itoa(http.StatusNotFound), "URL not found", nil, nil, &utils.UnsuccessfulResponse{Status: http.StatusNotFound, Err: "url not found"}),
	}
)
