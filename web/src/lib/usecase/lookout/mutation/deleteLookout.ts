import { gql } from "@apollo/client/core";
import { mutation } from "svelte-apollo";

const lookoutDeleteMutation = gql`
  mutation DeleteLookout(
    $id: Int!
  ) {
    deleteLookout(
      id: $id
    ) {
      id
    }
  }
`;

export const useLookoutDeleteMutation = () => mutation(lookoutDeleteMutation);
