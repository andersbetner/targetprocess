package targetprocess

type queryParams struct {
	Where       string
	Include     string
	Exclude     string
	Append      string
	Skip        int32
	Take        int32
	InnerTake   int32
	OrderBy     string
	OrderByDesc string
	format      string
}

type jsonBodyParams struct {
	ID               int32  `json:"Id,omitempty"`
	CreateDate       string `json:"CreateDate,omitempty"`
	Name             string `json:"Name,omitempty"`
	Description      string `json:"Description,omitempty"`
	Tags             string `json:"Tags,omitempty"`
	PlannedStartDate string `json:"PlannedStartDate,omitempty"`
	PlannedEndDate   string `json:"PlannedEndDate,omitempty"`
	Project          struct {
		ID int `json:"Id,omitempty"`
	} `json:"Project,omitempty"`
}

type UserStoriesJSONBodyParams = jsonBodyParams
type UserStoriesQueryParams = queryParams
type BugsQueryParams = queryParams
type BugsJSONBodyParams = jsonBodyParams
type EpicsQueryParams = queryParams
type EpicsJSONBodyParams = jsonBodyParams
type FeaturesQueryParams = queryParams
type FeaturesJSONBodyParams = jsonBodyParams
type TasksQueryParams = queryParams
type TasksJSONBodyParams = jsonBodyParams
type RequestersQueryParams = queryParams
type RequestersJSONBodyParams = jsonBodyParams
type UsersQueryParams = queryParams
type UsersJSONBodyParams = jsonBodyParams
type AssignablesQueryParams = queryParams
type AssignabesJSONBodyParams = jsonBodyParams
type GeneralsQueryParams = queryParams
type GeneralsJSONBodyParams = jsonBodyParams
type ProjectsQueryParams = queryParams
type ProjectsJSONBodyParams = jsonBodyParams

type UserStoriesGetResponse struct {
	Next  string `json:"Next"`
	Items []struct {
		ResourceType string  `json:"ResourceType"`
		ID           int     `json:"Id"`
		Name         string  `json:"Name"`
		Effort       float64 `json:"Effort"`
		BugsCount    int     `json:"Bugs-Count"`
		Project      struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Project"`
	} `json:"Items"`
}

type UserStoriesPostResponse struct {
	ResourceType        string      `json:"ResourceType"`
	ID                  int         `json:"Id"`
	Name                string      `json:"Name"`
	Description         interface{} `json:"Description"`
	StartDate           interface{} `json:"StartDate"`
	EndDate             interface{} `json:"EndDate"`
	CreateDate          string      `json:"CreateDate"`
	ModifyDate          string      `json:"ModifyDate"`
	LastCommentDate     interface{} `json:"LastCommentDate"`
	Tags                string      `json:"Tags"`
	NumericPriority     float64     `json:"NumericPriority"`
	Effort              float64     `json:"Effort"`
	EffortCompleted     float64     `json:"EffortCompleted"`
	EffortToDo          float64     `json:"EffortToDo"`
	Progress            float64     `json:"Progress"`
	TimeSpent           float64     `json:"TimeSpent"`
	TimeRemain          float64     `json:"TimeRemain"`
	LastStateChangeDate string      `json:"LastStateChangeDate"`
	PlannedStartDate    interface{} `json:"PlannedStartDate"`
	PlannedEndDate      interface{} `json:"PlannedEndDate"`
	InitialEstimate     float64     `json:"InitialEstimate"`
	EntityType          struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"EntityType"`
	Project struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Process      struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Process"`
	} `json:"Project"`
	Owner struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		FirstName    string `json:"FirstName"`
		LastName     string `json:"LastName"`
		Login        string `json:"Login"`
	} `json:"Owner"`
	LastCommentedUser interface{} `json:"LastCommentedUser"`
	LinkedTestPlan    interface{} `json:"LinkedTestPlan"`
	Release           interface{} `json:"Release"`
	Iteration         interface{} `json:"Iteration"`
	TeamIteration     interface{} `json:"TeamIteration"`
	Team              struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"Team"`
	Priority struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Importance   int    `json:"Importance"`
	} `json:"Priority"`
	EntityState struct {
		ResourceType    string  `json:"ResourceType"`
		ID              int     `json:"Id"`
		Name            string  `json:"Name"`
		NumericPriority float64 `json:"NumericPriority"`
	} `json:"EntityState"`
	ResponsibleTeam struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
	} `json:"ResponsibleTeam"`
	Feature      interface{}   `json:"Feature"`
	CustomFields []interface{} `json:"CustomFields"`
}

type BugsGetResponse struct {
	Next  string `json:"Next"`
	Items []struct {
		ResourceType string  `json:"ResourceType"`
		ID           int     `json:"Id"`
		Name         string  `json:"Name"`
		Effort       float64 `json:"Effort"`
		Project      struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Project"`
	} `json:"Items"`
}

