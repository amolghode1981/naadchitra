SUBDIRS = meghdoot

all: rpc

rpc:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir $(t); \
	done
rpc-clean:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir clean; \
	done
