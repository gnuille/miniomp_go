sources := $(wildcard *.pgo)
psources:= $(patsubst %.pgo,%.go,$(sources))
EXTRAE_ON?=0

install: $(psources)
	go install

%.go: %.pgo
	cpp -DEXTRAE_ON=$(EXTRAE_ON) -CC -E -P $< $@

clean:
	rm -f *.go

taskloop:
	$(MAKE) install
	export OMP_NUM_THREADS=8
	export OMP_PROC_BIND=true
	go run examples/taskloop.go
	$(MAKE) clean

extrae:
	$(MAKE) clean
	$(MAKE) install EXTRAE_ON=1
	bash share/run_extrae.sh
	$(MAKE) clean

