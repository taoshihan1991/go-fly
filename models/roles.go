package models
type Role struct{
	Id string `json:"role_id"`
	Name string `json:"role_name"`
}
func FindRoles()[]Role{
	var roles []Role
	DB.Order("id desc").Find(&roles)
	return roles
}