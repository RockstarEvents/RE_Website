FROM oven/bun AS builder

WORKDIR /app

COPY package.json ./

RUN bun install --frozen-lockfile

COPY . .

ENV NEXT_TELEMETRY_DISABLED 1

RUN bun run build

FROM oven/bun AS runner

WORKDIR /app

ENV NODE_ENV production

ENV NEXT_TELEMETRY_DISABLED 1

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json

RUN adduser --system --uid 1001 nextjs

RUN chown nextjs:bun .next

USER nextjs

EXPOSE 3000

ENV HOSTNAME "0.0.0.0"

CMD ["bun", "run", "start"]