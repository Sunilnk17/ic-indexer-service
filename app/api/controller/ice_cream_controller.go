package controller

import (
	"ic-indexer-service/app/api/dataaccessor"
	"ic-indexer-service/app/api/preprocessor"
	"ic-indexer-service/app/config/locales/local_config"
	"ic-indexer-service/app/model/request"
	"ic-indexer-service/app/processor"
	"ic-indexer-service/app/response_handler"
	"log"
	"net/http"
)

// swagger:operation GET /icecream getIcecream
// ---
// summary: Get Icecream
// description: Get Icecream
// parameters:
// - name: name
//   in: query
//   description: name of icecream
//   type: string
//   required: false
//   Responses:
//     default: body:genericError
//     200: body:genericModel
// swagger:route GET /icecream getIcecream
//
// Get Icecream
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//
//     Responses:
//       default: body:genericError
//       200: body:genericModel
func GetIcecream(w http.ResponseWriter, r *http.Request) {

	var icecreamRequest request.IcecreamFilter

	err := preprocessor.DecodeAndValidateRequestParams(r, &icecreamRequest)

	if err != nil {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	indexerService := processor.GetNewIcecreamIndexerService(getIcecreamIndexerParams())

	icecreamResponse, icecreamErr := indexerService.Search(r.Context(), icecreamRequest)

	if icecreamErr != nil {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_companyId_or_sectorId")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	response_handler.WriteResponseMapAsJson(w, r, http.StatusOK, getResultAsMap(icecreamResponse))

}

func getIcecreamIndexerParams() (dataaccessor.IceCreamDataAccessor, processor.UpdateHandler) {
	return dataaccessor.NewIceCreamDataAccessor(), processor.NewUpdateHandler()
}

func getResultAsMap(data interface{}) map[string]interface{} {
	return map[string]interface{}{"data": data}
}

// swagger:operation PUT /icecream SaveOrUpdateIcecream
// ---
// summary: SaveOrUpdate an Icecream
// description: SaveOrUpdate an Icecream
// parameters:
// - name: icecreamClientRequest
//   in: body
//   description: icecream request
//   required: true
//
//   schema:
//     "$ref": "#/definitions/icecreamClientRequest"
//   Responses:
//     default: body:genericError
//     200: body:genericModel
//
// swagger:route PUT /icecream SaveOrUpdateIcecream
//
// SaveOrUpdate an Icecream
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//
//     Responses:
//       default: body:genericError
//       200: body:genericModel
func UpdateIcecream(w http.ResponseWriter, r *http.Request) {
	var icecreamRequest request.IcecreamIndexRequest

	err := preprocessor.DecodeAndValidateRequestParams(r, &icecreamRequest)

	if err != nil {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	indexerService := processor.GetNewIcecreamIndexerService(getIcecreamIndexerParams())
	updatingError := indexerService.PartialUpdate(r.Context(), icecreamRequest)
	if updatingError != nil {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_companyId_or_sectorId")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}
	var responseMessage string
	responseMessage = local_config.GetTranslationMessage(r.Context(), "indexed_successfully")
	log.Print(r.Context(), responseMessage)
	response_handler.WriteResponseMapAsJson(w, r, http.StatusOK, getResultAsMap(responseMessage))
	return
}
