FROM --platform=${BUILDPLATFORM} golang:latest AS base

ARG TARGETOS
ARG TARGETARCH
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}

WORKDIR /app
COPY src .
RUN make vet

FROM base as run
ENTRYPOINT ["/bin/bash", "-c"]
CMD make run

FROM base as build
RUN make build

FROM scratch as bin
COPY --from=build /app/out ./
