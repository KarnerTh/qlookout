import { gql } from "@apollo/client/core";
import { mutation } from "svelte-apollo";

const lookoutCreateMutation = gql`
  mutation CreateLookout(
    $name: String!
    $query: String!
    $cron: String!
    $notifyLocal: Boolean!
    $notifyMail: Boolean!
  ) {
    createLookout(
      data: {
        name: $name
        query: $query
        cron: $cron
        notifyMail: $notifyMail
        notifyLocal: $notifyLocal
      }
    ) {
      id
    }
  }
`;

export const useLookoutCreateMutation = () => mutation(lookoutCreateMutation);
