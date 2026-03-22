package codegen

import (
	"fmt"
	"sort"
	"strings"

	"github.com/foomo/gotsrpc/v2/config"
	"github.com/foomo/gotsrpc/v2/internal/model"
)

func valueIsHTTPResponseWriter(v *model.Value) bool {
	return (v.StructType != nil && v.StructType.Name == "ResponseWriter" && v.StructType.Package == "net/http") ||
		(v.Scalar != nil && v.Scalar.Name == "ResponseWriter" && v.Scalar.Package == "net/http")
}

func valueIsHTTPRequest(v *model.Value) bool {
	return (v.IsPtr && v.StructType != nil && v.StructType.Name == "Request" && v.StructType.Package == "net/http") ||
		(v.IsPtr && v.Scalar != nil && v.Scalar.Name == "Request" && v.Scalar.Package == "net/http")
}

func valueIsContext(v *model.Value) bool {
	return (v.StructType != nil && v.StructType.Name == "Context" && v.StructType.Package == "context") ||
		(v.Scalar != nil && v.Scalar.Name == "Context" && v.Scalar.Package == "context")
}

func valueGoType(v *model.Value, aliases map[string]string, packageName string) (t string) {
	if v.IsPtr {
		t = "*"
	}
	switch {
	case v.Array != nil:
		t += "[]" + valueGoType(v.Array.Value, aliases, packageName)
	case len(v.GoScalarType) > 0:
		t += v.GoScalarType
	case v.StructType != nil:
		if packageName != v.StructType.Package && aliases[v.StructType.Package] != "" {
			t += aliases[v.StructType.Package] + "."
		}
		t += v.StructType.Name
	case v.Map != nil:
		t += `map[` + valueGoType(v.Map.Key, aliases, packageName) + `]` + valueGoType(v.Map.Value, aliases, packageName)
	case v.Scalar != nil:
		if packageName != v.Scalar.Package && aliases[v.Scalar.Package] != "" {
			t += aliases[v.Scalar.Package] + "."
		}
		t += v.Scalar.Name
	case v.IsInterface:
		t += "any"
	default:
		fmt.Println("WARN: can't resolve goType")
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
		if _, ok := aliases[packageName]; !ok {
			packageParts := strings.Split(packageName, "/")
			beautifulAlias := packageParts[len(packageParts)-1]
			uglyAlias := r.Replace(packageName)
			alias := uglyAlias
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

func extractImports(fields []*model.Field, fullPackageName string, aliases map[string]string) {
	for _, f := range fields {
		extractImportValue(f.Value, fullPackageName, aliases)
	}
}

func extractImportValue(value *model.Value, fullPackageName string, aliases map[string]string) {
	switch {
	case value.StructType != nil:
		extractImport(value.StructType.Package, fullPackageName, aliases)
	case value.Array != nil:
		extractImportValue(value.Array.Value, fullPackageName, aliases)
	case value.Map != nil:
		extractImportValue(value.Map.Key, fullPackageName, aliases)
		extractImportValue(value.Map.Value, fullPackageName, aliases)
	case value.Scalar != nil:
		extractImport(value.Scalar.Package, fullPackageName, aliases)
	}
}

func renderTSRPCServiceProxies(services model.ServiceList, fullPackageName string, packageName string, config *config.Target, unions map[string][]string, g *Code) error {
	aliases := map[string]string{
		"time":                        "time",
		"net/http":                    "http",
		"io":                          "io",
		"github.com/foomo/gotsrpc/v2": "gotsrpc",
	}
	for _, service := range services {
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

	g.L(renderImports(aliases, packageName))

	renderInit(unions, aliases, packageName, g)

	for _, service := range services {
		if !config.IsTSRPC(service.Name) {
			continue
		}

		servicePointer := "*"
		if service.IsInterface {
			servicePointer = ""
		}

		proxyName := service.Name + "GoTSRPCProxy"

		g.L("const (")
		for _, method := range service.Methods {
			g.L(proxyName + method.Name + " = \"" + method.Name + "\"")
		}
		g.L(")")

		g.L(`
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
	        if r.Method == http.MethodOptions {
				return
			} else if r.Method != http.MethodPost {
		        gotsrpc.ErrorMethodNotAllowed(w)
		        return
	        }
			defer io.Copy(io.Discard, r.Body) // Drain Request Body
		`)

		g.L("funcName := gotsrpc.GetCalledFunc(r, p.EndPoint)")
		g.L("callStats, _ := gotsrpc.GetStatsForRequest(r)")
		g.L("callStats.Func = funcName")
		g.L("callStats.Package = \"" + fullPackageName + "\"")
		g.L("callStats.Service = \"" + service.Name + "\"")

		g.L(`switch funcName {`)

		g.Ind(4)

		for _, method := range service.Methods {
			g.L("case " + proxyName + method.Name + ":")
			g.Ind(1)
			var (
				callArgs         []string
				isContextRequest bool
				isSessionRequest bool
			)
			g.L("var (")
			g.Ind(1)
			g.L("args []any")
			g.L("rets []any")
			g.Ind(-1)
			g.L(")")
			if len(method.Args) > 0 {
				var args []string
				var argsDecls []string

				skipArgI := 0

				nonHTTPRelatedArgs := goMethodArgsWithoutHTTPContextRelatedArgs(method)

				isSessionRequest = len(method.Args)-len(nonHTTPRelatedArgs) == 2
				isContextRequest = len(method.Args) > 0 && valueIsContext(method.Args[0].Value)

				for _, arg := range nonHTTPRelatedArgs {
					argName := "arg_" + arg.Name
					argsDecls = append(argsDecls, argName+"  "+valueGoType(arg.Value, aliases, packageName))
					args = append(args, "&"+argName)
					callArgs = append(callArgs, argName)
					skipArgI++
				}
				if len(args) > 0 {
					g.L("var (")
					for _, argDecl := range argsDecls {
						g.L(argDecl)
					}
					g.L(")")
					g.L("args = []any{" + strings.Join(args, ", ") + "}")
					g.L("if err := gotsrpc.LoadArgs(&args, callStats, r); err != nil {")
					g.Ind(1)
					g.L("gotsrpc.ErrorCouldNotLoadArgs(w)")
					g.L("return")
					g.Ind(-1)
					g.L("}")
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
			g.L("executionStart := time.Now()")

			if isSessionRequest {
				g.L("rw := gotsrpc.ResponseWriter{ResponseWriter: w}")
				callArgs = append([]string{"&rw", "r"}, callArgs...)
			} else if isContextRequest {
				callArgs = append([]string{"r.Context()"}, callArgs...)
			}
			if len(returnValueNames) > 0 {
				g.App(strings.Join(returnValueNames, ", ") + " := ")
			}
			g.App("p.service." + method.Name + "(" + strings.Join(callArgs, ", ") + ")")
			g.NL()
			g.L("callStats.Execution = time.Since(executionStart)")
			if isSessionRequest {
				g.L("if rw.Status() == http.StatusOK {").Ind(1)
			}
			g.L("rets = []any{" + strings.Join(returnValueNames, ", ") + "}")
			g.L("if err := gotsrpc.Reply(rets, callStats, r, w); err != nil {")
			g.Ind(1)
			g.L("gotsrpc.ErrorCouldNotReply(w)")
			g.L("return")
			g.Ind(-1)
			g.L("}")
			if isSessionRequest {
				g.Ind(-1).L("}")
			}
			g.L("gotsrpc.Monitor(w, r, args, rets, callStats)")
			g.L("return")
			g.Ind(-1)
		}
		g.L("default:")
		g.Ind(1).L("gotsrpc.ClearStats(r)")
		g.Ind(1).L("gotsrpc.ErrorFuncNotFound(w)")
		g.Ind(-2).L("}") // close switch
		g.Ind(-1).L("}") // close ServeHttp
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

func newMethodSignature(method *model.Method, aliases map[string]string, fullPackageName string) goMethod {
	var args []string
	var params []string
	params = append(params, "ctx go_context.Context")
	for _, a := range goMethodArgsWithoutHTTPContextRelatedArgs(method) {
		args = append(args, a.Name)
		params = append(params, a.Name+" "+valueGoType(a.Value, aliases, fullPackageName))
	}
	var rets []string
	var returns []string
	for i, r := range method.Return {
		name := r.Name
		if len(name) == 0 {
			name = fmt.Sprintf("ret%s_%d", method.Name, i)
		}
		rets = append(rets, "&"+name)
		returns = append(returns, name+" "+valueGoType(r.Value, aliases, fullPackageName))
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

func renderTSRPCServiceClients(services model.ServiceList, fullPackageName string, packageName string, config *config.Target, g *Code) error {
	aliases := map[string]string{
		"github.com/pkg/errors":       "pkg_errors",
		"github.com/foomo/gotsrpc/v2": "gotsrpc",
		"net/http":                    "go_net_http",
		"context":                     "go_context",
	}

	for _, service := range services {
		if !config.IsTSRPC(service.Name) {
			continue
		}
		for _, m := range service.Methods {
			extractImports(m.Args, fullPackageName, aliases)
			extractImports(m.Return, fullPackageName, aliases)
		}
	}

	g.L(renderImports(aliases, packageName))

	for _, service := range services {
		if !config.IsTSRPC(service.Name) {
			continue
		}

		interfaceName := service.Name + "GoTSRPCClient"
		clientName := "HTTP" + interfaceName

		g.L(`type ` + interfaceName + ` interface { `)
		for _, method := range service.Methods {
			ms := newMethodSignature(method, aliases, fullPackageName)
			g.L(ms.renderSignature())
		}

		g.L(`} `)

		g.L(`
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
		g.NL()

		for _, method := range service.Methods {
			ms := newMethodSignature(method, aliases, fullPackageName)
			g.L(`func (tsc *` + clientName + `) ` + ms.renderSignature() + ` {`)
			g.L(`rpcArgs := []any{` + strings.Join(ms.args, ", ") + `}`)
			g.L(`rpcReply := []any{` + strings.Join(ms.rets, ", ") + `}`)
			g.L(`rpcErr := tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "` + method.Name + `", rpcArgs, rpcReply)`)
			g.L(`if rpcErr != nil {`)
			g.Ind(1).L(`clientErr = pkg_errors.WithMessage(rpcErr, "failed to call ` + packageName + `.` + service.Name + `GoTSRPCProxy ` + method.Name + `")`).Ind(-1)
			g.L(`}`)
			g.L(`return`)
			g.L(`}`)
			g.NL()
		}
	}
	return nil
}

func renderGoRPCServiceProxies(services model.ServiceList, fullPackageName string, packageName string, config *config.Target, g *Code) error {
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

	g.L(renderImports(aliases, packageName))

	for _, service := range services {
		if !config.IsGoRPC(service.Name) {
			continue
		}

		servicePointer := "*"
		if service.IsInterface {
			servicePointer = ""
		}

		proxyName := service.Name + "GoRPCProxy"
		g.L(`type (`)
		g.L(`
        ` + proxyName + ` struct {
        	server *gorpc.Server
	        service  ` + servicePointer + service.Name + `
	        callStatsHandler gotsrpc.GoRPCCallStatsHandlerFun
        }
		`)
		for _, method := range service.Methods {
			g.L(ucfirst(service.Name+method.Name) + `Request struct {`)
			for _, a := range goMethodArgsWithoutHTTPContextRelatedArgs(method) {
				g.L(ucfirst(a.Name) + ` ` + valueGoType(a.Value, aliases, fullPackageName))
			}
			g.L(`}`)
			g.L(ucfirst(service.Name+method.Name) + `Response struct {`)
			for i, r := range method.Return {
				name := r.Name
				if len(name) == 0 {
					name = fmt.Sprintf("ret%s_%d", method.Name, i)
				}
				g.L(ucfirst(name) + ` ` + valueGoType(r.Value, aliases, fullPackageName))
			}
			g.L(`}`)
			g.NL()
		}
		g.L(`)`)
		g.NL()
		g.L(`func init() {`)
		for _, method := range service.Methods {
			g.L(`gob.Register(` + ucfirst(service.Name+method.Name) + `Request{})`)
			g.L(`gob.Register(` + ucfirst(service.Name+method.Name) + `Response{})`)
		}
		g.L(`}`)
		g.L(`
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
		g.NL()
		g.L(`func (p *` + proxyName + `) handler(clientAddr string, request any) (response any) {`)
		g.L(`start := time.Now()`)
		g.NL()
		g.L(`reqType := reflect.TypeOf(request).String()`)
		g.L(`funcNameParts := strings.Split(reqType, ".")`)
		g.L(`funcName := funcNameParts[len(funcNameParts)-1]`)
		g.NL()
		g.L(`switch funcName {`)
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
			g.L(`case "` + service.Name + method.Name + `Request":`)
			if len(nonHTTPRelatedMethodArgs) > 0 {
				g.L(`req := request.(` + service.Name + method.Name + `Request)`)
			}
			if len(rets) > 0 {
				g.L(strings.Join(rets, ", ") + ` := p.service.` + method.Name + `(` + strings.Join(argParams, ", ") + `)`)
			} else {
				g.L(`p.service.` + method.Name + `(` + strings.Join(argParams, ", ") + `)`)
			}
			g.L(`response = ` + service.Name + method.Name + `Response{` + strings.Join(retParams, ", ") + `}`)
		}
		g.L(`default:`)
		g.L(`fmt.Println("Unknown request type", reflect.TypeOf(request).String())`)
		g.L(`}`)
		g.NL()
		g.L(`if p.callStatsHandler != nil {`)
		g.L(`p.callStatsHandler(&gotsrpc.CallStats{`)
		g.L(`Func: funcName,`)
		g.L(`Package: "` + fullPackageName + `",`)
		g.L(`Service: "` + service.Name + `",`)
		g.L(`Execution: time.Since(start),`)
		g.L(`})`)
		g.L(`}`)
		g.NL()
		g.L(`return`)
		g.L(`}`)
	}
	return nil
}

func renderGoRPCServiceClients(services model.ServiceList, fullPackageName string, packageName string, config *config.Target, g *Code) error {
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

	g.L(renderImports(aliases, packageName))

	for _, service := range services {
		if !config.IsGoRPC(service.Name) {
			continue
		}

		clientName := service.Name + "GoRPCClient"
		g.L(`
        type ` + clientName + ` struct {
        	Client *gorpc.Client
        }
		`)
		g.L(`
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
		g.NL()
		for _, method := range service.Methods {
			var (
				args   []string
				params []string
			)
			for _, a := range goMethodArgsWithoutHTTPContextRelatedArgs(method) {
				args = append(args, ucfirst(a.Name)+`: `+a.Name)
				params = append(params, a.Name+" "+valueGoType(a.Value, aliases, fullPackageName))
			}
			var (
				rets    []string
				returns []string
			)
			for i, r := range method.Return {
				name := r.Name
				if len(name) == 0 {
					name = fmt.Sprintf("ret%s_%d", method.Name, i)
				}
				rets = append(rets, "rpcResp."+ucfirst(name))
				returns = append(returns, name+" "+valueGoType(r.Value, aliases, fullPackageName))
			}
			returns = append(returns, "clientErr error")
			g.L(`func (tsc *` + clientName + `) ` + method.Name + `(` + strings.Join(params, ", ") + `) (` + strings.Join(returns, ", ") + `) {`)
			g.L(`rpcReq := ` + service.Name + method.Name + `Request{` + strings.Join(args, ", ") + `}`)
			if len(rets) > 0 {
				g.L(`rpcRes, rpcErr := tsc.Client.Call(rpcReq)`)
			} else {
				g.L(`_, rpcErr := tsc.Client.Call(rpcReq)`)
			}
			g.L(`if rpcErr != nil {`)
			g.L(`clientErr = rpcErr`)
			g.L(`return`)
			g.L(`}`)
			if len(rets) > 0 {
				g.L(`rpcResp := rpcRes.(` + service.Name + method.Name + `Response)`)
				g.L(`return ` + strings.Join(rets, ", ") + `, nil`)
			} else {
				g.L(`return nil`)
			}
			g.L(`}`)
			g.NL()
		}
	}
	return nil
}

func RenderGoTSRPCProxies(services model.ServiceList, longPackageName, packageName string, config *config.Target, unions map[string][]string) (gocode string, err error) {
	g := NewCode("	")
	err = renderTSRPCServiceProxies(services, longPackageName, packageName, config, unions, g)
	if err != nil {
		return
	}
	gocode = g.String()
	return
}

func RenderGoTSRPCClients(services model.ServiceList, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := NewCode("	")
	err = renderTSRPCServiceClients(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.String()
	return
}

func RenderGoRPCProxies(services model.ServiceList, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := NewCode("	")
	err = renderGoRPCServiceProxies(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.String()
	return
}

func RenderGoRPCClients(services model.ServiceList, longPackageName, packageName string, config *config.Target) (gocode string, err error) {
	g := NewCode("	")
	err = renderGoRPCServiceClients(services, longPackageName, packageName, config, g)
	if err != nil {
		return
	}
	gocode = g.String()
	return
}

func goMethodArgsWithoutHTTPContextRelatedArgs(m *model.Method) (filteredArgs []*model.Field) {
	filteredArgs = []*model.Field{}
	for argI, arg := range m.Args {
		if argI == 0 && valueIsHTTPResponseWriter(arg.Value) {
			continue
		}
		if argI == 1 && valueIsHTTPRequest(arg.Value) {
			continue
		}
		if argI == 0 && valueIsContext(arg.Value) {
			continue
		}
		filteredArgs = append(filteredArgs, arg)
	}
	return
}

func renderInit(unions map[string][]string, aliases map[string]string, packageName string, g *Code) {
	if len(unions) > 0 {
		g.L("func init() {")
		g.Ind(1)
		var strs []string
		for pkg, us := range unions {
			for _, name := range us {
				var str string
				if packageName != pkg && aliases[pkg] != "" {
					str += aliases[pkg] + "."
				}
				str += name
				strs = append(strs, str)
			}
		}
		sort.Strings(strs)
		for _, str := range strs {
			g.L("gotsrpc.MustRegisterUnionExt(" + str + "{})")
		}
		g.Ind(-1)
		g.L("}")
	}
}

func renderImports(aliases map[string]string, packageName string) string {
	imports := ""
	for importPath, alias := range aliases {
		imports += alias + " \"" + importPath + "\"\n"
	}
	return `
		// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2 - DO NOT EDIT.

		package ` + packageName + `

		import (
			` + imports + `
		)
	`
}
