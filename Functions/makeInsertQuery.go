package Functions

// CreateInsertString creates a insert query string
func CreateInsertString(table string, columns []string, values []string) string {
	query := "INSERT INTO " + table + " ("
	for i := 0; i < len(columns); i++ {
		query += columns[i]
		if i != len(columns)-1 {
			query += ", "
		}
	}
	query += ") VALUES ("
	for i := 0; i < len(values); i++ {
		query += "'" + values[i] + "'"
		if i != len(values)-1 {
			query += ", "
		}
	}
	query += ")"
	return query
}

// CreateInsertStringWithID creates some insert query string
func CreateMutilpleInsertString(table string, column []string, value [][]string) []string {
	var queries []string
	for i := 0; i < len(value); i++ {
		queries = append(queries, CreateInsertString(table, column, value[i]))
	}
	return queries
}
