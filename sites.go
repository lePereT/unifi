package unifi

// A Site is a physical location with UniFi devices managed by a UniFi
// Controller.

type Site struct {
	ID               string   `json:"_id"`
	AttrHiddenID     string   `json:"attr_hidden_id,omitempty"`
	AttrNoDelete     bool     `json:"attr_no_delete,omitempty"`
	Desc             string   `json:"desc"`
	Health           []health `json:"health"`
	Name             string   `json:"name"`
	NumNewAlarms     int      `json:"num_new_alarms"`
	LocationAccuracy float64  `json:"location_accuracy,omitempty"`
	LocationLat      float64  `json:"location_lat,omitempty"`
	LocationLng      float64  `json:"location_lng,omitempty"`
}

type health struct {
	NumAdopted      int    `json:"num_adopted,omitempty"`
	NumAp           int    `json:"num_ap,omitempty"`
	NumDisabled     int    `json:"num_disabled,omitempty"`
	NumDisconnected int    `json:"num_disconnected,omitempty"`
	NumGuest        int    `json:"num_guest,omitempty"`
	NumPending      int    `json:"num_pending,omitempty"`
	NumUser         int    `json:"num_user,omitempty"`
	RxBytesR        int    `json:"rx_bytes-r,omitempty"`
	Status          string `json:"status"`
	Subsystem       string `json:"subsystem"`
	TxBytesR        int    `json:"tx_bytes-r,omitempty"`
	NumGw           int    `json:"num_gw,omitempty"`
	NumSw           int    `json:"num_sw,omitempty"`
}

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

	return v.Sites, err
}
