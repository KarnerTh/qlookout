import { gql } from "@apollo/client/core";
import { mutation } from "svelte-apollo";

const ruleMutation = gql`
  mutation CreateRule(
    $lookoutId: Int!
    $columnName: String!
    $columnType: String!
    $rowIndex: Int!
    $exactValue: String
    $lessThan: String
    $greaterThan: String
    $shouldBeNull: Boolean
  ) {
    createRule(
      data: {
        lookoutId: $lookoutId
        columnName: $columnName
        columnType: $columnType
        rowIndex: $rowIndex
        exactValue: $exactValue
        lessThan: $lessThan
        greaterThan: $greaterThan
        shouldBeNull: $shouldBeNull
      }
    ) {
      id
    }
  }
`;

export const useRuleCreateMutation = () => mutation(ruleMutation);
