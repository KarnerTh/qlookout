<script lang="ts">
  import "../app.css";
  import NavigationHeader from "../lib/navigation/NavigationHeader.svelte";
  import NavigationItem from "../lib/navigation/NavigationItem.svelte";
</script>

<div
  class="min-h-screen flex flex-col flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800"
>
  <div class="fixed flex flex-col top-0 left-0 w-64 bg-white h-full border-r">
    <div class="flex items-center justify-center h-14 border-b">
      <div class="font-bold">Query-Lookout</div>
    </div>
    <div class="overflow-y-auto overflow-x-hidden flex-grow">
      <ul class="flex flex-col py-4 space-y-1">
        <NavigationHeader title="Menu" />
        <NavigationItem icon="dashboard" title="Dashboard" link="/dashboard" />
        <NavigationItem icon="lookouts" title="Lookouts" link="/lookout" />
        <NavigationItem
          icon="notification"
          title="Notifications"
          link="/notifications"
          badge="12"
        />
        <NavigationHeader title="Settings" />
        <NavigationItem icon="settings" title="Settings" link="/settings" />
        <NavigationItem icon="logout" title="Logout" link="/logout" />
      </ul>
    </div>
  </div>
  <div class="fixed right-0 left-64 bg-green-50">
    <slot />
  </div>
</div>

<style lang="postcss">
  :global(html) {
    background-color: theme(colors.gray.100);
  }
</style>
