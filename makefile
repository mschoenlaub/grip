# project configuration
name := grip
buildDir := build
packages := recovery logging message send slogger $(name)
orgPath := github.com/mschoenlaub
projectPath := $(orgPath)/$(name)
# end project configuration

lintBin := ./bin/golangci-lint run


#   include test files and give linters 40s to run to avoid timeouts
lintArgs := --tests --deadline=1m
lintArgs += --skip-dirs="$(buildDir)" --skip-dirs="buildscripts"
#  add and configure additional linters
lintArgs += --enable="goimports" --enable="misspell"
#  there are a lot of logging methods that don't have doc strings, and probably shouldn't
lintArgs += --exclude="exported method Grip\..*should have comment or be unexported.*"
lintArgs += --exclude="exported function (Catch|Log|Default|Emergency|Alert|Critical|Error|Warning|Notice|Info|Debug).* should have comment.*"
lintArgs += --exclude="exported func.*InternalLogger returns unexported type.*"
lintArgs += --exclude="exported method (Log|SystemInfo|InternalSender)\..+ should have comment"
# end lint suppressions


######################################################################
##
## Everything below this point is generic, and does not contain
## project specific configuration. (with one noted case in the "build"
## target for library-only projects)
##
######################################################################


# start dependency installation tools
#   implementation details for being able to lazily install dependencies
gopath := $(shell go env GOPATH)
lintDeps := $(addprefix $(gopath)/src/,$(lintDeps))
srcFiles := makefile $(shell find . -name "*.go" -not -path "./$(buildDir)/*" -not -name "*_test.go")
testSrcFiles := makefile $(shell find . -name "*.go" -not -path "./$(buildDir)/*")
testOutput := $(foreach target,$(packages),$(buildDir)/output.$(target).test)
raceOutput := $(foreach target,$(packages),$(buildDir)/output.$(target).race)
testBin := $(foreach target,$(packages),$(buildDir)/test.$(target))
raceBin := $(foreach target,$(packages),$(buildDir)/race.$(target))
coverageOutput := $(foreach target,$(packages),$(buildDir)/output.$(target).coverage)
coverageHtmlOutput := $(foreach target,$(packages),$(buildDir)/output.$(target).coverage.html)
$(gopath)/src/%:
	@-[ ! -d $(gopath) ] && mkdir -p $(gopath) || true
	go get $(subst $(gopath)/src/,,$@)
$(buildDir)/.lintSetup:$(lintDeps) $(buildDir)
	@-curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.21.0 && touch $@
# end dependency installation tools


# userfacing targets for basic build and development operations
lint:$(buildDir)/output.lint
build:$(deps) $(srcFiles) $(gopath)/src/$(projectPath)
	go build $(subst $(name),,$(subst -,/,$(foreach pkg,$(packages),./$(pkg))))
build-race:$(deps) $(srcFiles) $(gopath)/src/$(projectPath)
	go build -race $(subst -,/,$(foreach pkg,$(packages),./$(pkg)))
test:$(testOutput)
race:$(raceOutput)
coverage:$(coverageOutput)
coverage-html:$(coverageHtmlOutput)
phony := lint build build-race race test benchmark-send coverage coverage-html deps
.PRECIOUS:$(testOutput) $(raceOutput) $(coverageOutput) $(coverageHtmlOutput)
.PRECIOUS:$(foreach target,$(packages),$(buildDir)/test.$(target))
.PRECIOUS:$(foreach target,$(packages),$(buildDir)/race.$(target))
.PRECIOUS:$(foreach target,$(packages),$(buildDir)/output.$(target).lint)
.PRECIOUS:$(buildDir)/output.lint
# end front-ends


# implementation details for building the binary and creating a
# convienent link in the working directory
$(gopath)/src/$(orgPath):
	@mkdir -p $@
$(gopath)/src/$(projectPath):$(gopath)/src/$(orgPath)
	@[ -L $@ ] || ln -s $(shell pwd) $@
$(name):$(buildDir)/$(name)
	@[ -L $@ ] || ln -s $< $@
$(buildDir)/$(name):$(gopath)/src/$(projectPath) $(srcFiles) $(deps)
	go build -o $@ main/$(name).go
