# Go checksum

Command line and library to manage files integrity checksums

As lib :
- Support plenty of algorithms 
- Dynamic algorithm creation and usage 
- Ease hashing usage of readers and files
- Can create binaries like sha*sum with few lines of code 
- Create easily checksum files
- Detect checksum files 
- Validate files against detected checksum

+ As a binary :
- mimic `sha1sum` but for all algo
- process directory trees to 
  - `set` sum in a **sumfile** or in the **filename**
  - `set watch` tree for new file and set sum
  - `check` that sum is valid
  - `remove` sum from **sumfile** or from **filename**
  - `list` files that will be processed


configuration file :
```yaml
pattern: (?i)\.*$         # pattern for files to process
patternIsExclusive: true  # use pattern as inclusive or exclusive
hash: sha1
strategy: sumfile        # `sumfile` or `filename` store sum to a sumefile and directly in the source filename
```

```bash
# work like `sha1sum`
fi sum -H sha1 testfile.txt


# list files that will be hashed and checked
fi -c ./fi.yaml list test/
# set sum to filesum or filename
fi set test/
# watch for new files in the tree and set sum
fi set watch test/
# remove filesum or sum from filename
fi remove test/
# check that sum matches file
fi check test/
```

Supported algorithms : Blake2b-256, Blake2b-384, Blake2b-512, Blake2s-256, Ripemd160, Md4, Md5, Sha1, Sha256, Sha384, Sha3-224, Sha3-256, Sha3-384, Sha3-512, Sha512, Sha512-224, Sha512-256, Crc32-ieee, Crc64-iso, Crc64-ecma, Adler32, Fnv32, Fnv32a, Fnv64, Fnv64a, Fnv128, Fnv128a, Xor8, Fletch16, Luhn, Sum16, Sum32, Sum64, Crc8, Crc16-ccitt,