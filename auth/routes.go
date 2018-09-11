package auth

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/crusttech/crust/config"
)

func mountRoutes(r chi.Router, opts *config.HTTP, mounts ...func(r chi.Router)) {
	if opts.Logging {
		r.Use(middleware.Logger)
	}
	if opts.Metrics {
		r.Use(metrics{}.Middleware("crm"))
	}

	for _, mount := range mounts {
		mount(r)
	}
}

func mountSystemRoutes(r chi.Router, opts *config.HTTP) {
	if opts.Metrics {
		r.Handle("/metrics", metrics{}.Handler())
	}
	r.Mount("/debug", middleware.Profiler())
}

func printRoutes(r chi.Router, opts *config.HTTP) {
	var printRoutes func(chi.Routes, string, string)
	printRoutes = func(r chi.Routes, indent string, prefix string) {
		routes := r.Routes()
		for _, route := range routes {
			if route.SubRoutes != nil && len(route.SubRoutes.Routes()) > 0 {
				fmt.Printf(indent+"%s - with %d handlers, %d subroutes\n", route.Pattern, len(route.Handlers), len(route.SubRoutes.Routes()))
				printRoutes(route.SubRoutes, indent+"\t", prefix+route.Pattern[:len(route.Pattern)-2])
			} else {
				for key, fn := range route.Handlers {
					fmt.Printf("%s%s\t%s -> %s\n", indent, key, prefix+route.Pattern, runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
				}
			}
		}
	}
	printRoutes(r, "", "")
}
