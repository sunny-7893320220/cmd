package config

var DefaultConfig = []byte(`
application : "mongo_connection"
env : "local"
port : 8080

mysql :
    host : "localhost"
    port : 3306
    database : "test"
    username : ""
    password : ""
`)
