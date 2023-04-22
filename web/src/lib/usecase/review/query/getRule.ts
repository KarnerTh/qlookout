import { gql } from "@apollo/client/core";
import { query } from "svelte-apollo";
import type { RuleModel } from "../ruleModel";

const ruleQuery = gql`
  query Rule($id: Int!) {
    rule(id: $id) {
      id
      lookoutId
      columnName
      columnType
      rowIndex
      exactValue
      greaterThan
      lessThan
      shouldBeNull
    }
  }
`;

export const useRule = (id: number) =>
  query<{ rule: RuleModel }>(ruleQuery, {
    variables: { id },
  });
