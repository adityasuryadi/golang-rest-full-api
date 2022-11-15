package simple

type Database struct {
	Name string
}

// alias untuk google wire karena google wire tidak mendukung multiple param dengan tipe data yang sama
type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"}) //paksa konversi ke alias
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongDB"})
}

type DatabaseRepository struct {
	DataBasePostgreSQL *DatabasePostgreSQL
	DataBaseMongoDB    *DatabaseMongoDB
}

func NewDataBaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DataBasePostgreSQL: postgreSQL,
		DataBaseMongoDB:    mongoDB,
	}
}
