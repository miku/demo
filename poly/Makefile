.PHONY: demo
demo: data.json
	go run main.go < data.json

data.json:
	python gen.py -n 1000 > data.json

clean:
	rm -f data.json