type BugsPostResponse struct {
	Items []struct {
		ResourceType        string      `json:"ResourceType"`
		ID                  int         `json:"Id"`
		Name                string      `json:"Name"`
		Description         string      `json:"Description"`
		StartDate           interface{} `json:"StartDate"`
		EndDate             interface{} `json:"EndDate"`
		CreateDate          string      `json:"CreateDate"`
		ModifyDate          string      `json:"ModifyDate"`
		LastCommentDate     interface{} `json:"LastCommentDate"`
		Tags                string      `json:"Tags"`
		NumericPriority     int         `json:"NumericPriority"`
		Effort              int         `json:"Effort"`
		EffortCompleted     int         `json:"EffortCompleted"`
		EffortToDo          int         `json:"EffortToDo"`
		Progress            int         `json:"Progress"`
		TimeSpent           int         `json:"TimeSpent"`
		TimeRemain          int         `json:"TimeRemain"`
		LastStateChangeDate string      `json:"LastStateChangeDate"`
		PlannedStartDate    interface{} `json:"PlannedStartDate"`
		PlannedEndDate      interface{} `json:"PlannedEndDate"`
		Units               string      `json:"Units"`
		EntityType          struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"EntityType"`
		Project struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
			Process      struct {
				ResourceType string `json:"ResourceType"`
				ID           int    `json:"Id"`
				Name         string `json:"Name"`
			} `json:"Process"`
		} `json:"Project"`
		LastEditor struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			FirstName    string `json:"FirstName"`
			LastName     string `json:"LastName"`
			Login        string `json:"Login"`
		} `json:"LastEditor"`
		Owner struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			FirstName    string `json:"FirstName"`
			LastName     string `json:"LastName"`
			Login        string `json:"Login"`
		} `json:"Owner"`
		LastCommentedUser interface{} `json:"LastCommentedUser"`
		LinkedTestPlan    interface{} `json:"LinkedTestPlan"`
		Release           interface{} `json:"Release"`
		Iteration         interface{} `json:"Iteration"`
		TeamIteration     interface{} `json:"TeamIteration"`
		Team              interface{} `json:"Team"`
		Priority          struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
			Importance   int    `json:"Importance"`
		} `json:"Priority"`
		EntityState struct {
			ResourceType    string `json:"ResourceType"`
			ID              int    `json:"Id"`
			Name            string `json:"Name"`
			NumericPriority int    `json:"NumericPriority"`
		} `json:"EntityState"`
		ResponsibleTeam interface{} `json:"ResponsibleTeam"`
		Build           interface{} `json:"Build"`
		UserStory       interface{} `json:"UserStory"`
		Feature         interface{} `json:"Feature"`
		Severity        struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
			Importance   int    `json:"Importance"`
		} `json:"Severity"`
		CustomFields []struct {
			Name  string      `json:"Name"`
			Type  string      `json:"Type"`
			Value interface{} `json:"Value"`
		} `json:"CustomFields"`
	} `json:"Items"`
}

type EpicsGetResponse struct {
	Items []struct {
		ResourceType  string  `json:"ResourceType"`
		ID            int     `json:"Id"`
		Name          string  `json:"Name"`
		Effort        float64 `json:"Effort"`
		FeaturesCount int     `json:"Features-Count"`
		Project       struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Project"`
	} `json:"Items"`
}

