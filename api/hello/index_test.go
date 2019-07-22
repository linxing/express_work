package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../../routers"
	mockrouter "../../routers/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {

	t.Run("should return http code 200 if api address was right", func(t *testing.T){

		router := routers.SetupRouter(mockrouter.Servlet{
			Hello: new(Servlet),
		})
		w := httptest.NewRecorder()
	
		req, err := http.NewRequest("GET", "/api/index", nil)
		require.NoError(t, err)
	
		router.ServeHTTP(w, req)
	
		assert.Equal(t, 200, w.Code)
	})

	t.Run("should return http.code 404 if api address was not found", func(t *testing.T){

		router := routers.SetupRouter(mockrouter.Servlet{
			Hello: new(Servlet),
		})
		w := httptest.NewRecorder()
	
		req, err := http.NewRequest("GET", "/api/index/xx", nil)
		require.NoError(t, err)
	
		router.ServeHTTP(w, req)
	
		assert.Equal(t, 404, w.Code)
	})
}
