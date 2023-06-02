package database

const CreateIfNotExists = `CREATE TABLE IF NOT EXISTS urls(short_id TEXT PRIMARY KEY, url TEXT);`
const InsertNewUrl = `INSERT INTO urls(short_id, url) VALUES(?,?);`
const Select = `SELECT url FROM urls WHERE short_id = ?;`
