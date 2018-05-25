package targetprocess

type Endpoint struct {
	client   *TPClient
	Endpoint string
}

func (c *TPClient) UserStories() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "UserStories/",
	}
	return ep
}

func (c *TPClient) Bugs() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "Bugs/",
	}
	return ep
}

func (c *TPClient) Epics() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "Epics/",
	}
	return ep
}

func (c *TPClient) Features() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "Features/",
	}
	return ep
}

func (c *TPClient) Tasks() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "Tasks/",
	}
	return ep
}

func (c *TPClient) Requesters() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "Requesters/",
	}
	return ep
}

func (c *TPClient) Users() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "Users/",
	}
	return ep
}

func (c *TPClient) Assignables() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "Assignables/",
	}
	return ep
}

func (c *TPClient) Generals() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "Generals/",
	}
	return ep
}

func (c *TPClient) Projects() *Endpoint {
	ep := &Endpoint{
		client:   c,
		Endpoint: "Projects/",
	}
	return ep
}

// Get queries the endpoint for existing data
func (c *Endpoint) Get(q queryParams) (interface{}, BadRequest) {
	q.format = "json"
	return nil, BadRequest{}
}

// Post updates or creates objects
func (c *Endpoint) Post(q queryParams, j jsonBodyParams) {
	c.client.SetHeader("Content-type", "application/json")
	q.format = "json"

}

// Delete removes the given ID from the endpoint
func (c *Endpoint) Delete(id string) {
	return
}
