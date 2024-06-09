package response

type OutletsResponse struct {
	Id                int    `json:"id"`
	OutletName        string `json:"outlet_name"`
	Slug              string `json:"slug"`
	Address           string `json:"address"`
	Latitude          string `json:"lat"`
	Longitude         string `json:"lon"`
	OperationTime     string `json:"operation_time"`
	Contact           string `json:"contact"`
	GofoodLink        string `json:"gofood_link"`
	ShopeefoodLink    string `json:"shopeefood_link"`
	GrabfoodLink      string `json:"grabfood_link"`
	TravelokaEatsLink string `json:"travelokaeats_link"`
}
