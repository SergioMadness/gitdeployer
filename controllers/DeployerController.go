package controllers

import "net/http"

type DeployerController struct {
}

func CreateDeployer() *DeployerController {
	return new(DeployerController)
}

func (*DeployerController) Deploy(w http.ResponseWriter, r *http.Request) {

}
