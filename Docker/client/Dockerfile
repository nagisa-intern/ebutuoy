FROM node:10.9.0

ENV NODE_ENV=production
ENV HOST 0.0.0.0

RUN mkdir -p /client
COPY ../ebutuoy-nuxt /client
WORKDIR /client

EXPOSE 3000

RUN yarn install --production
RUN yarn run build

CMD ["yarn", "start"]