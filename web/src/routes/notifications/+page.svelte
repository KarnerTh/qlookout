<script lang="ts">
  import Button from "$lib/components/button/Button.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import NotificationListItem from "$lib/components/notification/NotificationListItem.svelte";
  import { notifications } from "$lib/stores/notification.store";
</script>

<div class="flex flex-col max-h-full">
  <PageHeader title="Notifications" />
  <div class="flex items-center pb-3">
    <span class="grow" />
    <Button
      title="Clear"
      leadingIcon="trash"
      on:click={() => {
        notifications.clear();
      }}
    />
  </div>

  <div class="overflow-auto mb-2">
    {#if $notifications.length === 0}
      <div class="flex items-center justify-center pt-5">No notifications</div>
    {:else}
      {#each $notifications as notification (notification)}
        <div class="mb-2">
          <NotificationListItem {notification} />
        </div>
      {/each}
    {/if}
  </div>
</div>
