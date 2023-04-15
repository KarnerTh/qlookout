import type { Icon } from "$lib/components/icons";

export const getBoolIcon = (value: boolean | undefined): Icon => {
  return value ? "check" : "x_mark";
};
