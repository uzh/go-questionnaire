package quest

import "github.com/zew/go-questionnaire/pkg/trl"

var labelsVerfuegbarNicht = []trl.S{
	{
		"de": "Verfügbar",
	},
	{
		"de": "Nicht verfügbar",
	},
	{
		"de": "Von dieser Alternative abraten<sup>1</sup>",
	},
}

var q1Names = []string{
	"q1_1_15_6",
	"q1_3_07_6",
	"q1_4_06_6",
}

var q2Names = []string{
	"q2_1_15_7",
	"q2_3_07_7",
	"q2_4_06_7",
}

/*
var q1Labels = []trl.S{
	{"de": "<b>1</b> Euro sofort und       <b>15</b>&nbsp;Euro  in <b>6</b>&nbsp;Monaten"},
	{"de": "<b>3</b> Euro sofort und &nbsp;<b>7</b>&nbsp;Euro   in <b>6</b>&nbsp;Monaten"},
	{"de": "<b>4</b> Euro sofort und &nbsp;<b>1</b>&nbsp;Euro   in <b>6</b>&nbsp;Monaten"},
}

var q2Labels = []trl.S{
	{"de": "<b>1</b> Euro sofort und       <b>15</b>&nbsp;Euro  in <b>7</b>&nbsp;Monaten"},
	{"de": "<b>3</b> Euro sofort und &nbsp;<b>7</b>&nbsp;Euro   in <b>7</b>&nbsp;Monaten"},
	{"de": "<b>4</b> Euro sofort und &nbsp;<b>1</b>&nbsp;Euro   in <b>7</b>&nbsp;Monaten"},
}

*/

var q4Labels = []trl.S{
	{"de": "Es sollte keine solche Verpflichtung geben"},
	{"de": "Jede(r) Erwerbstätige sollte verpflichtet sein, mindestens &nbsp;<b>2.5%</b> zu sparen"},
	{"de": "Jede(r) Erwerbstätige sollte verpflichtet sein, mindestens &nbsp;<b>5%</b> zu sparen"},
	{"de": "Jede(r) Erwerbstätige sollte verpflichtet sein, mindestens <b>10%</b> zu sparen"},
	{"de": "Jede(r) Erwerbstätige sollte verpflichtet sein, mindestens <b>20%</b> zu sparen"},
	{"de": "Jede(r) Erwerbstätige sollte verpflichtet sein, mindestens <b>30%</b> zu sparen"},
}

var labelsOneToTen1 = []trl.S{
	{"de": "0<br>gar nicht bereit"},
	{"de": "1"},
	{"de": "2"},
	{"de": "3"},
	{"de": "4"},
	{"de": "5"},
	{"de": "6"},
	{"de": "7"},
	{"de": "8"},
	{"de": "9"},
	{"de": "10<br>Absolut bereit"},
}
var labelsOneToTen2 = []trl.S{
	{"de": "0<br>Risikovermeiden"},
	{"de": "1"},
	{"de": "2"},
	{"de": "3"},
	{"de": "4"},
	{"de": "5"},
	{"de": "6"},
	{"de": "7"},
	{"de": "8"},
	{"de": "9"},
	{"de": "10<br>Risikobereit"},
}
var labelsOneToTen3 = []trl.S{
	{"de": "0<br>Teile nicht gern"},
	{"de": "1"},
	{"de": "2"},
	{"de": "3"},
	{"de": "4"},
	{"de": "5"},
	{"de": "6"},
	{"de": "7"},
	{"de": "8"},
	{"de": "9"},
	{"de": "10<br>Teile gern"},
}

/*
	2021-02 - doch keine Ziffern
*/
var labelsBereitGarNicht = []trl.S{
	{"de": "Gar nicht bereit <div class='vspacer-24'> &nbsp; </div>"},
	{"de": ""},
	{"de": ""},
	{"de": ""},
	{"de": ""},
	{"de": ""},
	{"de": "Sehr bereit <div class='vspacer-24'> &nbsp; </div>"},
}
var labelsRiskobereit = []trl.S{
	{"de": "Gar nicht risikobereit <div class='vspacer-24'> &nbsp; </div>"},
	{"de": ""},
	{"de": ""},
	{"de": ""},
	{"de": ""},
	{"de": ""},
	{"de": "Sehr risikobereit <div class='vspacer-24'> &nbsp; </div>"},
}

var labelsZustimmung = []trl.S{
	{"de": "&nbsp;<br>Stimme gar nicht zu <div class='vspacer-24'> &nbsp; </div>"},
	{"de": ""},
	{"de": ""},
	{"de": ""},
	{"de": ""},
	{"de": ""},
	{"de": "&nbsp;<br>Stimme komplett zu <div class='vspacer-24'> &nbsp; </div>"},
}

var labelsLeftRight = []trl.S{
	{"de": "Li&shy;nk&shy;s <div class='vspacer-24'> &nbsp; </div> "}, // 1
	{"de": ""},
	{"de": ""}, // 3
	{"de": ""},
	{"de": ""}, // 5
	{"de": ""},
	{"de": ""}, // 7
	{"de": ""},
	{"de": ""}, // 9
	{"de": ""},
	{"de": "Re&shy;cht&shy;s <div class='vspacer-24'> &nbsp; </div> "}, // 11
	// {"de": "We&shy;iß ni&shy;cht"},
	// {"de": "Weiß<br>nicht"},
}
