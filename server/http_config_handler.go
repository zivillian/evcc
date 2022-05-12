package server

import (
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util/templates"
)

type VehicleConfig struct {
	Name   string      `json:"name"`
	Fields []FormField `json:"fields"`
}

type FormField struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Help        string `json:"help"`
	Required    bool   `json:"required"`
	Mask        bool   `json:"mask"`
	Advanced    bool   `json:"advanced"`
}

func productName(prd templates.Product, lang string) string {
	result := prd.Brand + " " + prd.Description.String(lang)
	return strings.TrimSpace(result)
}

func toFormFields(params []templates.Param, lang string) []FormField {
	var fields []FormField
	for _, p := range params {
		if !p.Advanced {
			fields = append(fields, FormField{
				Name:        p.Name,
				Description: p.Description.String(lang),
				Help:        p.Help.String(lang),
				Required:    p.Required,
				Mask:        p.Mask,
			})
		}
	}
	return fields
}

func vehicleTemplatesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpls := templates.ByClass(templates.Vehicle)
		var response []VehicleConfig
		for _, tpl := range tpls {
			for _, prd := range tpl.Products {
				response = append(response, VehicleConfig{
					Name:   productName(prd, "de"),
					Fields: toFormFields(tpl.Params, "de"),
				})
			}
		}
		sort.Slice(response, func(i, j int) bool {
			return response[i].Name < response[j].Name
		})
		jsonResult(w, response)
	}
}

func sponsorTokenHandler(cu api.ConfigUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		jsonWrite(w, cu.SponsorToken())
	}
}

func siteTitleHandler(cu api.ConfigUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		title, err := cu.SiteTitle()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		jsonWrite(w, title)
	}
}

func setSiteTitleHandler(cu api.ConfigUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newTitle, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		cu.SetSiteTitle(string(newTitle))
		w.WriteHeader(http.StatusOK)
		jsonWrite(w, "Tadda!")
	}
}
