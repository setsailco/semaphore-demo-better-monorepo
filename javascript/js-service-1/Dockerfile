FROM node:16.15.1-alpine

ENV NODE_ENV=production

WORKDIR /app

COPY pnpm-lock.yaml pnpm-workspace.yaml javascript/js-service-1/server.js javascript/js-service-1/app.js /app/

RUN corepack enable && pnpm install --prod --frozen-lockfile --filter ...js-service-1

CMD ["node", "server.js"]
