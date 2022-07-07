FROM node:16.15.1-alpine

ENV NODE_ENV=production

WORKDIR /app

COPY pnpm-lock.yaml pnpm-workspace.yaml javascript/js-service-2/server.js javascript/js-service-2/app.js /app/

RUN corepack enable && pnpm install --prod --frozen-lockfile --filter ...js-service-2

CMD ["node", "server.js"]
