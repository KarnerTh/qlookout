import { gql } from "@apollo/client/core";
import { mutation } from "svelte-apollo";

const ruleUpdateMutation = gql`
  mutation UpdateRule(
    $id: Int!
    $columnName: String
    $columnType: String
    $rowIndex: Int!
    $exactValue: String
    $lessThan: String
    $greaterThan: String
    $shouldBeNull: Boolean
  ) {
    updateRule(
      id: $id
      data: {
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

export const useRuleUpdateMutation = () => mutation(ruleUpdateMutation);
