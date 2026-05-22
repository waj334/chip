SIGOC ?= sigoc
stack_size ?= 4096
all:
	$(SIGOC) build --work --dump-ir --cpu stm32h747xi -g -O0 --float hardfp --stack-size=$(stack_size) -o./bin/firmware.elf .
clean:
	rm -rf ./bin