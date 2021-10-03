package api

// Returns content for each given section relative to provided customer identifier.
func (api Application) GetContent(customerID string, sections *[]map[string]interface{}) ([]map[string]string, error) {

	profile, err := api.rest.GetOctyProfile(customerID)
	if err != nil {
		return nil, err
	}

	templates, err := api.rest.GetTemplates()
	if err != nil {
		return nil, err
	}

	templateMaps, defaultMaps, err := api.content.SelectOctyTemplate(profile, &templates, sections)
	if err != nil {
		return nil, err
	}

	var returnSections []map[string]string
	for _, section := range *defaultMaps {
		var defaultValue string
		for _, s := range *sections {
			if s["SectionID"] == section["SectionID"] {
				defaultValue = s["DefaultValue"].(string)
			}
		}
		returnSections = append(returnSections, map[string]string{
			"sectionID": section["SectionID"],
			"content":   defaultValue,
		})
	}

	generatedMessages, failedMessages, failedTemplates, err := api.rest.GenerateContent(*templateMaps)
	if err != nil {
		return nil, err
	}
	returnSections = append(returnSections, api.content.MapSectionContent(sections, templateMaps, &generatedMessages, "generated")...)
	returnSections = append(returnSections, api.content.MapSectionContent(sections, templateMaps, &failedMessages, "failed")...)
	returnSections = append(returnSections, api.content.MapSectionContent(sections, templateMaps, &failedTemplates, "failed")...)

	return returnSections, err

}
