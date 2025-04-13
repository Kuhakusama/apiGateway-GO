// implementa os endpoints para o serviços WEB(implementação estrutura WEB)
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kuhakusama/apiGateway-GO/internal/dto"
	"github.com/kuhakusama/apiGateway-GO/internal/service"
)

type AccountHandler struct {
	AccountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{AccountService: accountService}
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request){ //interface basica e padronozida
	var input dto.CreateAccountInput
	err := json.NewDecoder(r.Body).Decode(&input) //decodifica o json para o struct, acessa a variavel input para fazer alteração
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest) //retorna erro 400
		return
	}

	output, err := h.AccountService.CreateAccount(input) //chama o service para criar a conta
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError) //retorna erro 500
		return
	}

	w.Header().Set("Content-Type", "application/json") //define o tipo de conteudo da resposta
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output) //codifica o struct para json e retorna na resposta, dto -> Json
}

func (h *AccountHandler) Get(w http.ResponseWriter, r *http.Request){
	apiKey := r.Header.Get("X-AP-key")
	if apiKey == "" {
		http.Error(w, "API Key is required", http.StatusUnauthorized) //retorna erro 400
		return
	}

	output, err := h.AccountService.FindByApiKey(apiKey) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}