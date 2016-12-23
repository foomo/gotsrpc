package gotsrpc

import (
	"fmt"
	"strings"

	"github.com/foomo/gotsrpc/config"
)

func (v *Value) isHTTPResponseWriter() bool {
	return v.StructType != nil && v.StructType.Name == "ResponseWriter" && v.StructType.Package == "net/http"
}

func (v *Value) isHTTPRequest() bool {
	return v.IsPtr && v.StructType != nil && v.StructType.Name == "Request" && v.StructType.Package == "net/http"
}

func (v *Value) goType(aliases map[string]string, packageName string) (t string) {
	if v.IsPtr {
		t = "*"
	}
	switch true {
	case v.Array != nil:
		t += "[]" + v.Array.Value.goType(aliases, packageName)
	case len(v.GoScalarType) > 0:
		t += v.GoScalarType
	case v.StructType != nil:
		if packageName != v.StructType.Package {
			t += aliases[v.StructType.Package] + "."
		}
		t += v.StructType.Name
	case v.Map != nil:
		t += `map[` + v.Map.KeyType + `]` + v.Map.Value.goType(aliases, packageName)
	case v.Scalar != nil:
		// TODO this is a hack to retrieve string types
		if packageName != v.Scalar.Package {
			t += aliases[v.Scalar.Package] + "."
		}
		t += v.Scalar.Name[len(v.Scalar.Package)+1:]
	default:
		// TODO
		fmt.Println("WARN: can't resolve goType")
	}

	return
}

func (v *Value) emptyLiteral(aliases map[string]string) (e string) {
	e = ""
	if v.IsPtr {
		e += "&"
	}
	switch true {
	case len(v.GoScalarType) > 0:
		switch v.GoScalarType {
		case "string":
			e += "\"\""
		case "float":
			return "float(0.0)"
		case "float32":
			return "float32(0.0)"
		case "float64":
			return "float64(0.0)"
		case "int":
			return "int(0)"
		case "int8":
			return "int8(0)"
		case "int16":
			return "int16(0)"
		case "int32":
			return "int32(0)"
		case "int64":
			return "int64(0)"
		case "uint":
			return "uint(0)"
		case "uint8":
			return "uint8(0)"
		case "uint16":
			return "uint16(0)"
		case "uint32":
			return "uint32(0)"
		case "uint64":
			return "uint64(0)"
		case "bool":
			return "false"
		}
	case v.Array != nil:
		e += "[]"
		if v.Array.Value.IsPtr {
			e += "*"
		}
		l := v.Array.Value.emptyLiteral(aliases)
		if len(v.Array.Value.GoScalarType) == 0 {
			if v.Array.Value.IsPtr {
				l = strings.TrimPrefix(l, "&")
			}
			l = strings.TrimSuffix(l, "{}")
		} else {
			l = v.Array.Value.GoScalarType
		}
		e += l + "{}"
	case v.StructType != nil:
		alias := aliases[v.StructType.Package]
		if alias != "" {
			e += alias + "."
		}
		e += v.StructType.Name + "{}"

	}
	return
}

func lcfirst(str string) string {
	return strfirst(str, strings.ToLower)
}

func ucfirst(str string) string {
	return strfirst(str, strings.ToUpper)
}

func strfirst(str string, strfunc func(string) string) string {
	res := ""
	for i, char := range str {
		if i == 0 {
			res += strfunc(string(char))
		} else {
			res += string(char)
		}

	}
	return res
}

func extractImports(fields []*Field, fullPackageName string, aliases map[string]string) {

	r := strings.NewReplacer(".", "_", "/", "_", "-", "_")

	extractImport := func(st *StructType) {
		if st.Package != fullPackageName {
			alias, ok := aliases[st.Package]
			if !ok {
				packageParts := strings.Split(st.Package, "/")
				beautifulAlias := packageParts[len(packageParts)-1]
				uglyAlias := r.Replace(st.Package)
				alias = beautifulAlias
				for _, otherAlias := range aliases {
					if otherAlias == beautifulAlias {
						alias = uglyAlias
						break
					}
				}
				aliases[st.Package] = alias
			}
		}
	}

	for _, f := range fields {
		if f.Value.StructType != nil {
			extractImport(f.Value.StructType)
		} else if f.Value.Array != nil && f.Value.Array.Value.StructType != nil {
			extractImport(f.Value.Array.Value.StructType)
		}
	}
}

