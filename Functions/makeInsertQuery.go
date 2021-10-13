package Functions

/**
 * @function MakeInsertQueryString - Create a query to insert a new record
 * @param {string} table - The table name
 * @param {[]string} columns - The columns names
 * @param {[]string} values - The values to insert
 * @returns {string} - The query
 */
func MakeInsertQueryString(table string, columns []string, values []string) string {
	query := "INSERT INTO " + table + " ("
	for i := 0; i < len(columns); i++ {
		query += columns[i]
		if i != len(columns)-1 {
			query += ", "
		}
	}
	query += ") VALUES ("
	for i := 0; i < len(values); i++ {
		query += values[i]
		if i != len(values)-1 {
			query += ", "
		}
	}
	query += ")"
	return query
}

func MakeMultipleInsertQueryString(table string, column []string, value [][]string) []string {
	var queries []string
	for i := 0; i < len(value); i++ {
		queries = append(queries, MakeInsertQueryString(table, column, value[i]))
	}
	return queries
}