type EpicsPostResponse struct {
	ResourceType        string      `json:"ResourceType"`
	ID                  int         `json:"Id"`
	Name                string      `json:"Name"`
	Description         interface{} `json:"Description"`
	StartDate           interface{} `json:"StartDate"`
	EndDate             interface{} `json:"EndDate"`
	CreateDate          string      `json:"CreateDate"`
	ModifyDate          string      `json:"ModifyDate"`
	LastCommentDate     interface{} `json:"LastCommentDate"`
	Tags                string      `json:"Tags"`
	NumericPriority     float64     `json:"NumericPriority"`
	Effort              float64     `json:"Effort"`
	EffortCompleted     float64     `json:"EffortCompleted"`
	EffortToDo          float64     `json:"EffortToDo"`
	Progress            float64     `json:"Progress"`
	TimeSpent           float64     `json:"TimeSpent"`
	TimeRemain          float64     `json:"TimeRemain"`
	LastStateChangeDate string      `json:"LastStateChangeDate"`
	PlannedStartDate    interface{} `json:"PlannedStartDate"`
	PlannedEndDate      interface{} `json:"PlannedEndDate"`
	InitialEstimate     float64     `json:"InitialEstimate"`
	EntityType          struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"EntityType"`
	Project struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Process      struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Process"`
	} `json:"Project"`
	Owner struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		FirstName    string `json:"FirstName"`
		LastName     string `json:"LastName"`
		Login        string `json:"Login"`
	} `json:"Owner"`
	LastCommentedUser interface{} `json:"LastCommentedUser"`
	LinkedTestPlan    interface{} `json:"LinkedTestPlan"`
	Release           interface{} `json:"Release"`
	Iteration         interface{} `json:"Iteration"`
	TeamIteration     interface{} `json:"TeamIteration"`
	Team              struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"Team"`
	Priority struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Importance   int    `json:"Importance"`
	} `json:"Priority"`
	EntityState struct {
		ResourceType    string  `json:"ResourceType"`
		ID              int     `json:"Id"`
		Name            string  `json:"Name"`
		NumericPriority float64 `json:"NumericPriority"`
	} `json:"EntityState"`
	ResponsibleTeam struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
	} `json:"ResponsibleTeam"`
	CustomFields []interface{} `json:"CustomFields"`
}

type FeaturesGetResponse struct {
	Items []struct {
		ResourceType     string  `json:"ResourceType"`
		ID               int     `json:"Id"`
		Name             string  `json:"Name"`
		Effort           float64 `json:"Effort"`
		UserStoriesCount int     `json:"UserStories-Count"`
		Project          struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Project"`
	} `json:"Items"`
}

type FeaturesPostResponse struct {
	ResourceType        string      `json:"ResourceType"`
	ID                  int         `json:"Id"`
	Name                string      `json:"Name"`
	Description         interface{} `json:"Description"`
	StartDate           interface{} `json:"StartDate"`
	EndDate             interface{} `json:"EndDate"`
	CreateDate          string      `json:"CreateDate"`
	ModifyDate          string      `json:"ModifyDate"`
	LastCommentDate     interface{} `json:"LastCommentDate"`
	Tags                string      `json:"Tags"`
	NumericPriority     float64     `json:"NumericPriority"`
	Effort              float64     `json:"Effort"`
	EffortCompleted     float64     `json:"EffortCompleted"`
	EffortToDo          float64     `json:"EffortToDo"`
	Progress            float64     `json:"Progress"`
	TimeSpent           float64     `json:"TimeSpent"`
	TimeRemain          float64     `json:"TimeRemain"`
	LastStateChangeDate string      `json:"LastStateChangeDate"`
	PlannedStartDate    interface{} `json:"PlannedStartDate"`
	PlannedEndDate      interface{} `json:"PlannedEndDate"`
	InitialEstimate     float64     `json:"InitialEstimate"`
	EntityType          struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"EntityType"`
	Project struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Process      struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Process"`
	} `json:"Project"`
	Owner struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		FirstName    string `json:"FirstName"`
		LastName     string `json:"LastName"`
		Login        string `json:"Login"`
	} `json:"Owner"`
	LastCommentedUser interface{} `json:"LastCommentedUser"`
	LinkedTestPlan    interface{} `json:"LinkedTestPlan"`
	Release           interface{} `json:"Release"`
	Iteration         interface{} `json:"Iteration"`
	TeamIteration     interface{} `json:"TeamIteration"`
	Team              struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"Team"`
	Priority struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Importance   int    `json:"Importance"`
	} `json:"Priority"`
	EntityState struct {
		ResourceType    string  `json:"ResourceType"`
		ID              int     `json:"Id"`
		Name            string  `json:"Name"`
		NumericPriority float64 `json:"NumericPriority"`
	} `json:"EntityState"`
	ResponsibleTeam struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
	} `json:"ResponsibleTeam"`
	Epic         interface{}   `json:"Epic"`
	CustomFields []interface{} `json:"CustomFields"`
}

type TasksGetResponse struct {
	Next  string `json:"Next"`
	Items []struct {
		ResourceType string  `json:"ResourceType"`
		ID           int     `json:"Id"`
		Name         string  `json:"Name"`
		Effort       float64 `json:"Effort"`
		Project      struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Project"`
	} `json:"Items"`
}

