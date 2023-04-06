import type { TableRow } from "$lib/table/table-data"
import type { PageLoad } from "./$types"

export interface LookoutPageData {
  lookouts: TableRow[]
}

export const load: PageLoad<LookoutPageData> = () => {
  return {
    lookouts: [
      {
        id: 1,
        data: [
          { type: "text", value: "test 1" },
          { type: "text", value: "000000" },
          { type: "number", value: 2 },
          { type: "boolean", value: true },
          { type: "boolean", value: false },
        ],
      },
      {
        id: 2,
        data: [
          { type: "text", value: "test 2" },
          { type: "text", value: "000000" },
          { type: "number", value: 1 },
          { type: "boolean", value: true },
          { type: "boolean", value: true },
        ],
      },
      {
        id: 3,
        data: [
          { type: "text", value: "test works" },
          { type: "text", value: "000000" },
          { type: "number", value: 1 },
          { type: "boolean", value: false },
          { type: "boolean", value: false },
        ],
      },
    ]
  }
}