$(buildDir)/$(name).race:$(gopath)/src/$(projectPath) $(srcFiles) $(deps)
	go build -race -o $@ main/$(name).go

# convenience targets for runing tests and coverage tasks on a
# specific package.
makeArgs := --no-print-directory
race-%:$(buildDir)/output.%.race
	@grep -s -q -e "^PASS" $< && ! grep -s -q "^WARNING: DATA RACE" $<
test-%:$(buildDir)/output.%.test
	@grep -s -q -e "^PASS" $<
benchmark-send:
	@mkdir -p build
	go test -v -bench=$(if $(RUN_BENCH),$(RUN_BENCH),BenchmarkAllSenders) ./send/ ./send/benchmark/ -run=^^$$
coverage-%:$(buildDir)/output.%.coverage
	@grep -s -q -e "^PASS" $(buildDir)/output.$*.test
html-coverage-%:$(buildDir)/output.%.coverage $(buildDir)/output.%.coverage.html
	@grep -s -q -e "^PASS" $(buildDir)/output.$*.test
lint-%:$(buildDir)/output.%.lint
	@grep -v -s -q "^--- FAIL" $<
# end convienence targets


# start test and coverage artifacts
#    tests have compile and runtime deps. This varable has everything
#    that the tests actually need to run. (The "build" target is
#    intentional and makes these targets rerun as expected.)
testArgs := -test.v --test.timeout=5m
ifneq (,$(RUN_TEST))
testArgs += -test.run='$(RUN_TEST)'
endif
ifneq (,$(RUN_CASE))
testArgs += -testify.m='$(RUN_CASE)'
endif
#    to avoid vendoring the coverage tool, install it as needed
coverDeps := $(if $(DISABLE_COVERAGE),,golang.org/x/tools/cmd/cover)
coverDeps := $(addprefix $(gopath)/src/,$(coverDeps))
#    implementation for package coverage and test running,mongodb to produce
#    and save test output.
$(buildDir)/test.%:$(testSrcFiles) $(coverDeps)
	go test $(if $(DISABLE_COVERAGE),,-covermode=count) -c -o $@ ./$(subst -,/,$*)
$(buildDir)/output.%.test: .FORCE
	@mkdir -p $(buildDir)
	go test $(if $(DISABLE_COVERAGE),,-covermode=count) $(testArgs) ./$(subst -,/,$*) | tee $@
$(buildDir)/output.$(name).test: .FORCE
	@mkdir -p $(buildDir)
	go test $(testArgs) ./ | tee $@

$(buildDir)/output.%.race: .FORCE
	@mkdir -p $(buildDir)
	go test -race $(testArgs) ./$(subst -,/,$*) 2>&1 | tee $@
$(buildDir)/output.$(name).race: .FORCE
	@mkdir -p $(buildDir)
	go test -race $(testArgs) ./ 2>&1 | tee $@
#  targets to generate gotest output from the linter.
$(buildDir)/output.%.lint:$(buildDir)/.lintSetup $(testSrcFiles) .FORCE
	@./$(lintBin) $(lintArgs) --skip-dirs='(^|/)$*($|/)' > $@
$(buildDir)/output.lint:$(buildDir)/.lintSetup .FORCE
	@./$(lintBin) $(lintArgs) > $@
#  targets to process and generate coverage reports
$(buildDir)/output.%.coverage: .FORCE $(coverDeps)
	@mkdir -p $(buildDir)
	go test $(testArgs) -test.coverprofile=$@ | tee $(subst coverage,test,$@)
	@-[ -f $@ ] && go tool cover -func=$@ | sed 's%$(projectPath)/%%' | column -t
$(buildDir)/output.%.coverage.html:$(buildDir)/output.%.coverage $(coverDeps)
	@mkdir -p $(buildDir)
	go tool cover -html=$< -o $@
# end test and coverage artifacts

# clean and other utility targets
clean:
	rm -rf $(name) $(lintDeps) $(buildDir)/test.* $(buildDir)/coverage.* $(buildDir)/race.*
phony += clean
# end dependency targets


# configure phony targets
.FORCE:
.PHONY:$(phony) .FORCE
