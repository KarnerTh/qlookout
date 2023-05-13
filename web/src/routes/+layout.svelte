<script lang="ts">
  import "../app.css";
  import NavigationItem from "../lib/components/navigation/NavigationItem.svelte";
  import logo from "$lib/assets/logo.svg";
  import { setupGraphQl } from "$lib/graphql_setup";
  import { useNewNotificationSubscription } from "$lib/usecase/notify/subscription/newNotification";
  import { notifications } from "$lib/stores/notification.store";

  setupGraphQl();

  // add notifications to store
  const notificationSubscription = useNewNotificationSubscription();
  $: {
    if ($notificationSubscription.data) {
      notifications.add($notificationSubscription.data.newNotification);
    }
  }
</script>

<div
  class="min-h-screen flex flex-col flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800"
>
  <div class="fixed flex flex-col top-0 left-0 w-64 bg-white h-full border-r">
    <div class="flex flex-col items-center justify-center border-b">
      <img alt="logo" class="h-24 my-2" src={logo} />
      <div class="font-bold mb-2">Query-Lookout</div>
    </div>
    <div class="overflow-y-auto overflow-x-hidden flex-grow">
      <ul class="flex flex-col py-4 space-y-1">
        <!-- <NavigationHeader title="Menu" /> -->
        <!-- <NavigationItem icon="dashboard" title="Dashboard" link="/dashboard" /> -->
        <NavigationItem icon="lookouts" title="Lookouts" link="/lookout" />
        <NavigationItem
          icon="notification"
          title="Notifications"
          link="/notifications"
          badge={$notifications.length
            ? $notifications.length.toString()
            : undefined}
        />
        <!-- <NavigationHeader title="Settings" /> -->
        <!-- <NavigationItem icon="settings" title="Settings" link="/settings" /> -->
        <!-- <NavigationItem icon="logout" title="Logout" link="/logout" /> -->
      </ul>
    </div>
  </div>
  <div class="fixed right-1 left-64 top-0 bottom-0 m-3">
    <slot />
  </div>
</div>

<style lang="postcss">
  :global(html) {
    background-color: theme(colors.gray.100);
  }
</style>
