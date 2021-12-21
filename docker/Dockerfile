FROM --platform=${BUILDPLATFORM} golang:latest AS base

ARG TARGETOS
ARG TARGETARCH
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}

WORKDIR /app
COPY . .

FROM base as run
ENTRYPOINT ["/bin/bash", "-c"]

FROM base as build
RUN make

FROM scratch as bin
COPY --from=build /app/out ./
