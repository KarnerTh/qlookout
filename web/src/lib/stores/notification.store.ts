import type { NotificationModel } from "$lib/usecase/notify/notificationModel";
import { writable } from "svelte/store";
import type { StoreType } from "./storeUtil";

interface NotificationStore {
  subscribe: StoreType<NotificationModel[]>;
  add: (entry: NotificationModel) => void;
  clear: () => void;
}

const createNotificationStore = (): NotificationStore => {
  const { subscribe, update, set } = writable<NotificationModel[]>([]);

  const add = (entry: NotificationModel): void => {
    update((list) => [entry, ...list]);
  };

  const clear = (): void => {
    set([]);
  };

  return {
    subscribe,
    add,
    clear,
  };
};

export const notifications = createNotificationStore();
