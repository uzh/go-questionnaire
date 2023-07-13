package kneb1

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func yesNo(page qst.WrappedPageT, lbl trl.S, inputName, outlineNumber string, indent bool) {

	twoAnswers(
		page,
		lbl,
		inputName, outlineNumber, indent,
		//
		[]string{
			"yes",
			"no",
		},
		[]trl.S{
			{
				"de": "ja",
				"en": "yes",
			},
			{
				"de": "nein",
				"en": "no",
			},
		},
	)

}

func meOrTogether(page qst.WrappedPageT, lbl trl.S, inputName, outlineNumber string, indent bool) {

	twoAnswers(
		page,
		lbl,
		inputName, outlineNumber, indent,
		//
		[]string{
			"1",
			"2",
		},
		[]trl.S{
			{
				"de": "eher ich",
				"en": "rather me",
			},
			{
				"de": "ich gemeinsam mit Partner(in) / anderem Haushaltsmitglied",
				"en": "me together with partner or other household member(s)",
			},
		},
	)

}

func twoAnswers(
	page qst.WrappedPageT,
	lbl trl.S,
	inputName, outlineNumber string,
	indent bool,
	//
	radioValues []string,
	labels []trl.S,
) {

	gr := page.AddGroup()
	gr.Cols = 8 + 1
	gr.WidthMax("40rem")
	gr.BottomVSpacers = 2

	if indent {
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Margin = "0 0 0 1.2rem"
	}
	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = lbl.Outline(outlineNumber)
		inp.ColSpan = gr.Cols
		inp.ColSpanLabel = 1
		inp.ColSpanControl = 0
	}

	// vertical spacer
	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = trl.S{"de": "", "en": ""}
		inp.ColSpan = 1
		inp.ColSpanLabel = 1
		inp.ColSpanControl = 0
	}
	for idx, label := range labels {
		rad := gr.AddInput()
		rad.Type = "radio"
		rad.Name = inputName
		rad.ValueRadio = radioValues[idx]

		rad.ColSpan = gr.Cols
		rad.ColSpan = 2
		if idx > 0 {
			rad.ColSpan = gr.Cols - 2 - 1
		}
		rad.ColSpanLabel = 1
		rad.ColSpanControl = 6

		rad.Label = label
		rad.ControlFirst()
	}
}

var radioVals4 = []string{
	"1", "2",
	"3", "4",
}
var radioVals5 = []string{
	"1", "2",
	"3",
	"4", "5",
}
var radioVals6 = []string{
	"1", "2",
	"3",
	"4", "5",
	"6",
}

var radioVals7 = []string{
	"1", "2", "3",
	"4",
	"5", "6", "7",
}

var radioVals11 = []string{
	"1", "2", "3", "4", "5",
	"6",
	"7", "8", "9", "10", "11",
}

var radioVals12 = []string{
	"1", "2", "3", "4", "5",
	"6",
	"7", "8", "9", "10", "11",
	"12",
}

// no first col labels
var columnTemplate5a = []float32{
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0.4, 1,
}

var columnTemplate6 = []float32{
	2, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0.4, 1,
}

// no first col labels
var columnTemplate6a = []float32{
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0.4, 1,
}

// no first col labels
var columnTemplate4 = []float32{

	0, 1,

	0, 1,

	0, 1,

	0, 1,

	// 0.4, 1,
}

// no first col labels
var columnTemplate5 = []float32{

	0, 1,
	0, 1,
	0, 1, // middle
	0, 1,
	0, 1,

	// 0.4, 1,
}

// no first col labels
var colsStockMarket = []float32{
	0, 2,
	0, 2,
	0, 2,
	0, 1.5,
	0, 1,
	0, 1,
	// 0.4, 1,
}

// no first col labels
var colsAssetClasses = []float32{
	0, 2,
	0, 2,
	0, 2,
	0, 2,
	0, 1,
	0, 1,
	// 0.4, 1,
}

