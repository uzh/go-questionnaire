package pat

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// PersonalQuestions1 - numbered 8-15
func PersonalQuestions2(q *qst.QuestionnaireT, vE VariableElements) error {

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.3rem 0 0"

	{
		// page := q.AddPage()
		page := q.AddPage()
		// page.Label = trl.S{"de": "POP page"}
		// page.Short = trl.S{"de": "Stiftungen 1"}
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		{
			gr := page.AddGroup()
			gr.Cols = 12
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 12
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Bitte geben Sie Ihr Geschlecht an:
					</p>
				`, vE.NumberingStart+0),
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q8"
				rad.ValueRadio = "male"
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Männlich",
				}
				rad.StyleLbl = lblStyleRight
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q8"
				rad.ValueRadio = "female"
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Weiblich",
				}
				rad.StyleLbl = lblStyleRight
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q8"
				rad.ValueRadio = "diverse"
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Divers",
				}
				rad.StyleLbl = lblStyleRight
			}
		}

		//
		//
		{
			gr := page.AddGroup()
			gr.Cols = 12
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Bitte geben Sie Ihr Geburtsjahr an:
					</p>
				`, vE.NumberingStart+1),
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q9"
				inp.ColSpan = 4
				inp.ColSpanControl = 1

				inp.Min = 1900
				inp.Max = 2010
				inp.Step = 1
				inp.MaxChars = 5
			}
		}

		// separate header - since the states are vertically shown
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					In welchem Bundesland befindet sich Ihr Hauptwohnsitz?
					</p>
				`, vE.NumberingStart+2),
				}

			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 8
			gr.Vertical(8)
			for _, stt := range trl.FederalStatesGermanyISOs2 {
				lbl := stt.S
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q10"
				rad.ValueRadio = strings.ToLower(stt.Key)
				rad.ColSpan = 4
				rad.ColSpan = 1 // for vertical
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = lbl
				rad.StyleLbl = lblStyleRight
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 8
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Welcher ist Ihr höchster Bildungsabschluss?
					</p>
				`, vE.NumberingStart+3),
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q11"
				rad.ValueRadio = "kein_abschluss"
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Ohne Abschluss",
				}
				rad.StyleLbl = lblStyleRight
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q11"
				rad.ValueRadio = "hauptschule"
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Hauptschule",
				}
				rad.StyleLbl = lblStyleRight
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q11"
				rad.ValueRadio = "realschule"
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Realschule",
				}
				rad.StyleLbl = lblStyleRight
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q11"
				rad.ValueRadio = "abitur"
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Abitur",
				}
				rad.StyleLbl = lblStyleRight
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q11"
				rad.ValueRadio = "hochschule"
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Hochschulabschluss",
				}
				rad.StyleLbl = lblStyleRight
			}
		}

	}

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		{
			gr := page.AddGroup()
			gr.Cols = 12
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 12
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Wie ist Ihr Familienstand?
					</p>
				`, vE.NumberingStart+4),
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q12"
				rad.ValueRadio = "single"
				rad.ColSpan = 6
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Alleinstehend",
				}
				rad.StyleLbl = lblStyleRight
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q12"
				rad.ValueRadio = "engaged"
				rad.ColSpan = 6
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Partnerschaft ohne Ehe",
				}
				rad.StyleLbl = lblStyleRight
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q12"
				rad.ValueRadio = "married"
				rad.ColSpan = 6
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Verheiratet",
				}
				rad.StyleLbl = lblStyleRight
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 8
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Wie viel Geld verdienen Sie persönlich monatlich 
					nach Abzug von Steuern und Sozialversicherungsbeiträgen?
					</p>
				`, vE.NumberingStart+5),
				}
			}

			keyVals := []string{
				"0:0 Euro",
				"upto500:bis unter 500 Euro",
				"upto1000:500 bis unter 1000 Euro",
				"upto1500:1000 bis unter 1500 Euro",
				"upto2000:1500 bis unter 2000 Euro",
				"upto3000:2000 bis unter 3000 Euro",
				"upto4000:3000 bis unter 4000 Euro",
				"upto5000:4000 bis unter 5000 Euro",
				"over5000:Mehr als 5000 Euro",
			}

			for _, kv := range keyVals {
				sp := strings.Split(kv, ":")
				key := sp[0]
				val := sp[1]
				lbl := trl.S{"de": val}
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q13"
				rad.ValueRadio = key
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = lbl
				rad.StyleLbl = lblStyleRight
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 8
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Mit welcher Partei fühlen Sie sich 
					aufgrund Ihrer Werte und Überzeugungen am ehesten verbunden?
					</p>
				`, vE.NumberingStart+6),
				}
			}

			keyVals := []string{
				"cducsu:CDU/CSU",
				"linke:Die Linke",
				"spd:SPD",
				"gruene:Bündnis 90/Die Grünen",
				"fdp:FDP",
				"afd:AfD",
				"other:Andere",
			}

			for _, kv := range keyVals {
				sp := strings.Split(kv, ":")
				key := sp[0]
				val := sp[1]
				lbl := trl.S{"de": val}
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q14"
				rad.ValueRadio = key
				rad.ColSpan = 4
				rad.ColSpanLabel = 4
				rad.ColSpanControl = 1
				rad.Label = lbl
				rad.StyleLbl = lblStyleRight
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q14_other_text"
				inp.ColSpan = 4
				inp.ColSpanControl = 1
				inp.MaxChars = 14
			}

		}

		//
		{
			grStPage78 := css.NewStylesResponsive(nil)
			grStPage78.Desktop.StyleGridContainer.GapRow = "0.1rem"
			grStPage78.Desktop.StyleGridContainer.GapColumn = "0.01rem"

			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				labelsOneToSeven4,
				[]string{"q15"},
				radioVals7,
				[]trl.S{},
			)
			gb.MainLabel = trl.S{
				"de": fmt.Sprintf(`
					<p>
					<b>Frage&nbsp;%v.</b>
					Wie stufen Sie sich politisch ein?
					</p>
				`, vE.NumberingStart+7),
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = grStPage78
			gr.BottomVSpacers = 4
		}

	}

	return nil
}