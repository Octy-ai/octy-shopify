package content

import (
	"encoding/json"
	"testing"
)

func TestSelectOctyTemplate(t *testing.T) {

	content := New()

	tempaltes := []map[string]interface{}{
		{
			"TemplateID": "12345",
			"Metadata": map[string]interface{}{
				"churn_pred":      "high",
				"rfm_score_upper": float64(222),
				"rfm_score_lower": float64(111),
				"section_id":      "footer",
			},
		},
		{
			"TemplateID": "54321",
			"Metadata": map[string]interface{}{
				"churn_pred":      "high",
				"rfm_score_upper": float64(444),
				"rfm_score_lower": float64(333),
				"section_id":      "header",
			},
		},
		{
			"TemplateID": "67890",
			"Metadata": map[string]interface{}{
				"churn_pred":      "low",
				"rfm_score_upper": float64(444),
				"rfm_score_lower": float64(222),
				"section_id":      "body",
			},
		},
		{
			"TemplateID": "09876",
			"Metadata": map[string]interface{}{
				"churn_pred":      "mid",
				"rfm_score_upper": float64(333),
				"rfm_score_lower": float64(111),
				"section_id":      "footer",
			},
		},
	}

	sections := []map[string]interface{}{
		{
			"SectionID":    "header",
			"DefaultValue": "Hi",
		},
		{
			"SectionID":    "body",
			"DefaultValue": "Welcome to our store",
		},
	}

	tables := []struct {
		profile      map[string]interface{}
		tempaltes    []map[string]interface{}
		sections     []map[string]interface{}
		templateMaps []map[string]string
		defaultMaps  []map[string]string
	}{

		{
			map[string]interface{}{ //profile
				"ChurnProbability": "high",
				"Rfm":              float64(342),
			}, tempaltes, sections, []map[string]string{ // return values
				{
					"SectionID":  "header",
					"TemplateID": "54321",
				},
			},
			[]map[string]string{
				{
					"SectionID":  "body",
					"TemplateID": "none",
				},
			},
		}, // test 1

		{
			map[string]interface{}{ //profile
				"ChurnProbability": "low",
				"Rfm":              float64(111),
			}, tempaltes, sections, nil,
			[]map[string]string{
				{
					"SectionID":  "header",
					"TemplateID": "none",
				},
				{
					"SectionID":  "body",
					"TemplateID": "none",
				},
			},
		}, // test 2

	}

	for _, table := range tables {

		templateMaps, defaultMaps, err := content.SelectOctyTemplate(table.profile, &table.tempaltes, &table.sections)
		if err != nil {
			t.Errorf("Error occurred when running this test: %v", err)
		}

		if mapToJsonString(templateMaps) != mapToJsonString(&table.templateMaps) {
			t.Errorf("Incorrect template recommended, got: %v, want: %v.", mapToJsonString(templateMaps), mapToJsonString(&table.templateMaps))
		}

		if mapToJsonString(defaultMaps) != mapToJsonString(&table.defaultMaps) {
			t.Errorf("Incorrect default map returned, got: %v, want: %v.", mapToJsonString(defaultMaps), mapToJsonString(&table.defaultMaps))
		}

	}

}

func TestMapSectionContent(t *testing.T) {

	content := New()
	sections := []map[string]interface{}{
		{
			"SectionID":    "header",
			"DefaultValue": "Hi",
		},
		{
			"SectionID":    "body",
			"DefaultValue": "Welcome to our store",
		},
	}

	tables := []struct {
		sections       []map[string]interface{}
		templateMaps   []map[string]string
		messages       []map[string]interface{}
		messagesType   string
		returnSections []map[string]string
	}{
		{
			sections,
			[]map[string]string{
				{
					"SectionID":  "header",
					"TemplateID": "54321",
				},
				{
					"SectionID":  "body",
					"TemplateID": "none",
				},
			},
			[]map[string]interface{}{
				{
					"TemplateID": "54321",
					"Content":    "We don't want to loose you! Can we talk about this?",
				},
			},
			"generated",
			[]map[string]string{
				{
					"sectionID": "header",
					"content":   "We don't want to loose you! Can we talk about this?",
				},
			},
		}, // test 1
		{
			sections,
			[]map[string]string{
				{
					"SectionID":  "header",
					"TemplateID": "54321",
				},
				{
					"SectionID":  "body",
					"TemplateID": "none",
				},
			},
			[]map[string]interface{}{
				{
					"TemplateID": "54321",
					"Content":    "We don't want to loose you! Can we talk about this?",
				},
			},
			"failed",
			[]map[string]string{
				{
					"sectionID": "header",
					"content":   "Hi",
				},
			},
		}, // test 2
	}

	for _, table := range tables {

		returnSections := content.MapSectionContent(
			&table.sections,
			&table.templateMaps,
			&table.messages,
			table.messagesType,
		)

		if mapToJsonString(&returnSections) != mapToJsonString(&table.returnSections) {
			t.Errorf("Incorrect return value, got: %v, want: %v.", mapToJsonString(&returnSections), mapToJsonString(&table.returnSections))
		}
	}

}

func mapToJsonString(sliceMap *[]map[string]string) string {
	JsonStr, _ := json.Marshal(sliceMap)
	return string(JsonStr)
}
