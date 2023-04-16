<script lang="ts">
  import { iconData } from "$lib/components/icons";
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

<PageHeader title="Lookouts" />

<div class="overflow-x-auto">
  <div class="flex items-center pb-3">
    <div class="relative">
      <div
        class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none"
      >
        <span class="pr-2">
          {@html iconData.search}
        </span>
      </div>
      <input
        type="text"
        id="table-search"
        class="block p-2 pl-10 text-sm text-gray-900 focus:outline-none border border-gray-300 rounded-lg w-80 bg-gray-50 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        placeholder="Search for items"
      />
    </div>

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
  {/if}
</div>
