package targetprocess

// TPClient contains the base client for accessing the API
type TPClient struct {
	url             string
	basicAuth       bool
	tokenAuth       bool
	accessTokenAuth bool
	auth            string
}
