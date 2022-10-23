import * as grpc from "@grpc/grpc-js";
import { GreeterClient } from "./protos/helloworld_grpc_pb";
import { HelloReply, HelloRequest } from "./protos/helloworld_pb";

const client = new GreeterClient(
  "localhost:8080",
  grpc.ChannelCredentials.createInsecure()
);

const r = new HelloRequest();
r.setName("node");

client.sayHello(r, (err: grpc.ServiceError | null, response?: HelloReply) => {
  if (err) {
    console.log("cannot sayHello: ", err.message);
    return;
  }
  console.log("Received: ", response?.getMessage());
});
