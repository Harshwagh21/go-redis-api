package domain

type CacheRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
	TTL   int    `json:"ttl"`
}

type HashRequest struct {
	Key    string            `json:"key" binding:"required"`
	Fields map[string]string `json:"fields" binding:"required"`
}

type ListRequest struct {
	Key    string   `json:"key" binding:"required"`
	Values []string `json:"values" binding:"required"`
}

type SetRequest struct {
	Key     string   `json:"key" binding:"required"`
	Members []string `json:"members" binding:"required"`
}

type SortedSetMember struct {
	Score  float64 `json:"score"`
	Member string  `json:"member"`
}

type SortedSetRequest struct {
	Key     string            `json:"key" binding:"required"`
	Members []SortedSetMember `json:"members" binding:"required"`
}

type GeoRequest struct {
	Key       string  `json:"key" binding:"required"`
	Name      string  `json:"name" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type StreamRequest struct {
	Stream string            `json:"stream" binding:"required"`
	Fields map[string]string `json:"fields" binding:"required"`
}

type PubSubRequest struct {
	Channel string `json:"channel" binding:"required"`
	Message string `json:"message" binding:"required"`
}
