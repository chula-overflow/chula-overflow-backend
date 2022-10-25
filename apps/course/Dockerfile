FROM node:18-alpine As build

WORKDIR /app

ENV NODE_ENV production

COPY --chown=node:node apps/course .

COPY apps/proto /proto

RUN npm i -g @nestjs/cli && npm install && npm cache clean --force && npx @nestjs/cli build

###################
# PRODUCTION
###################

FROM node:18-alpine As production

USER node

WORKDIR /app

COPY --chown=node:node apps/proto /proto
COPY --chown=node:node --from=build /app/node_modules ./node_modules
COPY --chown=node:node --from=build /app/dist ./dist

CMD [ "node", "dist/main.js" ]
