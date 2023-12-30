#! /bin/bash
files=( "$@" )

for file in "${files[@]}" ; do
  temp="${file}".tmp
  title="$(basename $file ".md")"
  printf -- '---\ntitle: '${title}'\ndate: 2023-01-01T01:01:01-08:00\ndraft: true\n---\n' > "${temp}"
  awk 'BEGIN{toggle=0;} {if ($0 ~ /^---/) { toggle = (toggle+1)%2 } ; if ((toggle == 0) && !($0 ~ /^---/)) {print $0;}}' "${file}" >> "${temp}"
  mv "${temp}" "${file}"
done
