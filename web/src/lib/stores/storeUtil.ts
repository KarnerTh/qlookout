import type { Subscriber, Unsubscriber } from "svelte/store";

export type StoreType<T> = (
  this: void,
  run: Subscriber<T>,
  invalidate?: (value?: T) => void
) => Unsubscriber;
