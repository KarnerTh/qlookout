<script lang="ts">
  import FormInputField from "$lib/components/form/FormInputField.svelte";
  import type { RuleType } from "../ruleType";

  export let type: RuleType;
</script>

{#if type === "exact"}
  <FormInputField name="exactValue" required label="Exact value" />
{:else if type === "null"}
  <input type="hidden" name="shouldBeNull" value={true} />
{:else if type === "less"}
  <FormInputField name="lessThan" required label="Less than" />
{:else if type === "greater"}
  <FormInputField name="greaterThan" required label="Greater than" />
{:else if type === "between"}
  <FormInputField name="greaterThan" required label="Greater than" />
  <FormInputField name="lessThan" required label="Less than" />
{/if}
