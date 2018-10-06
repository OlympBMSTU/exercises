package db

import (
	"fmt"

	"github.com/HustonMmmavr/excercieses/entities"
	"github.com/jackc/pgx"
)

func SaveExcerciese(excerciese entities.ExcercieseEntity, poll *pgx.ConnPool) {
	_, err := poll.Exec(insert_excercise, excerciese.GetAuthorId(),
		excerciese.GetFileName(),
		excerciese.GetRightAnswer(),
		excerciese.GetLevel(),
	)

	fmt.Println(err)

}

func GetExcerciese(id uint) {

}

func GetExcercieseList() {

}
