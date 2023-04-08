import { gql } from "@apollo/client/core";
import { query } from "svelte-apollo";
import type { LookoutConfigModel } from "../lookoutConfigModel";

const lookoutsQuery = gql`
  query Lookouts {
    lookouts {
      id
      name
      cron
      notifyLocal
      notifyMail
    }
  }
`;

export const getLookouts = () =>
  query<{ lookouts: LookoutConfigModel[] }>(lookoutsQuery);