func renderTSRPCServiceProxies(services map[string]*Service, fullPackageName string, packageName string, config *config.Target, g *code) error {
	aliases := map[string]string{
		"time":                     "time",
		"net/http":                 "http",
		"github.com/foomo/gotsrpc": "gotsrpc",
	}

	for _, service := range services {
		// Check if we should render this service as ts rcp
		// Note: remove once there's a separate gorcp generator
		if !config.IsTSRPC(service.Name) {
			continue
		}
		for _, m := range service.Methods {
			extractImports(m.Args, fullPackageName, aliases)
		}
	}

	imports := ""
	for packageName, alias := range aliases {
		imports += alias + " \"" + packageName + "\"\n"
	}

	g.l(`
        // this file was auto generated by gotsrpc https://github.com/foomo/gotsrpc
        package ` + packageName + `
        import (
			` + imports + `
        )
    `)
	for endpoint, service := range services {
		// Check if we should render this service as ts rcp
		// Note: remove once there's a separate gorcp generator
		if !config.IsTSRPC(service.Name) {
			continue
		}

		proxyName := service.Name + "GoTSRPCProxy"
		g.l(`
        type ` + proxyName + ` struct {
	        EndPoint string
			allowOrigin []string
	        service  *` + service.Name + `
        }

        func NewDefault` + proxyName + `(service *` + service.Name + `, allowOrigin []string) *` + proxyName + ` {
	        return &` + proxyName + `{
		        EndPoint: "` + endpoint + `",
				allowOrigin : allowOrigin,
		        service:  service,
	        }
        }


        func New` + proxyName + `(service *` + service.Name + `, endpoint string, allowOrigin []string) *` + proxyName + ` {
	        return &` + proxyName + `{
		        EndPoint: endpoint,
				allowOrigin : allowOrigin,
		        service:  service,
	        }
        }

        // ServeHTTP exposes your service
        func (p *` + proxyName + `) ServeHTTP(w http.ResponseWriter, r *http.Request) {

			for _, origin := range p.allowOrigin {
				// todo we have to compare this with the referer ... and only send one
				w.Header().Add("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Credentials", "true")
	        if r.Method != http.MethodPost {
				if r.Method == http.MethodOptions {
					return
				}
		        gotsrpc.ErrorMethodNotAllowed(w)
		        return
	        }
		`)
		needsArgs := false
		for _, method := range service.Methods {
			if len(method.Args) > 0 {
				needsArgs = true
				break
			}
		}
		if needsArgs {
			g.l(`var args []interface{}`)
		}

		g.l("funcName := gotsrpc.GetCalledFunc(r, p.EndPoint)")
		g.l("callStats := gotsrpc.GetStatsForRequest(r)")
		g.l("if callStats != nil {").ind(1)
		g.l("callStats.Func = funcName")
		g.l("callStats.Package = \"" + fullPackageName + "\"")
		g.l("callStats.Service = \"" + service.Name + "\"")
		g.ind(-1).l("}")

		g.l(`switch funcName {`)

		// indenting into switch cases
		g.ind(4)

		for _, method := range service.Methods {
			// a case for each method
			g.l("case \"" + method.Name + "\":")
			g.ind(1)
			callArgs := []string{}
			isSessionRequest := false
			if len(method.Args) > 0 {

				args := []string{}

				skipArgI := 0

				for argI, arg := range method.Args {

					if argI == 0 && arg.Value.isHTTPResponseWriter() {
						continue
					}
					if argI == 1 && arg.Value.isHTTPRequest() {
						isSessionRequest = true
						continue
					}

					args = append(args, arg.Value.emptyLiteral(aliases))
					switch arg.Value.GoScalarType {
					case "int", "int8", "int16", "int32", "int64",
						"uint", "uint8", "uint16", "uint32", "uint64":
						callArgs = append(callArgs, fmt.Sprint(arg.Value.GoScalarType+"(args[", skipArgI, "].(float64))"))
					default:
						// assert
						callArgs = append(callArgs, fmt.Sprint("args[", skipArgI, "].("+arg.Value.goType(aliases, fullPackageName)+")"))
					}
					skipArgI++
				}
				g.l("args = []interface{}{" + strings.Join(args, ", ") + "}")
				g.l("err := gotsrpc.LoadArgs(args, callStats, r)")
				g.l("if err != nil {")
				g.ind(1)
				g.l("gotsrpc.ErrorCouldNotLoadArgs(w)")
				g.l("return")
				g.ind(-1)
				g.l("}")

			}
			returnValueNames := []string{}
			for retI, retField := range method.Return {
				retArgName := retField.Name
				if len(retArgName) == 0 {
					retArgName = "ret"
					if retI > 0 {
						retArgName += "_" + fmt.Sprint(retI)
					}
				}
				returnValueNames = append(returnValueNames, lcfirst(method.Name)+ucfirst(retArgName))
			}
			g.l("executionStart := time.Now()")
			if len(returnValueNames) > 0 {
				g.app(strings.Join(returnValueNames, ", ") + " := ")
			}
			if isSessionRequest {
				callArgs = append([]string{"w", "r"}, callArgs...)
			}
			g.app("p.service." + method.Name + "(" + strings.Join(callArgs, ", ") + ")")
			g.nl()

			g.l("if callStats != nil {")
			g.ind(1).l("callStats.Execution = time.Now().Sub(executionStart)").ind(-1)
			g.l("}")
			g.l("gotsrpc.Reply([]interface{}{" + strings.Join(returnValueNames, ", ") + "}, callStats, r, w)")
			g.l("return")
			g.ind(-1)
		}
		g.l("default:")
		g.ind(1).l("http.Error(w, \"404 - not found \" + r.URL.Path, http.StatusNotFound)")
		g.ind(-2).l("}") // close switch
		g.ind(-1).l("}") // close ServeHttp
	}
	return nil
}