type TasksPostResponse struct {
	ResourceType        string      `json:"ResourceType"`
	ID                  int         `json:"Id"`
	Name                string      `json:"Name"`
	Description         interface{} `json:"Description"`
	StartDate           interface{} `json:"StartDate"`
	EndDate             interface{} `json:"EndDate"`
	CreateDate          string      `json:"CreateDate"`
	ModifyDate          string      `json:"ModifyDate"`
	LastCommentDate     interface{} `json:"LastCommentDate"`
	Tags                string      `json:"Tags"`
	NumericPriority     float64     `json:"NumericPriority"`
	Effort              float64     `json:"Effort"`
	EffortCompleted     float64     `json:"EffortCompleted"`
	EffortToDo          float64     `json:"EffortToDo"`
	Progress            float64     `json:"Progress"`
	TimeSpent           float64     `json:"TimeSpent"`
	TimeRemain          float64     `json:"TimeRemain"`
	LastStateChangeDate string      `json:"LastStateChangeDate"`
	PlannedStartDate    interface{} `json:"PlannedStartDate"`
	PlannedEndDate      interface{} `json:"PlannedEndDate"`
	EntityType          struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"EntityType"`
	Project struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Process      struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Process"`
	} `json:"Project"`
	Owner struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		FirstName    string `json:"FirstName"`
		LastName     string `json:"LastName"`
		Login        string `json:"Login"`
	} `json:"Owner"`
	LastCommentedUser interface{} `json:"LastCommentedUser"`
	LinkedTestPlan    interface{} `json:"LinkedTestPlan"`
	Release           interface{} `json:"Release"`
	Iteration         interface{} `json:"Iteration"`
	TeamIteration     interface{} `json:"TeamIteration"`
	Team              interface{} `json:"Team"`
	Priority          struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Importance   int    `json:"Importance"`
	} `json:"Priority"`
	EntityState struct {
		ResourceType    string  `json:"ResourceType"`
		ID              int     `json:"Id"`
		Name            string  `json:"Name"`
		NumericPriority float64 `json:"NumericPriority"`
	} `json:"EntityState"`
	ResponsibleTeam interface{} `json:"ResponsibleTeam"`
	UserStory       struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"UserStory"`
	CustomFields []interface{} `json:"CustomFields"`
}
type RequestersGetResponse struct {
	Items []struct {
		ResourceType  string `json:"ResourceType"`
		Kind          string `json:"Kind"`
		ID            int    `json:"Id"`
		FirstName     string `json:"FirstName"`
		LastName      string `json:"LastName"`
		Email         string `json:"Email"`
		RequestsCount int    `json:"Requests-Count"`
	} `json:"Items"`
}

type RequestersPostResponse struct {
	ResourceType    string        `json:"ResourceType"`
	Kind            string        `json:"Kind"`
	ID              int           `json:"Id"`
	FirstName       string        `json:"FirstName"`
	LastName        string        `json:"LastName"`
	Email           string        `json:"Email"`
	Login           string        `json:"Login"`
	CreateDate      string        `json:"CreateDate"`
	ModifyDate      string        `json:"ModifyDate"`
	DeleteDate      interface{}   `json:"DeleteDate"`
	IsActive        bool          `json:"IsActive"`
	IsAdministrator bool          `json:"IsAdministrator"`
	Phone           interface{}   `json:"Phone"`
	Notes           interface{}   `json:"Notes"`
	SourceType      string        `json:"SourceType"`
	Company         interface{}   `json:"Company"`
	CustomFields    []interface{} `json:"CustomFields"`
}

type UsersGetReponse struct {
	Items []struct {
		ResourceType  string  `json:"ResourceType"`
		Kind          string  `json:"Kind"`
		ID            int     `json:"Id"`
		FirstName     string  `json:"FirstName"`
		LastName      string  `json:"LastName"`
		Email         string  `json:"Email"`
		TimesSpentSum float64 `json:"Times-Spent-Sum"`
	} `json:"Items"`
}

type UsersPostResponse struct {
	ResourceType              string      `json:"ResourceType"`
	Kind                      string      `json:"Kind"`
	ID                        int         `json:"Id"`
	FirstName                 string      `json:"FirstName"`
	LastName                  string      `json:"LastName"`
	Email                     string      `json:"Email"`
	Login                     string      `json:"Login"`
	CreateDate                string      `json:"CreateDate"`
	ModifyDate                string      `json:"ModifyDate"`
	DeleteDate                interface{} `json:"DeleteDate"`
	IsActive                  bool        `json:"IsActive"`
	IsAdministrator           bool        `json:"IsAdministrator"`
	LastLoginDate             interface{} `json:"LastLoginDate"`
	WeeklyAvailableHours      float64     `json:"WeeklyAvailableHours"`
	CurrentAllocation         int         `json:"CurrentAllocation"`
	CurrentAvailableHours     float64     `json:"CurrentAvailableHours"`
	AvailableFrom             string      `json:"AvailableFrom"`
	AvailableFutureAllocation int         `json:"AvailableFutureAllocation"`
	AvailableFutureHours      float64     `json:"AvailableFutureHours"`
	IsObserver                bool        `json:"IsObserver"`
	IsContributor             bool        `json:"IsContributor"`
	Locale                    interface{} `json:"Locale"`
	Skills                    interface{} `json:"Skills"`
	ActiveDirectoryName       interface{} `json:"ActiveDirectoryName"`
	Role                      struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"Role"`
	CustomFields []interface{} `json:"CustomFields"`
}

