package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage11AC1(q *QuestionnaireT, page *pageT) error {
	return pdsPage11(q, page, 0)
}
func pdsPage11AC2(q *QuestionnaireT, page *pageT) error {
	return pdsPage11(q, page, 1)
}
func pdsPage11AC3(q *QuestionnaireT, page *pageT) error {
	return pdsPage11(q, page, 2)
}

func pdsPage11(q *QuestionnaireT, page *pageT, acIdx int) error {

	ac := PDSAssetClasses[acIdx]
	ac = onlySelectedTranchTypes(q, ac)
	rn := rune(65 + acIdx) // ascii 65 is A; 97 is a

	page.ValidationFuncName = "pdsRange"

	page.Label = trl.S{
		"en": fmt.Sprintf("%v: Portfolio changes  (past 3 months)", ac.Lbl["en"]),
		"de": fmt.Sprintf("%v: Portfolio changes  (past 3 months)", ac.Lbl["de"]),
	}.Outline(fmt.Sprintf("%c1.", rn))
	page.Short = trl.S{
		"en": "Portfolio<br>changes",
		"de": "Portfolio<br>changes",
	}
	page.Short = trl.S{
		"en": fmt.Sprintf("%v<br>changes", ac.Short["en"]),
		"de": fmt.Sprintf("%v<br>changes", ac.Short["de"]),
	}

	page.CounterProgress = fmt.Sprintf("%c1", rn)

	page.WidthMax("58rem")
	if len(ac.TrancheTypes) == 2 {
		page.WidthMax("42rem")
	}
	if len(ac.TrancheTypes) == 1 {
		page.WidthMax("34rem")
	}

	// dynamically recreate the groups
	page.Groups = nil

	restrictedTextMultiCols(page, ac, rT1)

	lblDuration := trl.S{
		"en": "Average time to close a deal in weeks",
		"de": "Durchschnittl. Zeit bis Abschluss in Wochen",
	}.Outline("b.)")

	dropdownsLabelsTop(
		page,
		ac,
		"q11b_closing_time",
		lblDuration,
		mCh5,
	)

	restrictedTextMultiCols(page, ac, rT2)

	restrictedTextMultiCols(page, ac, rT3)

	restrictedTextMultiCols(page, ac, rT4)

	page11fghInputs := []string{
		"q11f_esg",
		"q11g_ratch",
		"q11h_degrees",
	}

	page11fghTypes := []string{
		"range-pct",
		"range-pct",
		"range-pct",
	}

	page11fghLbls := []trl.S{
		{
			"en": `<bb>Share ESG KPIs</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) with explicit ESG targets in the credit documentation? </span>`,
			"de": `<bb>Share ESG KPIs</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) with explicit ESG targets in the credit documentation? </span>`,
		},
		{
			"en": `<bb>Share ESG ratchets</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) with ESG ratchets? </span>`,
			"de": `<bb>Share ESG ratchets</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) with ESG ratchets? </span>`,
		},
		{
			"en": `<bb>Share 1.5°C target</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?</span>`,
			"de": `<bb>Share 1.5°C target</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?</span>`,
		},
	}

	{

		// 4cols layout

		for i := 0; i < len(page11fghLbls); i++ {
			rn := rune(102 + i) // 102 is f
			page11fghLbls[i] = page11fghLbls[i].Outline(fmt.Sprintf("%c.)", rn))
		}

		createRows(
			page,
			ac,
			page11fghInputs,
			page11fghTypes,
			page11fghLbls,
			[]*rangeConf{
				&sliderPctZeroHundredWide, // &sliderPctZeroHundredMiddle,
				&sliderPctZeroHundredWide,
				&sliderPctZeroHundredWide,
			},
		)

	}

	return nil
}