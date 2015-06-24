package main

func GivenGroupIdFindUsers(id int) (result string, err error) {

	query := "SELECT u.name FROM groups AS g INNER JOIN user_group_relations AS ugr ON g.id = ugr.group_id AND  g.id = $1 INNER JOIN users AS u ON u.id = ugr.user_id"

	rows, err := Db.Query(query, id)
	for rows.Next() {
		var st string
		rows.Scan(&st)
		result += st + "\n"
	}
	return
}

func GivenIdFindGroups(id int) (result string, err error) {

	query := "SELECT g.name FROM users AS u INNER JOIN user_group_relations AS ugr ON u.id = ugr.user_id AND u.id = $1 INNER JOIN groups AS g ON g.id = ugr.group_id"

	rows, err := Db.Query(query, id)
	for rows.Next() {
		var st string
		rows.Scan(&st)
		result += st + "\n"
	}
	return
}
