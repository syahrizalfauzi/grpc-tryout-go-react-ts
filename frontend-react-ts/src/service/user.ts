import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { UserClient } from "../proto/user.client";

const ENVOY_URL = "http://localhost:8080";

const transport = new GrpcWebFetchTransport({
  baseUrl: ENVOY_URL,
});

export const userService = new UserClient(transport);
