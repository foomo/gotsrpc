package service

import (
	"net/http"
)

type Service interface {
	VariantA(w http.ResponseWriter, r *http.Request, i1 Base) (r1 Base)
	VariantB(w http.ResponseWriter, r *http.Request, i1 BCustomType) (r1 BCustomType)
	VariantC(w http.ResponseWriter, r *http.Request, i1 BCustomTypes) (r1 BCustomTypes)
	VariantD(w http.ResponseWriter, r *http.Request, i1 BCustomTypesMap) (r1 BCustomTypesMap)

	VariantE(w http.ResponseWriter, r *http.Request, i1 *Base) (r1 *Base)

	VariantF(w http.ResponseWriter, r *http.Request, i1 []*Base) (r1 []*Base)

	VariantG(w http.ResponseWriter, r *http.Request, i1 map[string]*Base) (r1 map[string]*Base)

	VariantH(w http.ResponseWriter, r *http.Request, i1 Base, i2 *Base, i3 []*Base, i4 map[string]Base) (r1 Base, r2 *Base, r3 []*Base, r4 map[string]Base)
}
