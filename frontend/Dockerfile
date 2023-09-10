FROM oven/bun:latest as dev

WORKDIR /app

ADD ./package.json .
ADD ./bun.lockb .
RUN bun i 
ADD . .

#############################
FROM oven/bun:latest as builder

WORKDIR /app

ADD ./package.json .
ADD ./bun.lockb .
RUN bun i 
ADD . .
RUN bun run generate
#############################
FROM joseluisq/static-web-server:2 as prod
COPY --from=builder /app/.output/public /pages

EXPOSE 3000/tcp
CMD ["--root" ,"/pages" ,"--page-fallback" ,"/pages/200.html", "--compression", "--log-level", "trace", "--port", "3000"]
