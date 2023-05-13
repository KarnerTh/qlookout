import type { NotificationModel } from "$lib/usecase/notify/notificationModel";
import { writable } from "svelte/store";
import type { StoreType } from "./storeUtil";
import { browser } from "$app/environment";

interface NotificationStore {
  subscribe: StoreType<NotificationModel[]>;
  add: (entry: NotificationModel) => void;
  clear: () => void;
}

const createNotificationStore = (): NotificationStore => {
  let storedNotifications: NotificationModel[] | undefined;

  if (browser) {
    storedNotifications = !!localStorage.notifications
      ? JSON.parse(localStorage.notifications)
      : null;
  }

  const { subscribe, update, set } = writable<NotificationModel[]>(
    storedNotifications || []
  );

  const add = (entry: NotificationModel): void => {
    update((list) => [entry, ...list]);
  };

  const clear = (): void => {
    set([]);
  };

  if (browser) {
    subscribe((value) => (localStorage.notifications = JSON.stringify(value)));
  }

  return {
    subscribe,
    add,
    clear,
  };
};

export const notifications = createNotificationStore();
