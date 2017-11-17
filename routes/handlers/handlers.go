package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"goji.io/pat"
)

func Search(w http.ResponseWriter, r *http.Request) {
	client := http.DefaultClient

	searchText := pat.Param(r, "searchText")
	start := r.URL.Query().Get("start")

	walmartSearchBaseURL := "http://api.walmartlabs.com/v1/search"
	walmartAPIKey := "u63audhfqxgm76swmmdwvy98"

	escaptedST := url.QueryEscape(searchText)

	URL := fmt.Sprintf(walmartSearchBaseURL+"?query=%s&format=json&apiKey=%s&start=%s", escaptedST, walmartAPIKey, start)
	fmt.Println(URL)

	// http://api.walmartlabs.com/v1/search?query=ipod&format=json&apiKey=u63audhfqxgm76swmmdwvy98
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var wr walmartResponse
	err = json.Unmarshal(body, &wr)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	type respItem struct {
		ItemID           int
		Name             string
		Price            float64
		Categories       string
		ShortDescription string
		LongDescription  string
		ThumbnailImage   string
		ProductURL       string
		MediumImage      string
		LargeImage       string
	}

	var response struct {
		TotalResults int
		Start        int
		NumItems     int
		Items        []respItem
	}

	response.TotalResults = wr.TotalResults
	response.Start = wr.Start
	response.NumItems = wr.NumItems
	response.Items = make([]respItem, 0)

	for _, i := range wr.Items {
		response.Items = append(response.Items, respItem{
			ItemID:           i.ItemID,
			Name:             i.Name,
			Price:            i.SalePrice,
			Categories:       i.CategoryPath,
			ShortDescription: i.ShortDescription,
			LongDescription:  i.LongDescription,
			ThumbnailImage:   i.ThumbnailImage,
			ProductURL:       i.ProductURL,
			MediumImage:      i.MediumImage,
			LargeImage:       i.LargeImage,
		})
	}

	json.NewEncoder(w).Encode(response)
}

type walmartResponse struct {
	Query         string `json:"query"`
	Sort          string `json:"sort"`
	ResponseGroup string `json:"responseGroup"`
	TotalResults  int    `json:"totalResults"`
	Start         int    `json:"start"`
	NumItems      int    `json:"numItems"`
	Items         []struct {
		ItemID                int     `json:"itemId"`
		ParentItemID          int     `json:"parentItemId"`
		Name                  string  `json:"name"`
		Msrp                  float64 `json:"msrp"`
		SalePrice             float64 `json:"salePrice"`
		Upc                   string  `json:"upc"`
		CategoryPath          string  `json:"categoryPath"`
		ShortDescription      string  `json:"shortDescription"`
		LongDescription       string  `json:"longDescription"`
		ThumbnailImage        string  `json:"thumbnailImage"`
		MediumImage           string  `json:"mediumImage"`
		LargeImage            string  `json:"largeImage"`
		ProductTrackingURL    string  `json:"productTrackingUrl"`
		StandardShipRate      float64 `json:"standardShipRate"`
		Marketplace           bool    `json:"marketplace"`
		ModelNumber           string  `json:"modelNumber"`
		ProductURL            string  `json:"productUrl"`
		CustomerRating        string  `json:"customerRating"`
		NumReviews            int     `json:"numReviews"`
		CustomerRatingImage   string  `json:"customerRatingImage"`
		CategoryNode          string  `json:"categoryNode"`
		Bundle                bool    `json:"bundle"`
		Stock                 string  `json:"stock"`
		AddToCartURL          string  `json:"addToCartUrl"`
		AffiliateAddToCartURL string  `json:"affiliateAddToCartUrl"`
		GiftOptions           struct {
			AllowGiftWrap    bool `json:"allowGiftWrap"`
			AllowGiftMessage bool `json:"allowGiftMessage"`
			AllowGiftReceipt bool `json:"allowGiftReceipt"`
		} `json:"giftOptions"`
		ImageEntities []struct {
			ThumbnailImage string `json:"thumbnailImage"`
			MediumImage    string `json:"mediumImage"`
			LargeImage     string `json:"largeImage"`
			EntityType     string `json:"entityType"`
		} `json:"imageEntities"`
		OfferType                string `json:"offerType"`
		IsTwoDayShippingEligible bool   `json:"isTwoDayShippingEligible"`
		AvailableOnline          bool   `json:"availableOnline"`
	} `json:"items"`
}
