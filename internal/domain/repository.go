package domain

//fala como o acesso aos dados devem ser realizados (é uma impplementação do repository pattern)
type AccountRepository interface {
	Save(account *Account) error
	FindByApiKey(apiKey string) (*Account, error) //1 recebe e o 2 retorna
	FindByID(id string) (*Account, error)
	Update(account *Account) error
}