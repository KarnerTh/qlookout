<script lang="ts">
  import Table from "$lib/components/table/Table.svelte";
  import { afterNavigate, goto } from "$app/navigation";
  import { useLookouts } from "$lib/usecase/lookout/query/getLookouts";
  import { convertLookoutConfigModelToTableData } from "$lib/usecase/lookout/lookoutConfigModel";
  import LoadingSpinner from "$lib/components/loading/LoadingSpinner.svelte";
  import Button from "$lib/components/button/Button.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";

  const lookouts = useLookouts();

  $: tableData = $lookouts.data
    ? $lookouts.data.lookouts.map(convertLookoutConfigModelToTableData)
    : [];

  const onRowClicked = (id: number): void => {
    goto(`/lookout/${id}`);
  };

  afterNavigate(() => {
    if (history.state.refetch) {
      lookouts.refetch();
    }
  });
</script>

<div class="flex flex-col max-h-full">
  <PageHeader title="Lookouts" />

  <div class="flex items-center pb-3">
    <span class="grow" />
    <Button
      title="Create"
      leadingIcon="plus"
      on:click={() => goto("/lookout/create")}
    />
  </div>
  {#if $lookouts.loading}
    <div class="flex justify-center pt-9">
      <LoadingSpinner />
    </div>
  {:else}
    <div class="overflow-auto mb-2">
      <Table
        on:rowClicked={(event) => onRowClicked(event.detail.id)}
        columns={[
          "Lookout name",
          "Cron",
          "# Rules",
          "Notify local",
          "Notify mail",
        ]}
        rows={tableData}
      />
    </div>
  {/if}
</div>
