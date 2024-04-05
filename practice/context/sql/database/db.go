package db

import "database/sql"

var Conn *sql.DB

// quando abrimos uma conexão com um banco de dados, é retornado na verdade um pool de conexões,
// que podem ser usadas entre várias funções ao mesmo tempo, desde que tenha alguma disponível
// [
//   conn,
//   conn,
//   conn,
//   conn,
//   conn,
//   conn,
//   conn,
// ]
