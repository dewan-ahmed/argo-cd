package smoke

import (
	"fmt"
	"testing"
	"time"

	"github.com/argoproj/argo-cd/util/errors"
	"gotest.tools/assert"
)

func TestSyncApp(t *testing.T) {
	expected := "Synced"
	// Create the guestbook app if not already created
	_, err := RunCmd("argocd", "app", "create", "guestbook", "--repo", "https://github.com/argoproj/argocd-example-apps.git", "--path", "guestbook", "--dest-server", "https://kubernetes.default.svc", "--dest-namespace", "default")
	errors.CheckError(err)
	time.Sleep(2 * time.Second)
	// Sync the app guestbook
	_, err = RunCmd("argocd", "app", "sync", "guestbook")
	output, err := RunCmd("argocd", "app", "get", "guestbook")
	errors.CheckError(err)
	outputAsString := BytesToString(output)
	pass := OutputContains(outputAsString, expected)
	fmt.Println(pass)
	assert.Equal(t, true, pass)
}

func BytesToString(data []byte) string {
	return string(data[:])
}
