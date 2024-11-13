package models

type Profile struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

func (profile *Profile) ProfileToResponse() interface{} {
	respProfile := map[string]interface{}{
		"id":           profile.Id,
		"name":         profile.Name,
		"introduction": profile.Introduction,
	}
	return respProfile
}
