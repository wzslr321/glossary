package cli

type actions struct {
	Add       string
	Build string
	Configure string
}


const (
	aAdd       *actions = "kurwa"
	aBuild     string = "build"
	aConfigure string = "configure"
)

type actionsRegistry struct {
	Add       *actions
	Build     *actions
	Configure *actions
}

func RegisterActions() *actionsRegistry {
	add := &actions{add: aAdd}
	build := &actions{build: aBuild}
	cfg := &actions{configure: aConfigure}

	AR := &actionsRegistry{Add: add, Build: build, Configure: cfg}

	return AR
}
