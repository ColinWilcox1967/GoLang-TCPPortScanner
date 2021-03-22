package access

import "errors"

var (
	//ErrProblemSettingUpAccessLayer layer initialisation problem
	ErrProblemSettingUpAccessLayer = errors.New("Access Layer : Problem Setting Up")

	//ErrProblemLoadingAccessMapFile Problem loading access mapping file
	ErrProblemLoadingAccessMappingFile = errors.New ("Access Layer : Problem loading access mapping file")

	//ErrControlIndexOutOfRange
	ErrControlIndexOutOfRange = errors.New("Acess Layer : Specified control index out of range")
)
