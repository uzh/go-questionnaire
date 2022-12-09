package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// const firstColLbl = float32(4)
// const firstColLbl = float32(2)
const firstColLbl = float32(3)

var suffixWeeks = trl.S{
	"en": "weeks",
	"de": "Wochen",
}

var suffixYears = trl.S{
	"en": "years",
	"de": "Jahre",
}

var suffixEBITDA = trl.S{
	"en": "x EBITDA",
	"de": "x EBITDA",
}
var suffixPercent = trl.S{
	"en": "%",
	"de": "%",
}

var suffixMillionEuro = trl.S{
	// capitalizemytitle.com/how-to-abbreviate-million/
	// "en": "million €",
	// "en": "MM €",
	"en": "mn €",
	"de": "Mio €",
}

var placeHolderNum = trl.S{
	"en": "#",
	"de": "#",
}

var placeHolderMillion = trl.S{
	"en": "million Euro",
	"de": "Millionen Euro",
}

func rangesRowLabelsLeft(
	page *qst.WrappedPageT,
	inputName string,
	lbl trl.S,
	// sfx trl.S,
	cf rangeConf,
	// rangeType string,
) {

	numCols := firstColLbl + float32(len(trancheTypeNamesAC1))
	idxLastCol := len(trancheTypeNamesAC1) - 1

	{
		gr := page.AddGroup()
		gr.Cols = numCols

		// row0 - headers
		for idx1 := 0; idx1 < len(trancheTypeNamesAC1)+1; idx1++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			if idx1 == 0 {
				inp.ColSpan = firstColLbl
			}
			if idx1 > 0 {
				ttLbl := allLbls["ac1-tranche-types"][idx1-1]
				// inp.Label = ttLbl.Bold()
				inp.Label = ttLbl
			}
			inp.LabelVertical()

			inp.StyleLbl = trancheNameStyle
		}

		// row1
		for idx1, trancheType := range trancheTypeNamesAC1 {

			ttPref := trancheType[:3]

			{
				inp := gr.AddInput()
				inp.Type = "range"
				inp.DynamicFuncParamset = cf.RangeType

				inp.Name = fmt.Sprintf("%v_%v", ttPref, inputName)

				// 0%-100% in 5% brackets
				inp.Min = cf.Min
				inp.Max = cf.Max
				inp.Step = cf.Step

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.WidthMax = "90%"

				if idx1 == 0 {
					inp.Label = lbl
					inp.LabelPadRight()
					inp.ColSpan = firstColLbl + 1
					inp.ColSpanLabel = firstColLbl
					inp.ColSpanControl = 1
				}

				if idx1 == idxLastCol {
					inp.Suffix = cf.Suffix
				}
				inp.Suffix = cf.Suffix

			}
		}

	}
}
