func (mux *Mux) Handle(path string, prefix bool, handler http.Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	mux.trie.Set(splitpath(path), muxEntry{prefix, handler})
}

func (mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := mux.lookup(r.URL.Path)
	if !ok {
		http.NotFound(w, r)
		return
	}

	handler.ServeHTTP(w, r)
}
