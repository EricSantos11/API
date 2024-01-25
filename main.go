package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)


type UUIDResponse []string

func obterUUIDs(count int) (UUIDResponse, error) {
	url := fmt.Sprintf("https://www.uuidtools.com/api/generate/v4/count/%d", count)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var uuids UUIDResponse
	err = json.Unmarshal(body, &uuids)
	if err != nil {
		return nil, err
	}

	return uuids, nil
}

func removerNumeros(uuid string) string {
	re := regexp.MustCompile("\\d")
	return re.ReplaceAllString(uuid, "")
}

func compararUUIDs(original, modificado string) {
	fmt.Printf("Original: %s / Modificado: %s\n", original, modificado)
}

func main() {
	// Pega UUIDs
	uuidsOriginais, err := obterUUIDs(20)
	if err != nil {
		fmt.Printf("Erro ao obter UUIDs: %v\n", err)
		return
	}

	// Remover números e compara
	for _, uuidOriginal := range uuidsOriginais {
		uuidModificado := removerNumeros(uuidOriginal)
		compararUUIDs(uuidOriginal, uuidModificado)
	}
}