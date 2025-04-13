package domain //sempre sera o nome da pasta
//orientado a estrutura de dados
import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
) 

type Account struct { //tipo c
	ID string 
	Name string 
	Email string 
	ApiKey string 
	Balance float64 
	mu sync.Mutex //para garantir que o acesso a variavel Balance seja seguro em ambientes concorrentes, race conditions
	CreatedAt time.Time 
	UpdatedAt time.Time 
}
//assim como existe um struct, existe uma função construtora (não metodo pois não é POO)

func GenerateApiKey() string { //não é necessário exportar, pois não será utilizado fora do pacote
	b := make([]byte, 16) //cria um slice de bytes com 16 posições
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewAccount(name, email string) *Account { //*Account é um ponteiro para a struct Account 
	account := &Account{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		ApiKey: GenerateApiKey(),
		Balance: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

//func (a *Account) = como se fosse um metodo (o ponteiro esta para o a para podermos acessar)
func (a *Account) AddBalance(amount float64) {
	a.mu.Lock() //bloqueia a var balance enquanto o processo estiver em execução
	defer a.mu.Unlock() //vai ser exececutado somente quando a funçao terminar ()

	a.Balance += amount
	a.UpdatedAt = time.Now()
}