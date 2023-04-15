<script lang="ts">
  import type { TableRow } from "./table-data";
  import { createEventDispatcher } from "svelte";
  import { getBoolIcon } from "$lib/util/boolEmojiUtil";

  const dispatch = createEventDispatcher<{
    rowClicked: { id: number };
    checkBoxClicked: { id: number };
    checkBoxAllClicked: { selected: boolean };
  }>();

  export let columns: string[];
  export let rows: TableRow[];
  export let selectedIds: number[];

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
    class="text-xs text-gray-700 uppercase bg-gray-200 dark:bg-gray-700 dark:text-gray-400"
  >
    <tr>
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
      {#each columns as column}
        <th scope="col" class="px-6 py-3"> {column} </th>
      {/each}
    </tr>
  </thead>
  <tbody>
    {#each rows as rowEntry}
      <tr
        class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
        on:click={() => onRowClicked(rowEntry.id)}
      >
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
        {#each rowEntry.data as value}
          <td class="px-6 py-4">
            {#if value.type === "boolean"}
              {getBoolIcon(value.value)}
            {:else}
              {value.value}
            {/if}
          </td>
        {/each}
      </tr>
    {/each}
  </tbody>
</table>
