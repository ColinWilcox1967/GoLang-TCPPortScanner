package physical

import "errors"

var (
	//ErrProblemSettingUpPhysicalLayer layer initialisation problem
	ErrProblemSettingUpPhysicalLayer = errors.New("Physical Layer : Problem Setting Up")

	//ErrProblemLoadingPhysicalMapFile problem reading contgents of map file
	ErrProblemLoadingPhysicalMapFile = errors.New("Physical Layer : Problem Loading Map File")
)
