// Copyright 2024 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: BUSL-1.1

package v3

import (
    "context"
    drBase "github.com/pb33f/doctor/model/high/base"
    v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
    "github.com/pb33f/libopenapi/orderedmap"
)

type Operation struct {
    Value        *v3.Operation
    ExternalDocs *drBase.ExternalDoc
    Parameters   []*Parameter
    RequestBody  *RequestBody
    Responses    *Responses
    Callbacks    *orderedmap.Map[string, *Callback]
    Security     []*drBase.SecurityRequirement
    Servers      []*Server
    drBase.Foundation
}

func (o *Operation) Walk(ctx context.Context, operation *v3.Operation) {
    o.Value = operation

    if operation.ExternalDocs != nil {
        ed := &drBase.ExternalDoc{}
        ed.Parent = o
        ed.PathSegment = "externalDocs"
        ed.Value = operation.ExternalDocs
        o.ExternalDocs = ed
    }

    if operation.Parameters != nil {
        for i, parameter := range operation.Parameters {
            p := &Parameter{}
            p.PathSegment = "parameters"
            p.Parent = o
            p.IsIndexed = true
            p.Index = i
            p.Walk(ctx, parameter)
            o.Parameters = append(o.Parameters, p)
        }
    }

    if operation.RequestBody != nil {
        rb := &RequestBody{}
        rb.Parent = o
        rb.PathSegment = "requestBody"
        rb.Walk(ctx, operation.RequestBody)
        o.RequestBody = rb
    }

    if operation.Responses != nil {
        r := &Responses{}
        r.Parent = o
        r.PathSegment = "responses"
        r.Walk(ctx, operation.Responses)
        o.Responses = r
    }

    if operation.Callbacks != nil {
        callbacks := orderedmap.New[string, *Callback]()
        for callbackPairs := operation.Callbacks.First(); callbackPairs != nil; callbackPairs = callbackPairs.Next() {
            c := &Callback{}
            c.Parent = o
            c.PathSegment = "callbacks"
            c.Key = callbackPairs.Key()
            c.Walk(ctx, callbackPairs.Value())
            callbacks.Set(callbackPairs.Key(), c)
        }
        o.Callbacks = callbacks
    }

    if operation.Security != nil {
        for i, security := range operation.Security {
            s := &drBase.SecurityRequirement{}
            s.Parent = o
            s.IsIndexed = true
            s.Index = i
            s.Walk(ctx, security)
            o.Security = append(o.Security, s)
        }
    }

    if operation.Servers != nil {
        for i, server := range operation.Servers {
            s := &Server{}
            s.Parent = o
            s.IsIndexed = true
            s.Index = i
            s.Walk(ctx, server)
            o.Servers = append(o.Servers, s)
        }
    }
}
