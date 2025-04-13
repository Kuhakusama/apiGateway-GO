package dto

import (
	"time"

	"github.com/kuhakusama/apiGateway-GO/internal/domain"
)

type CreateAccountInput struct { //Dto depende do service
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	ApiKey	  string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email) //converte o dto para o domain
}

func FromAccount(account *domain.Account) AccountOutput { //domain para dto
	return AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		ApiKey:    account.ApiKey,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}