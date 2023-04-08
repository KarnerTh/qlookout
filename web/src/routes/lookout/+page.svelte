<script lang="ts">
  import Dropdown from "$lib/components/dropdown/Dropdown.svelte";
  import DropdownItem from "$lib/components/dropdown/DropdownItem.svelte";
  import { iconData } from "$lib/components/icons";
  import Table from "$lib/components/table/Table.svelte";
  import { goto } from "$app/navigation";
  import { getLookouts } from "$lib/usecase/lookout/query/getLookouts";
  import { convertLookoutConfigModelToTableData } from "$lib/usecase/lookout/lookoutConfigModel";
  import LoadingSpinner from "$lib/components/loading/LoadingSpinner.svelte";

  let selectedIds: number[] = [];
  const lookouts = getLookouts();

  $: tableData = $lookouts.data
    ? $lookouts.data.lookouts.map(convertLookoutConfigModelToTableData)
    : [];

  const onRowClicked = (id: number): void => {
    goto(`/lookout/${id}`);
  };

  const onCheckboxClicked = (id: number): void => {
    if (selectedIds.includes(id)) {
      selectedIds = selectedIds.filter((item) => item !== id);
    } else {
      selectedIds = [...selectedIds, id];
    }
  };

  const onCheckboxAllClicked = (selected: boolean): void => {
    if (selected) {
      selectedIds = tableData.map((item) => item.id);
    } else {
      selectedIds = [];
    }
  };
</script>

<h1 class="mb-3 text-4xl font-extrabold tracking-tight">Lookouts</h1>

<div class="overflow-x-auto">
  <div class="flex items-center justify-between pb-3">
    <Dropdown title="Action">
      <DropdownItem title="Set active" />
      <DropdownItem title="Set inactive" />
    </Dropdown>

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
  </div>
  {#if $lookouts.loading}
    <div class="flex justify-center pt-9">
      <LoadingSpinner />
    </div>
  {:else}
    <Table
      on:rowClicked={(event) => onRowClicked(event.detail.id)}
      on:checkBoxAllClicked={(event) =>
        onCheckboxAllClicked(event.detail.selected)}
      on:checkBoxClicked={(event) => onCheckboxClicked(event.detail.id)}
      {selectedIds}
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
