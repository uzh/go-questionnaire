package qst

// compposite inputs combine challenging HTML and
// multiple inputs in complicated ways
//
//
//
// parameters
//   dynamic questionnaire - filled with response values
//   dynamic sequence - from random order
//   param set index -  statically determined - from a slice of param sets
//
// returns
//   rendered HTML of the group
//   slice of input names
//   error
type compositFuncT func(*QuestionnaireT, int, int) (string, []string, error)

var compositeFuncs = map[string]compositFuncT{
	"PoliticalFoundations": PoliticalFoundations,
}
