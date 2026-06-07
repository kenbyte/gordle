package main

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./gordle.db")
	if err != nil {
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS players (
    	id TEXT PRIMARY KEY,
    	name TEXT NOT NULL,
    	game_history TEXT,
    	creation_date TEXT NOT NULL,
    	score INTEGER NOT NULL DEFAULT 0,
    	rank INTEGER NOT NULL DEFAULT 0
	);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func addPlayer(db *sql.DB, p *Player) error {
	sqlStatement := `
    INSERT INTO players (
        id,
        name,
        game_history,
        creation_date,
        score,
        rank
    ) VALUES (?, ?, ?, ?, ?, ?)
	`

	historyJSON, err := json.Marshal(p.GameHistory)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		sqlStatement,
		p.ID.String(),
		p.Name,
		string(historyJSON),
		p.CreationDate,
		p.Score,
		p.Rank,
	)

	return err
}

func getAllPlayers(db *sql.DB) ([]Player, error) {
	sqlStatement := `
    SELECT
        id,
        name,
        game_history,
        creation_date,
        score,
        rank
    FROM players
	`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []Player

	for rows.Next() {
		var pl Player

		var idStr string
		var historyJSON string

		err := rows.Scan(
			&idStr,
			&pl.Name,
			&historyJSON,
			&pl.CreationDate,
			&pl.Score,
			&pl.Rank,
		)
		if err != nil {
			return nil, err
		}

		pl.ID, err = uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(
			[]byte(historyJSON),
			&pl.GameHistory,
		)
		if err != nil {
			return nil, err
		}

		players = append(players, pl)
	}

	return players, rows.Err()
}

func getPlayerByName(db *sql.DB, n string) (Player, error) {
	sqlStatement := `
    SELECT
        id,
        name,
        game_history,
        creation_date,
        score,
        rank
    FROM players
    WHERE name = ?
	`

	row := db.QueryRow(sqlStatement, n)

	var pl Player
	var idStr string
	var historyJSON string

	err := row.Scan(
		&idStr,
		&pl.Name,
		&historyJSON,
		&pl.CreationDate,
		&pl.Score,
		&pl.Rank,
	)
	if err != nil {
		return pl, err
	}

	pl.ID, err = uuid.Parse(idStr)
	if err != nil {
		return pl, err
	}

	err = json.Unmarshal(
		[]byte(historyJSON),
		&pl.GameHistory,
	)
	if err != nil {
		return pl, err
	}

	return pl, nil
}

func deletePlayerByName(db *sql.DB, name string) error {

	sqlStatement := `
    DELETE FROM players
    WHERE name = ?
	`

	_, err := db.Exec(sqlStatement, name)

	return err
}

func initGameTable(db *sql.DB) error {

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS games (
		id TEXT PRIMARY KEY,
		start_time TEXT NOT NULL,
		is_finished INTEGER NOT NULL,
		is_single INTEGER NOT NULL,
		word TEXT NOT NULL
	);
	`

	_, err := db.Exec(createTableSQL)

	return err
}

func addGame(db *sql.DB, g *Game) error {

	sqlStatement := `
	INSERT INTO games (
		id,
		start_time,
		is_finished,
		is_single,
		word
	) VALUES (?, ?, ?, ?, ?)
	`

	_, err := db.Exec(
		sqlStatement,

		g.ID.String(),

		g.StartTime.Format(time.RFC3339),

		g.IsFinished,

		g.IsSingle,

		g.Word,
	)

	return err
}

func getAllGames(db *sql.DB) ([]Game, error) {

	sqlStatement := `
	SELECT
		id,
		start_time,
		is_finished,
		is_single,
		word
	FROM games
	`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var games []Game

	for rows.Next() {

		var g Game

		var idStr string
		var startTimeStr string

		err := rows.Scan(
			&idStr,
			&startTimeStr,
			&g.IsFinished,
			&g.IsSingle,
			&g.Word,
		)

		if err != nil {
			return nil, err
		}

		g.ID, err = uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		g.StartTime, err = time.Parse(
			time.RFC3339,
			startTimeStr,
		)

		if err != nil {
			return nil, err
		}

		games = append(games, g)
	}

	return games, rows.Err()
}

func getGameByID(
	db *sql.DB,
	id uuid.UUID,
) (Game, error) {

	sqlStatement := `
	SELECT
		id,
		start_time,
		is_finished,
		is_single,
		word
	FROM games
	WHERE id = ?
	`

	row := db.QueryRow(
		sqlStatement,
		id.String(),
	)

	var g Game

	var idStr string
	var startTimeStr string

	err := row.Scan(
		&idStr,
		&startTimeStr,
		&g.IsFinished,
		&g.IsSingle,
		&g.Word,
	)

	if err != nil {
		return g, err
	}

	g.ID, err = uuid.Parse(idStr)
	if err != nil {
		return g, err
	}

	g.StartTime, err = time.Parse(
		time.RFC3339,
		startTimeStr,
	)

	return g, err
}

func updateGame(db *sql.DB, g *Game) error {

	sqlStatement := `
	UPDATE games
	SET
		start_time = ?,
		is_finished = ?,
		is_single = ?,
		word = ?
	WHERE id = ?
	`

	_, err := db.Exec(
		sqlStatement,

		g.StartTime.Format(time.RFC3339),

		g.IsFinished,

		g.IsSingle,

		g.Word,

		g.ID.String(),
	)

	return err
}

func deleteGame(
	db *sql.DB,
	id uuid.UUID,
) error {

	sqlStatement := `
	DELETE FROM games
	WHERE id = ?
	`

	_, err := db.Exec(
		sqlStatement,
		id.String(),
	)

	return err
}
