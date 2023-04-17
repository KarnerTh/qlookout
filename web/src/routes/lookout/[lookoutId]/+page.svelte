<script lang="ts">
  import { afterNavigate, goto } from "$app/navigation";
  import { page } from "$app/stores";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import LoadingSpinner from "$lib/components/loading/LoadingSpinner.svelte";
  import { useLookout } from "$lib/usecase/lookout/query/getLookout";
  import LookoutBaseInfo from "./LookoutBaseInfo.svelte";
  import LookoutQueryInfo from "./LookoutQueryInfo.svelte";
  import LookoutRuleInfo from "./LookoutRuleInfo.svelte";

  const lookout = useLookout(+$page.params.lookoutId);

  afterNavigate(() => {
    if (history.state.refetch) {
      lookout.refetch();
    }
  });
</script>

<PageHeader title={"Lookout Detail"} backAction={() => goto("/lookout")} />

{#if $lookout.loading}
  <div class="flex justify-center pt-9">
    <LoadingSpinner />
  </div>
{:else if $lookout.data}
  <div class="flex flex-wrap gap-2">
    <div class="max-w-sm">
      <LookoutBaseInfo lookout={$lookout.data.lookout} />
    </div>

    <div class="grow">
      <LookoutQueryInfo query={$lookout.data.lookout.query} />
    </div>
  </div>

  <div class="mt-2">
    <LookoutRuleInfo rules={$lookout.data.lookout.rules} />
  </div>
{/if}