// no first col labels
var columnTemplate7 = []float32{

	0, 1,

	0, 1,
	0, 1,

	0, 1, // middle

	0, 1,
	0, 1,

	0, 1,

	// 0.4, 1,
}

// no first col labels
var columnTemplate11 = []float32{

	0, 1,

	0, 1,
	0, 1,
	0, 1,
	0, 1,

	0, 1, // middle

	0, 1,
	0, 1,
	0, 1,
	0, 1,

	0, 1,

	// 0.4, 1,
}

// no first col labels
var columnTemplate12 = []float32{

	0, 1,

	0, 1,
	0, 1,
	0, 1,
	0, 1,

	0, 1, // middle

	0, 1,
	0, 1,
	0, 1,
	0, 1,

	0, 1,

	0, 2.5, // dont know - guessed

	// 0.4, 1,
}

func labelsGoodBad6() []trl.S {

	tm := []trl.S{
		{
			"de": "Sehr gut",
			"en": "Very good",
		},
		{
			"de": "Gut",
			"en": "Good",
		},
		{
			"de": "Durchschnittlich",
			"en": "Fair",
		},
		{
			"de": "Schlecht",
			"en": "Bad",
		},
		{
			"de": "Sehr schlecht",
			"en": "Very bad",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func labelsYesNo5() []trl.S {

	tm := []trl.S{
		{
			"de": "Ja",
			"en": "Yes",
		},
		{
			"de": "Nein",
			"en": "No",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Nicht anwendbar",
			"en": "Not applicable",
		},
		{
			"de": "Keine Antwort",
			"en": "No answer",
		},
	}

	return tm

}

func labelsHighLow6() []trl.S {

	tm := []trl.S{
		{
			"de": "Sehr hoch",
			"en": "Very high",
		},
		{
			"de": "Ziemlich hoch",
			"en": "Quite high",
		},
		{
			"de": "Durchschnittlich",
			"en": "About average",
		},
		{
			"de": "Ziemlich gering",
			"en": "Quite low",
		},
		{
			"de": "Sehr gering",
			"en": "Very low",
		},
		{
			"de": "Keine Antwort",
			"en": "No answer",
		},
	}

	return tm

}

func assetVariance5() []trl.S {

	tm := []trl.S{
		{
			"de": "Sparbuch",
			"en": "Savings accounts",
		},
		{
			"de": "Aktien",
			"en": "Stocks",
		},
		{
			"de": "Anleihen",
			"en": "Bonds",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func trueFalse4() []trl.S {

	tm := []trl.S{
		{
			"de": "Wahr",
			"en": "True",
		},
		{
			"de": "Falsch",
			"en": "False",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine Antwort",
			"en": "No answer",
		},
	}

	return tm

}

func diversification5() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "Increase",
		},
		{
			"de": "xxx",
			"en": "Decrease",
		},
		{
			"de": "xxx",
			"en": "Stay the same",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func stockMarketFuncs6() []trl.S {

	tm := []trl.S{
		{
			"de": "Der Aktienmarkt hilft Dividendenerträge vorauszusagen.",
			"en": "The stock market helps to predict stock earnings.",
		},
		{
			"de": "Der Aktienmarkt führt zu einer Preiserhöhung für Aktien.",
			"en": "The stock market results in an increase in the price of stocks. ",
		},
		{
			"de": "Der Aktienmarkt bringt Aktienkäufer und -verkäufer zusammen.",
			"en": "The stock market brings people who want to buy stocks together with those who want to sell stocks. ",
		},
		{
			"de": "Keine der Antworten",
			"en": "None of the above",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func mutualFunds6() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "Once one invests in a mutual fund, one cannot withdraw the money in the first year.",
		},
		{
			"de": "xxx",
			"en": "Mutual funds can invest in several assets, for example invest in both stocks and bonds.",
		},
		{
			"de": "xxx",
			"en": "Mutual funds pay a guaranteed rate of return which depends on their past performance.",
		},
		{
			"de": "Keine der Antworten",
			"en": "None of the above",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func highestReturn5() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "Savings accounts",
		},
		{
			"de": "xxx",
			"en": "Stocks",
		},
		{
			"de": "xxx",
			"en": "Bonds",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func compounding5() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "More than 102 Euros",
		},
		{
			"de": "xxx",
			"en": "Exactly 102 Euros",
		},
		{
			"de": "xxx",
			"en": "Less than 102",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func realRate5() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "More than today",
		},
		{
			"de": "xxx",
			"en": "Exactly the same",
		},
		{
			"de": "xxx",
			"en": "Less than today",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func compoundingMulti5() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "More than 200 Euros",
		},
		{
			"de": "xxx",
			"en": "Exactly 200 Euros",
		},
		{
			"de": "xxx",
			"en": "Less than 200",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func inflationIndexIncome5() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "More than today",
		},
		{
			"de": "xxx",
			"en": "The same",
		},
		{
			"de": "xxx",
			"en": "Less than today",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func interestBondPrice5() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "They rise",
		},
		{
			"de": "xxx",
			"en": "They fall",
		},
		{
			"de": "xxx",
			"en": "They stay the same",
		},
		{
			"de": "Weiß nicht",
			"en": "Do not know",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

func riskPreference7() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "1. Not at all willing to take risks",
		},
		{
			"de": " &nbsp; ",
			"en": "2. &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": "3.  &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": "4.  &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": "5.  &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": "6.  &nbsp; ",
		},
		{
			"de": "xxx",
			"en": "7. Very willing to take risks",
		},
	}

	return tm

}

func agreeToDisagree7() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "Disagree completely",
		},
		{
			"de": " &nbsp; ",
			"en": "Disagree strongly",
		},
		{
			"de": " &nbsp; ",
			"en": "Disagree a little",
		},
		{
			"de": " &nbsp; ",
			"en": "Neither agree nor disagree",
		},
		{
			"de": " &nbsp; ",
			"en": "Agree a little",
		},
		{
			"de": " &nbsp; ",
			"en": "Agree strongly",
		},
		{
			"de": "xxx",
			"en": "Agree completely",
		},
	}

	return tm

}

func agreeToDisagree5() []trl.S {

	tm := []trl.S{
		{
			"de": " &nbsp; ",
			"en": "Disagree strongly",
		},
		{
			"de": " &nbsp; ",
			"en": "Disagree a little",
		},
		{
			"de": " &nbsp; ",
			"en": "Neither agree nor disagree",
		},
		{
			"de": " &nbsp; ",
			"en": "Agree a little",
		},
		{
			"de": " &nbsp; ",
			"en": "Agree strongly",
		},
	}

	return tm

}

func optimistic7() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "Not at all optimistic",
		},
		{
			"de": " &nbsp; ",
			"en": "Strongly non-optimistic",
		},
		{
			"de": " &nbsp; ",
			"en": "A little non-optimistic",
		},
		{
			"de": " &nbsp; ",
			"en": "Indifferent",
		},
		{
			"de": " &nbsp; ",
			"en": "A little optimistic",
		},
		{
			"de": " &nbsp; ",
			"en": "Strongly optimistic",
		},
		{
			"de": "xxx",
			"en": "Very optimistic",
		},
	}

	return tm

}

