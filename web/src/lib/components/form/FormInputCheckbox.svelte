<script lang="ts">
  export let name: string;
  export let title: string;
  export let checked: boolean = false;
  export let description: string | null = null;
</script>

<div>
  <input
    type="checkbox"
    {name}
    id={name}
    value=""
    {checked}
    class="hidden peer"
  />
  <label
    for={name}
    class="inline-flex items-center justify-between w-full p-3 text-gray-500 bg-white border-2 border-gray-200 rounded-lg cursor-pointer dark:hover:text-gray-300 dark:border-gray-700 peer-checked:border-red-600 hover:text-gray-600 dark:peer-checked:text-gray-300 peer-checked:text-gray-600 hover:bg-gray-50 dark:text-gray-400 dark:bg-gray-800 dark:hover:bg-gray-700"
  >
    <div class="block">
      <div class="w-full text-sm font-semibold">{title}</div>
      <div class="w-full text-xs">
        {#if description}
          {description}
        {:else}
          {@html "&nbsp;"}
        {/if}
      </div>
    </div>
  </label>
</div>
