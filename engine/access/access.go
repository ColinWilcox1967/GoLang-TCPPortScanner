package access

//import "../../accesscontrolnetwork/graph"

// Access Control Details
type AccessControlDetails struct {
	id int
	label string
	active bool
}

type AccessControls struct {
	id int
	controls []AccessControlDetails
}

func processAccessMap () error {
	return nil
}

func loadAccessMapFile() error {
	return ErrProblemLoadingAccessMappingFile
}

func (a *AccessControlDetails)ID () int {
	return a.id
}

func (a *AccessControlDetails)Label () string {
	return a.label
}

func (a *AccessControlDetails)Status () bool {
	return a.active
}

func (a *AccessControlDetails)SetLabel (l string) {
	a.label = l
}

func (a* AccessControlDetails)SetID (n int) {
	a.id = n
}

func (a *AccessControlDetails)SetStatus (status bool) {
	a.active = status
}

func (ac *AccessControls)ID () int {
	return ac.id
}

func (ac *AccessControls)SetID (n int) {
	ac.id = n
}

func (ac *AccessControls)GetControl (n int) (error, AccessControlDetails) {
	if n <0 || n >= len(ac.controls) {
		return ErrControlIndexOutOfRange, AccessControlDetails{}
	}

	return nil, ac.controls[n]
}
func (ac *AccessControls)SetControl (n int, details AccessControlDetails)  {
	if n < 0 || n >= len(ac.controls) {
		var newControl AccessControlDetails

		newControl.SetID (details.id)
		newControl.SetLabel (details.Label ())
		newControl.SetStatus (details.Status ())

		ac.controls = append(ac.controls, newControl)
	}

	// change details in existsing control
	ac.controls[n].id = details.ID ()
	ac.controls[n].label = details.Label ()
	ac.controls[n].active = details.Status ()
}

// Setup engine access layer
func Setup(filePath string) error {
	return loadAccessMapFile()
}