func pessimistic7() []trl.S {

	tm := []trl.S{
		{
			"de": "xxx",
			"en": "Not at all pessimistic",
		},
		{
			"de": " &nbsp; ",
			"en": "Strongly non-pessimistic",
		},
		{
			"de": " &nbsp; ",
			"en": "A little non-pessimistic",
		},
		{
			"de": " &nbsp; ",
			"en": "Indifferent",
		},
		{
			"de": " &nbsp; ",
			"en": "A little pessimistic",
		},
		{
			"de": " &nbsp; ",
			"en": "Strongly pessimistic",
		},
		{
			"de": "xxx",
			"en": "Very pessimistic",
		},
	}

	return tm

}

func labelsRisk() []trl.S {

	return []trl.S{

		{
			"de": "<small>gar nicht risikobereit</small> 0",
			"en": "<small>no risk at all</small>         0",
		},

		// 1-4
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		// center
		{
			"de": " 5 ",
			"en": " 5 ",
		},

		// 6-9
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		//
		{
			"de": "<small>sehr risikobereit</small>   10",
			"en": "<small>very fond of risk</small>   10",
		},
	}

}

func labelsPositiveAspects() []trl.S {

	return []trl.S{

		{
			"de": "<small>nur an die positiven Seiten</small> 1",
			"en": "<small>positiv aspects only</small> 1",
		},

		// 2-3
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		// center
		{
			"de": " 4 ",
			"en": " 4 ",
		},

		// 5-6
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		{
			"de": "<small>nur an die negativen Seiten</small> 7",
			"en": "<small>negative aspects only</small>       7",
		},
	}

}

