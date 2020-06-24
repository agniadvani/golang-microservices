package polo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/agniadvani/golang-microservices/src/api/utilis/test_utilis"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/marco", nil)
	c := test_utilis.GetMockedContext(request, response)
	Polo(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
