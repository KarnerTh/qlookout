<script lang="ts">
  import { goto } from "$app/navigation";
  import type { NotificationModel } from "$lib/usecase/notify/notificationModel";
  import Card from "../card/Card.svelte";

  export let notification: NotificationModel;

  const formatDate = (date: Date): string => {
    return date.toLocaleDateString(undefined, {
      hour: "numeric",
      minute: "numeric",
      second: "numeric",
    });
  };
</script>

<Card cursor on:click={() => goto(`/lookout/${notification.lookoutId}`)}>
  <div class="flex gap-2">
    <div class="flex flex-col">
      <div class="font-bold">{notification.title}</div>
      <div class="">{notification.description}</div>
    </div>
    <div class="flex-grow" />
    <span> {formatDate(new Date(notification.timestamp))} </span>
  </div>
</Card>
