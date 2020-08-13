package models
type Role struct{
	Id string `json:"role_id"`
	Name string `json:"role_name"`
	Method string `json:"method"`
	Path string `json:"path"`
}
func FindRoles()[]Role{
	var roles []Role
	DB.Order("id desc").Find(&roles)
	return roles
}
func FindRole(id interface{})Role{
	var role Role
	DB.Where("id = ?", id).First(&role)
	return role
}