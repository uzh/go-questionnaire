package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func eachMonth2inQ(q *qst.QuestionnaireT) error {

	if q.Survey.MonthOfQuarter() != 2 {
		return nil
	}

	if q.Survey.Year == 2021 && q.Survey.Month == 8 {
		return nil
	}

	if q.Survey.Year == 2021 && q.Survey.Month == 11 {
		return nil
	}

	if q.Survey.Year == 2022 && q.Survey.Month == 2 {
		return nil
	}

	if q.Survey.Year == 2022 && q.Survey.Month == 5 {
		return nil
	}

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.2rem 0 0"

	/*
		SELECT
			frage_kurz,
			GROUP_CONCAT( DISTINCT antwort ORDER BY antwort) aw,
			count(*) anz
		FROM sonderfragen_ger
		WHERE survey_id = 202005
		GROUP BY frage_kurz
	*/

	page := q.AddPage()
	// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
	page.Label = trl.S{
		"de": "Sonderfrage: Inflation und Geldpolitik",
		"en": "Special: Inflation and monetary policy",
	}
	page.Short = trl.S{
		"de": "Inflation,<br>Geldpolitik",
		"en": "Inflation,<br>Mon. Policy",
	}
	page.WidthMax("48rem")
	page.ValidationFuncName = "fmt-m2-p6"
	page.ValidationFuncMsg = trl.S{
		"de": "Ihre Antworten auf Frage 1b addieren sich nicht zu 100%. Wirklich weiter?",
		"en": "Your answers to question 1b dont add up to 100%. Continue anyway?",
	}

	{
		gr := page.AddGroup()
		gr.Cols = 9
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Width = "70%"
		gr.Style.Mobile.StyleBox.Width = "100%"

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 9
			// inp.ColSpanLabel = 12
			inp.Label = trl.S{
				"de": `
					<b>1a.</b> &nbsp; Punktprognose der <b>jährlichen Inflationsrate im Euroraum</b>
					<br>
					(durchschnittliche jährliche Veränderung des HICP in Prozent)
					<!-- Anstieg des HICP von Jan bis Dez; Erwartungswert -->
				`,
				"en": `
					<b>1a.</b> &nbsp; Forecast of <b>annual inflation rate in the euro area</b>
					<br>
					(annual average change of the HICP, in percent)
					<!-- Avg. percentage change in HICP from Jan to Dec; -->
				`,
			}
		}

		for idx := range []int{0, 1, 2} {

			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("ppjinf_jp%v", idx) //"p1_y1"
			inp.Min = -10
			inp.Max = +20
			inp.Validator = "inRange20"
			inp.MaxChars = 5
			inp.Step = 0.01
			inp.Label = trl.S{
				"de": q.Survey.YearStr(idx),
				"en": q.Survey.YearStr(idx),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}

			inp.ColSpan = 3
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 2

			inp.StyleLbl = lblStyleRight
		}

	}

	//
	//
	//
	//
	//
	// gr2
	{

		// colspan := float32(2 + 4*3 + 2 + 2)

		gr := page.AddGroup()
		gr.Cols = 6
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Mobile.StyleGridContainer.GapRow = "0.02rem"
		// gr.WidthMax("30rem")
		gr.ColWidths("1.6fr    2.7fr 3.1fr 3.1fr 2.4fr    2.4fr  1.4fr")

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 7
			inp.Label = trl.S{
				"de": `<b>1b.</b> &nbsp; Wir möchten gerne von Ihnen erfahren, 
						für wie wahrscheinlich Sie bestimmte Ausprägungen 
						der durchschnittlichen jährlichen Inflationsrate 
						in den folgenden Jahren halten.
						
						<br>
						<i>
						Bitte stellen Sie sicher, 
						dass die Summen der Wahrscheinlichkeiten 
						in den Zeilen jeweils 100% ergeben.
						</i>

						
						`,
				"en": `<b>1b.</b> &nbsp; 
						How likely are specific future realizations of inflation? 
						
						Please give us your assessments for the annual average inflation rate 
						in the euro area.
						
						<br>
						<i>
						Please ensure that the probabilities 
						in every line add up to 100%.
						</i>
						`,
			}

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Mobile.StyleBox.Padding = "0 0 0.8rem 0"

		}
		// first row: labels
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = trl.S{
				"de": "&nbsp;",
				"en": "&nbsp;",
			}
		}

		labels := []trl.S{
			{
				"de": "unter <br>2&nbsp;Prozent",
				"en": "below <br>2&nbsp;percent",
			},
			{
				// "de": "zwischen 1&nbsp;u.&nbsp;2&nbsp;Prozent",
				"de": "zwischen 2 u.  <br>3&nbsp;Prozent",
				"en": "between  2 and <br>3&nbsp;percent",
			},
			{
				// "de": "zwischen 2&nbsp;u.&nbsp;3&nbsp;Prozent",
				"de": "zwischen 3 u.  <br>4&nbsp;Prozent",
				"en": "between  3 and <br>4&nbsp;percent",
			},
			{
				"de": "größer als 4&nbsp;Prozent",
				"en": "above  <br>4&nbsp;percent",
			},
		}

		// first row - cols 2-5
		for _, lbl := range labels {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = lbl
			inp.Style = css.ItemStartCA(inp.Style)
			inp.Style.Mobile.StyleBox.Padding = " 0 0.3rem 0 0" // prevent overlapping of columns in narrow mobile view

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Desktop.StyleText.LineHeight = 118

		}

		//
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = trl.S{
				"de": "&nbsp;&nbsp;&nbsp;&#931;", // greek SUM sign
				"en": "&nbsp;&nbsp;&nbsp;&#931;",
			}
			inp.Style = css.ItemCenteredMCA(inp.Style)
			inp.Style = css.ItemStartMA(inp.Style)
			inp.Style = css.ItemCenteredCA(inp.Style)

			inp.Style = css.TextStart(inp.Style)
			inp.Style = css.TextCACenter(inp.Style)

			inp.Style.Desktop.StyleText.FontSize = 120

		}
		{
			inp := gr.AddInput()
			inp.ColSpan = 1
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": "keine Ang.",
				"en": "no answer",
			}
			inp.Style = css.ItemCenteredMCA(inp.Style)
			inp.Style = css.ItemStartCA(inp.Style)
			inp.Style = css.TextStart(inp.Style)

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Desktop.StyleText.LineHeight = 118
		}

		//
		// second to fourth row: inputs
		for i := 2022; i <= 2024; i++ {

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Label = trl.S{
					"de": fmt.Sprint(i),
					"en": fmt.Sprint(i),
				}
			}

			inpNames := []string{
				"under2", "between2and3", "between3and4", "above4",
			}

			for _, inpname := range inpNames {
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("inf%v_%v", i, inpname)
				inp.Suffix = trl.S{"de": "%", "en": "%"}
				// inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.ColSpan = 1
				inp.ColSpanControl = 3
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0
				inp.MaxChars = 3
			}

			// last two cols
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Label = trl.S{
					"de": "100&nbsp;%",
					"en": "100&nbsp;%",
				}
				// inp.Style = css.ItemCenteredMCA(inp.Style)
				inp.Style = css.ItemStartMA(inp.Style)
				inp.Style = css.ItemStartCA(inp.Style)
				inp.Style = css.TextStart(inp.Style)

			}
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.ColSpan = 1
				inp.Name = fmt.Sprintf("inf%v__noanswer", i)
				inp.ColSpanControl = 1
				// inp.Style = css.ItemCenteredMCA(inp.Style)
				inp.Style = css.ItemStartMA(inp.Style)
				inp.Style = css.ItemStartCA(inp.Style)

			}
		}
	}

	// gr2
	colTemplate, colsRowFree, styleRowFree := colTemplateWithFreeRow()

	{
		rowLabelsEconomicAreasShort := []trl.S{
			{
				"de": "Konjunkturentwicklung im Eurogebiet",
				"en": "Development of GDP in the euro area",
			},
			{
				"de": "Entwicklung der Löhne im Eurogebiet",
				"en": "Development of wages in the euro area",
			},
			{
				"de": "Entwicklung der Energiepreise",
				"en": "Development of energy prices",
			},
			{
				"de": "Entwicklung der Rohstoffpreise (ohne Energiepreise)",
				"en": "Development of prices for raw materials (except energy) ",
			},
			{
				"de": "Veränderung der Wechselkurse (relativ zum Euro)",
				"en": "Changes in exchange rates (relative to the euro)",
			},
			{
				"de": "Geldpolitik der EZB",
				"en": "Monetary policy of the ECB",
			},
			{
				"de": "Internationale Handelskonflikte",
				"en": "International trade conflicts",
			},
			{
				"de": "Internationale Lieferengpässe",
				"en": "International supply bottlenecks",
			},
			{
				"de": "Corona-Pandemie",
				"en": "Covid pandemic",
			},
			{
				"de": "Grüne Transformation",
				"en": "Green transformation",
			},
			{
				"de": "Krieg in der Ukraine",
				"en": "War in the Ukraine",
			},
		}

		gb := qst.NewGridBuilderRadios(
			colTemplate,
			labelsPlusPlusMinusMinus(),
			// prefix ioi_ => impact on inflation
			//   but we stick to rev_ => revision
			[]string{
				"rev_bus_cycle_ea",
				"rev_wages_ea",
				"rev_energy_prices",
				"rev_commodity_prices",
				"rev_exch_rates",
				"rev_mp_ecb",
				"rev_trade_conflicts",
				"rev_supply_shortages",
				"rev_corona",
				"rev_green_trafo",
				"rev_war_ukraine",
			},
			radioVals6,
			rowLabelsEconomicAreasShort,
		)
		idxThreeMonthsBefore := trl.MonthsShift(int(q.Survey.Month), -3)
		monthMinus3 := trl.MonthByInt(idxThreeMonthsBefore)
		gb.MainLabel = trl.S{
			"de": fmt.Sprintf(`<b>2.</b> &nbsp; 
				Haben Entwicklungen in den folgenden Bereichen Sie zu einer Revision 
				Ihrer <b>Inflationsprognosen</b> für den Euroraum (ggü. %v 2022) bewogen 
				und wenn ja, nach oben (+) oder unten (-)?
				<br>
				<br>
				<b>Für die Jahre 2022, 2023 und 2024</b>
			`, monthMinus3.Tr("de"),
			),
			"en": fmt.Sprintf(`<b>2.</b> &nbsp;
				What are the main factors leading you to change your inflation forecasts
				for the euro area (in comparison to your forecasts as of %v 2022).
				(+) means increase in inflation forecast,
				(-) means decrease in inflation forecast.
				<br>
				<br>
				<b>For the years 2022, 2023, and 2024</b>
			`, monthMinus3.Tr("en")),
		}
		gr := page.AddGrid(gb)
		gr.OddRowsColoring = true
		gr.BottomVSpacers = 1

	}
	{

		//
		// row free input
		gr := page.AddGroup()
		gr.Cols = float32(len(improvedDeterioratedPlusMinus6()) + 1)
		gr.Cols = 7

		gr.Style = css.NewStylesResponsive(gr.Style)
		if gr.Style.Desktop.StyleGridContainer.TemplateColumns == "" {
			gr.Style.Desktop.StyleBox.Display = "grid"
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = styleRowFree
			// log.Printf("fmt special 2021-09: grid template - %v", stl)
		} else {
			return fmt.Errorf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", styleRowFree, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
		}

		gr.BottomVSpacers = 4

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "rev_free_label"
			// inp.MaxChars = 17
			inp.MaxChars = 15
			inp.ColSpan = 1
			inp.ColSpanLabel = 2.4
			inp.ColSpanControl = 4
			inp.Label = trl.S{
				"de": "Andere",
				"en": "Other",
			}
		}

		//
		for idx := 0; idx < len(improvedDeterioratedPlusMinus6()); idx++ {
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "rev_free"
			rad.ValueRadio = fmt.Sprint(idx + 1)
			rad.ColSpan = 1
			rad.ColSpanLabel = colsRowFree[2*(idx+1)]
			rad.ColSpanControl = colsRowFree[2*(idx+1)] + 1
		}

	}

	// // gr2a
	// {
	// 	gr := page.AddGroup()
	// 	gr.Cols = 1
	// 	gr.BottomVSpacers = 1
	// 	gr.BottomVSpacers = 0
	// 	{
	// 		inp := gr.AddInput()
	// 		inp.Type = "text"
	// 		inp.Name = "rev_free_label"
	// 		inp.MaxChars = 24
	// 		inp.ColSpan = 1
	// 		inp.ColSpanControl = 1
	// 		inp.Label = nil
	// 		inp.Placeholder = trl.S{"de": "Sonstige", "en": "Other"}
	// 	}
	// }

	// // gr2b
	// {
	// 	gb := qst.NewGridBuilderRadios(
	// 		columnTemplateLocal,
	// 		nil,
	// 		[]string{"rev_free"},
	// 		radioVals6,
	// 		[]trl.S{
	// 			{
	// 				"de": " &nbsp;  ", // -
	// 				"en": " &nbsp;  ", // -
	// 			},
	// 		},
	// 	)
	// 	gb.MainLabel = nil
	// 	gr := page.AddGrid(gb)
	// 	gr.OddRowsColoring = true
	// }

	// gr3
	{
		// 2019	18 Sep. 0.00
		latestECBRate, err := q.Survey.Param("main_refinance_rate_ecb")
		if err != nil {
			return fmt.Errorf("Set field 'main_refinance_rate_ecb' to `01.02.2018: 3.2%%` as in `main refinancing operations rate of the ECB (01.02.2018: 3.2%%)`; error was %v", err)
		}

		//
		//
		gr := page.AddGroup()
		gr.Cols = 12

		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Width = "70%"
		gr.Style.Mobile.StyleBox.Width = "100%"

		// row-1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.ColSpanLabel = 12
			inp.Label = trl.S{
				"de": fmt.Sprintf(
					`<b>3.</b> &nbsp; 
					Den <b>Hauptrefinanzierungssatz</b> der EZB (derzeit %v) erwarte ich 
					`, latestECBRate,
				),
				"en": fmt.Sprintf(
					`<b>3.</b> &nbsp; 
					I expect the <b>main refinancing facility rate</b> of the ECB (currently at %v) to be 
					`, latestECBRate,
				),
			}
		}

		lbls := []trl.S{
			{
				"de": "in 6&nbsp;Monaten",
				"en": "in 6&nbsp;months",
			},
			{
				"de": "Ende 2022",
				"en": "End of 2022",
			},
			{
				"de": "Ende 2023",
				"en": "End of 2023",
			},
			{
				"de": "Ende 2023",
				"en": "End of 2023",
			},
		}

		inputs := []string{
			"ezb6",
			"ezb2022",
			"ezb2023",
			"ezb2024",
		}

		// rows 2...5
		for i := 0; i < 4; i++ {
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = lbls[i]
				inp.StyleLbl = lblStyleRight
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%vmin", inputs[i]) // "ezb6min"
				inp.Min = -10
				inp.Max = +20
				inp.Validator = "inRange20"
				inp.MaxChars = 5
				inp.Step = 0.01

				inp.ColSpan = 5
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 2
				inp.Label = trl.S{
					"de": "zwischen&nbsp;",
					"en": "between&nbsp;",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.StyleLbl = lblStyleRight
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%vmax", inputs[i]) // "ezb6max"
				inp.Min = -10
				inp.Max = +20
				inp.Validator = "inRange20"
				inp.MaxChars = 5
				inp.Step = 0.01

				inp.ColSpan = 4
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 2
				inp.Label = trl.S{
					"de": "und",
					"en": "and",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
			}

		}

		//
		// row-6
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 5
			inp.Label = trl.S{
				"de": " &nbsp;",
				"en": " &nbsp;",
			}
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 7
			inp.Label = trl.S{
				"de": "[zentrales 90% Konfidenzintervall]",
				"en": "[central 90&nbsp;pct confidence interval]",
			}
			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)

			inp.StyleLbl.Desktop.StyleBox.Position = "relative"

			inp.StyleLbl.Desktop.StyleBox.Left = "2rem"
			inp.StyleLbl.Desktop.StyleBox.Top = "-0.2rem"
			inp.StyleLbl.Mobile.StyleBox.Left = "0"
			inp.StyleLbl.Mobile.StyleBox.Top = "0"

			inp.StyleLbl.Desktop.StyleText.FontSize = 90

		}

	}

	return nil

}
