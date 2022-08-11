// Code generated by jrpc. DO NOT EDIT.

package routes

import (
	context "context"

	taxi "github.com/jakewright/home-automation/libraries/go/taxi"
	def "github.com/jakewright/home-automation/services/dmx/def"
)

// taxiRouter is an interface implemented by taxi.Router
type taxiRouter interface {
	HandleFunc(method, path string, handler func(context.Context, taxi.Decoder) (interface{}, error))
}

type handler interface {
	GetMegaParProfile(ctx context.Context, body *def.GetMegaParProfileRequest) (*def.MegaParProfileResponse, error)
	UpdateMegaParProfile(ctx context.Context, body *def.UpdateMegaParProfileRequest) (*def.MegaParProfileResponse, error)
}

// Register adds the service's routes to the router
func Register(r taxiRouter, h handler) {
	r.HandleFunc("GET", "/mega-par-profile", func(ctx context.Context, decode taxi.Decoder) (interface{}, error) {
		body := &def.GetMegaParProfileRequest{}
		if err := decode(body); err != nil {
			return nil, err
		}

		if err := body.Validate(); err != nil {
			return nil, err
		}

		return h.GetMegaParProfile(ctx, body)
	})

	r.HandleFunc("PATCH", "/mega-par-profile", func(ctx context.Context, decode taxi.Decoder) (interface{}, error) {
		body := &def.UpdateMegaParProfileRequest{}
		if err := decode(body); err != nil {
			return nil, err
		}

		if err := body.Validate(); err != nil {
			return nil, err
		}

		return h.UpdateMegaParProfile(ctx, body)
	})

}
