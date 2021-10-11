# Go generics examples

## Slides

Normal markdown, to see them in a console, just use `present`

```sh
python3 -mvenv _present_
source _present_/bin/activate
pip install present
```

## Build go 1.18

```sh
make -C docker/
```

```sh
cat <<EOF > ~/bin/go2
#!/bin/sh

exec docker run --rm -v "${PWD}":/tmp local/golang2:latest /usr/local/go/bin/go "${1}" -gcflags=-G=3 /tmp/"${2}"
EOF
```

## Content

 * 01/ no generics
 * 02/ generic new2
 * 03/ type sets in constraints
 * 04/ interface constraints
 * 05/ generic types
 * 06/ befare of int
 * 07/ comparable types
 * 08/ more on type sets
 * 09/ composition types
 * 10/ mutually referencing type parameters
 * 11/ type conversions and untyped constants
 * 12/ interface types in union
 * 13/ empty type sets
 * 14/ (function argument) type inference
 * 15/ constraint type inference
 * 16/ pointer method example
 * 17/ type asserts and a reflec package
