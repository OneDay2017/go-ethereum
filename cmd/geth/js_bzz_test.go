package main

import (
	// "io/ioutil"
	"os"
	// "path"
	// "runtime"
	"testing"

	// "github.com/ethereum/go-ethereum/bzz"
	// "github.com/ethereum/go-ethereum/common/resolver"
	"github.com/ethereum/go-ethereum/eth"
)

func bzzREPL(t *testing.T, port string) (string, *testjethre, *eth.Ethereum) {
	return testREPL(t, func(c *eth.Config) {
		c.Bzz = true
		c.BzzPort = port
	})
}

func TestBzzUploadDownload(t *testing.T) {
	tmp, repl, ethereum := bzzREPL(t, "")
	if err := ethereum.Start(); err != nil {
		t.Fatalf("error starting ethereum: %v", err)
	}
	defer ethereum.Stop()
	defer os.RemoveAll(tmp)
	_ = repl
}

func TestBzzPutGet(t *testing.T) {
	tmp, repl, ethereum := bzzREPL(t, "")
	if err := ethereum.Start(); err != nil {
		t.Fatalf("error starting ethereum: %v", err)
	}
	defer ethereum.Stop()
	defer os.RemoveAll(tmp)
	if checkEvalJSON(t, repl, `hash = bzz.put("console.log(\"hello from console\")", "application/javascript")`, `"97f1b7c7ea12468fd37c262383b9aa862d0cfbc4fc7218652374679fc5cf40cd"`) != nil {
		return
	}
	want := `{"content":"console.log(\"hello from console\")","contentType":"application/javascript","size":"33","status":"0"}`
	if checkEvalJSON(t, repl, `bzz.get(hash)`, want) != nil {
		return
	}
}

// the server can be initialized only once per test session !
// until we implement a stoppable http server
// further http tests will need to make sure the correct server is running
func TestHTTP(t *testing.T) {
	tmp, repl, ethereum := bzzREPL(t, "8500")
	if err := ethereum.Start(); err != nil {
		t.Fatalf("error starting ethereum: %v", err)
	}
	defer ethereum.Stop()
	defer os.RemoveAll(tmp)
	if checkEvalJSON(t, repl, `hash = bzz.put("6*7", "application/javascript")`, `"97f1b7c7ea12468fd37c262383b9aa862d0cfbc4fc7218652374679fc5cf40cd"`) != nil {
		return
	}
	if checkEvalJSON(t, repl, `http.get("bzz://"+hash)`, `"6*7"`) != nil {
		return
	}

	if checkEvalJSON(t, repl, `http.loadScript("bzz://"+hash)`, `42`) != nil {
		return
	}
}
