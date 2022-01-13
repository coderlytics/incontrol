package config

var listeners []func(new Configuration)

// AddConfigChangeListener registers a new listener for configuration changes
func AddConfigChangeListener(listener func(new Configuration)) {
	listeners = append(listeners, listener)
}

// notify informs all listeners for the config change
func notify(new Configuration) {
	for _, listener := range listeners {
		listener(new)
	}
}