func labelsImportantSituations() []trl.S {

	return []trl.S{

		{
			"de": "<small>kleine Alltagssituationen</small> 1",
			"en": "<small>small everyday life</small> 1",
		},

		// 2-3
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		// center
		{
			"de": " 4 ",
			"en": " 4 ",
		},

		// 5-6
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		{
			"de": "<small>große, wichtige Situationen</small> 7",
			"en": "<small>big important occasions</small>     7",
		},
	}

}

func labelsReturns() []trl.S {

	return []trl.S{

		{
			"de": "<small>kleine Gewinne</small> 1",
			"en": "<small>small returns</small> 1",
		},

		// 2-3
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		// center
		{
			"de": " 4 ",
			"en": " 4 ",
		},

		// 5-6
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		{
			"de": "<small>große Gewinne</small> 7",
			"en": "<small>large returns</small> 7",
		},
	}
}

func labelsLosses() []trl.S {

	return []trl.S{

		{
			"de": "<small>kleine Verluste</small> <div>1</div>",
			"en": "<small>small losses</small>    <div>1</div>",
		},

		// 2-3
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		// center
		{
			"de": " 4 ",
			"en": " 4 ",
		},

		// 5-6
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		{
			"de": "<small>große Verluste</small> <div>7</div>",
			"en": "<small>large losses</small>   <div>7</div>",
		},
	}

}

func labelsFinRisk() []trl.S {

	return []trl.S{

		{
			"de": "<small>nicht bereit, ein Risiko einzugehen</small> <div>1</div>",
			"en": "<small>not prepared to take any risk</small>       <div>1</div>",
		},

		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		// center
		{
			"de": " 3 ",
			"en": " 3 ",
		},

		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		{
			"de": "<small>bereit, ein erhebliches Risiko einzugehen, um potenziell eine höhere Rendite zu erzielen</small> <div>5</div>",
			"en": "<small>prepared to take a considerable risk, for higher potential returns</small>                       <div>5</div>",
		},
	}

}

func labelsSelfKnowledge() []trl.S {

	return []trl.S{

		{
			"de": "<small>sehr gering</small> 0",
			"en": "<small>very small </small> 0",
		},

		// 1-4
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		// center
		{
			"de": " 5 ",
			"en": " 5 ",
		},

		// 6-9
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},
		{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		},

		//
		{
			"de": "<small>sehr hoch</small>   10",
			"en": "<small>very high</small>   10",
		},
	}

}