func renderTSRPCServiceClients(services map[string]*Service, fullPackageName string, packageName string, config *config.Target, g *code) error {
	aliases := map[string]string{
		"github.com/foomo/gotsrpc": "gotsrpc",
	}

	for _, service := range services {
		// Check if we should render this service as ts rcp
		// Note: remove once there's a separate gorcp generator
		if !config.IsTSRPC(service.Name) {
			continue
		}
		for _, m := range service.Methods {
			extractImports(m.Args, fullPackageName, aliases)
			extractImports(m.Return, fullPackageName, aliases)
		}
	}

	imports := ""
	for packageName, alias := range aliases {
		imports += alias + " \"" + packageName + "\"\n"
	}

	g.l(`
        // this file was auto generated by gotsrpc https://github.com/foomo/gotsrpc
        package ` + packageName + `
        import (
			` + imports + `
        )
    `)
	for endpoint, service := range services {
		// Check if we should render this service as ts rcp
		// Note: remove once there's a separate gorcp generator
		if !config.IsTSRPC(service.Name) {
			continue
		}

		clientName := service.Name + "GoTSRPCClient"
		g.l(`
        type ` + clientName + ` struct {
					URL string
	        EndPoint string
        }

        func NewDefault` + clientName + `(url string) *` + clientName + ` {
	        return New` + clientName + `(url, "` + endpoint + `")
        }

        func New` + clientName + `(url string, endpoint string) *` + clientName + ` {
	        return &` + clientName + `{
		        URL: url,
		        EndPoint: endpoint,
	        }
        }
		`)
		for _, method := range service.Methods {
			args := []string{}
			params := []string{}
			for _, a := range method.Args {
				args = append(args, a.Name)
				params = append(params, a.Name+" "+a.Value.goType(aliases, fullPackageName))
			}
			rets := []string{}
			returns := []string{}
			for i, r := range method.Return {
				name := r.Name
				if len(name) == 0 {
					name = fmt.Sprintf("ret%s_%d", method.Name, i)
				}
				rets = append(rets, "&"+name)
				returns = append(returns, name+" "+r.Value.goType(aliases, fullPackageName))
			}
			returns = append(returns, "clientErr error")
			g.l(`func (c *` + clientName + `) ` + method.Name + `(` + strings.Join(params, ", ") + `) (` + strings.Join(returns, ", ") + `) {`)
			g.l(`args := []interface{}{` + strings.Join(args, ", ") + `}`)
			g.l(`reply := []interface{}{` + strings.Join(rets, ", ") + `}`)
			g.l(`clientErr = gotsrpc.CallClient(c.URL, c.EndPoint, "` + method.Name + `", args, reply)`)
			g.l(`return`)
			g.l(`}`)
			g.nl()
		}
	}
	return nil
}

