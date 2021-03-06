// Copyright 2018 ekino.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package godim

// Initializer interface to implement if you need specific initialization
//
// OnInit will be called after Injection phase
type Initializer interface {
	OnInit() error
}

// Closer interface to implement if you need specific closing method
//
// OnClose will be called on close phase
type Closer interface {
	OnClose() error
}

// Identifier interface to implement if you want to name your service.
//
// Key is the key name that will reference it in the other service
type Identifier interface {
	Key() string
}
