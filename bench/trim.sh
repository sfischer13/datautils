DICT=/usr/share/dict/american-english

bench "sed -re 's/^\s+//' <$DICT >/dev/null" "trim -l <$DICT >/dev/null"
echo "---"
bench "sed -re 's/\s+$//' <$DICT >/dev/null" "trim -r <$DICT >/dev/null"
