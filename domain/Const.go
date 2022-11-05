package domain

type (
	Status string
	Gender string
	Role   string
)

const (
	ACTIVE        Status = "ACTIVE"
	DISABLED      Status = "DISABLED"
	PAUSE         Status = "PAUSE"
	TRIAL         Status = "TRIAL"
	MALE          Gender = "MALE"
	FEMALE        Gender = "FEMALE"
	GENERAL_ADMIN Role   = "GENERAL_ADMIN"
	ADMIN         Role   = "ADMIN"
	USER          Role   = "USER"
)