func renderGoRPCServiceProxies(services map[string]*Service, fullPackageName string, packageName string, config *config.Target, g *code) error {
	aliases := map[string]string{
		"fmt":                      "fmt",
		"time":                     "time",
		"strings":                  "strings",
		"reflect":                  "reflect",
		"crypto/tls":               "tls",
		"encoding/gob":             "gob",
		"github.com/valyala/gorpc": "gorpc",
		"github.com/foomo/gotsrpc": "gotsrpc",
	}

	for _, service := range services {
		if !config.IsGoRPC(service.Name) {
			continue
		}

		for _, m := range service.Methods {
			extractImports(m.Args, fullPackageName, aliases)
			extractImports(m.Return, fullPackageName, aliases)
		}
	}

	imports := ""
	for packageName, alias := range aliases {
		imports += alias + " \"" + packageName + "\"\n"
	}

	g.l(`
			// this file was auto generated by gotsrpc https://github.com/foomo/gotsrpc
			package ` + packageName + `
			import (
		` + imports + `
			)
	`)

	for _, service := range services {
		if !config.IsGoRPC(service.Name) {
			continue
		}

		proxyName := service.Name + "GoRPCProxy"
		// Types
		g.l(`type (`)
		// Proxy type
		g.l(`
        ` + proxyName + ` struct {
        	server *gorpc.Server
	        service  *` + service.Name + `
	        callStatsHandler gotsrpc.GoRPCCallStatsHandlerFun
        }
		`)
		// Request & Response types
		for _, method := range service.Methods {
			// Request type
			g.l(ucfirst(service.Name+method.Name) + `Request struct {`)
			for _, a := range method.Args {
				g.l(ucfirst(a.Name) + ` ` + a.Value.goType(aliases, fullPackageName))
			}
			g.l(`}`)
			// Response type
			g.l(ucfirst(service.Name+method.Name) + `Response struct {`)
			for i, r := range method.Return {
				name := r.Name
				if len(name) == 0 {
					name = fmt.Sprintf("ret%s_%d", method.Name, i)
				}
				g.l(ucfirst(name) + ` ` + r.Value.goType(aliases, fullPackageName))
			}
			g.l(`}`)
			g.nl()
		}
		g.l(`)`)
		g.nl()
		// Init
		g.l(`func init() {`)
		for _, method := range service.Methods {
			g.l(`gob.Register(` + ucfirst(service.Name+method.Name) + `Request{})`)
			g.l(`gob.Register(` + ucfirst(service.Name+method.Name) + `Response{})`)
		}
		g.l(`}`)
		// Constructor
		g.l(`
        func New` + proxyName + `(addr string, service *` + service.Name + `, tlsConfig *tls.Config) *` + proxyName + ` {
        	proxy :=  &` + proxyName + `{
		        service:  service,
	        }

        	if tlsConfig != nil {
        		proxy.server = gorpc.NewTLSServer(addr, proxy.handler, tlsConfig)
        	} else {
        		proxy.server = gorpc.NewTCPServer(addr, proxy.handler)
        	}

	        return proxy
        }

        func (p *` + proxyName + `) Start() error {
        	return p.server.Start()
        }

        func (p *` + proxyName + `) Stop() {
        	p.server.Stop()
        }

        func (p *` + proxyName + `) SetCallStatsHandler(handler gotsrpc.GoRPCCallStatsHandlerFun) {
					p.callStatsHandler = handler
				}
		`)
		g.nl()
		// Handler
		g.l(`func (p *` + proxyName + `) handler(clientAddr string, request interface{}) (response interface{}) {`)
		g.l(`start := time.Now()`)
		g.nl()
		g.l(`reqType := reflect.TypeOf(request).String()`)
		g.l(`funcNameParts := strings.Split(reqType, ".")`)
		g.l(`funcName := funcNameParts[len(funcNameParts)-1]`)
		g.nl()
		g.l(`switch funcName {`)
		for _, method := range service.Methods {
			argParams := []string{}
			for _, a := range method.Args {
				argParams = append(argParams, "req."+ucfirst(a.Name))
			}
			rets := []string{}
			retParams := []string{}
			for i, r := range method.Return {
				name := r.Name
				if len(name) == 0 {
					name = fmt.Sprintf("ret%s_%d", method.Name, i)
				}
				rets = append(rets, name)
				retParams = append(retParams, ucfirst(name)+`: `+name)
			}
			g.l(`case "` + service.Name + method.Name + `Request":`)
			if len(argParams) > 0 {
				g.l(`req := request.(` + service.Name + method.Name + `Request)`)
			}
			if len(rets) > 0 {
				g.l(strings.Join(rets, ", ") + ` := p.service.` + method.Name + `(` + strings.Join(argParams, ", ") + `)`)
			} else {
				g.l(`p.service.` + method.Name + `(` + strings.Join(argParams, ", ") + `)`)
			}
			g.l(`response = ` + service.Name + method.Name + `Response{` + strings.Join(retParams, ", ") + `}`)
		}
		g.l(`default:`)
		g.l(`fmt.Println("Unkown request type", reflect.TypeOf(request).String())`)
		g.l(`}`)
		g.nl()
		g.l(`if p.callStatsHandler != nil {`)
		g.l(`p.callStatsHandler(&gotsrpc.CallStats{`)
		g.l(`Func: funcName,`)
		g.l(`Package: "` + fullPackageName + `",`)
		g.l(`Service: "` + service.Name + `",`)
		g.l(`Execution: time.Since(start),`)
		g.l(`})`)
		g.l(`}`)
		g.nl()
		g.l(`return`)
		g.l(`}`)
	}
	return nil
}

