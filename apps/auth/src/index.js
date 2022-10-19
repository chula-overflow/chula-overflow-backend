const PROTO_PATH = __dirname + '/../../proto/auth.proto'

const grpc = require('@grpc/grpc-js')
const protoLoader = require('@grpc/proto-loader')

const crypto = require('crypto')

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
})

const auth = grpc.loadPackageDefinition(packageDefinition).auth

function login(call, cb) {
    const id = crypto.randomUUID()
    
    cb(null, {token: id})
}

function getServer() {
    var server = new grpc.Server()
    server.addService(auth.Auth.service, {
        login: login
    })
    return server
}

const server = getServer()

server.bindAsync("localhost:3001", grpc.ServerCredentials.createInsecure(), () => {
    server.start()
    console.log("Server started")
})