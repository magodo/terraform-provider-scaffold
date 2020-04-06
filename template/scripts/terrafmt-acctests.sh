#!/usr/bin/env bash

echo "==> Checking acceptance test terraform blocks are formatted..."

files=$(find ./pkg -type f -name "*_test.go")
error=false

for f in $files; do
  terrafmt diff -c -q -f "$f" || error=true
done

if ${error}; then
  echo "------------------------------------------------"
  echo ""
  echo "The preceding files contain terraform blocks that are not correctly formatted or contain errors."
  echo "You can fix this by running make tools and then terrafmt on them."
  echo ""
  echo "format a single file:"
  echo "$ terrafmt fmt -f ./{{.ProviderName}}/path/to/_test.go"
  echo ""
  echo "format all website files:"
  echo "$ find pkg | egrep \"_test.go\" | sort | while read f; do terrafmt fmt -f \$f; done"
  echo ""
  echo "on windows:"
  echo "$ Get-ChildItem -Path . -Recurse -Filter \"*_test.go\" | foreach {terrafmt fmt -f $_.fullName}"
  echo ""
  exit 1
fi

exit 0
