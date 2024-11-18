package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"github.com/markable/internal/dependency"
	"github.com/markable/internal/http/api"
	"github.com/markable/internal/http/transport"
	"github.com/markable/internal/pkg/config"
)

type echoHandler func(c echo.Context) error
type TearDownSuite func(tb testing.TB)

// SetupSuite sets up the test suite
func SetupSuite(tb testing.TB) (a *api.MarkAbleApi, e *echo.Echo, tearDown TearDownSuite) {
	opt := config.Options{
		ConfigFileSource: config.SourceEnv,
		ConfigFile:       "../../test.env",
	}
	cfg, err := dependency.NewConfig(opt)
	if err != nil {
		tb.Fatal("Error in initializing config:", err)
	}
	// Initializing the database
	db, err := dependency.NewDatabaseConfig(cfg)
	if err != nil {
		tb.Fatal("Error in initializing database:", err)
	}

	// Create a new Echo instance
	e = echo.New()
	// Set up the validator middleware
	e.Validator = &transport.CustomValidator{Validator: validator.New()}

	a, err = dependency.NewMarkableApi(cfg, db)
	if err != nil {
		tb.Fatal("Error in initializing MarkableApi:", err)
	}
	td := func(tb testing.TB) {
		// Cleanup the database
		db.Close()

	}
	return a, e, td

}

// SendRequest sends a request to the given handler
func SendRequest(e *echo.Echo, handler echoHandler, method, path string, pathParams map[string]string, queryParams map[string]string, body interface{}) (rec *httptest.ResponseRecorder, err error) {
	var req *http.Request
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
		// Create a request with a JSON body
		reqJSON, _ := json.Marshal(body)
		req = httptest.NewRequest(method, path, bytes.NewReader(reqJSON))
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	// Set the headers
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Create a response recorder
	rec = httptest.NewRecorder()

	// Call the handler function
	ctx := e.NewContext(req, rec)

	// If there are path parameters, set them in the context
	if pathParams != nil {
		for k, v := range pathParams {
			ctx.SetParamNames(k)
			ctx.SetParamValues(v)
		}
	}

	// If there are query parameters, set them in the context
	if queryParams != nil {
		for k, v := range queryParams {
			ctx.QueryParams().Add(k, v)
		}
	}

	// Call the handler
	err = handler(ctx)

	return rec, err
}

func ParseResponse(t *testing.T, rec *httptest.ResponseRecorder, resp interface{}) {
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Error unmarshaling response body: %v", err)
	}
}

func ParseEntityData(t *testing.T, data interface{}, entity interface{}) {
	eData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Error marshaling entity data: %v", err)
	}
	err = json.Unmarshal(eData, &entity)
	if err != nil {
		t.Fatalf("Error unmarshaling response body: %v", err)
	}
}
