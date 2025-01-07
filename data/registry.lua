---@class ToolDefinition
---@field owner string
---@field repo string
---@field url_template string
---@field version_override? string

---@class glug
---@field register_tool fun(name: string, definition: ToolDefinition)
glug = glug

glug.register_tool("atmos", {
  owner = "cloudposse",
  repo = "atmos",
  url_template = [[
    https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Version}}/{{.Name}}_{{.VersionNumber}}_{{.OS}}_{{.Arch}}
  ]],
})

glug.register_tool("buildx", {
  owner = "docker",
  repo = "buildx",
  url_template = [[
		{{$os := "linux"}}
		{{$arch := .Arch}}
    {{$ext := ""}}
		{{ if HasPrefix .OS "ming" -}}
			{{$os = "windows"}}
  		{{$ext = ".exe"}}
		{{- end -}}
		{{ if eq .OS "darwin" -}}
			{{$os = "darwin"}}
		{{- end -}}
		{{- if eq .Arch "armv6l" -}}
			{{$arch = "arm-v6"}}
		{{- else if eq .Arch "armv7l" -}}
			{{$arch = "arm-v7"}}
    {{- else if or (eq .Arch "aarch64") (eq .Arch "arm64") -}}
			{{$arch = "arm64"}}
		{{- end -}}
		https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Version}}/{{.Name}}-{{.Version}}.{{$os}}-{{$arch}}{{$ext}}
  ]],
})

glug.register_tool("docker-compose", {
  owner = "docker",
  repo = "compose",
  url_template = [[
		{{$os := "linux"}}
		{{$arch := .Arch}}
    {{$ext := ""}}
		{{ if HasPrefix .OS "ming" -}}
			{{$os = "windows"}}
  		{{$ext = ".exe"}}
		{{- end -}}
		{{ if eq .OS "darwin" -}}
			{{$os = "darwin"}}
		{{- end -}}
		{{- if eq .Arch "armv6l" -}}
			{{$arch = "armv6"}}
		{{- else if eq .Arch "armv7l" -}}
			{{$arch = "armv7"}}
    {{- else if or (eq .Arch "aarch64") (eq .Arch "arm64") -}}
			{{$arch = "aarch64"}}
		{{- else if eq .Arch "amd64" -}}
			{{$arch = "x86_64"}}
		{{- end -}}
		https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Version}}/{{.Name}}-{{$os}}-{{$arch}}
  ]],
})

glug.register_tool("fzf", {
  owner = "junegunn",
  repo = "fzf",
  url_template = [[
    {{ $os := "linux" }}
		{{ $arch := "amd64" }}
  	{{ $ext := ".tar.gz" }}
		{{ if HasPrefix .OS "ming" -}}
  		{{ $os = "windows" }}
  		{{ $ext = ".zip" }}
		{{- else if eq .OS "darwin" -}}
  		{{  $os = "darwin" }}
		{{- end -}}
		{{- if eq .Arch "armv6l" -}}
	  	{{ $arch = "armv6" }}
		{{- else if eq .Arch "armv7l" -}}
	  	{{ $arch = "armv7" }}
    {{- else if or (eq .Arch "aarch64") (eq .Arch "arm64") -}}
	  	{{ $arch = "arm64" }}
		{{- end -}}
    https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Version}}/{{.Name}}-{{.VersionNumber}}-{{$os}}_{{$arch}}{{$ext}}
  ]],
})

glug.register_tool("lazydocker", {
  owner = "jesseduffield",
  repo = "lazydocker",
  url_template = [[
		{{$os := "Linux"}}
		{{$arch := .Arch}}
    {{$ext := "tar.gz"}}
		{{ if HasPrefix .OS "ming" -}}
			{{$os = "Windows"}}
  		{{$ext = "zip"}}
		{{- end -}}
		{{ if eq .OS "darwin" -}}
			{{$os = "Darwin"}}
		{{- end -}}
		{{- if eq .Arch "armv6l" -}}
			{{$arch = "arm-v6"}}
		{{- else if eq .Arch "armv7l" -}}
			{{$arch = "arm-v7"}}
    {{- else if or (eq .Arch "aarch64") (eq .Arch "arm64") -}}
			{{$arch = "arm64"}}
		{{- else if eq .Arch "amd64" -}}
			{{$arch = "x86_64"}}
		{{- end -}}
		https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Version}}/{{.Name}}_{{.VersionNumber}}_{{$os}}_{{$arch}}.{{$ext}}
  ]],
})

