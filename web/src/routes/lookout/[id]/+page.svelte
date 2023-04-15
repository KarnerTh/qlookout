<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import Card from "$lib/components/card/Card.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import LoadingSpinner from "$lib/components/loading/LoadingSpinner.svelte";
  import { useLookout } from "$lib/usecase/lookout/query/getLookout";
    import { getBoolIcon } from "$lib/util/boolEmojiUtil";

  const lookout = useLookout(+$page.params.id);
</script>

<PageHeader title={"Lookout Detail"} backAction={() => goto("/lookout")} />

{#if $lookout.loading}
  <div class="flex justify-center pt-9">
    <LoadingSpinner />
  </div>
{:else}
  <div class="flex flex-wrap gap-2">
    <Card>
      <h5 class="mb-2 text-xl font-medium leading-tight text-neutral-800">
        {$lookout.data?.lookout.name}
      </h5>
      <p class="text-base text-neutral-600">
        {$lookout.data?.lookout.cron}
      </p>

      <h6 class="mb-2 mt-4 text-xl font-medium leading-tight text-neutral-800">
        Notifications
      </h6>
      <p class="text-base text-neutral-600">
        {getBoolIcon($lookout.data?.lookout.notifyLocal)} Local
      </p>
      <p class="text-base text-neutral-600">
        {getBoolIcon($lookout.data?.lookout.notifyMail)} Mail
      </p>
    </Card>
    <div class="grow">
      <Card>
        <h5 class="mb-2 text-xl font-medium leading-tight text-neutral-800">
          Query
        </h5>
        <div class=" max-h-80 overflow-auto">
          <code class="text-base text-neutral-600">
            {$lookout.data?.lookout.query}
          </code>
        </div>
      </Card>
    </div>
  </div>
  {JSON.stringify($lookout)}
{/if}
