#!/bin/bash

while true; do
  m=$((RANDOM % 10 + 1))
  n=$((RANDOM % 10 + 1))

  K=""
  for ((i=1; i<=m; i++)); do
    K="$K $((RANDOM % 100 + 1))"
  done
  K="${K:1}"  

  L=""
  for ((i=1; i<=n; i++)); do
    L="$L $((RANDOM % 100 + 1))"
  done
  L="${L:1}"

  echo "=== Новый цикл ==="
  echo "m = $m"
  echo "K = $K"
  echo "n = $n"
  echo "L = $L"

  input=$(cat <<EOF
$m
$K
$n
$L
EOF
)

  echo "$input" | /usr/bin/factor_dev 

  echo "---------------------"

  sleep $((RANDOM % 5 + 1))
done