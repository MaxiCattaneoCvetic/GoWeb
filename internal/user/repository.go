package user

import (
	"context"
	"goWeb/v2/internal/domain"
	"log"
)

// generamos una estructura para la base de datos
type DB struct {
	Users     []domain.User // sera un slice de user
	MaxUserID uint64
}

// generamos la interfaz para la base de datos
type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error // Metodo create, que tendra el context  y despues vamos a recibir por parametro un usuario. (un puntero de usuario) y puede devolver un error
		GetAll(ctx context.Context) ([]domain.User, error)
	}
	// estos campos que tendremos en repositorio lo manejamos en minuscula porque son campos privados
	repo struct {
		db  DB
		log *log.Logger // para manejar el log
	}
)

func NewRepo(db DB, log *log.Logger) Repository {
	return &repo{
		db:  db,
		log: log,
	}
}

func (r *repo) Create(ctx context.Context, user *domain.User) error {

	r.db.MaxUserID++                       // generamos un ID
	user.ID = r.db.MaxUserID               // setemos el ID generado al usuario  nuevo
	r.db.Users = append(r.db.Users, *user) // lo agregamos al slice de usuarios
	r.log.Println("Repository log: User created", user)
	return nil

}
func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	log.Println("Repository log: Get all users")
	return r.db.Users, nil
}
