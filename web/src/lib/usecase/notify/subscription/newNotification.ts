import { gql } from "@apollo/client/core";
import { subscribe } from "svelte-apollo";
import type { NotificationModel } from "../notificationModel";

const newNotificationSubscription = gql`
  subscription newNotificationSubscription {
    newNotification {
      title
      description
      timestamp
    }
  }
`;

export const useNewNotificationSubscription = () =>
  subscribe<{ newNotification: NotificationModel }>(
    newNotificationSubscription
  );
