<script lang="ts">
  import type { TableRow } from "./table-data";
  import { createEventDispatcher } from "svelte";
  import { iconData } from "../icons";

  const dispatch = createEventDispatcher<{
    rowClicked: { id: number };
    checkBoxClicked: { id: number };
    checkBoxAllClicked: { selected: boolean };
  }>();

  export let columns: string[];
  export let rows: TableRow[];
  export let selectedIds: number[] = [];
  export let showCheckBox: boolean = false;

  let allCheckboxActive = false;

  const onRowClicked = (id: number): void => {
    dispatch("rowClicked", { id });
  };

  const onCheckboxClicked = (id: number): void => {
    dispatch("checkBoxClicked", { id });
  };

  const onCheckBoxAllClicked = (): void => {
    allCheckboxActive = allCheckboxActive ? false : true;
    dispatch("checkBoxAllClicked", { selected: allCheckboxActive });
  };

  $: {
    allCheckboxActive = selectedIds.length === rows.length;
  }
</script>

<table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
  <thead
    class="text-xs text-gray-700 sticky top-0 uppercase bg-gray-200 dark:bg-gray-700 dark:text-gray-400"
  >
    <tr>
      {#if showCheckBox}
        <th scope="col" class="p-4">
          <div class="flex items-center">
            <input
              type="checkbox"
              class="w-4 h-4 accent-red-500"
              checked={allCheckboxActive}
              indeterminate={!allCheckboxActive && selectedIds.length > 0}
              on:click={() => onCheckBoxAllClicked()}
            />
          </div>
        </th>
      {/if}
      {#each columns as column}
        <th scope="col" class="px-6 py-3"> {column} </th>
      {/each}
    </tr>
  </thead>
  <tbody>
    {#each rows as rowEntry}
      <tr
        class="bg-white cursor-pointer border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
        on:click={() => onRowClicked(rowEntry.id)}
      >
        {#if showCheckBox}
          <td class="w-4 p-4">
            <div class="flex items-center">
              <input
                type="checkbox"
                class="w-4 h-4 accent-red-500"
                checked={selectedIds.includes(rowEntry.id)}
                on:click={(event) => {
                  event.stopPropagation();
                  onCheckboxClicked(rowEntry.id);
                }}
              />
            </div>
          </td>
        {/if}
        {#each rowEntry.data as value}
          <td class="px-6 py-4">
            {#if value.type === "boolean"}
              {@html iconData[value.value ? "check" : "x_mark"]}
            {:else}
              {value.value}
            {/if}
          </td>
        {/each}
      </tr>
    {/each}
  </tbody>
</table>
