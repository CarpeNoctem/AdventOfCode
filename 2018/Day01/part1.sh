offset=0; for line in $(cat input.txt); do offset=$(echo "$offset $line" | bc); done; echo $offset
