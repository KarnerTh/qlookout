import { gql } from "@apollo/client/core";
import { mutation } from "svelte-apollo";

const ruleDeleteMutation = gql`
  mutation DeleteRule(
    $id: Int!
  ) {
    deleteRule(
      id: $id
    ) {
      id
    }
  }
`;

export const useRuleDeleteMutation = () => mutation(ruleDeleteMutation);
