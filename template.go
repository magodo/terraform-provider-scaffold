// Code generated by go generate; DO NOT EDIT.
package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
	"path/filepath"
)

type dirMetaData struct {
	path     string
	fileMode os.FileMode
}

type fileMetaData struct {
	path     string
	fileMode os.FileMode
	content  string
}

var (
	templateDirs = []dirMetaData{
		dirMetaData{
			path: 		".github",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		".github/ISSUE_TEMPLATE",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		".github/workflows",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		"scripts",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		"website",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		"website/docs",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		"website/docs/d",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		"website/docs/r",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		"{{.ProviderName}}",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		"{{.ProviderName}}/internal",
			fileMode: 	2147484141,
		},
		dirMetaData{
			path: 		"{{.ProviderName}}/internal/provider",
			fileMode: 	2147484141,
		},
	}

	templateFiles = []fileMetaData{
		fileMetaData{
			path: 		".github/CODE_OF_CONDUCT.md",
			fileMode: 	420,
			content:	`# Code of Conduct

HashiCorp Community Guidelines apply to you when interacting with the community here on GitHub and contributing code.

Please read the full text at https://www.hashicorp.com/community-guidelines
`,
		},
		fileMetaData{
			path: 		".github/ISSUE_TEMPLATE/Bug_Report.md",
			fileMode: 	420,
			content:	fmt.Sprintf(`---
name: 🐛 Bug Report
about: If something isn't working as expected 🤔.

---

<!---
Please note the following potential times when an issue might be in Terraform core:

* [Configuration Language](https://www.terraform.io/docs/configuration/index.html) or resource ordering issues
* [State](https://www.terraform.io/docs/state/index.html) and [State Backend](https://www.terraform.io/docs/backends/index.html) issues
* [Provisioner](https://www.terraform.io/docs/provisioners/index.html) issues
* [Registry](https://registry.terraform.io/) issues
* Spans resources across multiple providers

If you are running into one of these scenarios, we recommend opening an issue in the [Terraform core repository](https://github.com/hashicorp/terraform/) instead.
--->

<!--- Please keep this note for the community --->

### Community Note

* Please vote on this issue by adding a 👍 [reaction](https://blog.github.com/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/) to the original issue to help the community and maintainers prioritize this request
* Please do not leave "+1" or "me too" comments, they generate extra noise for issue followers and do not help prioritize the request
* If you are interested in working on this issue or have submitted a pull request, please leave a comment

<!--- Thank you for keeping this note for the community --->

### Terraform (and AzureRM Provider) Version

<!--- Please run %[1]sterraform -v%[1]s to show the Terraform core version and provider version(s). If you are not running the latest version of Terraform or the provider, please upgrade because your issue may have already been fixed. [Terraform documentation on provider versioning](https://www.terraform.io/docs/configuration/providers.html#provider-versions). --->

### Affected Resource(s)

<!--- Please list the affected resources and data sources. --->

* %[1]sazurerm_XXXXX%[1]s

### Terraform Configuration Files

<!--- Information about code formatting: https://help.github.com/articles/basic-writing-and-formatting-syntax/#quoting-code --->

%[1]s%[1]s%[1]shcl
# Copy-paste your Terraform configurations here - for large Terraform configs,
# please use a service like Dropbox and share a link to the ZIP file. For
# security, you can also encrypt the files using our GPG public key: https://keybase.io/hashicorp
%[1]s%[1]s%[1]s

### Debug Output

<!---
Please provide a link to a GitHub Gist containing the complete debug output. Please do NOT paste the debug output in the issue; just paste a link to the Gist.

To obtain the debug output, see the [Terraform documentation on debugging](https://www.terraform.io/docs/internals/debugging.html).
--->

### Panic Output

<!--- If Terraform produced a panic, please provide a link to a GitHub Gist containing the output of the %[1]scrash.log%[1]s. --->

### Expected Behavior

<!--- What should have happened? --->

### Actual Behavior

<!--- What actually happened? --->

### Steps to Reproduce

<!--- Please list the steps required to reproduce the issue. --->

1. %[1]sterraform apply%[1]s

### Important Factoids

<!--- Are there anything atypical about your accounts that we should know? For example: Running in a Azure China/Germany/Government? --->

### References

<!---
Information about referencing Github Issues: https://help.github.com/articles/basic-writing-and-formatting-syntax/#referencing-issues-and-pull-requests

Are there any other GitHub issues (open or closed) or pull requests that should be linked here? Such as vendor documentation?
--->

* #0000
`, "`"),
		},
		fileMetaData{
			path: 		".github/ISSUE_TEMPLATE/Feature_Request.md",
			fileMode: 	420,
			content:	fmt.Sprintf(`---
name: 🚀 Feature Request
about: I have a suggestion (and might want to implement myself 🙂)!
title: Support for [thing]

---

<!--- Please keep this note for the community --->

### Community Note

* Please vote on this issue by adding a 👍 [reaction](https://blog.github.com/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/) to the original issue to help the community and maintainers prioritize this request
* Please do not leave "+1" or "me too" comments, they generate extra noise for issue followers and do not help prioritize the request
* If you are interested in working on this issue or have submitted a pull request, please leave a comment

<!--- Thank you for keeping this note for the community --->

### Description

<!--- Please leave a helpful description of the feature request here. --->

### New or Affected Resource(s)

<!--- Please list the new or affected resources and data sources. --->

* azurerm_XXXXX

### Potential Terraform Configuration

<!--- Information about code formatting: https://help.github.com/articles/basic-writing-and-formatting-syntax/#quoting-code --->

%[1]s%[1]s%[1]shcl
# Copy-paste your Terraform configurations here - for large Terraform configs,
# please use a service like Dropbox and share a link to the ZIP file. For
# security, you can also encrypt the files using our GPG public key.
%[1]s%[1]s%[1]s

### References

<!---
Information about referencing Github Issues: https://help.github.com/articles/basic-writing-and-formatting-syntax/#referencing-issues-and-pull-requests

Are there any other GitHub issues (open or closed) or pull requests that should be linked here? Vendor blog posts or documentation? For example:

* https://azure.microsoft.com/en-us/roadmap/virtual-network-service-endpoint-for-azure-cosmos-db/
--->

* #0000
`, "`"),
		},
		fileMetaData{
			path: 		".github/ISSUE_TEMPLATE/Question.md",
			fileMode: 	420,
			content:	`---
name: 💬 Question
about: If you have a question, please check out our other community resources!

---

Issues on GitHub are intended to be related to bugs or feature requests with provider codebase,
so we recommend using our other community resources instead of asking here 👍.

---

If you have a support request or question please submit them to one of these resources:

* [HashiCorp Community Forums](https://discuss.hashicorp.com/c/terraform-providers)
* [HashiCorp support](https://support.hashicorp.com) (Terraform Enterprise customers)
`,
		},
		fileMetaData{
			path: 		".github/ISSUE_TEMPLATE/config.yml",
			fileMode: 	420,
			content:	`blank_issues_enabled: false
`,
		},
		fileMetaData{
			path: 		".github/ISSUE_TEMPLATE.md",
			fileMode: 	420,
			content:	`<!---
Thanks for filing an issue 😄 ! Before you submit, please read the following:

Check the other issue templates if you are trying to submit a bug report, feature request, or question
Search open/closed issues before submitting since someone might have asked the same thing before!
-->
`,
		},
		fileMetaData{
			path: 		".github/SUPPORT.md",
			fileMode: 	420,
			content:	`# Support

Terraform is a mature project with a growing community. There are active, dedicated people willing to help you through various mediums.

Take a look at those mediums listed at https://www.terraform.io/community.html
`,
		},
		fileMetaData{
			path: 		".github/workflows/provider.yml",
			fileMode: 	420,
			content:	`name: Provider
on: 
  pull_request:
    branches: 
      - master

jobs:
  lintrest:
    name: lintrest
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go 1.13
        uses: actions/setup-go@v1
        with:
          go-version: 1.13
        id: go

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Install dependencies
        run: make tools

      - name: run lintrest
        run: |
          export PATH="$PATH:$HOME/go/bin"
          export GOPATH=$HOME/go
          GOGC=5 make lintrest
  tflint:
    name: tflint
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go 1.13
        uses: actions/setup-go@v1
        with:
          go-version: 1.13
        id: go

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Install dependencies
        run: make tools

      - name: run tflint
        run: |
          export PATH="$PATH:$HOME/go/bin"
          export GOPATH=$HOME/go
          GO111MODULE=off go get -u github.com/hashicorp/terraform
          make tflint
  test:
    name: test
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go 1.13
        uses: actions/setup-go@v1
        with:
          go-version: 1.13
        id: go

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Install dependencies
        run: make tools

      - name: make test
        run: |
          export PATH="$PATH:$HOME/go/bin"
          export GOPATH=$HOME/go
          make test
  depscheck:
    name: depscheck
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go 1.13
        uses: actions/setup-go@v1
        with:
          go-version: 1.13
        id: go

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Install dependencies
        run: make tools

      - name: make depscheck
        run: |
          export PATH="$PATH:$HOME/go/bin"
          export GOPATH=$HOME/go
          make depscheck
  website-lint:
    name: website-lint
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go 1.13
        uses: actions/setup-go@v1
        with:
          go-version: 1.13
        id: go

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Install dependencies
        run: make tools

      - name: make website-lint
        run: |
          export PATH="$PATH:$HOME/go/bin"
          export GOPATH=$HOME/go
          GO111MODULE=off go get -u github.com/hashicorp/terraform
          make website-lint
  website-test:
    name: website-test
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go 1.13
        uses: actions/setup-go@v1
        with:
          go-version: 1.13
        id: go

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Install dependencies
        run: make tools

      - name: make website-test
        run: |
          export PATH="$PATH:$HOME/go/bin"
          export GOPATH=$HOME/go
          make website-test
`,
		},
		fileMetaData{
			path: 		".gitignore",
			fileMode: 	420,
			content:	`server
test.bin
cmd/server/ver.go
`,
		},
		fileMetaData{
			path: 		"CHANGELOG.md",
			fileMode: 	420,
			content:	`## 0.0.1 (Unreleased)

NOTES:

* note1

FEATURES:

* feature1

ENHANCEMENTS:

* enhancement1

BUG FIXES:

* bugfix1
`,
		},
		fileMetaData{
			path: 		"GNUmakefile",
			fileMode: 	420,
			content:	`TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
WEBSITE_REPO=github.com/hashicorp/terraform-website
TESTTIMEOUT=180m
PKG_NAME={{.ProviderName}}

default: build

tools:
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	GO111MODULE=off go get -u github.com/bflad/tfproviderlint/cmd/tfproviderlint
	GO111MODULE=off go get -u github.com/katbyte/terrafmt

build: fmtcheck
	go install

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

goimports:
	@echo "==> Fixing imports code with goimports..."
	goimports -w $(PKG_NAME)/

lint:
	@echo "==> Checking source code against linters..."
	golangci-lint run ./...

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout $(TESTTIMEOUT)

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

depscheck:
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)
	@echo "==> Checking source code with go mod vendor..."
	@go mod vendor
	@git diff --compact-summary --exit-code -- vendor || \
		(echo; echo "Unexpected difference in vendor/ directory. Run 'go mod vendor' command or revert any go.mod/go.sum/vendor changes and commit."; exit 1)

tflint:
	tfproviderlint ./$(PKG_NAME)/...

whitespace:
	@echo "==> Fixing source code with whitespace linter..."
	golangci-lint run ./... --no-config --disable-all --enable=whitespace --fix

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-lint:
	@echo "==> Checking website against linters..."
	@misspell -error -source=text website/
	@echo "==> Checking documentation for errors..."
	@tfproviderdocs check -provider-name=$(PKG_NAME) -require-resource-subcategory \
		-allowed-resource-subcategories-file website/allowed-subcategories

terrafmt-lint:
	@sh -c "'$(CURDIR)/scripts/terrafmt-acctests.sh'"
	@sh -c "'$(CURDIR)/scripts/terrafmt-website.sh'"

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: build test testacc vet fmt fmtcheck lint tools test-compile website website-lint website-test
`,
		},
		fileMetaData{
			path: 		"LICENSE",
			fileMode: 	420,
			content:	`Mozilla Public License Version 2.0
==================================

1. Definitions
--------------

1.1. "Contributor"
    means each individual or legal entity that creates, contributes to
    the creation of, or owns Covered Software.

1.2. "Contributor Version"
    means the combination of the Contributions of others (if any) used
    by a Contributor and that particular Contributor's Contribution.

1.3. "Contribution"
    means Covered Software of a particular Contributor.

1.4. "Covered Software"
    means Source Code Form to which the initial Contributor has attached
    the notice in Exhibit A, the Executable Form of such Source Code
    Form, and Modifications of such Source Code Form, in each case
    including portions thereof.

1.5. "Incompatible With Secondary Licenses"
    means

    (a) that the initial Contributor has attached the notice described
        in Exhibit B to the Covered Software; or

    (b) that the Covered Software was made available under the terms of
        version 1.1 or earlier of the License, but not also under the
        terms of a Secondary License.

1.6. "Executable Form"
    means any form of the work other than Source Code Form.

1.7. "Larger Work"
    means a work that combines Covered Software with other material, in
    a separate file or files, that is not Covered Software.

1.8. "License"
    means this document.

1.9. "Licensable"
    means having the right to grant, to the maximum extent possible,
    whether at the time of the initial grant or subsequently, any and
    all of the rights conveyed by this License.

1.10. "Modifications"
    means any of the following:

    (a) any file in Source Code Form that results from an addition to,
        deletion from, or modification of the contents of Covered
        Software; or

    (b) any new file in Source Code Form that contains any Covered
        Software.

1.11. "Patent Claims" of a Contributor
    means any patent claim(s), including without limitation, method,
    process, and apparatus claims, in any patent Licensable by such
    Contributor that would be infringed, but for the grant of the
    License, by the making, using, selling, offering for sale, having
    made, import, or transfer of either its Contributions or its
    Contributor Version.

1.12. "Secondary License"
    means either the GNU General Public License, Version 2.0, the GNU
    Lesser General Public License, Version 2.1, the GNU Affero General
    Public License, Version 3.0, or any later versions of those
    licenses.

1.13. "Source Code Form"
    means the form of the work preferred for making modifications.

1.14. "You" (or "Your")
    means an individual or a legal entity exercising rights under this
    License. For legal entities, "You" includes any entity that
    controls, is controlled by, or is under common control with You. For
    purposes of this definition, "control" means (a) the power, direct
    or indirect, to cause the direction or management of such entity,
    whether by contract or otherwise, or (b) ownership of more than
    fifty percent (50%) of the outstanding shares or beneficial
    ownership of such entity.

2. License Grants and Conditions
--------------------------------

2.1. Grants

Each Contributor hereby grants You a world-wide, royalty-free,
non-exclusive license:

(a) under intellectual property rights (other than patent or trademark)
    Licensable by such Contributor to use, reproduce, make available,
    modify, display, perform, distribute, and otherwise exploit its
    Contributions, either on an unmodified basis, with Modifications, or
    as part of a Larger Work; and

(b) under Patent Claims of such Contributor to make, use, sell, offer
    for sale, have made, import, and otherwise transfer either its
    Contributions or its Contributor Version.

2.2. Effective Date

The licenses granted in Section 2.1 with respect to any Contribution
become effective for each Contribution on the date the Contributor first
distributes such Contribution.

2.3. Limitations on Grant Scope

The licenses granted in this Section 2 are the only rights granted under
this License. No additional rights or licenses will be implied from the
distribution or licensing of Covered Software under this License.
Notwithstanding Section 2.1(b) above, no patent license is granted by a
Contributor:

(a) for any code that a Contributor has removed from Covered Software;
    or

(b) for infringements caused by: (i) Your and any other third party's
    modifications of Covered Software, or (ii) the combination of its
    Contributions with other software (except as part of its Contributor
    Version); or

(c) under Patent Claims infringed by Covered Software in the absence of
    its Contributions.

This License does not grant any rights in the trademarks, service marks,
or logos of any Contributor (except as may be necessary to comply with
the notice requirements in Section 3.4).

2.4. Subsequent Licenses

No Contributor makes additional grants as a result of Your choice to
distribute the Covered Software under a subsequent version of this
License (see Section 10.2) or under the terms of a Secondary License (if
permitted under the terms of Section 3.3).

2.5. Representation

Each Contributor represents that the Contributor believes its
Contributions are its original creation(s) or it has sufficient rights
to grant the rights to its Contributions conveyed by this License.

2.6. Fair Use

This License is not intended to limit any rights You have under
applicable copyright doctrines of fair use, fair dealing, or other
equivalents.

2.7. Conditions

Sections 3.1, 3.2, 3.3, and 3.4 are conditions of the licenses granted
in Section 2.1.

3. Responsibilities
-------------------

3.1. Distribution of Source Form

All distribution of Covered Software in Source Code Form, including any
Modifications that You create or to which You contribute, must be under
the terms of this License. You must inform recipients that the Source
Code Form of the Covered Software is governed by the terms of this
License, and how they can obtain a copy of this License. You may not
attempt to alter or restrict the recipients' rights in the Source Code
Form.

3.2. Distribution of Executable Form

If You distribute Covered Software in Executable Form then:

(a) such Covered Software must also be made available in Source Code
    Form, as described in Section 3.1, and You must inform recipients of
    the Executable Form how they can obtain a copy of such Source Code
    Form by reasonable means in a timely manner, at a charge no more
    than the cost of distribution to the recipient; and

(b) You may distribute such Executable Form under the terms of this
    License, or sublicense it under different terms, provided that the
    license for the Executable Form does not attempt to limit or alter
    the recipients' rights in the Source Code Form under this License.

3.3. Distribution of a Larger Work

You may create and distribute a Larger Work under terms of Your choice,
provided that You also comply with the requirements of this License for
the Covered Software. If the Larger Work is a combination of Covered
Software with a work governed by one or more Secondary Licenses, and the
Covered Software is not Incompatible With Secondary Licenses, this
License permits You to additionally distribute such Covered Software
under the terms of such Secondary License(s), so that the recipient of
the Larger Work may, at their option, further distribute the Covered
Software under the terms of either this License or such Secondary
License(s).

3.4. Notices

You may not remove or alter the substance of any license notices
(including copyright notices, patent notices, disclaimers of warranty,
or limitations of liability) contained within the Source Code Form of
the Covered Software, except that You may alter any license notices to
the extent required to remedy known factual inaccuracies.

3.5. Application of Additional Terms

You may choose to offer, and to charge a fee for, warranty, support,
indemnity or liability obligations to one or more recipients of Covered
Software. However, You may do so only on Your own behalf, and not on
behalf of any Contributor. You must make it absolutely clear that any
such warranty, support, indemnity, or liability obligation is offered by
You alone, and You hereby agree to indemnify every Contributor for any
liability incurred by such Contributor as a result of warranty, support,
indemnity or liability terms You offer. You may include additional
disclaimers of warranty and limitations of liability specific to any
jurisdiction.

4. Inability to Comply Due to Statute or Regulation
---------------------------------------------------

If it is impossible for You to comply with any of the terms of this
License with respect to some or all of the Covered Software due to
statute, judicial order, or regulation then You must: (a) comply with
the terms of this License to the maximum extent possible; and (b)
describe the limitations and the code they affect. Such description must
be placed in a text file included with all distributions of the Covered
Software under this License. Except to the extent prohibited by statute
or regulation, such description must be sufficiently detailed for a
recipient of ordinary skill to be able to understand it.

5. Termination
--------------

5.1. The rights granted under this License will terminate automatically
if You fail to comply with any of its terms. However, if You become
compliant, then the rights granted under this License from a particular
Contributor are reinstated (a) provisionally, unless and until such
Contributor explicitly and finally terminates Your grants, and (b) on an
ongoing basis, if such Contributor fails to notify You of the
non-compliance by some reasonable means prior to 60 days after You have
come back into compliance. Moreover, Your grants from a particular
Contributor are reinstated on an ongoing basis if such Contributor
notifies You of the non-compliance by some reasonable means, this is the
first time You have received notice of non-compliance with this License
from such Contributor, and You become compliant prior to 30 days after
Your receipt of the notice.

5.2. If You initiate litigation against any entity by asserting a patent
infringement claim (excluding declaratory judgment actions,
counter-claims, and cross-claims) alleging that a Contributor Version
directly or indirectly infringes any patent, then the rights granted to
You by any and all Contributors for the Covered Software under Section
2.1 of this License shall terminate.

5.3. In the event of termination under Sections 5.1 or 5.2 above, all
end user license agreements (excluding distributors and resellers) which
have been validly granted by You or Your distributors under this License
prior to termination shall survive termination.

************************************************************************
*                                                                      *
*  6. Disclaimer of Warranty                                           *
*  -------------------------                                           *
*                                                                      *
*  Covered Software is provided under this License on an "as is"       *
*  basis, without warranty of any kind, either expressed, implied, or  *
*  statutory, including, without limitation, warranties that the       *
*  Covered Software is free of defects, merchantable, fit for a        *
*  particular purpose or non-infringing. The entire risk as to the     *
*  quality and performance of the Covered Software is with You.        *
*  Should any Covered Software prove defective in any respect, You     *
*  (not any Contributor) assume the cost of any necessary servicing,   *
*  repair, or correction. This disclaimer of warranty constitutes an   *
*  essential part of this License. No use of any Covered Software is   *
*  authorized under this License except under this disclaimer.         *
*                                                                      *
************************************************************************

************************************************************************
*                                                                      *
*  7. Limitation of Liability                                          *
*  --------------------------                                          *
*                                                                      *
*  Under no circumstances and under no legal theory, whether tort      *
*  (including negligence), contract, or otherwise, shall any           *
*  Contributor, or anyone who distributes Covered Software as          *
*  permitted above, be liable to You for any direct, indirect,         *
*  special, incidental, or consequential damages of any character      *
*  including, without limitation, damages for lost profits, loss of    *
*  goodwill, work stoppage, computer failure or malfunction, or any    *
*  and all other commercial damages or losses, even if such party      *
*  shall have been informed of the possibility of such damages. This   *
*  limitation of liability shall not apply to liability for death or   *
*  personal injury resulting from such party's negligence to the       *
*  extent applicable law prohibits such limitation. Some               *
*  jurisdictions do not allow the exclusion or limitation of           *
*  incidental or consequential damages, so this exclusion and          *
*  limitation may not apply to You.                                    *
*                                                                      *
************************************************************************

8. Litigation
-------------

Any litigation relating to this License may be brought only in the
courts of a jurisdiction where the defendant maintains its principal
place of business and such litigation shall be governed by laws of that
jurisdiction, without reference to its conflict-of-law provisions.
Nothing in this Section shall prevent a party's ability to bring
cross-claims or counter-claims.

9. Miscellaneous
----------------

This License represents the complete agreement concerning the subject
matter hereof. If any provision of this License is held to be
unenforceable, such provision shall be reformed only to the extent
necessary to make it enforceable. Any law or regulation which provides
that the language of a contract shall be construed against the drafter
shall not be used to construe this License against a Contributor.

10. Versions of the License
---------------------------

10.1. New Versions

Mozilla Foundation is the license steward. Except as provided in Section
10.3, no one other than the license steward has the right to modify or
publish new versions of this License. Each version will be given a
distinguishing version number.

10.2. Effect of New Versions

You may distribute the Covered Software under the terms of the version
of the License under which You originally received the Covered Software,
or under the terms of any subsequent version published by the license
steward.

10.3. Modified Versions

If you create software not governed by this License, and you want to
create a new license for such software, you may create and use a
modified version of this License if you rename the license and remove
any references to the name of the license steward (except to note that
such modified license differs from this License).

10.4. Distributing Source Code Form that is Incompatible With Secondary
Licenses

If You choose to distribute Source Code Form that is Incompatible With
Secondary Licenses under the terms of this version of the License, the
notice described in Exhibit B of this License must be attached.

Exhibit A - Source Code Form License Notice
-------------------------------------------

  This Source Code Form is subject to the terms of the Mozilla Public
  License, v. 2.0. If a copy of the MPL was not distributed with this
  file, You can obtain one at http://mozilla.org/MPL/2.0/.

If it is not possible or desirable to put the notice in a particular
file, then You may include the notice in a location (such as a LICENSE
file in a relevant directory) where a recipient would be likely to look
for such a notice.

You may add additional accurate notices of copyright ownership.

Exhibit B - "Incompatible With Secondary Licenses" Notice
---------------------------------------------------------

  This Source Code Form is "Incompatible With Secondary Licenses", as
  defined by the Mozilla Public License, v. 2.0.
`,
		},
		fileMetaData{
			path: 		"go.mod",
			fileMode: 	420,
			content:	`module {{.PkgPath}}

go 1.14
`,
		},
		fileMetaData{
			path: 		"main.go",
			fileMode: 	420,
			content:	`package main

import (
	"{{.PkgPath}}/{{.ProviderName}}"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	// remove date and time stamp from log output as the plugin SDK already adds its own
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return {{.ProviderName}}.Provider()
		},
	})
}
`,
		},
		fileMetaData{
			path: 		"scripts/changelog-links.sh",
			fileMode: 	493,
			content:	fmt.Sprintf(`#!/bin/bash

# This script rewrites [GH-nnnn]-style references in the CHANGELOG.md file to
# be Markdown links to the given github issues.
#
# This is run during releases so that the issue references in all of the
# released items are presented as clickable links, but we can just use the
# easy [GH-nnnn] shorthand for quickly adding items to the "Unrelease" section
# while merging things between releases.

set -e

if [[ ! -f CHANGELOG.md ]]; then
  echo "ERROR: CHANGELOG.md not found in pwd."
  echo "Please run this from the root of the terraform provider repository"
  exit 1
fi

if [[ %[1]suname%[1]s == "Darwin" ]]; then
  echo "Using BSD sed"
  SED="sed -i.bak -E -e"
else
  echo "Using GNU sed"
  SED="sed -i.bak -r -e"
fi

PROVIDER_URL="https://{{.PkgPath}}/issues"

$SED "s;GH-([0-9]+);\[#\1\]\($PROVIDER_URL\/\1\);g" -e 's/\[\[#(.+)([0-9])\)]$/(\[#\1\2))/g' CHANGELOG.md

rm CHANGELOG.md.bak
`, "`"),
		},
		fileMetaData{
			path: 		"scripts/gofmtcheck.sh",
			fileMode: 	493,
			content:	fmt.Sprintf(`#!/usr/bin/env bash

# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."

# This filter should match the search filter in ../GNUMakefile
gofmt_files=$(find . -name '*.go' | grep -v vendor | xargs gofmt -l)
if [ -n "${gofmt_files}" ]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \%[1]smake fmt\%[1]s to reformat code."
    exit 1
fi

exit 0`, "`"),
		},
		fileMetaData{
			path: 		"scripts/terrafmt-acctests.sh",
			fileMode: 	493,
			content:	`#!/usr/bin/env bash

echo "==> Checking acceptance test terraform blocks are formatted..."

files=$(find ./{{.ProviderName}} -type f -name "*_test.go")
error=false

for f in $files; do
  terrafmt diff -c -q -f "$f" || error=true
done

if ${error}; then
  exit 1
fi

exit 0
`,
		},
		fileMetaData{
			path: 		"scripts/terrafmt-website.sh",
			fileMode: 	493,
			content:	`#!/usr/bin/env bash

echo "==> Checking documentation terraform blocks are formatted..."

files=$(find ./website -type f -name "*.html.markdown")
error=false

for f in $files; do
  terrafmt diff -c -q "$f" || error=true
done

if ${error}; then
  echo "------------------------------------------------"
  echo ""
  echo "The preceding files contain terraform blocks that are not correctly formatted or contain errors."
  echo "You can fix this by running make tools and then terrafmt on them."
  echo ""
  echo "format a single file:"
  echo "$ terrafmt fmt ./website/path/to/file.html.markdown"
  echo ""
  echo "format all website files:"
  echo "$ find . | egrep html.markdown | sort | while read f; do terrafmt fmt \$f; done"
  echo ""
  echo "on windows:"
  echo "$ Get-ChildItem -Path . -Recurse -Filter \"*html.markdown\" | foreach {terrafmt fmt $_.fullName}"
  echo ""
  exit 1
fi

exit 0
`,
		},
		fileMetaData{
			path: 		"website/allowed-subcategories",
			fileMode: 	420,
			content:	``,
		},
		fileMetaData{
			path: 		"{{.ProviderName}}/internal/provider/provider.go",
			fileMode: 	420,
			content:	`package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{},

		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   map[string]*schema.Resource{},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		return nil, nil
	}
}
`,
		},
		fileMetaData{
			path: 		"{{.ProviderName}}/provider.go",
			fileMode: 	420,
			content:	`package {{.ProviderName}}

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"{{.PkgPath}}/{{.ProviderName}}/internal/provider"
)

func Provider() terraform.ResourceProvider {
	return provider.Provider()
}
`,
		},
	}
)

func GenScaffold(outdir string, data interface{}) error {
	// ensure outputdir
	if _, err := os.Stat(outdir); os.IsNotExist(err) {
		if err := os.Mkdir(outdir, os.ModePerm); err != nil {
			return err
		}
	}

	// prepare directories
	for _, di := range templateDirs {
		var buffer bytes.Buffer
		pt := template.Must(template.New("").Parse(di.path))
		if err := pt.Execute(&buffer, data); err != nil {
			return err
		}
		dir := filepath.Join(outdir, buffer.String())
		if err := os.MkdirAll(dir, di.fileMode); err != nil {
			return err
		}
	}

	// prepare files
	for _, fi := range templateFiles {
		var buffer bytes.Buffer
		pt := template.Must(template.New("").Parse(fi.path))
		if err := pt.Execute(&buffer, data); err != nil {
			return err
		}
		path := filepath.Join(outdir, buffer.String())
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fi.fileMode)
		if err != nil {
			return err
		}
		defer f.Close()

		t, err := template.New("").Parse(fi.content)
		if err != nil {
			return err
		}
		if err := t.Execute(f, data); err != nil {
			return err
		}
	}

	return nil
}