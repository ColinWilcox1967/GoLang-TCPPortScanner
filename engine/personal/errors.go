package personal

import "errors"

var (
	//ErrProblemSettingUpPersonalLayer layer initialisation problem
	ErrProblemSettingUpPersonalLayer = errors.New("Personal Layer : Problem Setting Up")

	//ErrProblemLoadingPersonalMappingFile Problem loading the personal mapping file
	ErrProblemLoadingPersonalMappingFile = errors.New("Personal Layer : Problem loading personal mapping file")
)
