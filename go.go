package gotsrpc

import (
	"fmt"
	"strings"

	"github.com/foomo/gotsrpc/v2/config"
)

func (v *Value) isHTTPResponseWriter() bool {
	return (v.StructType != nil && v.StructType.Name == "ResponseWriter" && v.StructType.Package == "net/http") ||
		(v.Scalar != nil && v.Scalar.Name == "ResponseWriter" && v.Scalar.Package == "net/http")
}

func (v *Value) isHTTPRequest() bool {
	return (v.IsPtr && v.StructType != nil && v.StructType.Name == "Request" && v.StructType.Package == "net/http") ||
		(v.IsPtr && v.Scalar != nil && v.Scalar.Name == "Request" && v.Scalar.Package == "net/http")
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
		if packageName != v.StructType.Package && aliases[v.StructType.Package] != "" {
			t += aliases[v.StructType.Package] + "."
		}
		t += v.StructType.Name
	case v.Map != nil:
		t += `map[` + v.Map.KeyGoType + `]` + v.Map.Value.goType(aliases, packageName)
	case v.Scalar != nil:
		if packageName != v.Scalar.Package && aliases[v.Scalar.Package] != "" {
			t += aliases[v.Scalar.Package] + "."
		}
		t += v.Scalar.Name
	case v.IsInterface:
		t += "interface{}"
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
	case v.Map != nil:
		e += "map[" + v.Map.KeyGoType + "]" + v.Map.Value.emptyLiteral(aliases)
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
	case v.Scalar != nil:
		alias := aliases[v.Scalar.Package]
		if alias != "" {
			e += alias + "."
		}
		e += v.Scalar.Name + "{}"
	case v.IsInterface:
		e += "interface{}{}"
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

func extractImport(packageName string, fullPackageName string, aliases map[string]string) {
	r := strings.NewReplacer(".", "_", "/", "_", "-", "_")
	if packageName != fullPackageName {
		alias, ok := aliases[packageName]
		if !ok {
			packageParts := strings.Split(packageName, "/")
			beautifulAlias := packageParts[len(packageParts)-1]
			uglyAlias := r.Replace(packageName)
			alias = uglyAlias //beautifulAlias
			for _, otherAlias := range aliases {
				if otherAlias == beautifulAlias {
					alias = uglyAlias
					break
				}
			}
			aliases[packageName] = alias
		}
	}
}

func extractImports(fields []*Field, fullPackageName string, aliases map[string]string) {
	for _, f := range fields {
		if f.Value.StructType != nil {
			extractImport(f.Value.StructType.Package, fullPackageName, aliases)
		} else if f.Value.Array != nil && f.Value.Array.Value.StructType != nil {
			extractImport(f.Value.Array.Value.StructType.Package, fullPackageName, aliases)
		} else if f.Value.Map != nil && f.Value.Map.Value.StructType != nil {
			extractImport(f.Value.Map.Value.StructType.Package, fullPackageName, aliases)
		} else if f.Value.Map != nil && f.Value.Map.Key.StructType != nil {
			extractImport(f.Value.Map.Key.StructType.Package, fullPackageName, aliases)
		} else if f.Value.Array != nil && f.Value.Array.Value.Scalar != nil {
			extractImport(f.Value.Array.Value.Scalar.Package, fullPackageName, aliases)
		} else if f.Value.Map != nil && f.Value.Map.Value.Scalar != nil {
			extractImport(f.Value.Map.Value.Scalar.Package, fullPackageName, aliases)
		} else if f.Value.Map != nil && f.Value.Map.Key.Scalar != nil {
			extractImport(f.Value.Map.Key.Scalar.Package, fullPackageName, aliases)
		} else if f.Value.Scalar != nil {
			extractImport(f.Value.Scalar.Package, fullPackageName, aliases)
		}
	}
}

func renderTSRPCServiceProxies(services ServiceList, fullPackageName string, packageName string, config *config.Target, unions map[string][]string, g *code) error {
	aliases := map[string]string{
		"time":                        "time",
		"net/http":                    "http",
		"io":                          "io",
		"io/ioutil":                   "ioutil",
		"github.com/foomo/gotsrpc/v2": "gotsrpc",
	}
	for _, service := range services {
		// Check if we should render this service as ts rpc
		// Note: remove once there's a separate gorcp generator
		if !config.IsTSRPC(service.Name) {
			continue
		}
		for _, m := range service.Methods {
			extractImports(m.Args, fullPackageName, aliases)
		}
	}

	for pkg := range unions {
		extractImport(pkg, fullPackageName, aliases)
	}

	g.l(renderImports(aliases, packageName))

	renderInit(unions, aliases, packageName, g)

	for _, service := range services {
		// Check if we should render this service as ts rcp
		// Note: remove once there's a separate gorcp generator
		if !config.IsTSRPC(service.Name) {
			continue
		}

		servicePointer := "*"
		if service.IsInterface {
			servicePointer = ""
		}

		proxyName := service.Name + "GoTSRPCProxy"

		g.l("const (")
		for _, method := range service.Methods {
			g.l(proxyName + method.Name + " = \"" + method.Name + "\"")
		}
		g.l(")")

		g.l(`
        type ` + proxyName + ` struct {
	        EndPoint string
	        service  ` + servicePointer + service.Name + `
        }

        func NewDefault` + proxyName + `(service ` + servicePointer + service.Name + `) *` + proxyName + ` {
	        return New` + proxyName + `(service, "` + service.Endpoint + `")
        }

        func New` + proxyName + `(service ` + servicePointer + service.Name + `, endpoint string) *` + proxyName + ` {
	        return &` + proxyName + `{
		        EndPoint: endpoint,
		        service:  service,
	        }
        }

        // ServeHTTP exposes your service
        func (p *` + proxyName + `) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	        if r.Method != http.MethodPost {
				if r.Method == http.MethodOptions {
					return
				}
		        gotsrpc.ErrorMethodNotAllowed(w)
		        return
	        }
			defer io.Copy(ioutil.Discard, r.Body) // Drain Request Body 
		`)

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
			g.l("case " + proxyName + method.Name + ":")
			g.ind(1)
			callArgs := []string{}
			isSessionRequest := false
			if len(method.Args) > 0 {
				args := []string{}
				argsDecls := []string{}

				skipArgI := 0

				nonHTTPRelatedArgs := goMethodArgsWithoutHTTPContextRelatedArgs(method)

				isSessionRequest = len(method.Args)-len(nonHTTPRelatedArgs) == 2

				for _, arg := range nonHTTPRelatedArgs {
					argName := "arg_" + arg.Name //strconv.Itoa(argI)

					//argsDecls = append(argsDecls, argName+" := "+arg.Value.emptyLiteral(aliases))
					argsDecls = append(argsDecls, argName+"  "+arg.Value.goType(aliases, packageName))
					args = append(args, "&"+argName)
					callArgs = append(callArgs, argName)
					skipArgI++
				}
				if len(args) > 0 {
					g.l("var (")
					for _, argDecl := range argsDecls {
						g.l(argDecl)
					}
					g.l(")")
					g.l("args := []interface{}{" + strings.Join(args, ", ") + "}")
					g.l("err := gotsrpc.LoadArgs(&args, callStats, r)")
					g.l("if err != nil {")
					g.ind(1)
					g.l("gotsrpc.ErrorCouldNotLoadArgs(w)")
					g.l("return")
					g.ind(-1)
					g.l("}")
				}
			}
			var returnValueNames []string
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
			if isSessionRequest {
				g.l("rw := gotsrpc.ResponseWriter{ResponseWriter: w}")
				callArgs = append([]string{"&rw", "r"}, callArgs...)
			}
			if len(returnValueNames) > 0 {
				g.app(strings.Join(returnValueNames, ", ") + " := ")
			}
			g.app("p.service." + method.Name + "(" + strings.Join(callArgs, ", ") + ")")
			g.nl()
			g.l("if callStats != nil {")
			g.ind(1).l("callStats.Execution = time.Now().Sub(executionStart)").ind(-1)
			g.l("}")
			if isSessionRequest {
				g.l("if rw.Status() == http.StatusOK {").ind(1)
			}
			g.l("gotsrpc.Reply([]interface{}{" + strings.Join(returnValueNames, ", ") + "}, callStats, r, w)")
			if isSessionRequest {
				g.ind(-1).l("}")
			}
			g.l("return")
			g.ind(-1)
		}
		g.l("default:")
		g.ind(1).l("gotsrpc.ClearStats(r)")
		g.ind(1).l("http.Error(w, \"404 - not found \" + r.URL.Path, http.StatusNotFound)")
		g.ind(-2).l("}") // close switch
		g.ind(-1).l("}") // close ServeHttp
	}
	return nil
}

