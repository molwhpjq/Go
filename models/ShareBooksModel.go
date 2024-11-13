package models

type ShareBooks struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Category  string `json:"category"`
}

func (sharebooks *ShareBooks) sharebooksToRespDesc() interface{} {
	respShareBooks := map[string]interface{}{
		"id":        sharebooks.Id,
		"title":     sharebooks.Title,
		"author":    sharebooks.Author,
		"publisher": sharebooks.Publisher,
		"category":  sharebooks.Category,
	}
	return respShareBooks
}
