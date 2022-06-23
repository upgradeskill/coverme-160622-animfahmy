package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGETUsers(t *testing.T) {
	t.Run("returns data user", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/user", nil)
		response := httptest.NewRecorder()

		user(response, request)

		got := response.Body.String()
		want := `[{"ID":3,"Name":"ethan","Grade":21},{"ID":1,"Name":"wick","Grade":22},{"ID":5,"Name":"bourne","Grade":23},{"ID":4,"Name":"bond","Grade":23}]`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
func TestGETUser3(t *testing.T) {
	t.Run("returns data user 3", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/user?id=3", nil)
		response := httptest.NewRecorder()

		user(response, request)

		got := response.Body.String()
		want := `{"ID":3,"Name":"ethan","Grade":21}`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
func TestPOSTUser(t *testing.T) {

	/*
			buff := new(bytes.Buffer)
		json.NewEncoder(buff).Encode(data)

	*/
	newdata := strings.NewReader("name=newdata&grade=99")

	t.Run("it records baru grade 99 on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/user", newdata)
		response := httptest.NewRecorder()

		user(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		fmt.Println(response.Body.String())

		// if len(store.winCalls) != 1 {
		// 	t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		// }

		// if store.winCalls[0] != player {
		// 	t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		// }
	})
}
func TestPUTUser(t *testing.T) {

	newdata := strings.NewReader("name=edi&grade=79")

	t.Run("it update id 3 on PUT", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPut, "/user?id=3", newdata)
		response := httptest.NewRecorder()

		user(response, request)

		// reader.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		fmt.Println(response.Body.String())

		// if len(store.winCalls) != 1 {
		// 	t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		// }

		// if store.winCalls[0] != player {
		// 	t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		// }
	})
}
func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
