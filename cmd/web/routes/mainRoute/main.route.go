package mainRoute

import (
	mainDomain "github.com/scrivx/go-api-web/cmd/web/domain/main"
	"github.com/scrivx/go-api-web/cmd/web/routes"
	"github.com/scrivx/go-api-web/internal/models"
)

func MainRender() {
	routes.GetMuxInstance().HandleFunc("GET " + models.RoutesInstance.MAIN, mainDomain.MainView)
}