package store

import "PyPs/models"
// FormQuery is a function that takes a struct and returns a string and a slice of values
func FormQuery(u *models.Task) (string, []interface{}) {

	var Query string
	var values []interface{}

	if u.Id < 0 {
		return "", nil
	}
	if u.Assignee != "" {
		Query += "assignee=?,"
		values = append(values, u.Assignee)
	}
	if u.Description != "" {
		Query += "description=?,"
		values = append(values, u.Description)
	}
	if u.Status != "" {
		Query += "status=?,"
		values = append(values, u.Status)
	}
	if u.ProjectId > 0 {
		Query += "project_id=?,"
		values = append(values, u.ProjectId)
	}
	if Query == "" {
		return "", nil
	}
	Query = Query[:len(Query)-1]
	return Query, values
	
}