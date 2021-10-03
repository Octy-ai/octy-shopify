package api

//business logic interface

type Customer interface {
	IsOctyProfileID(ID *string) (bool, error)
}

type Content interface {
	SelectOctyTemplate(profile map[string]interface{}, templates *[]map[string]interface{}, sections *[]map[string]interface{}) (*[]map[string]string, *[]map[string]string, error)
	MapSectionContent(sections *[]map[string]interface{},
		templateMaps *[]map[string]string,
		messages *[]map[string]interface{},
		messagesType string) []map[string]string
}
