test: b a c
	@echo "Running c so calling:"
	@LD_LIBRARY_PATH=. ./b
	@echo "Running go so calling:"
	@LD_LIBRARY_PATH=. go run a.go
	@echo "Running go exec calling:"
	@go run c.go

all: a.o liba.so b a
a.o: a.c
	@gcc -c -s -O2 -fPIC $<

liba.so: a.o
	@gcc -shared -s -O2 $< -o $@

b: b.c liba.so
	@gcc -L. -la -O2 -s $< -o $@

a: a.go liba.so
	@go build -ldflags "-s -w" -o $@ $<

c: c.go
	@go build -ldflags "-s -w" -o $@ $<

clean:
	@rm -f a.o liba.so b a c
