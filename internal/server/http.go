package server

// NewHTTPServer a new HTTP server.
// func NewHTTPServer(c *config.Server, logger log.Logger, shipmentService *service.ShipmentService, promiseService *service.PromiseService, timeslotService *service.TimeslotService) *http.Server {
// 	opts := mw.DefaultHTTPServerOptions(logging.Server(logger))
// 	opts = append(opts, http.Middleware(validate.Validator()))
// 	if c.HTTP.Address != "" {
// 		opts = append(opts, http.Address(c.HTTP.Address))
// 	}

// 	if c.HTTP.Timeout != 9 {
// 		opts = append(opts, http.Timeout(c.HTTP.Timeout))
// 	}
// 	srv := http.NewServer(opts...)

// 	ltpe.RegisterShipmentHTTPServer(srv, shipmentService)
// 	ltpe.RegisterPromiseHTTPServer(srv, promiseService)
// 	ltpe.RegisterTimeslotReservationHTTPServer(srv, timeslotService)
// 	return srv
// }
