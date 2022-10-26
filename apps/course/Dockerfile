###################
# BUILD FOR LOCAL DEVELOPMENT
###################

FROM node:18-alpine As development

WORKDIR /usr/src/app

COPY --chown=node:node apps/course/package*.json ./

RUN npm ci

COPY --chown=node:node apps/course .

USER node

###################
# BUILD FOR PRODUCTION
###################

FROM node:18-alpine As build

WORKDIR /usr/src/app

COPY --chown=node:node apps/course/package*.json ./

COPY --chown=node:node --from=development /usr/src/app/node_modules ./node_modules
COPY --chown=node:node apps/course .

RUN npx npm ci --only=production && npm cache clean --force && npx @nestjs/cli build

ENV NODE_ENV production

USER node

###################
# PRODUCTION
###################

FROM node:18-alpine As production

WORKDIR /app

COPY --chown=node:node apps/proto /proto
COPY --chown=node:node --from=build /usr/src/app/node_modules ./node_modules
COPY --chown=node:node --from=build /usr/src/app/dist ./dist

CMD [ "node", "dist/main.js" ]
