package serializer

import "encoding/json"

type ResponseJSON struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// TODO: correct marshalling
func Marshal(data interface{}) []byte {

	response := ResponseJSON{}

	switch data := data.(type) {
	case error:
		response.Error = data.Error()

	default:
		response.Data = data
		response.Success = false
	}

	marshaledResponse, _ := json.Marshal(response)
	// if err != nil {
	// 	return
	// }

	return marshaledResponse
}
