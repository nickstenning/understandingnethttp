func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	...
}

func ListenAndServe(addr string, handler Handler) error {
	...
}

func (mux *ServeMux) Handle(pattern string, handler Handler) {
	...
}
