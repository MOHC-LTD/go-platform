package http

// Option allows the client to be fully configurable
type Option interface {
	// configure takes the client and configures it with the option
	configure(client *Client) error
}
