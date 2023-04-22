<script lang="ts">
  import FormInputField from "$lib/components/form/FormInputField.svelte";
  import type { RuleModel } from "$lib/usecase/review/ruleModel";
  import type { RuleType } from "./ruleType";

  export let ruleType: RuleType | undefined;
  export let inputType: "text" | "number";
  export let value: RuleModel | undefined = undefined;
</script>

{#if ruleType === "exact"}
  <FormInputField
    name="exactValue"
    required
    value={value?.exactValue}
    label="Exact value"
    type={inputType}
  />
{:else if ruleType === "null"}
  <input name="shouldBeNull" value={true} type="hidden" />
{:else if ruleType === "less"}
  <FormInputField
    name="lessThan"
    required
    value={value?.lessThan}
    label="Less than"
    type={inputType}
  />
{:else if ruleType === "greater"}
  <FormInputField
    name="greaterThan"
    required
    value={value?.greaterThan}
    label="Greater than"
    type={inputType}
  />
{:else if ruleType === "between"}
  <FormInputField
    name="greaterThan"
    required
    value={value?.greaterThan}
    label="Greater than"
    type={inputType}
  />
  <FormInputField
    name="lessThan"
    required
    value={value?.lessThan}
    label="Less than"
    type={inputType}
  />
{/if}
