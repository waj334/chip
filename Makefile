TABLEGEN_CSP ?= tablegen-csp

generate-csp:
	$(TABLEGEN_CSP) -in=$(CURDIR)/arm/cortexm/platform/st/stm32h7x7/cm7/def/platform.td \
					-out=$(CURDIR)/arm/cortexm/platform/st/stm32h7x7/cm7 \
					-I=$(CURDIR)
