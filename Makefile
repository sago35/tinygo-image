smoketest: FORCE
	mkdir -p out
	tinygo build -o out/basic.uf2 --target wioterminal --size short ./examples/basic/

FORCE:
