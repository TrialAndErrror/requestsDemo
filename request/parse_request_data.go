package request

import (
	"encoding/json"
	"log"
	"strings"
)

func parseKeyValuePairs(params string) map[string]interface{} {
	paramsMap := make(map[string]interface{})
	paramParts := strings.Split(params, "&")
	for i := 0; i < len(paramParts); i++ {
		keyValue := strings.SplitN(paramParts[i], "=", 2)
		key, value := keyValue[0], keyValue[1]
		paramsMap[key] = value
	}
	return paramsMap
}

func parseQueryParams(endpoint string) map[string]interface{} {
	endpointPartsList := strings.SplitN(endpoint, "?", 2)
	if len(endpointPartsList) < 2 {
		log.Printf("Error parsing query params: %v", endpointPartsList)
		return map[string]interface{}{}
	}
	paramsString := endpointPartsList[1]
	return parseKeyValuePairs(paramsString)
}

func processJSONData(jsonDataString string) map[string]interface{} {
	jsonData := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonDataString), &jsonData)
	if err != nil {
		log.Println("Failed to parse request: ", err)
		return map[string]interface{}{}
	}
	return jsonData
}
