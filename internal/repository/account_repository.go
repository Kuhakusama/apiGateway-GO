package repository

import (
	"database/sql"
	"time"

	"github.com/kuhakusama/apiGateway-GO/internal/domain"
)

type AccountRepository struct {
	db *sql.DB //pacote, interface para banco de dados que depois é somente implementado
}

// FindByID implements domain.AccountRepository.
func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	panic("unimplemented")
}

// Update implements domain.AccountRepository.
func (r *AccountRepository) Update(account *domain.Account) error {
	panic("unimplemented")
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`) //stmt = statement
	if err != nil {
		return err
	}
	defer stmt.Close() //evita leaks de memoria

	_, err = stmt.Exec(
		account.ID, account.Name, account.Email, account.ApiKey, account.Balance, account.CreatedAt, account.UpdatedAt,
	) //retorna quantidade de linhas e um erro, _ = ignora o retorno, GO obriga o uso de todas as variaveis
	if err != nil {
		return err
	}
	return nil //nil = erro em branco ou erro nulo, tratamento de erros atraves dos states dos erros
}

func (r *AccountRepository) FindByApiKey(apikey string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updateAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts 
		WHERE api_key = $1
	`, apikey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.ApiKey,
		&account.Balance,
		&createdAt,
		&updateAt,
	) //& pois altera a memoria
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrAccountNotFound //não encontrou nada
		}
		return nil, err //erro diferente de não encontrado
	}
	account.CreatedAt = createdAt
	account.UpdatedAt = updateAt
	return &account, nil //encontrou o account
}

func (r *AccountRepository) FindById(id string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updateAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts
		WHERE id = $1 
	`, id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.ApiKey,
		&account.Balance,
		&createdAt,
		&updateAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrAccountNotFound //não encontrou nada
		}
		return nil, err //erro diferente de não encontrado
	}
	account.CreatedAt = createdAt
	account.UpdatedAt = updateAt
	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin() //inicia uma transação
	if err != nil {
		return err
	}
	defer tx.Rollback() //se der erro, desfaz a transação

	var currentBalance float64
	err = tx.QueryRow(`
		SELECT balance 
		FROM account
		WHERE id = $1
		FOR UPDATE
	`, account.ID).Scan(&currentBalance) //pega o saldo atual da contam, Scan pega o valor que ira ser recebido
	//FOR UPDATE = bloqueia a linha para que outros processos não possam acessá-la enquanto a transação estiver em andamento

	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE accounts
		SET balance = $1, updated_at = $2
		WHERE id = $3
	`, currentBalance+account.Balance, time.Now(), account.ID)
	if err != nil {
		return err
	}
	return tx.Commit() //se tudo der certo, confirma a transação
}
