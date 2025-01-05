package registry

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

type ToolRegistry map[string]ToolData

func BuildToolRegistry(configFile string) (ToolRegistry, error) {
	L := lua.NewState()
	defer L.Close()

	tools := make(map[string]ToolData)

	luaGlugNamespace := L.NewTable()
	L.SetField(luaGlugNamespace, "register_tool", L.NewFunction(func(L *lua.LState) int {
		toolName := L.CheckString(1)
		toolInfo := L.CheckTable(2)

		version := lua.LVAsString(toolInfo.RawGetString("version_override"))

		tools[toolName] = ToolData{
			Name:        toolName,
			Owner:       toolInfo.RawGetString("owner").String(),
			Repo:        toolInfo.RawGetString("repo").String(),
			UrlTemplate: toolInfo.RawGetString("url_template").String(),
			Version:     version,
		}

		return 0
	}))
	L.SetGlobal("glug", luaGlugNamespace)

	if err := L.DoFile(configFile); err != nil {
		return nil, fmt.Errorf("error processing config file: %w", err)
	}

	return tools, nil
}
