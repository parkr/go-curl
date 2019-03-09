build:
	go install github.com/parkr/go-curl/...

test: build
	go-curl https://ping.parkermoo.re
	go-curl -f https://ping.parkermoo.re || true
	go-curl -f https://plants.keltermoore.com/ping
