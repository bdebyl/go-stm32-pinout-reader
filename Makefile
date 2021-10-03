OUTFILE=go-stm32-pinout-reader

all: build

build:
	go build -o ${OUTFILE} .

install: build
	mv ${OUTFILE} $$HOME/.local/bin/${OUTFILE}

uninstall:
	rm $$HOME/.local/bin/${OUTFILE}
