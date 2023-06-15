import {
  ApolloClient,
  HttpLink,
  InMemoryCache,
  split,
} from "@apollo/client/core";
import { setClient } from "svelte-apollo";
import { getMainDefinition } from "@apollo/client/utilities";
import { browser } from "$app/environment";
import { WebSocketLink } from "@apollo/client/link/ws";
import { SubscriptionClient } from "subscriptions-transport-ws";

export const setupGraphQl = () => {
  const httpLink = new HttpLink({
    uri: "http://localhost:63001/query",
  });

  const wsLink = browser
    ? new WebSocketLink(
        new SubscriptionClient("ws://localhost:63001/query", {
          reconnect: true,
        })
      )
    : undefined;

  const link = wsLink
    ? split(
        ({ query }) => {
          const definition = getMainDefinition(query);
          return (
            definition.kind === "OperationDefinition" &&
            definition.operation === "subscription"
          );
        },
        wsLink,
        httpLink
      )
    : httpLink;

  const client = new ApolloClient({
    link: link,
    cache: new InMemoryCache(),
  });

  setClient(client);
};
