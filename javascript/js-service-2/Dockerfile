FROM node:16.15.1-alpine

ENV NODE_ENV=production

WORKDIR /app

COPY package.json package-lock.json server.js app.js /app/

RUN npm ci

CMD ["node", "server.js"]
