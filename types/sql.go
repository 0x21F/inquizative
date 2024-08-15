package types

import "database/sql"

const UP_TABLES string = `
	CREATE TABLE IF NOT EXISTS users (
		id BLOB PRIMARY KEY,
		email TEXT NOT NULL,
		public INT DEFAULT 0 
	);
	
	CREATE TABLE IF NOT EXISTS course (
		course_id BLOB PRIMARY KEY,
		user_id BLOB NOT NULL,
		course_name TEXT NOT NULL,
		FOREIGN KEY (user_id) 
			REFERENCES users (id) 
				ON DELETE CASCADE 
				ON UPDATE CASCADE 
	);

	CREATE TABLE IF NOT EXISTS topics (
		topic_id BLOB PRIMARY KEY,
		course_id BLOB NOT NULL,
		topic_name text NOT NULL,
		FOREIGN KEY (course_id) 
			REFERENCES course (course_id) 
				ON DELETE CASCADE
				ON UPDATE CASCADE
	);

	CREATE TABLE IF NOT EXISTS summaries (
		summary_id BLOB PRIMARY KEY,
		topic_id NOT NULL,
		content TEXT NOT NULL,
		filename TEXT NOT NULL,
		FOREIGN KEY (topic_id)  
			REFERENCES topics (topic_id) 
				ON DELETE CASCADE 
				ON UPDATE CASCADE 
	);

	CREATE TABLE IF NOT EXISTS flashcard_deck (
		deck_id BLOB PRIMARY KEY,
		user_id BLOB NOT NULL,
		topic_id BLOB NOT NULL,
		FOREIGN KEY (user_id)
			REFERENCES users (id)  
				ON DELETE CASCADE
				ON UPDATE CASCADE,
		FOREIGN KEY (topic_id) 
			REFERENCES topics (topic_id) 
				ON DELETE CASCADE 
				ON UPDATE CASCADE
	);

	CREATE TABLE IF NOT EXISTS flashcard (
		flashcard_id BLOB PRIMARY KEY,
		deck_id BLOB NOT NULL,
		front TEXT NOT NULL,
		content TEXT NOT NULL,
		add_notes TEXT NOT NULL,
		FOREIGN KEY (deck_id)
			REFERENCES flashcard_deck (deck_id) 
				ON DELETE CASCADE
				ON UPDATE CASCADE
	);
`

// good ol tommy drop tables
const DOWN_TABLES string = `
	DROP TABLE IF EXISTS users;
	
	DROP TABLE IF EXISTS course;

	DROP TABLE IF EXISTS topics;

	DROP TABLE IF EXISTS summaries;

	DROP TABLE IF EXISTS flashcard_deck;

	DROP TABLE IF EXISTS flashcard;
`

func UpTables(db *sql.DB) error {
	_, err := db.Exec(UP_TABLES)
	if err != nil {
		return err
	}

	return nil
}

func DownTables(db *sql.DB) error {
	_, err := db.Exec(DOWN_TABLES)
	if err != nil {
		return err
	}

	return nil
}
