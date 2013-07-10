type Router struct {
	mux         *triemux.Mux
	mongoUrl    string
	mongoDbName string
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.mux.ServeHTTP(w, r)
}

func (rt *Router) ReloadRoutes() {
	// save a reference to the previous mux in case we have to restore it
	oldmux := rt.mux
	defer func() {
		if r := recover(); r != nil {
			log.Println("router: recovered from panic in ReloadRoutes:", r)
			rt.mux = oldmux
			log.Println("router: original routes have been restored")
		}
	}()
	...
	sess, err := mgo.Dial(rt.mongoUrl)
	...
	newmux := triemux.NewMux()
	apps := loadApplications(db.C("applications"), newmux)
	loadRoutes(db.C("routes"), newmux, apps)
	...
	rt.mux = newmux // HL
	log.Printf("router: reloaded routes")
}
