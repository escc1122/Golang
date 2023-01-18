package main

import (
	"reflect"
	"strings"

	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
)

func addHeading1(doc *document.Document, text string) {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.Properties().SetBold(true)
	para.SetStyle("Heading1")
	run.AddText(text)
}

func addText(doc *document.Document, text string) {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(text)
}

func addIntroduction(doc *document.Document) {
	addHeading1(doc, "Introduction")
	addText(doc, "說明API 用途")
}

func addRequestParameters(doc *document.Document, data interface{}) {
	addHeading1(doc, "Request parameters")

	addRequestResponseParameter(doc, reflect.ValueOf(data))
}

func addResponseFields(doc *document.Document, data interface{}) {
	addHeading1(doc, "Response fields")

	addRequestResponseParameter(doc, reflect.ValueOf(data))
}

func formatType(t string) string {

	return strings.Replace(t, "main.", "", 1)
}

func addRequestResponseParameter(doc *document.Document, value reflect.Value) {
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	t := value.Type()

	addText(doc, formatType(value.Type().String()))

	var objectArray []reflect.Value

	{
		table := doc.AddTable()
		table.Properties().SetWidthPercent(100)
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)
		hdrRow := table.AddRow()
		{
			cell := hdrRow.AddCell()
			cellPara := cell.AddParagraph()
			cellPara.Properties().SetAlignment(wml.ST_JcLeft)
			run := cellPara.AddRun()
			run.Properties().SetBold(true)
			run.AddText("Attribute")
		}
		{
			cell := hdrRow.AddCell()
			cellPara := cell.AddParagraph()
			cellPara.Properties().SetAlignment(wml.ST_JcLeft)
			run := cellPara.AddRun()
			run.Properties().SetBold(true)
			run.AddText("Type")
		}
		{
			cell := hdrRow.AddCell()
			cellPara := cell.AddParagraph()
			cellPara.Properties().SetAlignment(wml.ST_JcLeft)
			run := cellPara.AddRun()
			run.Properties().SetBold(true)
			run.AddText("Required")
		}
		{
			cell := hdrRow.AddCell()
			cellPara := cell.AddParagraph()
			cellPara.Properties().SetAlignment(wml.ST_JcLeft)
			run := cellPara.AddRun()
			run.Properties().SetBold(true)
			run.AddText("Description")
		}

		for i := 0; i < value.NumField(); i++ {
			field := t.Field(i)
			valueField := value.Field(i)

			if valueField.Type().Kind() == reflect.Struct {
				bbbb := reflect.New(valueField.Type()).Interface()
				objectArray = append(objectArray, reflect.ValueOf(bbbb).Elem())
			}

			row := table.AddRow()
			{
				cell := row.AddCell()
				cell.AddParagraph().AddRun().AddText(field.Tag.Get("json"))
			}

			{
				cell := row.AddCell()
				cell.AddParagraph().AddRun().AddText(formatType(t.Field(i).Type.String()))
			}

			for j := 0; j < 2; j++ {
				cell := row.AddCell()
				cell.AddParagraph().AddRun()
				//cell := row.AddCell()
				//cell.AddParagraph().AddRun().AddText()
			}
		}

		doc.AddParagraph()

		for _, v := range objectArray {
			addRequestResponseParameter(doc, v)
		}
	}
}
