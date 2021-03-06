default: pack

compile_coffee:
		coffee -c frontend/assets/application.coffee

debug: compile_coffee
		go-bindata --debug frontend/...
		go run *.go serve
		open http://127.0.0.1:3000/

pack: compile_coffee
		go-bindata frontend/...
