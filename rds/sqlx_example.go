package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Person struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Deleted bool   `db:"deleted"`
}

type PersonScore struct {
	PersonID int    `db:"person_id"`
	Item     string `db:"item"`
	Score    int    `db:"score"`
}

func main() {
	dsn := "host=localhost port=5432 user=root dbname=person sslmode=disable password=root search_path=public"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("connection failed: %s", err)
	}

	// select persons
	persons := []Person{}
	err = db.Select(&persons, "SELECT id, name, deleted FROM person WHERE deleted = false")
	if err != nil {
		log.Fatalf("select failed: %s", err)
	}

	for _, p := range persons {
		fmt.Printf("%+v\n", p)
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatalln(err)
	}
	err = insert(tx)
	if err != nil {
		tx.Rollback()
		log.Fatalf("insert failed: %s", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully inserted person data.")

}

func insert(tx *sqlx.Tx) error {
	// insert person
	newPerson := &Person{
		Name:    "New person",
		Deleted: false,
	}

	// use NamedQuery insted of NamedExec to fetch ID by one query.
	// see: https://github.com/jmoiron/sqlx/issues/83
	rows, err := tx.NamedQuery("INSERT INTO person (name, deleted) VALUES (:name, :deleted) RETURNING id", newPerson)
	if err != nil {
		return err
	}
	if rows.Next() {
		rows.Scan(&newPerson.ID)
	}
	err = rows.Close()
	if err != nil {
		return err
	}

	// insert person_score
	personScores := []PersonScore{
		{PersonID: newPerson.ID, Item: "a", Score: 10},
		{PersonID: newPerson.ID, Item: "b", Score: 20},
	}
	query := `INSERT INTO person_score (person_id, item, score) VALUES (:person_id, :item, :score)`
	_, err = tx.NamedExec(query, personScores)
	if err != nil {
		return err
	}
	return nil
}
