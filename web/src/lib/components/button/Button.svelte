<script lang="ts">
  import { iconData, type Icon } from "../icons";

  export let title: string;
  export let leadingIcon: Icon | null = null;
  export let trailingIcon: Icon | null = null;
  export let type: "submit" | "reset" | "button" = "button";
</script>

<span class="rounded-md shadow-sm">
  <button
    class="flex justify-center items-center w-full px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:text-gray-500 focus:outline-none active:bg-gray-50 active:text-gray-800"
    {type}
    on:click
  >
    {#if leadingIcon}
      <span class="pr-2">
        {@html iconData[leadingIcon]}
      </span>
    {/if}
    <span>{title}</span>
    {#if trailingIcon}
      <span class="pl-2">
        {@html iconData[trailingIcon]}
      </span>
    {/if}
  </button>
</span>
