ROOT=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CACHE_DIR=$(ROOT)/../assets

BIN=$(CACHE_DIR)/ls_universal_signed
EXTRACTED_AMD_BIN=$(CACHE_DIR)/ls_amd64_signed
EXTRACTED_ARM_BIN=$(CACHE_DIR)/ls_arm64e_signed

.PHONY: check all
all: check $(BIN) $(EXTRACTED_AMD_BIN) $(EXTRACTED_ARM_BIN)

$(BIN):
	cp /bin/ls $(BIN)

$(EXTRACTED_AMD_BIN): $(BIN)
	cp $(BIN) $(BIN).tmp
	lipo $(BIN).tmp -thin x86_64 -output $(EXTRACTED_AMD_BIN)
	rm $(BIN).tmp

$(EXTRACTED_ARM_BIN): $(BIN)
	cp $(BIN) $(BIN).tmp
	lipo $(BIN).tmp -thin arm64e -output $(EXTRACTED_ARM_BIN)
	rm $(BIN).tmp

.PHONY: check-mac
check:
	@if [ $(shell uname) != "Darwin" ]; then echo "must run this on a mac" && exit 1; fi;

.PHONY: clean
clean:
	rm -f $(BIN)*