func renderGoRPCServiceClients(services map[string]*Service, fullPackageName string, packageName string, config *config.Target, g *code) error {
	aliases := map[string]string{
		"crypto/tls":               "tls",
		"github.com/valyala/gorpc": "gorpc",
	}

	for _, service := range services {
		if !config.IsGoRPC(service.Name) {
			continue
		}
		for _, m := range service.Methods {
			extractImports(m.Args, fullPackageName, aliases)
			extractImports(m.Return, fullPackageName, aliases)
		}
	}

	imports := ""
	for packageName, alias := range aliases {
		imports += alias + " \"" + packageName + "\"\n"
	}

	g.l(`
			// this file was auto generated by gotsrpc https://github.com/foomo/gotsrpc
			package ` + packageName + `
			import (
		` + imports + `
			)
	`)
	for _, service := range services {
		if !config.IsGoRPC(service.Name) {
			continue
		}
		clientName := service.Name + "GoRPCClient"
		// Client type
		g.l(`
        type ` + clientName + ` struct {
        	client *gorpc.Client
        }
		`)
		// Constructor
		g.l(`
        func New` + clientName + `(addr string, tlsConfig *tls.Config) *` + clientName + ` {
        	client := &` + clientName + `{}
        	if tlsConfig == nil {
						client.client = gorpc.NewTCPClient(addr)
					} else {
						client.client = gorpc.NewTLSClient(addr, tlsConfig)
					}
					client.Start()
					return client
        }

        func (c *` + clientName + `) Start() {
        	c.client.Start()
      	}

        func (c *` + clientName + `) Stop() {
        	c.client.Stop()
      	}
		`)
		g.nl()
		// Methods
		for _, method := range service.Methods {
			args := []string{}
			params := []string{}
			for _, a := range method.Args {
				args = append(args, ucfirst(a.Name)+`: `+a.Name)
				params = append(params, a.Name+" "+a.Value.goType(aliases, fullPackageName))
			}
			rets := []string{}
			returns := []string{}
			for i, r := range method.Return {
				name := r.Name
				if len(name) == 0 {
					name = fmt.Sprintf("ret%s_%d", method.Name, i)
				}
				rets = append(rets, "response."+ucfirst(name))
				returns = append(returns, name+" "+r.Value.goType(aliases, fullPackageName))
			}
			returns = append(returns, "clientErr error")
			g.l(`func (c *` + clientName + `) ` + method.Name + `(` + strings.Join(params, ", ") + `) (` + strings.Join(returns, ", ") + `) {`)
			g.l(`req := ` + service.Name + method.Name + `Request{` + strings.Join(args, ", ") + `}`)
			if len(rets) > 0 {
				g.l(`rpcCallRes, rpcCallErr := c.client.Call(req)`)
			} else {
				g.l(`_, rpcCallErr := c.client.Call(req)`)
			}
			g.l(`if rpcCallErr != nil {`)
			g.l(`clientErr = rpcCallErr`)
			g.l(`return`)
			g.l(`}`)
			if len(rets) > 0 {
				g.l(`response := rpcCallRes.(` + service.Name + method.Name + `Response)`)
				g.l(`return ` + strings.Join(rets, ", ") + `, nil`)
			} else {
				g.l(`return nil`)
			}
			g.l(`}`)
			g.nl()
		}
	}
	return nil
}

func RenderGoTSRPCProxies(services map[string]*Service, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := newCode("	")
	err = renderTSRPCServiceProxies(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.string()
	return
}

func RenderGoTSRPCClients(services map[string]*Service, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := newCode("	")
	err = renderTSRPCServiceClients(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.string()
	return
}

func RenderGoRPCProxies(services map[string]*Service, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := newCode("	")
	err = renderGoRPCServiceProxies(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.string()
	return
}

func RenderGoRPCClients(services map[string]*Service, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := newCode("	")
	err = renderGoRPCServiceClients(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.string()
	return
}
