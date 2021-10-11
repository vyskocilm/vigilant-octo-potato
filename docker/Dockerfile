FROM registry.opensuse.org/opensuse/golang:1.16.3 AS builder

# to enable Go testsuite under openSUSE
RUN cp /usr/etc/services /etc/services

ENV PATH /usr/local/go/bin:$PATH

RUN set -eux; \
    cd /usr/local; \
    git clone --single-branch --no-tags --depth 1 --branch master https://github.com/golang/go; \
    printf "DEV.GO2GO-REVISION=%s\n" $(git -C /usr/local/go rev-parse --short HEAD);

ENV GOARCH amd64
ENV GOOS linux

RUN set -eux; \
    cd /usr/local/go/src; \
    ./make.bash; \
    go install std; \
# go install -race std;
# remove system Go
# and a few intermediate / bootstrapping files the official binary release tarballs do not contain
     zypper --non-interactive remove go; \
     rm -rf \
        /usr/local/go/pkg/*/cmd \
        /usr/local/go/pkg/bootstrap \
        /usr/local/go/pkg/obj \
        /usr/local/go/pkg/tool/*/api \
        /usr/local/go/pkg/tool/*/go_bootstrap \
        /usr/local/go/src/cmd/dist/dist \
     ; \
     \
     printf "DEV.GO2GO-VERSION=%s\n" "$(go version)";

#FROM scratch
#COPY --from=builder /usr/local/go /usr/local/go

