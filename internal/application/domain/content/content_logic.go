package content

type Content struct {
}

func New() *Content {
	return &Content{}
}

// attributes used to comapre templates metdata and profile attributes
type requiredAttr struct {
	churnPred bool
	rfmScoreU bool
	rfmScoreL bool
}

// Compare each provided template metdata attributes against profile attributes to determine the best message
// template for each provided section, given the specified profile.
func (c Content) SelectOctyTemplate(profile map[string]interface{},
	templates *[]map[string]interface{},
	sections *[]map[string]interface{}) (*[]map[string]string, *[]map[string]string, error) {

	var templateMaps []map[string]string
	var defaultMaps []map[string]string

	for _, section := range *sections {

		sectionMap := make(map[string]string)
		sectionMap["SectionID"] = section["SectionID"].(string)
		sectionMap["TemplateID"] = "none"

		for _, template := range *templates {
			// check if the template has metadata attribute
			if !hasMetadataAttribute(template) {
				continue
			}
			if recommendTemplate(template["Metadata"].(map[string]interface{}),
				profile,
				section["SectionID"].(string)) {
				sectionMap["TemplateID"] = template["TemplateID"].(string)
				templateMaps = append(templateMaps, sectionMap)
				break
			}
		}

		if sectionMap["TemplateID"] == "none" {
			defaultMaps = append(defaultMaps, sectionMap)
		}

	}
	return &templateMaps, &defaultMaps, nil
}

// Map sections to generated or default content
func (c Content) MapSectionContent(sections *[]map[string]interface{},
	templateMaps *[]map[string]string,
	messages *[]map[string]interface{},
	messagesType string) []map[string]string {

	var returnSections []map[string]string

	for _, m := range *messages {
		for _, tm := range *templateMaps {
			if m["TemplateID"] == tm["TemplateID"] {
				for _, sec := range *sections {
					if sec["SectionID"] == tm["SectionID"] {
						switch {
						case messagesType == "failed":
							returnSections = append(returnSections, map[string]string{
								"sectionID": sec["SectionID"].(string),
								"content":   sec["DefaultValue"].(string),
							})
						case messagesType == "generated":
							returnSections = append(returnSections, map[string]string{
								"sectionID": sec["SectionID"].(string),
								"content":   m["Content"].(string),
							})
						default:
							break
						}
					}
				}
			}
		}
	}
	return returnSections
}

// <Private Functions>

// determine if template has a metadata attribute
func hasMetadataAttribute(template map[string]interface{}) bool {
	md := template["Metadata"].(map[string]interface{})
	return len(md) != 0
}

// determine if defined profile attributes match any template metadata tags
func recommendTemplate(metaData map[string]interface{},
	profile map[string]interface{}, sectionID string) bool {

	// In this basic implementation, we will simply attempt to match
	// profile rfm scores and churn probabilities to relevant template
	// metadata tags. In your own application, you could apply any degree of granularity
	// to find the ideal message template for this profile.

	// ensure template has matching required section identifier for this section
	// when setting section identifiers in tempalte metadata,
	// convention is to set the key name as 'section_id'.
	var sectionMatch bool = false
	for k, v := range metaData {
		if k == "section_id" {
			if v == sectionID {
				sectionMatch = true
			}
		}
	}
	if !sectionMatch {
		return false
	}

	// esnure profile has required values
	if profile["ChurnProbability"] == nil || profile["Rfm"] == nil {
		return false
	}

	// get values from profile and cast to type
	churnLabel := profile["ChurnProbability"].(string)
	rfmScore := profile["Rfm"].(float64)

	// used to track which required attributes have been found in template
	ra := requiredAttr{
		churnPred: false,
		rfmScoreU: false,
		rfmScoreL: false,
	}

	for k, v := range metaData {
		if k == "churn_pred" {
			if v == churnLabel {
				ra.churnPred = true
			}
			continue
		}

		if k == "rfm_score_upper" {
			// assert type
			if rfmScore <= v.(float64) {
				ra.rfmScoreU = true
			}
			continue
		}

		if k == "rfm_score_lower" {
			// assert type
			if rfmScore >= v.(float64) {
				ra.rfmScoreL = true
			}
			continue
		}

	}

	// assess if all require attributes where met between this profile and this template.
	if !ra.churnPred {
		return false
	}
	if !ra.rfmScoreL {
		return false
	}
	if !ra.rfmScoreU {
		return false
	}

	return true
}
