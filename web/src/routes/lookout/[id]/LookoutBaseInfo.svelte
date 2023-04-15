<script lang="ts">
  import Card from "$lib/components/card/Card.svelte";
  import type { LookoutConfigModel } from "$lib/usecase/lookout/lookoutConfigModel";

  export let lookout: LookoutConfigModel;

  let notifications: string[] = [];
  $: {
    notifications = [];
    if (lookout.notifyLocal) notifications.push("Local");
    if (lookout.notifyMail) notifications.push("Mail");
  }
</script>

<Card>
  <h5 class="mb-2 text-xl font-medium leading-tight text-neutral-800">
    {lookout.name}
  </h5>
  <p class="text-base text-neutral-600">
    {lookout.cron}
  </p>

  <h6 class="mb-2 mt-4 text-xl font-medium leading-tight text-neutral-800">
    Notifications
  </h6>
  <p class="text-base text-neutral-600 flex">{notifications.join(", ")}</p>
</Card>
