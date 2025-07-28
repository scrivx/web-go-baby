package mainDomain

import (
	"net/http"

	"github.com/scrivx/go-api-web/internal/models"
	"github.com/scrivx/go-api-web/internal/utils"
	"github.com/scrivx/go-api-web/ui"
)

func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)
	files := []string{
		"html/base.html",
		"html/pages/main/main.html",
	}
	
	utils.ParseTemplateFiles(w, "base", utils.EmplyStruct, utils.EmplyFuncMap, ui.Content, files...)
}

