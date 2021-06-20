package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trongtb88/wagerservice/src/business/entity"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (rst *rest) CreateWager(w http.ResponseWriter, r *http.Request) {
	var param entity.WagerParam
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rst.httpRespError(w, http.StatusBadRequest, err, "Can not read body request")
		return
	}
	if err := json.Unmarshal(body, &param); err != nil {
		rst.httpRespError(w, http.StatusBadRequest, err, "Can not parse body request")
		return
	}

	// Validation
	if param.SellingPercentage > 100 || param.SellingPercentage  <=0 {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid Selling Percentage")
		return
	}

	// Validation
	if param.Odds <= 0 {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid Odds")
		return
	}

	if param.TotalWagerValue < 0 {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid Total Wager Value")
		return
	}

	// Validation
	if param.Odds <= 0 {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid Odds")
		return
	}

	wager, err := rst.uc.Wager.CreateWager(r.Context(), param)
	if err != nil {
		rst.httpRespError(w, http.StatusUnprocessableEntity, err, "Can not create Wager : " + err.Error())
		return
	}
	rst.httpRespSuccess(w, http.StatusCreated, wager)
}

func (rst *rest) BuyWager(w http.ResponseWriter, r *http.Request) {
	var param entity.BuyWagerParam
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		rst.httpRespError(w, http.StatusBadRequest, err, "Can not read body request")
		return
	}
	if err := json.Unmarshal(body, &param); err != nil {
		rst.httpRespError(w, http.StatusBadRequest, err, "Can not parse body request")
		return
	}


	vars := mux.Vars(r)
	wagerIdPara, _ := vars["wager_id"]

	if len(wagerIdPara) == 0 {
		rst.httpRespError(w, http.StatusBadRequest, err, "Missing required para wager_id")
		return
	}

	 wagerId, err := strconv.ParseInt(wagerIdPara, 10, 32)

	if err != nil {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid wager_id")
		return
	}

	if wagerId < 0 {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid wager_id")
		return
	}

	if param.BuyingPrice < 0 {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid buying_price")
		return
	}

	wager, err  := rst.uc.Wager.GetWagerById(r.Context(), wagerId)
	if err != nil {
		rst.httpRespError(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	if wager.Id == 0 {
		rst.httpRespError(w, http.StatusNotFound, err, err.Error())
		return
	}

	if wager.CurrentSellingPrice > 0 && param.BuyingPrice > wager.CurrentSellingPrice {
		rst.httpRespError(w, http.StatusBadRequest, err, "Buying Price > Current Selling Price")
		return
	}

	param.WagerId = wagerId

	buyWager, err := rst.uc.Wager.CreateBuyWager(r.Context(), param)


	if err != nil {
		rst.httpRespError(w, http.StatusUnprocessableEntity, err, "Can not create Wager")
	}
	rst.httpRespSuccess(w, http.StatusCreated, buyWager)
}

func (rst *rest) GetWagers(w http.ResponseWriter, r *http.Request) {
	page  := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageNumber, err := strconv.Atoi(page)

	if err != nil {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid page")
		return
	}

	if pageNumber < 0 {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid page")
		return
	}

	limitNumber, err := strconv.Atoi(limit)

	if err != nil {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid limit")
		return
	}

	if limitNumber < 0 {
		rst.httpRespError(w, http.StatusBadRequest, err, "Invalid limit")
		return
	}

	wagers, err := rst.uc.Wager.GetWagers(r.Context(), pageNumber, limitNumber)

	if err != nil {
		rst.httpRespError(w, http.StatusUnprocessableEntity, err, "Can not get Wagers")
	}


	rst.httpRespSuccess(w, http.StatusOK, wagers)

}


