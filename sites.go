package unifi

// Sites returns all of the Sites managed by a UniFi Controller.
func (c *Client) Sites() ([]*Site, error) {
	var v struct {
		Sites []*Site `json:"data"`
	}

	req, err := c.newRequest(
		"GET",
		"/api/stat/sites",
		nil,
	)
	if err != nil {
		return nil, err
	}

	_, err = c.do(req, &v)

	req, err = c.newRequest(
		"GET",
		"/api/self/sites",
		nil,
	)
	if err != nil {
		return nil, err
	}

	_, err = c.do(req, &v)

	return v.Sites, err
}

// A Site is a physical location with UniFi devices managed by a UniFi
// Controller.
type Site struct {
	ID               string    `json:"_id"`
	AttrHiddenID     string    `json:"attr_hidden_id"`
	AttrNoDelete     bool      `json:"attr_no_delete"`
	Desc             string    `json:"desc"`
	Health           []*Health `json:"health"`
	Name             string    `json:"name"`
	NumNewAlarms     int       `json:"num_new_alarms"`
	LocationAccuracy float64   `json:"location_accuracy"`
	LocationLat      float64   `json:"location_lat"`
	LocationLng      float64   `json:"location_lng"`
	Role             string    `json:"role"`
}

// Health is a thing
type Health struct {
	NumAdopted      int    `json:"num_adopted"`
	NumAp           int    `json:"num_ap"`
	NumDisabled     int    `json:"num_disabled"`
	NumDisconnected int    `json:"num_disconnected"`
	NumGuest        int    `json:"num_guest"`
	NumPending      int    `json:"num_pending"`
	NumUser         int    `json:"num_user"`
	RxBytesR        int    `json:"rx_bytes-r"`
	Status          string `json:"status"`
	Subsystem       string `json:"subsystem"`
	TxBytesR        int    `json:"tx_bytes-r"`
	NumGw           int    `json:"num_gw"`
	NumSw           int    `json:"num_sw"`
}