type goMethod struct {
	name    string
	params  []string
	args    []string
	rets    []string
	returns []string
}

func newMethodSignature(method *Method, aliases map[string]string, fullPackageName string) goMethod {
	var args []string
	var params []string
	params = append(params, "ctx go_context.Context")
	for _, a := range goMethodArgsWithoutHTTPContextRelatedArgs(method) {
		args = append(args, a.Name)
		params = append(params, a.Name+" "+a.Value.goType(aliases, fullPackageName))
	}
	var rets []string
	var returns []string
	for i, r := range method.Return {
		name := r.Name
		if len(name) == 0 {
			name = fmt.Sprintf("ret%s_%d", method.Name, i)
		}
		rets = append(rets, "&"+name)
		returns = append(returns, name+" "+r.Value.goType(aliases, fullPackageName))
	}
	returns = append(returns, "clientErr error")

	return goMethod{
		name:    method.Name,
		params:  params,
		args:    args,
		rets:    rets,
		returns: returns,
	}
}

func (ms *goMethod) renderSignature() string {
	return ms.name + `(` + strings.Join(ms.params, ", ") + `) (` + strings.Join(ms.returns, ", ") + `)`
}

func renderTSRPCServiceClients(services ServiceList, fullPackageName string, packageName string, config *config.Target, g *code) error {
	aliases := map[string]string{
		"github.com/pkg/errors":       "pkg_errors",
		"github.com/foomo/gotsrpc/v2": "gotsrpc",
		"net/http":                    "go_net_http",
		"context":                     "go_context",
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

	g.l(renderImports(aliases, packageName))

	for _, service := range services {
		// Check if we should render this service as ts rcp
		// Note: remove once there's a separate gorcp generator
		if !config.IsTSRPC(service.Name) {
			continue
		}

		interfaceName := service.Name + "GoTSRPCClient"
		clientName := "HTTP" + interfaceName

		//Render Interface
		g.l(`type ` + interfaceName + ` interface { `)
		for _, method := range service.Methods {
			ms := newMethodSignature(method, aliases, fullPackageName)
			g.l(ms.renderSignature())
		}

		g.l(`} `)

		//Render Constructors
		g.l(`
        type ` + clientName + ` struct {
					URL string
	        EndPoint string
			Client gotsrpc.Client
        }

        func NewDefault` + interfaceName + `(url string) *` + clientName + ` {
	        return New` + interfaceName + `(url, "` + service.Endpoint + `")
        }

        func New` + interfaceName + `(url string, endpoint string) *` + clientName + ` {
			return New` + interfaceName + `WithClient(url, endpoint, nil) 
        }

        func New` + interfaceName + `WithClient(url string, endpoint string, client *go_net_http.Client) *` + clientName + ` {
	        return &` + clientName + `{
		        URL: url,
		        EndPoint: endpoint,
		        Client: gotsrpc.NewClientWithHttpClient(client),
	        }
		}`)

		for _, method := range service.Methods {
			ms := newMethodSignature(method, aliases, fullPackageName)
			g.l(`func (tsc *` + clientName + `) ` + ms.renderSignature() + ` {`)
			g.l(`args := []interface{}{` + strings.Join(ms.args, ", ") + `}`)
			g.l(`reply := []interface{}{` + strings.Join(ms.rets, ", ") + `}`)
			g.l(`clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "` + method.Name + `", args, reply)`)
			g.l(`if clientErr != nil {`)
			g.ind(1).l(`clientErr = pkg_errors.WithMessage(clientErr, "failed to call ` + packageName + `.` + service.Name + `GoTSRPCProxy ` + method.Name + `")`).ind(-1)
			g.l(`}`)
			g.l(`return`)
			g.l(`}`)
			g.nl()
		}
	}
	return nil
}

func renderGoRPCServiceProxies(services ServiceList, fullPackageName string, packageName string, config *config.Target, g *code) error {
	aliases := map[string]string{
		"fmt":                         "fmt",
		"time":                        "time",
		"strings":                     "strings",
		"reflect":                     "reflect",
		"crypto/tls":                  "tls",
		"encoding/gob":                "gob",
		"github.com/valyala/gorpc":    "gorpc",
		"github.com/foomo/gotsrpc/v2": "gotsrpc",
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

	g.l(renderImports(aliases, packageName))

	for _, service := range services {
		if !config.IsGoRPC(service.Name) {
			continue
		}

		servicePointer := "*"
		if service.IsInterface {
			servicePointer = ""
		}

		proxyName := service.Name + "GoRPCProxy"
		// Types
		g.l(`type (`)
		// Proxy type
		g.l(`
        ` + proxyName + ` struct {
        	server *gorpc.Server
	        service  ` + servicePointer + service.Name + `
	        callStatsHandler gotsrpc.GoRPCCallStatsHandlerFun
        }
		`)
		// Request & Response types
		for _, method := range service.Methods {
			// Request type
			g.l(ucfirst(service.Name+method.Name) + `Request struct {`)
			for _, a := range goMethodArgsWithoutHTTPContextRelatedArgs(method) {
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
        func New` + proxyName + `(addr string, service ` + servicePointer + service.Name + `, tlsConfig *tls.Config) *` + proxyName + ` {
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

		func (p *` + proxyName + `) Serve() error {
        	return p.server.Serve()
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
			nonHTTPRelatedMethodArgs := goMethodArgsWithoutHTTPContextRelatedArgs(method)
			diffNONHTTPRelatedMethodArgs := len(method.Args) - len(nonHTTPRelatedMethodArgs)
			for i := 0; i < diffNONHTTPRelatedMethodArgs; i++ {
				argParams = append(argParams, "nil")
			}
			for _, a := range nonHTTPRelatedMethodArgs {
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
			if len(nonHTTPRelatedMethodArgs) > 0 {
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

func renderGoRPCServiceClients(services ServiceList, fullPackageName string, packageName string, config *config.Target, g *code) error {
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

	g.l(renderImports(aliases, packageName))

	for _, service := range services {
		if !config.IsGoRPC(service.Name) {
			continue
		}

		clientName := service.Name + "GoRPCClient"
		// Client type
		g.l(`
        type ` + clientName + ` struct {
        	Client *gorpc.Client
        }
		`)
		// Constructor
		g.l(`
        func New` + clientName + `(addr string, tlsConfig *tls.Config) *` + clientName + ` {
        	client := &` + clientName + `{}
        	if tlsConfig == nil {
						client.Client = gorpc.NewTCPClient(addr)
					} else {
						client.Client = gorpc.NewTLSClient(addr, tlsConfig)
					}
					return client
        }

        func (tsc *` + clientName + `) Start() {
        	tsc.Client.Start()
      	}

        func (tsc *` + clientName + `) Stop() {
        	tsc.Client.Stop()
      	}
		`)
		g.nl()
		// Methods
		for _, method := range service.Methods {
			args := []string{}
			params := []string{}
			for _, a := range goMethodArgsWithoutHTTPContextRelatedArgs(method) {
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
			g.l(`func (tsc *` + clientName + `) ` + method.Name + `(` + strings.Join(params, ", ") + `) (` + strings.Join(returns, ", ") + `) {`)
			g.l(`req := ` + service.Name + method.Name + `Request{` + strings.Join(args, ", ") + `}`)
			if len(rets) > 0 {
				g.l(`rpcCallRes, rpcCallErr := tsc.Client.Call(req)`)
			} else {
				g.l(`_, rpcCallErr := tsc.Client.Call(req)`)
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

func RenderGoTSRPCProxies(services ServiceList, longPackageName, packageName string, config *config.Target, unions map[string][]string) (gocode string, err error) {
	g := newCode("	")
	err = renderTSRPCServiceProxies(services, longPackageName, packageName, config, unions, g)
	if err != nil {
		return
	}
	gocode = g.string()
	return
}

func RenderGoTSRPCClients(services ServiceList, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := newCode("	")
	err = renderTSRPCServiceClients(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.string()
	return
}

func RenderGoRPCProxies(services ServiceList, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := newCode("	")
	err = renderGoRPCServiceProxies(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.string()
	return
}

func RenderGoRPCClients(services ServiceList, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := newCode("	")
	err = renderGoRPCServiceClients(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.string()
	return
}

func goMethodArgsWithoutHTTPContextRelatedArgs(m *Method) (filteredArgs []*Field) {
	filteredArgs = []*Field{}
	for argI, arg := range m.Args {
		if argI == 0 && arg.Value.isHTTPResponseWriter() {
			continue
		}
		if argI == 1 && arg.Value.isHTTPRequest() {
			continue
		}
		filteredArgs = append(filteredArgs, arg)
	}
	return
}

func renderInit(unions map[string][]string, aliases map[string]string, packageName string, g *code) {
	if len(unions) > 0 {
		g.l("func init() {")
		g.ind(1)
		for pkg, us := range unions {
			for _, name := range us {
				var t string
				if packageName != pkg && aliases[pkg] != "" {
					t += aliases[pkg] + "."
				}
				t += name
				g.l("gotsrpc.MustRegisterUnionExt(" + t + "{})")
			}
		}
		g.ind(-1)
		g.l("}")
	}
}

func renderImports(aliases map[string]string, packageName string) string {
	imports := ""
	for importPath, alias := range aliases {
		imports += alias + " \"" + importPath + "\"\n"
	}
	return `
		// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

		package ` + packageName + `

		import (
			` + imports + `
		)
	`
}
