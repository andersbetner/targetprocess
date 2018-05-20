package targetprocess

type UserStoriesQueryParams struct {
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

type UserStoriesBodyParms struct {
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

type UserStoriesResponse struct {
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

type BugsQueryParams struct {
}

type BugsBodyParams struct {
}

type EpicsQueryParams struct {
}

type EpicsBodyParams struct {
}

type FeaturesQueryParams struct {
}

type FeaturesBodyParams struct {
}

type TasksQueryParams struct {
}

type TasksBodyParams struct {
}

type RequestersQueryParams struct {
}

type RequestersBodyParams struct {
}

type UsersQueryParams struct {
}

type UsersBodyParams struct {
}

type AssignablesQueryParams struct {
}

type AssignabesBodyParams struct {
}

type GeneralsQueryParams struct {
}

type GeneralsBodyParams struct {
}

type ProjectsQueryParams struct {
}

type ProjectsBodyParams struct {
}
