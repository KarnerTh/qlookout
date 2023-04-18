<script lang="ts">
  import FormInputField from "$lib/components/form/FormInputField.svelte";
  import type { RuleType } from "../ruleType";

  export let ruleType: RuleType;
  export let inputType: "text" | "number";
</script>

{#if ruleType === "exact"}
  <FormInputField
    name="exactValue"
    required
    label="Exact value"
    type={inputType}
  />
{:else if ruleType === "null"}
  <input name="shouldBeNull" value={true} type="hidden" />
{:else if ruleType === "less"}
  <FormInputField name="lessThan" required label="Less than" type={inputType} />
{:else if ruleType === "greater"}
  <FormInputField
    name="greaterThan"
    required
    label="Greater than"
    type={inputType}
  />
{:else if ruleType === "between"}
  <FormInputField
    name="greaterThan"
    required
    label="Greater than"
    type={inputType}
  />
  <FormInputField name="lessThan" required label="Less than" type={inputType} />
{/if}