glug.register_tool("lazygit", {
  owner = "jesseduffield",
  repo = "lazygit",
  url_template = [[
		{{$os := "Linux"}}
		{{$arch := .Arch}}
    {{$ext := "tar.gz"}}
		{{ if HasPrefix .OS "ming" -}}
			{{$os = "Windows"}}
  		{{$ext = "zip"}}
		{{- end -}}
		{{ if eq .OS "darwin" -}}
			{{$os = "Darwin"}}
		{{- end -}}
		{{- if eq .Arch "armv6l" -}}
			{{$arch = "arm-v6"}}
		{{- else if eq .Arch "armv7l" -}}
			{{$arch = "arm-v7"}}
    {{- else if or (eq .Arch "aarch64") (eq .Arch "arm64") -}}
			{{$arch = "arm64"}}
		{{- else if eq .Arch "amd64" -}}
			{{$arch = "x86_64"}}
		{{- end -}}
		https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Version}}/{{.Name}}_{{.VersionNumber}}_{{$os}}_{{$arch}}.{{$ext}}
  ]],
})

-- glug.register_tool("nvim", {
--   owner = "neovim",
--   repo = "neovim",
--   url_template = [[
-- 		{{$os := "linux"}}
-- 		{{$arch := .Arch}}
--     {{$ext := "tar.gz"}}
-- 		{{ if HasPrefix .OS "ming" -}}
-- 			{{$os = "win64"}}
--   		{{$ext = "zip"}}
-- 		{{- end -}}
-- 		{{ if eq .OS "darwin" -}}
-- 			{{$os = "macos"}}
-- 		{{- end -}}
-- 		{{- if eq .Arch "aarch64" -}}
-- 			{{$arch = "arm64"}}
-- 		{{- else if eq .Arch "x86_64" -}}
--       {{$arch = "x86_64"}}
-- 		{{- end -}}
-- 		https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Version}}/{{.Name}}-{{$os}}-{{$arch}}.{{$ext}}
--   ]],
-- })

glug.register_tool("op", {
  owner = "1password",
  repo = "op",
  version_override = "v2.30.3", --- https://app-updates.agilebits.com/product_history/CLI2
  url_template = [[
    {{$os := .OS}}
		{{$arch := .Arch}}
		{{- if eq .Arch "aarch64" -}}
			{{ $arch = "arm64" }}
		{{- else if eq .Arch "x86_64" -}}
			{{ $arch = "amd64" }}
  	{{- else if eq .Arch "armv7l" -}}
			{{ $arch = "arm" }}
		{{- end -}}
		{{ if HasPrefix .OS "ming" -}}
			{{$os = "windows"}}
		{{- end -}}
		https://cache.agilebits.com/dist/1P/op2/pkg/{{.Version}}/op_{{$os}}_{{$arch}}_{{.Version}}.zip
  ]],
})

glug.register_tool("tilt", {
  owner = "tilt-dev",
  repo = "tilt",
  url_template = [[
    {{$os := .OS}}
		{{$arch := .Arch}}
  	{{ $ext := "tar.gz" }}
		{{- if eq .Arch "aarch64" -}}
			{{ $arch = "arm64" }}
		{{- end -}}
		{{ if HasPrefix .OS "ming" -}}
			{{$os = "windows"}}
  	  {{$ext = "zip" }}
		{{- else if eq .OS "darwin" -}}
			{{$os = "mac"}}
		{{- end -}}
		https://github.com/{{.Owner}}/{{.Repo}}/releases/download/{{.Version}}/{{.Name}}.{{.VersionNumber}}.{{$os}}.{{$arch}}.{{$ext}}
  ]],
})