type AssignablesGetResponse struct {
	Items []struct {
		ResourceType string  `json:"ResourceType"`
		ID           int     `json:"Id"`
		Name         string  `json:"Name"`
		Effort       float64 `json:"Effort"`
		EntityType   struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"EntityType"`
		Project struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Project"`
	} `json:"Items"`
}

type AssignablesPostResponse struct {
	ResourceType        string      `json:"ResourceType"`
	ID                  int         `json:"Id"`
	Name                string      `json:"Name"`
	Description         interface{} `json:"Description"`
	StartDate           interface{} `json:"StartDate"`
	EndDate             interface{} `json:"EndDate"`
	CreateDate          string      `json:"CreateDate"`
	ModifyDate          string      `json:"ModifyDate"`
	LastCommentDate     interface{} `json:"LastCommentDate"`
	Tags                string      `json:"Tags"`
	NumericPriority     float64     `json:"NumericPriority"`
	Effort              float64     `json:"Effort"`
	EffortCompleted     float64     `json:"EffortCompleted"`
	EffortToDo          float64     `json:"EffortToDo"`
	Progress            float64     `json:"Progress"`
	TimeSpent           float64     `json:"TimeSpent"`
	TimeRemain          float64     `json:"TimeRemain"`
	LastStateChangeDate string      `json:"LastStateChangeDate"`
	PlannedStartDate    interface{} `json:"PlannedStartDate"`
	PlannedEndDate      interface{} `json:"PlannedEndDate"`
	EntityType          struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"EntityType"`
	Project struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Process      struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Process"`
	} `json:"Project"`
	Owner struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		FirstName    string `json:"FirstName"`
		LastName     string `json:"LastName"`
		Login        string `json:"Login"`
	} `json:"Owner"`
	LastCommentedUser interface{} `json:"LastCommentedUser"`
	LinkedTestPlan    interface{} `json:"LinkedTestPlan"`
	Release           interface{} `json:"Release"`
	Iteration         interface{} `json:"Iteration"`
	TeamIteration     interface{} `json:"TeamIteration"`
	Team              interface{} `json:"Team"`
	Priority          struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Importance   int    `json:"Importance"`
	} `json:"Priority"`
	EntityState struct {
		ResourceType    string  `json:"ResourceType"`
		ID              int     `json:"Id"`
		Name            string  `json:"Name"`
		NumericPriority float64 `json:"NumericPriority"`
	} `json:"EntityState"`
	ResponsibleTeam interface{}   `json:"ResponsibleTeam"`
	CustomFields    []interface{} `json:"CustomFields"`
}

