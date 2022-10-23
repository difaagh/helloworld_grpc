import * as grpc from "@grpc/grpc-js";
import { GreeterClient } from "./protos/helloworld_grpc_pb";
import { GreeterService, IGreeterServer } from "./protos/helloworld_grpc_pb";
import { HelloReply, HelloRequest } from "./protos/helloworld_pb";

describe("client_nodejs", () => {
  it("should getMessage", function (done) {
    const service: IGreeterServer = {
      sayHello(
        call: grpc.ServerUnaryCall<HelloRequest, HelloReply>,
        callback: grpc.sendUnaryData<HelloReply>
      ) {
        if (call.request) {
          const reply = new HelloReply();
          reply.setMessage("Hello " + call.request.getName());
          callback(null, reply);
        }
      },
    };
    const server = new grpc.Server();
    server.addService(GreeterService, service);
    server.bindAsync(
      "localhost:9000",
      grpc.ServerCredentials.createInsecure(),
      (err: Error | null, port: number) => {
        expect(err).toBe(null);
        expect(port).toBe(9000);
        server.start();
        const client = new GreeterClient(
          "localhost:9000",
          grpc.ChannelCredentials.createInsecure()
        );
        const r = new HelloRequest();
        r.setName("node");
        client.sayHello(
          r,
          (err: grpc.ServiceError | null, response?: HelloReply) => {
            expect(err).toBe(null);
            expect(response?.getMessage()).toBe("Hello node");
            client.close();
            server.forceShutdown();
            done();
          }
        );
      }
    );
  });
});
