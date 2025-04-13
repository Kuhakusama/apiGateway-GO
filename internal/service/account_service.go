package service

import (
	"github.com/kuhakusama/apiGateway-GO/internal/domain"
	"github.com/kuhakusama/apiGateway-GO/internal/dto"
)

//arquivo necessairio para ter acessos ao metodos

type AccountService struct {
	repository domain.AccountRepository
}
//vai recerbere um repository pois trablha com o banco de dados

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}	
}

func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input) 

	existingAccount, err := s.repository.FindByApiKey(account.ApiKey) //verifica se o account ja existe
	if err != nil && err!=domain.ErrAccountNotFound{
		return nil, err
	}//erro comun
	if existingAccount != nil {
		return nil, domain.ErrAccountAlreadyExists
	}//erro especifico de objeto duplicado, a outos detalhes a serem trabalhados

	err = s.repository.Save(account)
	if( err != nil) {
		return nil, err
	}
	
	output := dto.FromAccount(account)
	return &output, nil
}	

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.repository.Update(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil 
}

func (s *AccountService) FindByApiKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account) //só uma macete para não ter que fazer o cast e por causa da necessidade de retornar um ponteiro
	return &output, nil
}

func (s *AccountService) FindById(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByID(id)
	if err != nil{
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}