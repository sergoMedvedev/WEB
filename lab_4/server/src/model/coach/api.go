package coach

import (
	"lab_4/server/src/settings"
)

func GetCoaches() ([]Coach, error) {

	var requestCoaches []Coach = []Coach{}

	rows, err := settings.DBinst.Queryx("SELECT coach_id, first_name, last_name FROM coach")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var coach Coach
		if err := rows.Scan(&coach.CoachId, &coach.FirstName, &coach.LastName); err != nil {
			return nil, err
		}
		requestCoaches = append(requestCoaches, coach)
	}

	return requestCoaches, nil
}

func CreateCoach(coach Coach) error {
	_, err := settings.DBinst.Exec("INSERT INTO coach (first_name, last_name) VALUES ($1, $2)",
        coach.FirstName, coach.LastName)
	if err != nil {
		return  err
	}

	return nil
}
