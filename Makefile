# all our targets are phony (no files to check).
.PHONY: artifacts

curdir = $(PWD)
export curdir

# pass "version" argument and optional "skip-test=1"
artifacts:
	docker build . -f "./Dockerfile.artifacts" -t cardinalby/tlgstats:artifacts &&\
	docker run -it -v $(curdir)/artifacts:/artifacts/ --env VERSION=$(version) --env SKIP_TEST=${skip-test} cardinalby/tlgstats:artifacts