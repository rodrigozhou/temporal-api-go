// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/types"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"

	_ "go.temporal.io/api/workflowservice/v1"
)

const serviceFile = "../../proxy/service.go"

const ServiceHeader = `
// Code generated by proxygenerator; DO NOT EDIT.

package proxy

import (
	"context"

	"go.temporal.io/api/workflowservice/v1"
)

// WorkflowServiceProxyOptions provides options for configuring a WorkflowServiceProxyServer.
// Client is a WorkflowServiceClient used to forward requests received by the server to the 
// Temporal Frontend.
type WorkflowServiceProxyOptions struct {
	Client workflowservice.WorkflowServiceClient
}

type workflowServiceProxyServer struct {
	workflowservice.UnimplementedWorkflowServiceServer
	client workflowservice.WorkflowServiceClient
}

// NewWorkflowServiceProxyServer creates a WorkflowServiceServer suitable for registering with a gRPC Server. Requests will
// be forwarded to the passed in WorkflowService Client. gRPC interceptors can be added on the Server or Client to adjust
// requests and responses.
func NewWorkflowServiceProxyServer(options WorkflowServiceProxyOptions) (workflowservice.WorkflowServiceServer, error) {
	return &workflowServiceProxyServer{
		client: options.Client,
	}, nil
}
`

func generateService(cfg config) error {
	buf := &bytes.Buffer{}
	fmt.Fprint(buf, cfg.license)
	fmt.Fprint(buf, ServiceHeader)

	conf := &packages.Config{Mode: packages.NeedImports | packages.NeedTypes | packages.NeedTypesInfo}
	pkgs, err := packages.Load(conf, "go.temporal.io/api/workflowservice/v1")
	if err != nil {
		return err
	}

	pkg := pkgs[0]
	if len(pkg.Errors) > 0 {
		return fmt.Errorf("unable to load workflowservice: %v", pkg.Errors)
	}

	qual := func(other *types.Package) string {
		if other == pkg.Types {
			return "workflowservice"
		}
		return other.Path()
	}
	scope := pkg.Types.Scope()
	service := scope.Lookup("UnimplementedWorkflowServiceServer")
	if _, ok := service.(*types.TypeName); ok {
		for _, meth := range typeutil.IntuitiveMethodSet(service.Type(), nil) {
			if !meth.Obj().Exported() {
				continue
			}

			name := meth.Obj().Name()
			sig := meth.Obj().Type().(*types.Signature)
			fmt.Fprintf(buf, "\nfunc (s *workflowServiceProxyServer) %s %s %s {\n", name, types.TypeString(sig.Params(), qual), types.TypeString(sig.Results(), qual))
			params := make([]string, sig.Params().Len())
			for i := 0; i < sig.Params().Len(); i++ {
				params[i] = sig.Params().At(i).Name()
			}
			fmt.Fprintf(buf, "\treturn s.client.%s(%s)\n", name, strings.Join(params, ", "))
			fmt.Fprintf(buf, "}\n")
		}
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	if cfg.verifyOnly {
		currentSrc, err := os.ReadFile(serviceFile)
		if err != nil {
			return err
		}

		if !bytes.Equal(src, currentSrc) {
			return fmt.Errorf("generated file does not match existing file: %s", serviceFile)
		}

		return nil
	}

	return os.WriteFile(serviceFile, src, 0666)
}