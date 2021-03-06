package e2e

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/SIProjects/faucet-api/test/fixture"
	testutils "github.com/SIProjects/faucet-api/test/testutils"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func loadFixture(name string) (*fixture.Fixture, error) {
	data, err := ioutil.ReadFile("../../res/fixtures/" + name)
	if err != nil {
		return nil, err
	}

	var fx fixture.Fixture
	yaml.Unmarshal(data, &fx)
	return &fx, nil
}

func TestFixtures(t *testing.T) {
	asserts := assert.New(t)

	fixtures := [...]string{
		"health.yaml",
		"payouts/view.yaml",
		"queue/create.yaml",
		"queue/create_balance_fail.yaml",
	}

	for _, name := range fixtures {
		fx, err := loadFixture(name)
		asserts.NoError(err)
		sb, err := testutils.NewSandbox(fx)
		asserts.NoError(err)
		defer sb.Close()

		for _, query := range fx.Setup.DB {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			err = sb.App.DB.Exec(ctx, query)
			asserts.NoError(err)
		}

		body := strings.NewReader(fx.Request.Body)

		req, err := http.NewRequest(fx.Request.Method, fx.Request.Path, body)

		asserts.NoError(err)

		for _, h := range fx.Request.Headers {
			req.Header.Add(h.Key, h.Value)
		}

		rr := httptest.NewRecorder()

		sb.App.Server.Router.ServeHTTP(rr, req)

		asserts.Equal(
			fx.Response.Status, rr.Code,
			fmt.Sprintf(
				"Fixture %s should respond with code %d", name,
				fx.Response.Status,
			),
		)

		if fx.Response.Body != nil {
			body, err := ioutil.ReadAll(rr.Body)
			asserts.NoError(err)
			asserts.Equal(
				*fx.Response.Body, string(body),
				fmt.Sprintf(
					"Fixture %s should have the correct body", name,
				),
			)
		}

		for _, pending := range fx.Cache.PendingResults {
			_, ok := sb.Cache.Pending[pending]
			asserts.True(ok)
		}

		for _, dbCheck := range fx.DatabaseChecks {
			rs, err := sb.App.DB.Conn.Query(context.Background(), dbCheck.Query)

			asserts.NoError(err)

			rowCount := 0
			for rs.Next() {
				rowCount++
			}

			asserts.NoError(err)

			asserts.Equal(
				rowCount, dbCheck.Rows,
				fmt.Sprintf(
					"Fixture %s should have the correct rows for query '%s'",
					name, dbCheck.Query,
				),
			)
		}
	}
}
