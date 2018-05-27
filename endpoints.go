package targetprocess

type Endpoint struct {
	client *TPClient
	target string
	query  queryParams
	json   jsonBodyParams
}

const (
	enpoints = map[int]string{
		1: "UserStories",
		2: "Bugs",
		3: "Epics",
		4: "Features",
		5: "Tasks",
		6: "Requesters",
		7: "Users",
		8: "Assignables",
		9: "Generals",
		10: "Projects",
	}
)

// Response is the returned results from targetprocess with the sent
// json and query parameters
type Response struct {
	QueryParams queryParams
	JSONParams  jsonBodyParams
	Results     interface{}
}

// UserStories query userstories
func (c *TPClient) UserStories(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "UserStories/",
		query:  q,
		json:   j,
		endpoint: 
	}
	return ep
}

func (c *TPClient) Bugs(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "Bugs/",
		query:  q,
		json:   j,
	}
	return ep
}

func (c *TPClient) Epics(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "Epics/",
		query:  q,
		json:   j,
	}
	return ep
}

func (c *TPClient) Features(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "Features/",
		query:  q,
		json:   j,
	}
	return ep
}

func (c *TPClient) Tasks(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "Tasks/",
		query:  q,
		json:   j,
	}
	return ep
}

func (c *TPClient) Requesters(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "Requesters/",
		query:  q,
		json:   j,
	}
	return ep
}

func (c *TPClient) Users(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "Users/",
		query:  q,
		json:   j,
	}
	return ep
}

func (c *TPClient) Assignables(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "Assignables/",
		query:  q,
		json:   j,
	}
	return ep
}

func (c *TPClient) Generals(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "Generals/",
		query:  q,
		json:   j,
	}
	return ep
}

func (c *TPClient) Projects(q queryParams, j jsonBodyParams) *Endpoint {
	ep := &Endpoint{
		client: c,
		target: c.HostAddress + "Projects/",
		query:  q,
		json:   j,
	}
	return ep
}

// Get queries the endpoint for existing data
func (e *Endpoint) Get() (interface{}, BadRequest) {
	e.query.format = "json"
	return nil, BadRequest{}
}

// Post updates or creates objects
func (e *Endpoint) Post() (interface{}, BadRequest) {
	e.client.SetHeader("Content-type", "application/json")
	e.query.format = "json"
	return nil, BadRequest{}
}

// Delete removes the given ID from the endpoint
func (e *Endpoint) Delete(id string) {
	return
}
