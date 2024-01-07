build:
	go build -o ch 

run-leader: build
	./ch leader -r=1

run-worker:
	./ch worker