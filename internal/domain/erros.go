package domain

import "errors"

var(
	ErrAccountNotFound = errors.New("account not found") //erro de conta não encontrada
	ErrDuplicateApiKey = errors.New("duplicate api key") //erro de chave duplicada
	ErrInvoiceNotFound = errors.New("invoice not found") //erro de fatura não encontrada
	ErrUnathorizedAcess = errors.New("unathorized access") //erro de acesso não autorizado
	ErrAccountAlreadyExists = errors.New("account already exists") //erro de conta já existe
)