type GeneralsGetResponse struct {
	Items []struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		CreateDate   string `json:"CreateDate"`
		EntityType   struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"EntityType"`
		Owner struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			FirstName    string `json:"FirstName"`
			LastName     string `json:"LastName"`
			Login        string `json:"Login"`
		} `json:"Owner"`
	} `json:"Items"`
}

type GeneralsPostResponse struct {
	ResourceType    string      `json:"ResourceType"`
	ID              int         `json:"Id"`
	Name            string      `json:"Name"`
	Description     interface{} `json:"Description"`
	StartDate       interface{} `json:"StartDate"`
	EndDate         interface{} `json:"EndDate"`
	CreateDate      string      `json:"CreateDate"`
	ModifyDate      string      `json:"ModifyDate"`
	LastCommentDate interface{} `json:"LastCommentDate"`
	Tags            string      `json:"Tags"`
	NumericPriority float64     `json:"NumericPriority"`
	EntityType      struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"EntityType"`
	Project struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
		Process      struct {
			ResourceType string `json:"ResourceType"`
			ID           int    `json:"Id"`
			Name         string `json:"Name"`
		} `json:"Process"`
	} `json:"Project"`
	Owner struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		FirstName    string `json:"FirstName"`
		LastName     string `json:"LastName"`
		Login        string `json:"Login"`
	} `json:"Owner"`
	LastCommentedUser interface{}   `json:"LastCommentedUser"`
	LinkedTestPlan    interface{}   `json:"LinkedTestPlan"`
	CustomFields      []interface{} `json:"CustomFields"`
}

type ProjectsGetResponse struct {
	Items []struct {
		ResourceType     string  `json:"ResourceType"`
		ID               int     `json:"Id"`
		Name             string  `json:"Name"`
		Effort           float64 `json:"Effort"`
		UserStoriesCount int     `json:"UserStories-Count"`
	} `json:"Items"`
}

type ProjectsPostResponse struct {
	ResourceType        string      `json:"ResourceType"`
	ID                  int         `json:"Id"`
	Name                string      `json:"Name"`
	Description         interface{} `json:"Description"`
	StartDate           interface{} `json:"StartDate"`
	EndDate             interface{} `json:"EndDate"`
	CreateDate          string      `json:"CreateDate"`
	ModifyDate          string      `json:"ModifyDate"`
	LastCommentDate     interface{} `json:"LastCommentDate"`
	Tags                string      `json:"Tags"`
	NumericPriority     float64     `json:"NumericPriority"`
	Effort              float64     `json:"Effort"`
	EffortCompleted     float64     `json:"EffortCompleted"`
	EffortToDo          float64     `json:"EffortToDo"`
	IsActive            bool        `json:"IsActive"`
	IsProduct           bool        `json:"IsProduct"`
	Abbreviation        string      `json:"Abbreviation"`
	MailReplyAddress    interface{} `json:"MailReplyAddress"`
	Color               interface{} `json:"Color"`
	Progress            float64     `json:"Progress"`
	PlannedStartDate    interface{} `json:"PlannedStartDate"`
	PlannedEndDate      interface{} `json:"PlannedEndDate"`
	LastStateChangeDate string      `json:"LastStateChangeDate"`
	IsPrivate           interface{} `json:"IsPrivate"`
	Units               string      `json:"Units"`
	EntityType          struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"EntityType"`
	Process struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		Name         string `json:"Name"`
	} `json:"Process"`
	LastEditor struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		FirstName    string `json:"FirstName"`
		LastName     string `json:"LastName"`
		Login        string `json:"Login"`
	} `json:"LastEditor"`
	Owner struct {
		ResourceType string `json:"ResourceType"`
		ID           int    `json:"Id"`
		FirstName    string `json:"FirstName"`
		LastName     string `json:"LastName"`
		Login        string `json:"Login"`
	} `json:"Owner"`
	LastCommentedUser interface{} `json:"LastCommentedUser"`
	Project           interface{} `json:"Project"`
	LinkedTestPlan    interface{} `json:"LinkedTestPlan"`
	Program           interface{} `json:"Program"`
	EntityState       struct {
		ResourceType    string  `json:"ResourceType"`
		ID              int     `json:"Id"`
		Name            string  `json:"Name"`
		NumericPriority float64 `json:"NumericPriority"`
	} `json:"EntityState"`
	Company      interface{}   `json:"Company"`
	CustomFields []interface{} `json:"CustomFields"`
}

type BadRequest struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
	Type    string `json:"Type"`
	Details struct {
		Items []struct {
			Type    string `json:"Type"`
			Message struct {
				Token string `json:"Token"`
				Data  struct {
					Format   string `json:"format"`
					Line     string `json:"line"`
					Position string `json:"position"`
				} `json:"Data"`
				Value string `json:"Value"`
			} `json:"Message"`
		} `json:"Items"`
	} `json:"Details"`
	ErrorID string `json:"ErrorId"`
}
