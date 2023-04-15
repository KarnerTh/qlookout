import { gql } from "@apollo/client/core";
import { query } from "svelte-apollo";
import type { LookoutConfigModel } from "../lookoutConfigModel";

const lookoutQuery = gql`
  query Lookout($id: Int!) {
    lookout(id: $id) {
      id
      name
      cron
      query
      notifyLocal
      notifyMail
    }
  }
`;

export const useLookout = (id: number) =>
  query<{ lookout: LookoutConfigModel }>(lookoutQuery, { variables: { id } });
