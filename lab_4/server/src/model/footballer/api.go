package footballer

import "lab_4/server/src/settings"

func GetFootballers() ([]Footballer, error) {

	var requestCoaches []Footballer = []Footballer{}

	rows, err := settings.DBinst.Queryx("SELECT  footballer_id, first_name, last_name, position, rating FROM footballer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var foolballer Footballer
		if err := rows.Scan(&foolballer.FootballerId, &foolballer.FirstName, &foolballer.LastNane, &foolballer.Position, &foolballer.Rating); err != nil {
			return nil, err
		}
		requestCoaches = append(requestCoaches, foolballer)
	}

	return requestCoaches, nil
}
