package helpers

import (
	"domcermak/ctc/assignments/03/cmd/common"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"testing"

	"domcermak/ctc/assignments/03/cmd/server"
)

func Expect(expected, actual interface{}, t *testing.T) {
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Fatalf("Expected %v, but got %v", expected, actual)
	}
}

func ExpectProduct(expected, actual server.Product, t *testing.T) {
	Expect(expected.Id, actual.Id, t)
	Expect(expected.Name, actual.Name, t)
	Expect(expected.Price, actual.Price, t)
	Expect(expected.Amount, actual.Amount, t)
}

func ExpectProducts(expected, actual []server.Product, t *testing.T) {
	Expect(len(expected), len(actual), t)
	for i, product := range actual {
		ExpectProduct(expected[i], product, t)
	}
}

func ExpectIdInRequestVars(expected server.Id, r *http.Request, t *testing.T) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		t.Fatal(err)
	}
	Expect(expected, id, t)
}

func ExpectJsonRequestBody(expected map[string]interface{}, r *http.Request, t *testing.T) {
	defer r.Body.Close()

	actual := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&actual); err != nil {
		t.Fatal(err)
	}
	expectedKeys, actualKeys := common.SortedKeys(expected), common.SortedKeys(actual)

	Expect(expectedKeys, actualKeys, t)
	for _, key := range expectedKeys {
		Expect(expected[key], actual[key], t)
	}
}
