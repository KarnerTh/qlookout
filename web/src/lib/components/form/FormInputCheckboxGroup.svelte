<script lang="ts">
  import FormInputCheckbox from "./FormInputCheckbox.svelte";

  export let label: string;
  export let items: { name: string; title: string; description?: string, initialValue?: boolean }[];
</script>

<div class="w-full px-3 mb-6 md:mb-0">
  <span
    class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
  >
    {label}
  </span>
  <ul class="grid w-full gap-6 md:grid-cols-2">
    {#each items as item}
      <li>
        <FormInputCheckbox
          name={item.name}
          title={item.title}
          checked={item.initialValue}
          description={item.description}
        />
      </li>
    {/each}
  </ul>
</div>
