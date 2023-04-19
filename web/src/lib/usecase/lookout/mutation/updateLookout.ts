import { gql } from "@apollo/client/core";
import { mutation } from "svelte-apollo";

const lookoutUpdateMutation = gql`
  mutation UpdateLookout(
    $id: Int!
    $name: String
    $query: String
    $cron: String
    $notifyLocal: Boolean
    $notifyMail: Boolean
  ) {
    updateLookout(
      id: $id
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

export const useLookoutUpdateMutation = () => mutation(lookoutUpdateMutation);