func labelsInterestOverTwoYears() []trl.S {

	return []trl.S{

		{
			"de": "<small>höher als 102&nbsp;€</small>      ",
			"en": "<small>higher than 102&nbsp;€</small>    ",
		},
		{
			"de": "<small>genau   102&nbsp;€</small>        ",
			"en": "<small>exactly 102&nbsp;€</small>        ",
		},
		{
			"de": "<small>niedriger als 102&nbsp;€</small>  ",
			"en": "<small>lower than 102&nbsp;€</small>     ",
		},
		// -1
		{
			"de": "<small>weiß nicht</small>                ",
			"en": "<small>dont know </small>                ",
		},
		// -2
		{
			"de": "<small>keine Angabe</small>              ",
			"en": "<small>no answer   </small>              ",
		},
	}

}

func labelsStimmeZuOderNicht() []trl.S {

	return []trl.S{
		{
			"de": "<small>stimme zu</small>        ",
			"en": "<small>agree</small>            ",
		},
		{
			"de": "<small>stimme nicht zu</small>  ",
			"en": "<small>dont agree</small>       ",
		},
		{
			"de": "<small>weiß nicht</small>        ",
			"en": "<small>dont know </small>        ",
		},
		{
			"de": "<small>keine Angabe</small>      ",
			"en": "<small>no answer   </small>      ",
		},
	}

}

func labelsStockMarketPurpose() []trl.S {

	return []trl.S{

		{
			"de": "Der Aktienmarkt hilft Aktiengewinne vorherzusagen.",
			"en": "todo",
		},
		{
			"de": "Der Aktienmarkt führt zu einer Erhöhung der Aktienpreise.",
			"en": "todo",
		},
		{
			"de": "Der Aktienmarkt führt Käufer und Verkäufer von Aktien zusammen.",
			"en": "todo",
		},
		{
			"de": "nichts davon",
			"en": "todo",
		},
		{
			"de": "weiß nicht",
			"en": "todo",
		},
		{
			"de": "keine Angabe",
			"en": "todo",
		},
	}

}

func labelsAssetClassVola() []trl.S {

	return []trl.S{

		{
			"de": "Sparbücher",
			"en": "todo",
		},
		{
			"de": "Termineinlagen",
			"en": "todo",
		},
		{
			"de": "Festverzinsliche Wertpapiere",
			"en": "todo",
		},
		{
			"de": "Aktien",
			"en": "todo",
		},
		{
			"de": "weiß nicht",
			"en": "todo",
		},
		{
			"de": "keine Angabe",
			"en": "todo",
		},
	}

}

func labelsExpectedValueRisk() []trl.S {

	return []trl.S{
		{
			"de": "die Autoreparatur",
			"en": "todo",
		},
		{
			"de": "der Austausch der Klimaanlage",
			"en": "todo",
		},
		{
			"de": "man kann es nicht im voraus wissen",
			"en": "todo",
		},
		{
			"de": "ich weiß nicht",
			"en": "todo",
		},
	}

}

func labelsPensionAdvice() []trl.S {

	return []trl.S{

		{
			"de": "dieses Jahr",
			"en": "todo",
		},
		{
			"de": "im letzten Jahr",
			"en": "todo",
		},
		{
			"de": "in den letzten 5 Jahren",
			"en": "todo",
		},
		{
			"de": "vor mehr als 5 Jahren",
			"en": "todo",
		},
		{
			"de": "noch nie",
			"en": "todo",
		},
	}

}

func labelsFrequency() []trl.S {

	return []trl.S{

		{
			"de": "oft",
			"en": "todo",
		},
		{
			"de": "manchmal",
			"en": "todo",
		},
		{
			"de": "selten",
			"en": "todo",
		},
		{
			"de": "niemals",
			"en": "todo",
		},
	}

}

func labelsWhoIsCompetent() []trl.S {

	return []trl.S{

		{
			"de": "Mein Partner/ Meine Partnerin",
			"en": "todo",
		},
		{
			"de": "Familie",
			"en": "todo",
		},
		{
			"de": "Freunde und Bekannte",
			"en": "todo",
		},
		{
			"de": "Bankberater/in",
			"en": "todo",
		},
	}

}
