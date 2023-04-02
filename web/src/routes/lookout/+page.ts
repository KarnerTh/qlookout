import type { PageLoad } from "./$types"

export interface LookoutItem {
  id: number;
  title: string;
}

export const load: PageLoad = () => {
  return {
    lookouts: [
      { id: 1, title: "First works" },
      { id: 2, title: "Second works" },
      { id: 3, title: "Third works" },
    ]
  }
}
