package feed

import (
	"io/ioutil"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/deislabs/porter/pkg/context"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	tc := context.NewTestContext(t)
	tc.AddTestFile("testdata/atom-template.xml", "template.xml")

	tc.FileSystem.Create("bin/v1.2.3/helm-darwin-amd64")
	tc.FileSystem.Create("bin/v1.2.3/helm-linux-amd64")
	tc.FileSystem.Create("bin/v1.2.3/helm-windows-amd64.exe")

	// Force the up3 timestamps to stay the same for each test run
	up3, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
	tc.FileSystem.Chtimes("bin/v1.2.3/helm-darwin-amd64", up3, up3)
	tc.FileSystem.Chtimes("bin/v1.2.3/helm-linux-amd64", up3, up3)
	tc.FileSystem.Chtimes("bin/v1.2.3/helm-windows-amd64.exe", up3, up3)

	tc.FileSystem.Create("bin/v1.2.4/helm-darwin-amd64")
	tc.FileSystem.Create("bin/v1.2.4/helm-linux-amd64")
	tc.FileSystem.Create("bin/v1.2.4/helm-windows-amd64.exe")

	up4, _ := time.Parse("2006-Jan-02", "2013-Feb-04")
	tc.FileSystem.Chtimes("bin/v1.2.4/helm-darwin-amd64", up4, up4)
	tc.FileSystem.Chtimes("bin/v1.2.4/helm-linux-amd64", up4, up4)
	tc.FileSystem.Chtimes("bin/v1.2.4/helm-windows-amd64.exe", up4, up4)

	tc.FileSystem.Create("bin/v1.2.3/exec-darwin-amd64")
	tc.FileSystem.Create("bin/v1.2.3/exec-linux-amd64")
	tc.FileSystem.Create("bin/v1.2.3/exec-windows-amd64.exe")

	up2, _ := time.Parse("2006-Jan-02", "2013-Feb-02")
	tc.FileSystem.Chtimes("bin/v1.2.3/exec-darwin-amd64", up2, up2)
	tc.FileSystem.Chtimes("bin/v1.2.3/exec-linux-amd64", up2, up2)
	tc.FileSystem.Chtimes("bin/v1.2.3/exec-windows-amd64.exe", up2, up2)

	tc.FileSystem.Create("bin/canary/exec-darwin-amd64")
	tc.FileSystem.Create("bin/canary/exec-linux-amd64")
	tc.FileSystem.Create("bin/canary/exec-windows-amd64.exe")

	up10, _ := time.Parse("2006-Jan-02", "2013-Feb-10")
	tc.FileSystem.Chtimes("bin/canary/exec-darwin-amd64", up10, up10)
	tc.FileSystem.Chtimes("bin/canary/exec-linux-amd64", up10, up10)
	tc.FileSystem.Chtimes("bin/canary/exec-windows-amd64.exe", up10, up10)

	opts := GenerateOptions{
		AtomFile:        "atom.xml",
		SearchDirectory: "bin",
		TemplateFile:    "template.xml",
	}
	f := NewMixinFeed(tc.Context)
	err := f.Generate(opts)
	require.NoError(t, err)
	err = f.Save(opts)
	require.NoError(t, err)

	b, err := tc.FileSystem.ReadFile("atom.xml")
	require.NoError(t, err)
	gotXml := string(b)

	b, err = ioutil.ReadFile("testdata/atom.xml")
	require.NoError(t, err)
	wantXml := string(b)

	assert.Equal(t, wantXml, gotXml)
}

func TestGenerate_ExistingFeed(t *testing.T) {
	tc := context.NewTestContext(t)
	tc.AddTestFile("testdata/atom-template.xml", "template.xml")
	tc.AddTestFile("testdata/atom-existing.xml", "atom.xml")

	tc.FileSystem.Create("bin/v1.2.4/helm-darwin-amd64")
	tc.FileSystem.Create("bin/v1.2.4/helm-linux-amd64")
	tc.FileSystem.Create("bin/v1.2.4/helm-windows-amd64.exe")

	up4, _ := time.Parse("2006-Jan-02", "2013-Feb-04")
	tc.FileSystem.Chtimes("bin/v1.2.4/helm-darwin-amd64", up4, up4)
	tc.FileSystem.Chtimes("bin/v1.2.4/helm-linux-amd64", up4, up4)
	tc.FileSystem.Chtimes("bin/v1.2.4/helm-windows-amd64.exe", up4, up4)

	tc.FileSystem.Create("bin/canary/exec-darwin-amd64")
	tc.FileSystem.Create("bin/canary/exec-linux-amd64")
	tc.FileSystem.Create("bin/canary/exec-windows-amd64.exe")

	up10, _ := time.Parse("2006-Jan-02", "2013-Feb-10")
	tc.FileSystem.Chtimes("bin/canary/exec-darwin-amd64", up10, up10)
	tc.FileSystem.Chtimes("bin/canary/exec-linux-amd64", up10, up10)
	tc.FileSystem.Chtimes("bin/canary/exec-windows-amd64.exe", up10, up10)

	opts := GenerateOptions{
		AtomFile:        "atom.xml",
		SearchDirectory: "bin",
		TemplateFile:    "template.xml",
	}
	f := NewMixinFeed(tc.Context)
	err := f.Generate(opts)
	require.NoError(t, err)
	err = f.Save(opts)
	require.NoError(t, err)

	b, err := tc.FileSystem.ReadFile("atom.xml")
	require.NoError(t, err)
	gotXml := string(b)

	b, err = ioutil.ReadFile("testdata/atom.xml")
	require.NoError(t, err)
	wantXml := string(b)

	assert.Equal(t, wantXml, gotXml)
}

func TestMixinEntries_Sort(t *testing.T) {
	up2, _ := time.Parse("2006-Jan-02", "2013-Feb-02")
	up3, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
	up4, _ := time.Parse("2006-Jan-02", "2013-Feb-04")

	entries := MixinEntries{
		{
			Files: []*MixinFile{
				{Updated: up3},
			},
		},
		{
			Files: []*MixinFile{
				{Updated: up2},
			},
		},
		{
			Files: []*MixinFile{
				{Updated: up4},
			},
		},
	}

	sort.Sort(sort.Reverse(entries))

	assert.Equal(t, up4, entries[0].Files[0].Updated)
	assert.Equal(t, up3, entries[1].Files[0].Updated)
	assert.Equal(t, up2, entries[2].Files[0].Updated)
}
