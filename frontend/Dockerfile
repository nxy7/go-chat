FROM oven/bun:latest as dev

RUN npm i -g pnpm

WORKDIR /app

ADD ./package.json .
ADD ./pnpm-lock.yaml .
RUN pnpm i 
ADD . .

#############################
FROM oven/bun:latest as builder

WORKDIR /app

ADD ./package.json .
ADD ./pnpm-lock.yaml .
RUN bun i 
ADD . .
RUN bun run build
#############################
FROM joseluisq/static-web-server:2 as prod
COPY --from=builder /app/build /pages

EXPOSE 3000/tcp
CMD ["--root" ,"/pages" ,"--page-fallback" ,"/pages/200.html", "--compression", "--log-level", "trace", "--port", "3000"]
