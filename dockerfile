# 1 choose a compiler OS
FROM golang:alpine AS builder
# 2 (optional) label the compiler image
LABEL stage=builder
# 3 (optional) install any compiler-only dependencies
RUN apk add --no-cache gcc libc-dev
WORKDIR /workspace
# 4 copy all the source files
COPY . .
# 5 build the GO program
RUN CGO_ENABLED=0 GOOS=linux go build -a
# 6 choose a runtime OS
FROM redis:alpine AS final
# 7 
WORKDIR /

COPY --from=builder /workspace/negigo .
# 9 execute the program upon start 
CMD [ "./negigo" ]