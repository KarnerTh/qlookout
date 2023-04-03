<script lang="ts">
  import { page } from "$app/stores";
  import { iconData, type Icon } from "../icons";

  export let link: string;
  export let title: string;
  export let icon: Icon;
  export let badge: string | undefined = undefined;

  $: isSelected = $page.url.pathname.includes(link);
</script>

<li>
  <a
    href={link}
    class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 border-l-4 border-transparent hover:border-red-400 pr-6"
    class:text-red-500={isSelected}
    class:hover:text-gray-800={!isSelected}
    class:hover:text-red-500={isSelected}
  >
    <span class="inline-flex justify-center items-center ml-4">
      {@html iconData[icon]}
    </span>
    <span class="ml-2 text-sm tracking-wide truncate">{title}</span>
    {#if badge}
      <span
        class="px-2 py-0.5 ml-auto text-xs font-medium tracking-wide text-red-500 bg-red-50 rounded-full"
      >
        {badge}
      </span>
    {/if}
  </a>
</li>
