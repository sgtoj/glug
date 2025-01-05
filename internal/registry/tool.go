package registry

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"runtime"
	"strings"
)

type ToolData struct {
	Name        string
	Owner       string
	Repo        string
	Url         string
	UrlTemplate string
	Version     string
}

type ToolDataCtx struct {
	Name          string
	Owner         string
	Repo          string
	Arch          string
	OS            string
	Version       string
	VersionNumber string
}

var templateFns = template.FuncMap{
	"HasPrefix": strings.HasPrefix,
}

func (d *ToolData) GetUrl() (string, error) {
	if d.Url != "" {
		return d.Url, nil
	}

	toolVersion, err := d.GetVersion()
	if err != nil {
		return "", fmt.Errorf("failed to get version: %w", err)
	}

	toolCtx := ToolDataCtx{
		Name:          d.Name,
		Owner:         d.Owner,
		Repo:          d.Repo,
		Arch:          runtime.GOARCH,
		OS:            runtime.GOOS,
		Version:       toolVersion,
		VersionNumber: strings.TrimPrefix(toolVersion, "v"),
	}

	t, err := template.New(d.Name + "-url").Funcs(templateFns).Parse(d.UrlTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse url template: %w", err)
	}

	var toolUrlBuf bytes.Buffer
	if err := t.Execute(&toolUrlBuf, toolCtx); err != nil {
		return "", fmt.Errorf("failed to generate url: %w", err)
	}

	d.Url = strings.TrimSpace(toolUrlBuf.String())
	return d.Url, nil
}

func (d *ToolData) GetVersion() (string, error) {
	if d.Version != "" {
		return d.Version, nil
	}

	toolHostUrl := fmt.Sprintf("https://github.com/%s/%s/releases/latest", d.Owner, d.Repo)

	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, _ []*http.Request) error {
		return http.ErrUseLastResponse
	}

	req, err := http.NewRequest(http.MethodHead, toolHostUrl, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for latest version: %w", err)
	}
	req.Header.Set("User-Agent", "glug/0.0.0") // todo: get version from config

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed web request to get latest version: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusMovedPermanently && res.StatusCode != http.StatusFound {
		return "", fmt.Errorf("unexpected status returned requesting latest version: %d", res.StatusCode)
	}

	loc := res.Header.Get("Location")
	if len(loc) == 0 {
		return "", fmt.Errorf("unable to determine latest release")
	}

	d.Version = loc[strings.LastIndex(loc, "/")+1:]
	return d.Version, nil
}
