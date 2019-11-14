package client

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	bblfsh "github.com/bblfsh/go-client/v4"
	"github.com/stretchr/testify/require"
)

// TestK8s is a simple test that infinitely bombs given endpoint with parse requests
func TestK8s(t *testing.T) {
	address := os.Getenv("ADDRESS")
	fmt.Println(address)
	cli := newClient(t, address)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for {
				testNativeParseRequestCustom(t, cli, "go", "package main")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func newClient(t testing.TB, endpoint string) *bblfsh.Client {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cli, err := bblfsh.NewClientContext(ctx, endpoint)
	if err == context.DeadlineExceeded {
		t.Skip("bblfshd is not running")
	}
	require.Nil(t, err)
	return cli
}

func testNativeParseRequestCustom(t *testing.T, cli *bblfsh.Client, lang, content string) {
	res, err := cli.NewParseRequest().Mode(bblfsh.Native).Language(lang).Content(content).Do()
	require.NoError(t, err)

	require.Equal(t, 0, len(res.Errors))
	require.NotNil(t, res)
}
