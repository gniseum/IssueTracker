# pull official base image
FROM node:latest

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm install axios
RUN npm install react-bootstrap@next bootstrap@5.1.0

# add app
COPY . ./

# start app
CMD ["npm", "start"]

