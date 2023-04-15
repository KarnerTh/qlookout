import { gql } from "@apollo/client/core";
import { query } from "svelte-apollo";
import type { LookoutConfigDetailModel } from "../lookoutConfigModel";

const lookoutQuery = gql`
  query Lookout($id: Int!) {
    lookout(id: $id) {
      id
      name
      cron
      query
      notifyLocal
      notifyMail
      rules {
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
  }
`;

export const useLookout = (id: number) =>
  query<{ lookout: LookoutConfigDetailModel }>(lookoutQuery, {
    variables: { id },
  });
