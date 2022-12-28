package repositorio

import (
	"api/internal/modelo"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositoUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repo Usuarios) Criar(usuario modelo.Usuario) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO usuarios(nome,nick,email,senha) VALUES (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}
	UltimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(UltimoIDInserido), nil
}
